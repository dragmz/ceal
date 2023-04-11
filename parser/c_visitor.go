// Code generated from ../C.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // C
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by CParser.
type CVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by CParser#include.
	VisitInclude(ctx *IncludeContext) interface{}

	// Visit a parse tree produced by CParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by CParser#struct.
	VisitStruct(ctx *StructContext) interface{}

	// Visit a parse tree produced by CParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by CParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by CParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by CParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by CParser#DeclarationStmt.
	VisitDeclarationStmt(ctx *DeclarationStmtContext) interface{}

	// Visit a parse tree produced by CParser#DefinitionStmt.
	VisitDefinitionStmt(ctx *DefinitionStmtContext) interface{}

	// Visit a parse tree produced by CParser#AssignmentStmt.
	VisitAssignmentStmt(ctx *AssignmentStmtContext) interface{}

	// Visit a parse tree produced by CParser#CallStmt.
	VisitCallStmt(ctx *CallStmtContext) interface{}

	// Visit a parse tree produced by CParser#IfStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by CParser#ReturnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by CParser#BlockStmt.
	VisitBlockStmt(ctx *BlockStmtContext) interface{}

	// Visit a parse tree produced by CParser#AndExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by CParser#MulDivExpr.
	VisitMulDivExpr(ctx *MulDivExprContext) interface{}

	// Visit a parse tree produced by CParser#MemberExpr.
	VisitMemberExpr(ctx *MemberExprContext) interface{}

	// Visit a parse tree produced by CParser#ConstantExpr.
	VisitConstantExpr(ctx *ConstantExprContext) interface{}

	// Visit a parse tree produced by CParser#VariableExpr.
	VisitVariableExpr(ctx *VariableExprContext) interface{}

	// Visit a parse tree produced by CParser#CallExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by CParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by CParser#EqNeqExpr.
	VisitEqNeqExpr(ctx *EqNeqExprContext) interface{}

	// Visit a parse tree produced by CParser#AddSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by CParser#OrExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by CParser#MinusExpr.
	VisitMinusExpr(ctx *MinusExprContext) interface{}

	// Visit a parse tree produced by CParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by CParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by CParser#elseif.
	VisitElseif(ctx *ElseifContext) interface{}

	// Visit a parse tree produced by CParser#else.
	VisitElse(ctx *ElseContext) interface{}

	// Visit a parse tree produced by CParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by CParser#call_expr.
	VisitCall_expr(ctx *Call_exprContext) interface{}

	// Visit a parse tree produced by CParser#muldiv.
	VisitMuldiv(ctx *MuldivContext) interface{}

	// Visit a parse tree produced by CParser#addsub.
	VisitAddsub(ctx *AddsubContext) interface{}

	// Visit a parse tree produced by CParser#eqneq.
	VisitEqneq(ctx *EqneqContext) interface{}
}
