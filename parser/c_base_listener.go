// Code generated from ../C.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // C
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseCListener is a complete listener for a parse tree produced by CParser.
type BaseCListener struct{}

var _ CListener = &BaseCListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseCListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseCListener) ExitProgram(ctx *ProgramContext) {}

// EnterInclude is called when production include is entered.
func (s *BaseCListener) EnterInclude(ctx *IncludeContext) {}

// ExitInclude is called when production include is exited.
func (s *BaseCListener) ExitInclude(ctx *IncludeContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseCListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseCListener) ExitFunction(ctx *FunctionContext) {}

// EnterStruct is called when production struct is entered.
func (s *BaseCListener) EnterStruct(ctx *StructContext) {}

// ExitStruct is called when production struct is exited.
func (s *BaseCListener) ExitStruct(ctx *StructContext) {}

// EnterField is called when production field is entered.
func (s *BaseCListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseCListener) ExitField(ctx *FieldContext) {}

// EnterType is called when production type is entered.
func (s *BaseCListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseCListener) ExitType(ctx *TypeContext) {}

// EnterParams is called when production params is entered.
func (s *BaseCListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseCListener) ExitParams(ctx *ParamsContext) {}

// EnterParam is called when production param is entered.
func (s *BaseCListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseCListener) ExitParam(ctx *ParamContext) {}

// EnterDeclaration is called when production Declaration is entered.
func (s *BaseCListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production Declaration is exited.
func (s *BaseCListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterDefinition is called when production Definition is entered.
func (s *BaseCListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production Definition is exited.
func (s *BaseCListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterAssignment is called when production Assignment is entered.
func (s *BaseCListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production Assignment is exited.
func (s *BaseCListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterCall is called when production Call is entered.
func (s *BaseCListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production Call is exited.
func (s *BaseCListener) ExitCall(ctx *CallContext) {}

// EnterIf is called when production If is entered.
func (s *BaseCListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production If is exited.
func (s *BaseCListener) ExitIf(ctx *IfContext) {}

// EnterReturn is called when production Return is entered.
func (s *BaseCListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production Return is exited.
func (s *BaseCListener) ExitReturn(ctx *ReturnContext) {}

// EnterBlock is called when production Block is entered.
func (s *BaseCListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production Block is exited.
func (s *BaseCListener) ExitBlock(ctx *BlockContext) {}

// EnterMulDivExpr is called when production MulDivExpr is entered.
func (s *BaseCListener) EnterMulDivExpr(ctx *MulDivExprContext) {}

// ExitMulDivExpr is called when production MulDivExpr is exited.
func (s *BaseCListener) ExitMulDivExpr(ctx *MulDivExprContext) {}

// EnterMemberExpr is called when production MemberExpr is entered.
func (s *BaseCListener) EnterMemberExpr(ctx *MemberExprContext) {}

// ExitMemberExpr is called when production MemberExpr is exited.
func (s *BaseCListener) ExitMemberExpr(ctx *MemberExprContext) {}

// EnterConstantExpr is called when production ConstantExpr is entered.
func (s *BaseCListener) EnterConstantExpr(ctx *ConstantExprContext) {}

// ExitConstantExpr is called when production ConstantExpr is exited.
func (s *BaseCListener) ExitConstantExpr(ctx *ConstantExprContext) {}

// EnterVariableExpr is called when production VariableExpr is entered.
func (s *BaseCListener) EnterVariableExpr(ctx *VariableExprContext) {}

// ExitVariableExpr is called when production VariableExpr is exited.
func (s *BaseCListener) ExitVariableExpr(ctx *VariableExprContext) {}

// EnterCallExpr is called when production CallExpr is entered.
func (s *BaseCListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production CallExpr is exited.
func (s *BaseCListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterEqNeqExpr is called when production EqNeqExpr is entered.
func (s *BaseCListener) EnterEqNeqExpr(ctx *EqNeqExprContext) {}

// ExitEqNeqExpr is called when production EqNeqExpr is exited.
func (s *BaseCListener) ExitEqNeqExpr(ctx *EqNeqExprContext) {}

// EnterAddSubExpr is called when production AddSubExpr is entered.
func (s *BaseCListener) EnterAddSubExpr(ctx *AddSubExprContext) {}

// ExitAddSubExpr is called when production AddSubExpr is exited.
func (s *BaseCListener) ExitAddSubExpr(ctx *AddSubExprContext) {}

// EnterMinusExpr is called when production MinusExpr is entered.
func (s *BaseCListener) EnterMinusExpr(ctx *MinusExprContext) {}

// ExitMinusExpr is called when production MinusExpr is exited.
func (s *BaseCListener) ExitMinusExpr(ctx *MinusExprContext) {}

// EnterElseif is called when production elseif is entered.
func (s *BaseCListener) EnterElseif(ctx *ElseifContext) {}

// ExitElseif is called when production elseif is exited.
func (s *BaseCListener) ExitElseif(ctx *ElseifContext) {}

// EnterElse is called when production else is entered.
func (s *BaseCListener) EnterElse(ctx *ElseContext) {}

// ExitElse is called when production else is exited.
func (s *BaseCListener) ExitElse(ctx *ElseContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseCListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseCListener) ExitArgs(ctx *ArgsContext) {}

// EnterCall_expr is called when production call_expr is entered.
func (s *BaseCListener) EnterCall_expr(ctx *Call_exprContext) {}

// ExitCall_expr is called when production call_expr is exited.
func (s *BaseCListener) ExitCall_expr(ctx *Call_exprContext) {}

// EnterMuldiv is called when production muldiv is entered.
func (s *BaseCListener) EnterMuldiv(ctx *MuldivContext) {}

// ExitMuldiv is called when production muldiv is exited.
func (s *BaseCListener) ExitMuldiv(ctx *MuldivContext) {}

// EnterAddsub is called when production addsub is entered.
func (s *BaseCListener) EnterAddsub(ctx *AddsubContext) {}

// ExitAddsub is called when production addsub is exited.
func (s *BaseCListener) ExitAddsub(ctx *AddsubContext) {}

// EnterEqneq is called when production eqneq is entered.
func (s *BaseCListener) EnterEqneq(ctx *EqneqContext) {}

// ExitEqneq is called when production eqneq is exited.
func (s *BaseCListener) ExitEqneq(ctx *EqneqContext) {}
