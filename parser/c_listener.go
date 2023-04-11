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

	// EnterDeclarationStmt is called when entering the DeclarationStmt production.
	EnterDeclarationStmt(c *DeclarationStmtContext)

	// EnterDefinitionStmt is called when entering the DefinitionStmt production.
	EnterDefinitionStmt(c *DefinitionStmtContext)

	// EnterAssignmentStmt is called when entering the AssignmentStmt production.
	EnterAssignmentStmt(c *AssignmentStmtContext)

	// EnterCallStmt is called when entering the CallStmt production.
	EnterCallStmt(c *CallStmtContext)

	// EnterIfStmt is called when entering the IfStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterBlockStmt is called when entering the BlockStmt production.
	EnterBlockStmt(c *BlockStmtContext)

	// EnterAndExpr is called when entering the AndExpr production.
	EnterAndExpr(c *AndExprContext)

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

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterEqNeqExpr is called when entering the EqNeqExpr production.
	EnterEqNeqExpr(c *EqNeqExprContext)

	// EnterAddSubExpr is called when entering the AddSubExpr production.
	EnterAddSubExpr(c *AddSubExprContext)

	// EnterOrExpr is called when entering the OrExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterMinusExpr is called when entering the MinusExpr production.
	EnterMinusExpr(c *MinusExprContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

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

	// ExitDeclarationStmt is called when exiting the DeclarationStmt production.
	ExitDeclarationStmt(c *DeclarationStmtContext)

	// ExitDefinitionStmt is called when exiting the DefinitionStmt production.
	ExitDefinitionStmt(c *DefinitionStmtContext)

	// ExitAssignmentStmt is called when exiting the AssignmentStmt production.
	ExitAssignmentStmt(c *AssignmentStmtContext)

	// ExitCallStmt is called when exiting the CallStmt production.
	ExitCallStmt(c *CallStmtContext)

	// ExitIfStmt is called when exiting the IfStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitBlockStmt is called when exiting the BlockStmt production.
	ExitBlockStmt(c *BlockStmtContext)

	// ExitAndExpr is called when exiting the AndExpr production.
	ExitAndExpr(c *AndExprContext)

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

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitEqNeqExpr is called when exiting the EqNeqExpr production.
	ExitEqNeqExpr(c *EqNeqExprContext)

	// ExitAddSubExpr is called when exiting the AddSubExpr production.
	ExitAddSubExpr(c *AddSubExprContext)

	// ExitOrExpr is called when exiting the OrExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitMinusExpr is called when exiting the MinusExpr production.
	ExitMinusExpr(c *MinusExprContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

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
