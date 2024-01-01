// Code generated from ./neva.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // neva
import "github.com/antlr4-go/antlr/v4"

// BasenevaListener is a complete listener for a parse tree produced by nevaParser.
type BasenevaListener struct{}

var _ nevaListener = &BasenevaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasenevaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasenevaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasenevaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasenevaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasenevaListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasenevaListener) ExitProg(ctx *ProgContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasenevaListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasenevaListener) ExitStmt(ctx *StmtContext) {}

// EnterCompilerDirectives is called when production compilerDirectives is entered.
func (s *BasenevaListener) EnterCompilerDirectives(ctx *CompilerDirectivesContext) {}

// ExitCompilerDirectives is called when production compilerDirectives is exited.
func (s *BasenevaListener) ExitCompilerDirectives(ctx *CompilerDirectivesContext) {}

// EnterCompilerDirective is called when production compilerDirective is entered.
func (s *BasenevaListener) EnterCompilerDirective(ctx *CompilerDirectiveContext) {}

// ExitCompilerDirective is called when production compilerDirective is exited.
func (s *BasenevaListener) ExitCompilerDirective(ctx *CompilerDirectiveContext) {}

// EnterCompilerDirectivesArgs is called when production compilerDirectivesArgs is entered.
func (s *BasenevaListener) EnterCompilerDirectivesArgs(ctx *CompilerDirectivesArgsContext) {}

// ExitCompilerDirectivesArgs is called when production compilerDirectivesArgs is exited.
func (s *BasenevaListener) ExitCompilerDirectivesArgs(ctx *CompilerDirectivesArgsContext) {}

// EnterCompiler_directive_arg is called when production compiler_directive_arg is entered.
func (s *BasenevaListener) EnterCompiler_directive_arg(ctx *Compiler_directive_argContext) {}

// ExitCompiler_directive_arg is called when production compiler_directive_arg is exited.
func (s *BasenevaListener) ExitCompiler_directive_arg(ctx *Compiler_directive_argContext) {}

// EnterImportStmt is called when production importStmt is entered.
func (s *BasenevaListener) EnterImportStmt(ctx *ImportStmtContext) {}

// ExitImportStmt is called when production importStmt is exited.
func (s *BasenevaListener) ExitImportStmt(ctx *ImportStmtContext) {}

// EnterImportDef is called when production importDef is entered.
func (s *BasenevaListener) EnterImportDef(ctx *ImportDefContext) {}

// ExitImportDef is called when production importDef is exited.
func (s *BasenevaListener) ExitImportDef(ctx *ImportDefContext) {}

// EnterImportPath is called when production importPath is entered.
func (s *BasenevaListener) EnterImportPath(ctx *ImportPathContext) {}

// ExitImportPath is called when production importPath is exited.
func (s *BasenevaListener) ExitImportPath(ctx *ImportPathContext) {}

// EnterEntityRef is called when production entityRef is entered.
func (s *BasenevaListener) EnterEntityRef(ctx *EntityRefContext) {}

// ExitEntityRef is called when production entityRef is exited.
func (s *BasenevaListener) ExitEntityRef(ctx *EntityRefContext) {}

// EnterTypeStmt is called when production typeStmt is entered.
func (s *BasenevaListener) EnterTypeStmt(ctx *TypeStmtContext) {}

// ExitTypeStmt is called when production typeStmt is exited.
func (s *BasenevaListener) ExitTypeStmt(ctx *TypeStmtContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (s *BasenevaListener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (s *BasenevaListener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterTypeParams is called when production typeParams is entered.
func (s *BasenevaListener) EnterTypeParams(ctx *TypeParamsContext) {}

// ExitTypeParams is called when production typeParams is exited.
func (s *BasenevaListener) ExitTypeParams(ctx *TypeParamsContext) {}

// EnterTypeParamList is called when production typeParamList is entered.
func (s *BasenevaListener) EnterTypeParamList(ctx *TypeParamListContext) {}

// ExitTypeParamList is called when production typeParamList is exited.
func (s *BasenevaListener) ExitTypeParamList(ctx *TypeParamListContext) {}

// EnterTypeParam is called when production typeParam is entered.
func (s *BasenevaListener) EnterTypeParam(ctx *TypeParamContext) {}

// ExitTypeParam is called when production typeParam is exited.
func (s *BasenevaListener) ExitTypeParam(ctx *TypeParamContext) {}

// EnterTypeExpr is called when production typeExpr is entered.
func (s *BasenevaListener) EnterTypeExpr(ctx *TypeExprContext) {}

// ExitTypeExpr is called when production typeExpr is exited.
func (s *BasenevaListener) ExitTypeExpr(ctx *TypeExprContext) {}

// EnterTypeInstExpr is called when production typeInstExpr is entered.
func (s *BasenevaListener) EnterTypeInstExpr(ctx *TypeInstExprContext) {}

// ExitTypeInstExpr is called when production typeInstExpr is exited.
func (s *BasenevaListener) ExitTypeInstExpr(ctx *TypeInstExprContext) {}

// EnterTypeArgs is called when production typeArgs is entered.
func (s *BasenevaListener) EnterTypeArgs(ctx *TypeArgsContext) {}

// ExitTypeArgs is called when production typeArgs is exited.
func (s *BasenevaListener) ExitTypeArgs(ctx *TypeArgsContext) {}

// EnterTypeLitExpr is called when production typeLitExpr is entered.
func (s *BasenevaListener) EnterTypeLitExpr(ctx *TypeLitExprContext) {}

// ExitTypeLitExpr is called when production typeLitExpr is exited.
func (s *BasenevaListener) ExitTypeLitExpr(ctx *TypeLitExprContext) {}

// EnterEnumTypeExpr is called when production enumTypeExpr is entered.
func (s *BasenevaListener) EnterEnumTypeExpr(ctx *EnumTypeExprContext) {}

// ExitEnumTypeExpr is called when production enumTypeExpr is exited.
func (s *BasenevaListener) ExitEnumTypeExpr(ctx *EnumTypeExprContext) {}

// EnterArrTypeExpr is called when production arrTypeExpr is entered.
func (s *BasenevaListener) EnterArrTypeExpr(ctx *ArrTypeExprContext) {}

// ExitArrTypeExpr is called when production arrTypeExpr is exited.
func (s *BasenevaListener) ExitArrTypeExpr(ctx *ArrTypeExprContext) {}

// EnterStructTypeExpr is called when production structTypeExpr is entered.
func (s *BasenevaListener) EnterStructTypeExpr(ctx *StructTypeExprContext) {}

// ExitStructTypeExpr is called when production structTypeExpr is exited.
func (s *BasenevaListener) ExitStructTypeExpr(ctx *StructTypeExprContext) {}

// EnterStructFields is called when production structFields is entered.
func (s *BasenevaListener) EnterStructFields(ctx *StructFieldsContext) {}

// ExitStructFields is called when production structFields is exited.
func (s *BasenevaListener) ExitStructFields(ctx *StructFieldsContext) {}

// EnterStructField is called when production structField is entered.
func (s *BasenevaListener) EnterStructField(ctx *StructFieldContext) {}

// ExitStructField is called when production structField is exited.
func (s *BasenevaListener) ExitStructField(ctx *StructFieldContext) {}

// EnterUnionTypeExpr is called when production unionTypeExpr is entered.
func (s *BasenevaListener) EnterUnionTypeExpr(ctx *UnionTypeExprContext) {}

// ExitUnionTypeExpr is called when production unionTypeExpr is exited.
func (s *BasenevaListener) ExitUnionTypeExpr(ctx *UnionTypeExprContext) {}

// EnterNonUnionTypeExpr is called when production nonUnionTypeExpr is entered.
func (s *BasenevaListener) EnterNonUnionTypeExpr(ctx *NonUnionTypeExprContext) {}

// ExitNonUnionTypeExpr is called when production nonUnionTypeExpr is exited.
func (s *BasenevaListener) ExitNonUnionTypeExpr(ctx *NonUnionTypeExprContext) {}

// EnterInterfaceStmt is called when production interfaceStmt is entered.
func (s *BasenevaListener) EnterInterfaceStmt(ctx *InterfaceStmtContext) {}

// ExitInterfaceStmt is called when production interfaceStmt is exited.
func (s *BasenevaListener) ExitInterfaceStmt(ctx *InterfaceStmtContext) {}

// EnterInterfaceDef is called when production interfaceDef is entered.
func (s *BasenevaListener) EnterInterfaceDef(ctx *InterfaceDefContext) {}

// ExitInterfaceDef is called when production interfaceDef is exited.
func (s *BasenevaListener) ExitInterfaceDef(ctx *InterfaceDefContext) {}

// EnterInPortsDef is called when production inPortsDef is entered.
func (s *BasenevaListener) EnterInPortsDef(ctx *InPortsDefContext) {}

// ExitInPortsDef is called when production inPortsDef is exited.
func (s *BasenevaListener) ExitInPortsDef(ctx *InPortsDefContext) {}

// EnterOutPortsDef is called when production outPortsDef is entered.
func (s *BasenevaListener) EnterOutPortsDef(ctx *OutPortsDefContext) {}

// ExitOutPortsDef is called when production outPortsDef is exited.
func (s *BasenevaListener) ExitOutPortsDef(ctx *OutPortsDefContext) {}

// EnterPortsDef is called when production portsDef is entered.
func (s *BasenevaListener) EnterPortsDef(ctx *PortsDefContext) {}

// ExitPortsDef is called when production portsDef is exited.
func (s *BasenevaListener) ExitPortsDef(ctx *PortsDefContext) {}

// EnterPortDef is called when production portDef is entered.
func (s *BasenevaListener) EnterPortDef(ctx *PortDefContext) {}

// ExitPortDef is called when production portDef is exited.
func (s *BasenevaListener) ExitPortDef(ctx *PortDefContext) {}

// EnterConstStmt is called when production constStmt is entered.
func (s *BasenevaListener) EnterConstStmt(ctx *ConstStmtContext) {}

// ExitConstStmt is called when production constStmt is exited.
func (s *BasenevaListener) ExitConstStmt(ctx *ConstStmtContext) {}

// EnterConstDef is called when production constDef is entered.
func (s *BasenevaListener) EnterConstDef(ctx *ConstDefContext) {}

// ExitConstDef is called when production constDef is exited.
func (s *BasenevaListener) ExitConstDef(ctx *ConstDefContext) {}

// EnterConstVal is called when production constVal is entered.
func (s *BasenevaListener) EnterConstVal(ctx *ConstValContext) {}

// ExitConstVal is called when production constVal is exited.
func (s *BasenevaListener) ExitConstVal(ctx *ConstValContext) {}

// EnterBool is called when production bool is entered.
func (s *BasenevaListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BasenevaListener) ExitBool(ctx *BoolContext) {}

// EnterNil is called when production nil is entered.
func (s *BasenevaListener) EnterNil(ctx *NilContext) {}

// ExitNil is called when production nil is exited.
func (s *BasenevaListener) ExitNil(ctx *NilContext) {}

// EnterArrLit is called when production arrLit is entered.
func (s *BasenevaListener) EnterArrLit(ctx *ArrLitContext) {}

// ExitArrLit is called when production arrLit is exited.
func (s *BasenevaListener) ExitArrLit(ctx *ArrLitContext) {}

// EnterListItems is called when production listItems is entered.
func (s *BasenevaListener) EnterListItems(ctx *ListItemsContext) {}

// ExitListItems is called when production listItems is exited.
func (s *BasenevaListener) ExitListItems(ctx *ListItemsContext) {}

// EnterStructLit is called when production structLit is entered.
func (s *BasenevaListener) EnterStructLit(ctx *StructLitContext) {}

// ExitStructLit is called when production structLit is exited.
func (s *BasenevaListener) ExitStructLit(ctx *StructLitContext) {}

// EnterStructValueFields is called when production structValueFields is entered.
func (s *BasenevaListener) EnterStructValueFields(ctx *StructValueFieldsContext) {}

// ExitStructValueFields is called when production structValueFields is exited.
func (s *BasenevaListener) ExitStructValueFields(ctx *StructValueFieldsContext) {}

// EnterStructValueField is called when production structValueField is entered.
func (s *BasenevaListener) EnterStructValueField(ctx *StructValueFieldContext) {}

// ExitStructValueField is called when production structValueField is exited.
func (s *BasenevaListener) ExitStructValueField(ctx *StructValueFieldContext) {}

// EnterCompStmt is called when production compStmt is entered.
func (s *BasenevaListener) EnterCompStmt(ctx *CompStmtContext) {}

// ExitCompStmt is called when production compStmt is exited.
func (s *BasenevaListener) ExitCompStmt(ctx *CompStmtContext) {}

// EnterCompDef is called when production compDef is entered.
func (s *BasenevaListener) EnterCompDef(ctx *CompDefContext) {}

// ExitCompDef is called when production compDef is exited.
func (s *BasenevaListener) ExitCompDef(ctx *CompDefContext) {}

// EnterCompBody is called when production compBody is entered.
func (s *BasenevaListener) EnterCompBody(ctx *CompBodyContext) {}

// ExitCompBody is called when production compBody is exited.
func (s *BasenevaListener) ExitCompBody(ctx *CompBodyContext) {}

// EnterCompNodesDef is called when production compNodesDef is entered.
func (s *BasenevaListener) EnterCompNodesDef(ctx *CompNodesDefContext) {}

// ExitCompNodesDef is called when production compNodesDef is exited.
func (s *BasenevaListener) ExitCompNodesDef(ctx *CompNodesDefContext) {}

// EnterCompNodesDefBody is called when production compNodesDefBody is entered.
func (s *BasenevaListener) EnterCompNodesDefBody(ctx *CompNodesDefBodyContext) {}

// ExitCompNodesDefBody is called when production compNodesDefBody is exited.
func (s *BasenevaListener) ExitCompNodesDefBody(ctx *CompNodesDefBodyContext) {}

// EnterCompNodeDef is called when production compNodeDef is entered.
func (s *BasenevaListener) EnterCompNodeDef(ctx *CompNodeDefContext) {}

// ExitCompNodeDef is called when production compNodeDef is exited.
func (s *BasenevaListener) ExitCompNodeDef(ctx *CompNodeDefContext) {}

// EnterNodeInst is called when production nodeInst is entered.
func (s *BasenevaListener) EnterNodeInst(ctx *NodeInstContext) {}

// ExitNodeInst is called when production nodeInst is exited.
func (s *BasenevaListener) ExitNodeInst(ctx *NodeInstContext) {}

// EnterNodeDIArgs is called when production nodeDIArgs is entered.
func (s *BasenevaListener) EnterNodeDIArgs(ctx *NodeDIArgsContext) {}

// ExitNodeDIArgs is called when production nodeDIArgs is exited.
func (s *BasenevaListener) ExitNodeDIArgs(ctx *NodeDIArgsContext) {}

// EnterCompNetDef is called when production compNetDef is entered.
func (s *BasenevaListener) EnterCompNetDef(ctx *CompNetDefContext) {}

// ExitCompNetDef is called when production compNetDef is exited.
func (s *BasenevaListener) ExitCompNetDef(ctx *CompNetDefContext) {}

// EnterConnDefList is called when production connDefList is entered.
func (s *BasenevaListener) EnterConnDefList(ctx *ConnDefListContext) {}

// ExitConnDefList is called when production connDefList is exited.
func (s *BasenevaListener) ExitConnDefList(ctx *ConnDefListContext) {}

// EnterConnDef is called when production connDef is entered.
func (s *BasenevaListener) EnterConnDef(ctx *ConnDefContext) {}

// ExitConnDef is called when production connDef is exited.
func (s *BasenevaListener) ExitConnDef(ctx *ConnDefContext) {}

// EnterSenderSide is called when production senderSide is entered.
func (s *BasenevaListener) EnterSenderSide(ctx *SenderSideContext) {}

// ExitSenderSide is called when production senderSide is exited.
func (s *BasenevaListener) ExitSenderSide(ctx *SenderSideContext) {}

// EnterPortAddr is called when production portAddr is entered.
func (s *BasenevaListener) EnterPortAddr(ctx *PortAddrContext) {}

// ExitPortAddr is called when production portAddr is exited.
func (s *BasenevaListener) ExitPortAddr(ctx *PortAddrContext) {}

// EnterIoNodePortAddr is called when production ioNodePortAddr is entered.
func (s *BasenevaListener) EnterIoNodePortAddr(ctx *IoNodePortAddrContext) {}

// ExitIoNodePortAddr is called when production ioNodePortAddr is exited.
func (s *BasenevaListener) ExitIoNodePortAddr(ctx *IoNodePortAddrContext) {}

// EnterPortDirection is called when production portDirection is entered.
func (s *BasenevaListener) EnterPortDirection(ctx *PortDirectionContext) {}

// ExitPortDirection is called when production portDirection is exited.
func (s *BasenevaListener) ExitPortDirection(ctx *PortDirectionContext) {}

// EnterNormalNodePortAddr is called when production normalNodePortAddr is entered.
func (s *BasenevaListener) EnterNormalNodePortAddr(ctx *NormalNodePortAddrContext) {}

// ExitNormalNodePortAddr is called when production normalNodePortAddr is exited.
func (s *BasenevaListener) ExitNormalNodePortAddr(ctx *NormalNodePortAddrContext) {}

// EnterConnReceiverSide is called when production connReceiverSide is entered.
func (s *BasenevaListener) EnterConnReceiverSide(ctx *ConnReceiverSideContext) {}

// ExitConnReceiverSide is called when production connReceiverSide is exited.
func (s *BasenevaListener) ExitConnReceiverSide(ctx *ConnReceiverSideContext) {}

// EnterConnReceivers is called when production connReceivers is entered.
func (s *BasenevaListener) EnterConnReceivers(ctx *ConnReceiversContext) {}

// ExitConnReceivers is called when production connReceivers is exited.
func (s *BasenevaListener) ExitConnReceivers(ctx *ConnReceiversContext) {}
