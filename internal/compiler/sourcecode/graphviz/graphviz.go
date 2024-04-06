package graphviz

import (
	"fmt"
	"io"
	"strings"

	"github.com/nevalang/neva/internal/compiler/sourcecode"
)

type Node struct {
	Path     string
	MetaText string
}

func (n Node) Label() string { return n.MetaText }

func (n Node) Name() string { return fmt.Sprint(n.Path, "/", n.MetaText) }

// func MakeConstNode(c sourcecode.Const) Node {
// 	fmt.Printf("%#v\n", c)
// 	return Node{
// 		Node: c.Ref.Name,
// 		Port: fmt.Sprint("$", c.Ref.Name),
// 	}
// }

type Edge struct {
	Send Node
	Recv Node

	// HeadCluster is set for compound connections to deferred clusters.
	HeadCluster string
}

func (e Edge) Label() string {
	return ""
}

type Cluster struct {
	b *ClusterBuilder

	// Path is constructed by taking the name of the entities (usually components)
	// concatenated with '/'. The main purpose is to generate a unique node name
	// for the nodes in the path.
	Path string

	Index    int
	Nodes    map[Node]struct{}
	Clusters map[string]*Cluster
}

func (c *Cluster) Label() string {
	i := strings.LastIndexByte(c.Path, '/')
	if i == -1 {
		return c.Path
	}
	return c.Path[i+1:]
}

func (c *Cluster) AddCluster(name string) *Cluster {
	if c.Clusters == nil {
		c.Clusters = map[string]*Cluster{}
	}
	cluster, ok := c.Clusters[name]
	if !ok {
		cluster = &Cluster{b: c.b, Index: c.b.genNextId()}
		cluster.Path = fmt.Sprint(c.Path, "/", name)
		c.Clusters[name] = cluster
	}
	return cluster
}

func (c *Cluster) GraphvizName() string {
	return fmt.Sprintf("cluster_%d", c.Index)
}

func (c *Cluster) AddNode(n Node) {
	if c.Nodes == nil {
		c.Nodes = map[Node]struct{}{}
	}
	n.Path = c.Path
	c.Nodes[n] = struct{}{}
}

type ClusterBuilder struct {
	Main   *Cluster
	Edges  map[Edge]struct{}
	nextId int
}

func (b *ClusterBuilder) genNextId() (id int) {
	id = b.nextId
	b.nextId++
	return id
}

func (b *ClusterBuilder) init() {
	b.Main = &Cluster{b: b, Index: b.genNextId()}
}

func (c *ClusterBuilder) AddEdge(e Edge) {
	if c.Edges == nil {
		c.Edges = map[Edge]struct{}{}
	}
	c.Edges[e] = struct{}{}
}

func (b *ClusterBuilder) AddEntity(parent *Cluster, name string, e sourcecode.Entity) error {
	if parent == nil {
		b.init()
		parent = b.Main
	}
	switch k := e.Kind; k {
	case "component_entity":
		// c := parent.AddCluster(e.Component.Meta.Text)
		if err := b.AddComponent(parent, name, e.Component); err != nil {
			return err
		}
	case "const_entity":
		if err := b.AddConst(parent, name, e.Const); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unexpected entity kind: %q", k)
	}
	return nil
}

func (b *ClusterBuilder) AddComponent(parent *Cluster, name string, c sourcecode.Component) error {
	sub := parent.AddCluster(name)
	for n := range c.Nodes {
		sub.AddCluster(n)
	}
	for _, c := range c.Net {
		if err := b.AddConnection(sub, c); err != nil {
			return err
		}
	}
	return nil
}

func (b *ClusterBuilder) AddConnection(parent *Cluster, c sourcecode.Connection) error {
	switch {
	case c.ArrayBypass != nil:
		var e Edge
		e.Send = Node{"", c.ArrayBypass.SenderOutport.Meta.Text}
		e.Recv = Node{"", c.ArrayBypass.ReceiverInport.Meta.Text}
		parent.AddNode(e.Send)
		parent.AddNode(e.Recv)
		b.AddEdge(e)
	case c.Normal != nil:
		send := Node{"", c.Normal.SenderSide.Meta.Text}
		parent.AddNode(send)
		for _, r := range c.Normal.ReceiverSide.Receivers {
			recv := Node{"", r.Meta.Text}
			parent.AddNode(recv)
			b.AddEdge(Edge{send, recv, ""})
		}
		for _, d := range c.Normal.ReceiverSide.DeferredConnections {
			sub := parent.AddCluster("<deferred>")
			deferredSend := Node{"", c.Meta.Text}
			if err := b.AddConnection(sub, d); err != nil {
				return err
			}
			b.AddEdge(Edge{send, deferredSend, sub.GraphvizName()})
		}
	default:
		return fmt.Errorf("bad connection: missing normal or array_bypass spec")
	}
	return nil
}

func (b *ClusterBuilder) AddConst(parent *Cluster, name string, c sourcecode.Const) error {
	parent.AddNode(Node{"", name})
	return nil
}

func (b *ClusterBuilder) Build(w io.Writer) error {
	if _, err := fmt.Fprintln(w, "digraph G {"); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "  compound = true;"); err != nil {
		return err
	}
	if err := b.recursiveBuild(w, "  ", b.Main); err != nil {
		return err
	}
	for e := range b.Edges {
		if _, err := fmt.Fprintf(w, "  %q -> %q", e.Send.Name(), e.Recv.Name()); err != nil {
			return err
		}
		if e.Label() != "" || e.HeadCluster != "" {
			if _, err := fmt.Fprintf(w, "["); err != nil {
				return err
			}
			if label := e.Label(); label != "" {
				if _, err := fmt.Fprintf(w, "label = %q;", label); err != nil {
					return err
				}
			}
			if e.HeadCluster != "" {
				if _, err := fmt.Fprintf(w, "lhead = %q;", e.HeadCluster); err != nil {
					return err
				}
			}
			if _, err := fmt.Fprintf(w, "]"); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintln(w, ";"); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintln(w, "}"); err != nil {
		return err
	}
	return nil
}

func (b *ClusterBuilder) recursiveBuild(w io.Writer, indent string, c *Cluster) error {
	fprintlnIndent := func(f string, a ...any) error {
		_, err := fmt.Fprintln(w, indent, fmt.Sprintf(f, a...))
		return err
	}
	if err := fprintlnIndent("subgraph %s {", c.GraphvizName()); err != nil {
		return err
	}
	if label := c.Label(); label != "" {
		if err := fprintlnIndent("  label = %q;", label); err != nil {
			return err
		}
	}
	for n := range c.Nodes {
		if err := fprintlnIndent("  %q [label = \"%s\";];", n.Name(), n.Label()); err != nil {
			return err
		}
	}
	for _, sub := range c.Clusters {
		if err := b.recursiveBuild(w, indent+"  ", sub); err != nil {
			return err
		}
	}
	return fprintlnIndent("}")
}
