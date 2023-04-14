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

	// Visit a parse tree produced by CParser#AssignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by CParser#AssignSumDiffStmt.
	VisitAssignSumDiffStmt(ctx *AssignSumDiffStmtContext) interface{}

	// Visit a parse tree produced by CParser#CallStmt.
	VisitCallStmt(ctx *CallStmtContext) interface{}

	// Visit a parse tree produced by CParser#IfStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by CParser#ReturnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by CParser#BlockStmt.
	VisitBlockStmt(ctx *BlockStmtContext) interface{}

	// Visit a parse tree produced by CParser#GotoStmt.
	VisitGotoStmt(ctx *GotoStmtContext) interface{}

	// Visit a parse tree produced by CParser#LabelStmt.
	VisitLabelStmt(ctx *LabelStmtContext) interface{}

	// Visit a parse tree produced by CParser#ForStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by CParser#WhileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by CParser#DoWhileStmt.
	VisitDoWhileStmt(ctx *DoWhileStmtContext) interface{}

	// Visit a parse tree produced by CParser#SwitchStmt.
	VisitSwitchStmt(ctx *SwitchStmtContext) interface{}

	// Visit a parse tree produced by CParser#BreakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by CParser#ContinueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by CParser#AndExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by CParser#ConstantExpr.
	VisitConstantExpr(ctx *ConstantExprContext) interface{}

	// Visit a parse tree produced by CParser#BitAndExpr.
	VisitBitAndExpr(ctx *BitAndExprContext) interface{}

	// Visit a parse tree produced by CParser#EqNeqExpr.
	VisitEqNeqExpr(ctx *EqNeqExprContext) interface{}

	// Visit a parse tree produced by CParser#BitOrExpr.
	VisitBitOrExpr(ctx *BitOrExprContext) interface{}

	// Visit a parse tree produced by CParser#OrExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by CParser#AssignExpr.
	VisitAssignExpr(ctx *AssignExprContext) interface{}

	// Visit a parse tree produced by CParser#GroupExpr.
	VisitGroupExpr(ctx *GroupExprContext) interface{}

	// Visit a parse tree produced by CParser#MulDivExpr.
	VisitMulDivExpr(ctx *MulDivExprContext) interface{}

	// Visit a parse tree produced by CParser#AssignSumDiffExpr.
	VisitAssignSumDiffExpr(ctx *AssignSumDiffExprContext) interface{}

	// Visit a parse tree produced by CParser#MemberExpr.
	VisitMemberExpr(ctx *MemberExprContext) interface{}

	// Visit a parse tree produced by CParser#PostIncDecExpr.
	VisitPostIncDecExpr(ctx *PostIncDecExprContext) interface{}

	// Visit a parse tree produced by CParser#PreIncDecExpr.
	VisitPreIncDecExpr(ctx *PreIncDecExprContext) interface{}

	// Visit a parse tree produced by CParser#VariableExpr.
	VisitVariableExpr(ctx *VariableExprContext) interface{}

	// Visit a parse tree produced by CParser#CallExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by CParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by CParser#AddSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by CParser#BitXorExpr.
	VisitBitXorExpr(ctx *BitXorExprContext) interface{}

	// Visit a parse tree produced by CParser#MinusExpr.
	VisitMinusExpr(ctx *MinusExprContext) interface{}

	// Visit a parse tree produced by CParser#assign_expr.
	VisitAssign_expr(ctx *Assign_exprContext) interface{}

	// Visit a parse tree produced by CParser#alt.
	VisitAlt(ctx *AltContext) interface{}

	// Visit a parse tree produced by CParser#const.
	VisitConst(ctx *ConstContext) interface{}

	// Visit a parse tree produced by CParser#asdexpr.
	VisitAsdexpr(ctx *AsdexprContext) interface{}

	// Visit a parse tree produced by CParser#asd.
	VisitAsd(ctx *AsdContext) interface{}

	// Visit a parse tree produced by CParser#case.
	VisitCase(ctx *CaseContext) interface{}

	// Visit a parse tree produced by CParser#default.
	VisitDefault(ctx *DefaultContext) interface{}

	// Visit a parse tree produced by CParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by CParser#forCondition.
	VisitForCondition(ctx *ForConditionContext) interface{}

	// Visit a parse tree produced by CParser#forIter.
	VisitForIter(ctx *ForIterContext) interface{}

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

	// Visit a parse tree produced by CParser#incdec.
	VisitIncdec(ctx *IncdecContext) interface{}
}
