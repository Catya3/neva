package desugarer

import (
	"errors"
	"maps"

	"github.com/nevalang/neva/internal/compiler"
	src "github.com/nevalang/neva/pkg/sourcecode"
)

type handleNetResult struct {
	desugaredConnections []src.Connection     // desugared network
	virtualConstants     map[string]src.Const // constants that needs to be inserted in to make desugared network work
	virtualNodes         map[string]src.Node  // nodes that needs to be inserted in to make desugared network work
	usedNodePorts        nodePortsMap         // ports that were used in processed network
}

func (d Desugarer) handleNetwork( //nolint:funlen
	net []src.Connection,
	nodes map[string]src.Node,
	scope src.Scope,
) (handleNetResult, *compiler.Error) {
	nodesToInsert := map[string]src.Node{}
	desugaredConns := make([]src.Connection, 0, len(net))
	usedNodePorts := newNodePortsMap()
	constantsToInsert := map[string]src.Const{}

	for _, conn := range net {
		// array bypass connections don't need desugaring
		if arrBypass := conn.ArrayBypass; arrBypass != nil {
			usedNodePorts.set(
				arrBypass.SenderOutport.Node,
				arrBypass.SenderOutport.Port,
			)
			usedNodePorts.set(
				arrBypass.ReceiverInport.Node,
				arrBypass.ReceiverInport.Port,
			)
			desugaredConns = append(desugaredConns, conn)
			continue
		}

		// here only care normal connections
		normConn := conn.Normal

		// mark port address as used, if it's not const
		if normConn.SenderSide.PortAddr != nil {
			usedNodePorts.set(
				normConn.SenderSide.PortAddr.Node,
				normConn.SenderSide.PortAddr.Port,
			)
		}

		// check if it's actually nothing to desugar
		if normConn.SenderSide.Const == nil &&
			len(normConn.SenderSide.Selectors) == 0 &&
			len(normConn.ReceiverSide.DeferredConnections) == 0 {
			desugaredConns = append(desugaredConns, conn)
			continue
		}

		// === Actual desugaring starts here ===

		// desugar selectors if they are present
		if len(normConn.SenderSide.Selectors) != 0 {
			result, err := d.desugarStructSelectors(
				*conn.Normal,
				nodes,
				scope,
			)
			if err != nil {
				return handleNetResult{}, compiler.Error{
					Err:      errors.New("Cannot desugar struct selectors"),
					Location: &scope.Location,
					Meta:     &conn.Meta,
				}.Wrap(err)
			}
			nodesToInsert[result.nodeToInsertName] = result.nodeToInsert
			constantsToInsert[result.constToInsertName] = result.constToInsert
			conn = result.connToReplace
			desugaredConns = append(desugaredConns, result.connToInsert)
		}

		// handle const sender if needed
		// we add virtual nodes and constants if needed
		// and replace current connection with connection without const sender
		if normConn.SenderSide.Const != nil { //nolint:nestif
			if normConn.SenderSide.Const.Ref != nil {
				result, err := d.handleConstRefSender(conn, scope)
				if err != nil {
					return handleNetResult{}, err
				}
				nodesToInsert[result.emitterNodeName] = result.emitterNode
				conn = result.connectionWithoutConstSender
			} else if normConn.SenderSide.Const.Message != nil {
				result, err := d.handleLiteralSender(conn)
				if err != nil {
					return handleNetResult{}, err
				}
				constantsToInsert[result.constName] = *normConn.SenderSide.Const
				nodesToInsert[result.emitterNodeName] = result.emitterNode
				conn = result.connectionWithoutConstSender
			}
		}

		// the last thing we could desugar is deferred connections
		if len(normConn.ReceiverSide.DeferredConnections) == 0 {
			desugaredConns = append(desugaredConns, conn)
			continue
		}

		deferredConnsResult, err := d.handleDeferredConnections(
			*normConn,
			nodes,
			scope,
		)
		if err != nil {
			return handleNetResult{}, err
		}

		// handleThenConns recursively calls this function so it returns the same structure
		maps.Copy(usedNodePorts.m, deferredConnsResult.usedNodesPorts.m)
		maps.Copy(constantsToInsert, deferredConnsResult.virtualConstants)
		maps.Copy(nodesToInsert, deferredConnsResult.virtualNodes)

		// note that we discard original connection, it's desugared version is there
		desugaredConns = append(
			desugaredConns,
			deferredConnsResult.desugaredConnections...,
		)
	}

	return handleNetResult{
		desugaredConnections: desugaredConns,
		usedNodePorts:        usedNodePorts,
		virtualConstants:     constantsToInsert,
		virtualNodes:         nodesToInsert,
	}, nil
}
