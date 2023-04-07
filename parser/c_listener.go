// Code generated from ../C.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // C
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// CListener is a complete listener for a parse tree produced by CParser.
type CListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterInclude is called when entering the include production.
	EnterInclude(c *IncludeContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterStruct is called when entering the struct production.
	EnterStruct(c *StructContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterDeclaration is called when entering the Declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterDefinition is called when entering the Definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterAssignment is called when entering the Assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterCall is called when entering the Call production.
	EnterCall(c *CallContext)

	// EnterIf is called when entering the If production.
	EnterIf(c *IfContext)

	// EnterReturn is called when entering the Return production.
	EnterReturn(c *ReturnContext)

	// EnterBlock is called when entering the Block production.
	EnterBlock(c *BlockContext)

	// EnterMulDivExpr is called when entering the MulDivExpr production.
	EnterMulDivExpr(c *MulDivExprContext)

	// EnterMemberExpr is called when entering the MemberExpr production.
	EnterMemberExpr(c *MemberExprContext)

	// EnterConstantExpr is called when entering the ConstantExpr production.
	EnterConstantExpr(c *ConstantExprContext)

	// EnterVariableExpr is called when entering the VariableExpr production.
	EnterVariableExpr(c *VariableExprContext)

	// EnterCallExpr is called when entering the CallExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterEqNeqExpr is called when entering the EqNeqExpr production.
	EnterEqNeqExpr(c *EqNeqExprContext)

	// EnterAddSubExpr is called when entering the AddSubExpr production.
	EnterAddSubExpr(c *AddSubExprContext)

	// EnterMinusExpr is called when entering the MinusExpr production.
	EnterMinusExpr(c *MinusExprContext)

	// EnterElseif is called when entering the elseif production.
	EnterElseif(c *ElseifContext)

	// EnterElse is called when entering the else production.
	EnterElse(c *ElseContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterCall_expr is called when entering the call_expr production.
	EnterCall_expr(c *Call_exprContext)

	// EnterMuldiv is called when entering the muldiv production.
	EnterMuldiv(c *MuldivContext)

	// EnterAddsub is called when entering the addsub production.
	EnterAddsub(c *AddsubContext)

	// EnterEqneq is called when entering the eqneq production.
	EnterEqneq(c *EqneqContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitInclude is called when exiting the include production.
	ExitInclude(c *IncludeContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitStruct is called when exiting the struct production.
	ExitStruct(c *StructContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitDeclaration is called when exiting the Declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitDefinition is called when exiting the Definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitAssignment is called when exiting the Assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitCall is called when exiting the Call production.
	ExitCall(c *CallContext)

	// ExitIf is called when exiting the If production.
	ExitIf(c *IfContext)

	// ExitReturn is called when exiting the Return production.
	ExitReturn(c *ReturnContext)

	// ExitBlock is called when exiting the Block production.
	ExitBlock(c *BlockContext)

	// ExitMulDivExpr is called when exiting the MulDivExpr production.
	ExitMulDivExpr(c *MulDivExprContext)

	// ExitMemberExpr is called when exiting the MemberExpr production.
	ExitMemberExpr(c *MemberExprContext)

	// ExitConstantExpr is called when exiting the ConstantExpr production.
	ExitConstantExpr(c *ConstantExprContext)

	// ExitVariableExpr is called when exiting the VariableExpr production.
	ExitVariableExpr(c *VariableExprContext)

	// ExitCallExpr is called when exiting the CallExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitEqNeqExpr is called when exiting the EqNeqExpr production.
	ExitEqNeqExpr(c *EqNeqExprContext)

	// ExitAddSubExpr is called when exiting the AddSubExpr production.
	ExitAddSubExpr(c *AddSubExprContext)

	// ExitMinusExpr is called when exiting the MinusExpr production.
	ExitMinusExpr(c *MinusExprContext)

	// ExitElseif is called when exiting the elseif production.
	ExitElseif(c *ElseifContext)

	// ExitElse is called when exiting the else production.
	ExitElse(c *ElseContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitCall_expr is called when exiting the call_expr production.
	ExitCall_expr(c *Call_exprContext)

	// ExitMuldiv is called when exiting the muldiv production.
	ExitMuldiv(c *MuldivContext)

	// ExitAddsub is called when exiting the addsub production.
	ExitAddsub(c *AddsubContext)

	// ExitEqneq is called when exiting the eqneq production.
	ExitEqneq(c *EqneqContext)
}
