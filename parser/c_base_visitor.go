// Code generated from C.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // C
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseCVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitInclude(ctx *IncludeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitStruct(ctx *StructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitDeclarationStmt(ctx *DeclarationStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitDefinitionStmt(ctx *DefinitionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitPreIncDecStmt(ctx *PreIncDecStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitPostIncDecStmt(ctx *PostIncDecStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAssignSumDiffStmt(ctx *AssignSumDiffStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAsmStmt(ctx *AsmStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitCallStmt(ctx *CallStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitBlockStmt(ctx *BlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitGotoStmt(ctx *GotoStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitLabelStmt(ctx *LabelStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitDoWhileStmt(ctx *DoWhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitCommentStmt(ctx *CommentStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitSubscriptExpr(ctx *SubscriptExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitConstantExpr(ctx *ConstantExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitBitAndExpr(ctx *BitAndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitDotExpr(ctx *DotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitEqNeqExpr(ctx *EqNeqExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitBitOrExpr(ctx *BitOrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitConditionalExpr(ctx *ConditionalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAssignExpr(ctx *AssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitGroupExpr(ctx *GroupExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitMulDivExpr(ctx *MulDivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAssignSumDiffExpr(ctx *AssignSumDiffExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitPostIncDecExpr(ctx *PostIncDecExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitPreIncDecExpr(ctx *PreIncDecExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitBitXorExpr(ctx *BitXorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitMinusExpr(ctx *MinusExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitValue_access_expr(ctx *Value_access_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitPre_incdec_expr(ctx *Pre_incdec_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitPost_incdec_expr(ctx *Post_incdec_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitDot_expr(ctx *Dot_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitSubscript_expr(ctx *Subscript_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAssign_expr(ctx *Assign_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitConst(ctx *ConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAsd_expr(ctx *Asd_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAsd(ctx *AsdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitCase(ctx *CaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitDefault(ctx *DefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitForInit(ctx *ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitForCondition(ctx *ForConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitForIter(ctx *ForIterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitGlobal(ctx *GlobalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitElseif(ctx *ElseifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitElse(ctx *ElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitArgs(ctx *ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitCall_expr(ctx *Call_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitMuldiv(ctx *MuldivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitAddsub(ctx *AddsubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitEqneq(ctx *EqneqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCVisitor) VisitIncdec(ctx *IncdecContext) interface{} {
	return v.VisitChildren(ctx)
}
