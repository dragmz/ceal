// Code generated from C.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // C
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// CListener is a complete listener for a parse tree produced by CParser.
type CListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclarationStmt is called when entering the DeclarationStmt production.
	EnterDeclarationStmt(c *DeclarationStmtContext)

	// EnterDefinitionStmt is called when entering the DefinitionStmt production.
	EnterDefinitionStmt(c *DefinitionStmtContext)

	// EnterPreIncDecStmt is called when entering the PreIncDecStmt production.
	EnterPreIncDecStmt(c *PreIncDecStmtContext)

	// EnterPostIncDecStmt is called when entering the PostIncDecStmt production.
	EnterPostIncDecStmt(c *PostIncDecStmtContext)

	// EnterAssignStmt is called when entering the AssignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterAssignSumDiffStmt is called when entering the AssignSumDiffStmt production.
	EnterAssignSumDiffStmt(c *AssignSumDiffStmtContext)

	// EnterAsmStmt is called when entering the AsmStmt production.
	EnterAsmStmt(c *AsmStmtContext)

	// EnterCallStmt is called when entering the CallStmt production.
	EnterCallStmt(c *CallStmtContext)

	// EnterIfStmt is called when entering the IfStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterBlockStmt is called when entering the BlockStmt production.
	EnterBlockStmt(c *BlockStmtContext)

	// EnterGotoStmt is called when entering the GotoStmt production.
	EnterGotoStmt(c *GotoStmtContext)

	// EnterLabelStmt is called when entering the LabelStmt production.
	EnterLabelStmt(c *LabelStmtContext)

	// EnterForStmt is called when entering the ForStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterWhileStmt is called when entering the WhileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterDoWhileStmt is called when entering the DoWhileStmt production.
	EnterDoWhileStmt(c *DoWhileStmtContext)

	// EnterSwitchStmt is called when entering the SwitchStmt production.
	EnterSwitchStmt(c *SwitchStmtContext)

	// EnterBreakStmt is called when entering the BreakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the ContinueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterCommentStmt is called when entering the CommentStmt production.
	EnterCommentStmt(c *CommentStmtContext)

	// EnterAndExpr is called when entering the AndExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterSubscriptExpr is called when entering the SubscriptExpr production.
	EnterSubscriptExpr(c *SubscriptExprContext)

	// EnterConstantExpr is called when entering the ConstantExpr production.
	EnterConstantExpr(c *ConstantExprContext)

	// EnterBitAndExpr is called when entering the BitAndExpr production.
	EnterBitAndExpr(c *BitAndExprContext)

	// EnterDotExpr is called when entering the DotExpr production.
	EnterDotExpr(c *DotExprContext)

	// EnterEqNeqExpr is called when entering the EqNeqExpr production.
	EnterEqNeqExpr(c *EqNeqExprContext)

	// EnterBitOrExpr is called when entering the BitOrExpr production.
	EnterBitOrExpr(c *BitOrExprContext)

	// EnterConditionalExpr is called when entering the ConditionalExpr production.
	EnterConditionalExpr(c *ConditionalExprContext)

	// EnterOrExpr is called when entering the OrExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterAssignExpr is called when entering the AssignExpr production.
	EnterAssignExpr(c *AssignExprContext)

	// EnterGroupExpr is called when entering the GroupExpr production.
	EnterGroupExpr(c *GroupExprContext)

	// EnterMulDivExpr is called when entering the MulDivExpr production.
	EnterMulDivExpr(c *MulDivExprContext)

	// EnterAssignSumDiffExpr is called when entering the AssignSumDiffExpr production.
	EnterAssignSumDiffExpr(c *AssignSumDiffExprContext)

	// EnterPostIncDecExpr is called when entering the PostIncDecExpr production.
	EnterPostIncDecExpr(c *PostIncDecExprContext)

	// EnterPreIncDecExpr is called when entering the PreIncDecExpr production.
	EnterPreIncDecExpr(c *PreIncDecExprContext)

	// EnterCallExpr is called when entering the CallExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterAddSubExpr is called when entering the AddSubExpr production.
	EnterAddSubExpr(c *AddSubExprContext)

	// EnterBitXorExpr is called when entering the BitXorExpr production.
	EnterBitXorExpr(c *BitXorExprContext)

	// EnterMinusExpr is called when entering the MinusExpr production.
	EnterMinusExpr(c *MinusExprContext)

	// EnterInclude is called when entering the include production.
	EnterInclude(c *IncludeContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterStruct is called when entering the struct production.
	EnterStruct(c *StructContext)

	// EnterGlobal is called when entering the global production.
	EnterGlobal(c *GlobalContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterValue_access_expr is called when entering the value_access_expr production.
	EnterValue_access_expr(c *Value_access_exprContext)

	// EnterPre_incdec_expr is called when entering the pre_incdec_expr production.
	EnterPre_incdec_expr(c *Pre_incdec_exprContext)

	// EnterPost_incdec_expr is called when entering the post_incdec_expr production.
	EnterPost_incdec_expr(c *Post_incdec_exprContext)

	// EnterDot_expr is called when entering the dot_expr production.
	EnterDot_expr(c *Dot_exprContext)

	// EnterSubscript_expr is called when entering the subscript_expr production.
	EnterSubscript_expr(c *Subscript_exprContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterAssign_expr is called when entering the assign_expr production.
	EnterAssign_expr(c *Assign_exprContext)

	// EnterConst is called when entering the const production.
	EnterConst(c *ConstContext)

	// EnterAsd_expr is called when entering the asd_expr production.
	EnterAsd_expr(c *Asd_exprContext)

	// EnterAsd is called when entering the asd production.
	EnterAsd(c *AsdContext)

	// EnterCase is called when entering the case production.
	EnterCase(c *CaseContext)

	// EnterDefault is called when entering the default production.
	EnterDefault(c *DefaultContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterForCondition is called when entering the forCondition production.
	EnterForCondition(c *ForConditionContext)

	// EnterForIter is called when entering the forIter production.
	EnterForIter(c *ForIterContext)

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

	// EnterIncdec is called when entering the incdec production.
	EnterIncdec(c *IncdecContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclarationStmt is called when exiting the DeclarationStmt production.
	ExitDeclarationStmt(c *DeclarationStmtContext)

	// ExitDefinitionStmt is called when exiting the DefinitionStmt production.
	ExitDefinitionStmt(c *DefinitionStmtContext)

	// ExitPreIncDecStmt is called when exiting the PreIncDecStmt production.
	ExitPreIncDecStmt(c *PreIncDecStmtContext)

	// ExitPostIncDecStmt is called when exiting the PostIncDecStmt production.
	ExitPostIncDecStmt(c *PostIncDecStmtContext)

	// ExitAssignStmt is called when exiting the AssignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitAssignSumDiffStmt is called when exiting the AssignSumDiffStmt production.
	ExitAssignSumDiffStmt(c *AssignSumDiffStmtContext)

	// ExitAsmStmt is called when exiting the AsmStmt production.
	ExitAsmStmt(c *AsmStmtContext)

	// ExitCallStmt is called when exiting the CallStmt production.
	ExitCallStmt(c *CallStmtContext)

	// ExitIfStmt is called when exiting the IfStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitBlockStmt is called when exiting the BlockStmt production.
	ExitBlockStmt(c *BlockStmtContext)

	// ExitGotoStmt is called when exiting the GotoStmt production.
	ExitGotoStmt(c *GotoStmtContext)

	// ExitLabelStmt is called when exiting the LabelStmt production.
	ExitLabelStmt(c *LabelStmtContext)

	// ExitForStmt is called when exiting the ForStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitWhileStmt is called when exiting the WhileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitDoWhileStmt is called when exiting the DoWhileStmt production.
	ExitDoWhileStmt(c *DoWhileStmtContext)

	// ExitSwitchStmt is called when exiting the SwitchStmt production.
	ExitSwitchStmt(c *SwitchStmtContext)

	// ExitBreakStmt is called when exiting the BreakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the ContinueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitCommentStmt is called when exiting the CommentStmt production.
	ExitCommentStmt(c *CommentStmtContext)

	// ExitAndExpr is called when exiting the AndExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitSubscriptExpr is called when exiting the SubscriptExpr production.
	ExitSubscriptExpr(c *SubscriptExprContext)

	// ExitConstantExpr is called when exiting the ConstantExpr production.
	ExitConstantExpr(c *ConstantExprContext)

	// ExitBitAndExpr is called when exiting the BitAndExpr production.
	ExitBitAndExpr(c *BitAndExprContext)

	// ExitDotExpr is called when exiting the DotExpr production.
	ExitDotExpr(c *DotExprContext)

	// ExitEqNeqExpr is called when exiting the EqNeqExpr production.
	ExitEqNeqExpr(c *EqNeqExprContext)

	// ExitBitOrExpr is called when exiting the BitOrExpr production.
	ExitBitOrExpr(c *BitOrExprContext)

	// ExitConditionalExpr is called when exiting the ConditionalExpr production.
	ExitConditionalExpr(c *ConditionalExprContext)

	// ExitOrExpr is called when exiting the OrExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitAssignExpr is called when exiting the AssignExpr production.
	ExitAssignExpr(c *AssignExprContext)

	// ExitGroupExpr is called when exiting the GroupExpr production.
	ExitGroupExpr(c *GroupExprContext)

	// ExitMulDivExpr is called when exiting the MulDivExpr production.
	ExitMulDivExpr(c *MulDivExprContext)

	// ExitAssignSumDiffExpr is called when exiting the AssignSumDiffExpr production.
	ExitAssignSumDiffExpr(c *AssignSumDiffExprContext)

	// ExitPostIncDecExpr is called when exiting the PostIncDecExpr production.
	ExitPostIncDecExpr(c *PostIncDecExprContext)

	// ExitPreIncDecExpr is called when exiting the PreIncDecExpr production.
	ExitPreIncDecExpr(c *PreIncDecExprContext)

	// ExitCallExpr is called when exiting the CallExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitAddSubExpr is called when exiting the AddSubExpr production.
	ExitAddSubExpr(c *AddSubExprContext)

	// ExitBitXorExpr is called when exiting the BitXorExpr production.
	ExitBitXorExpr(c *BitXorExprContext)

	// ExitMinusExpr is called when exiting the MinusExpr production.
	ExitMinusExpr(c *MinusExprContext)

	// ExitInclude is called when exiting the include production.
	ExitInclude(c *IncludeContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitStruct is called when exiting the struct production.
	ExitStruct(c *StructContext)

	// ExitGlobal is called when exiting the global production.
	ExitGlobal(c *GlobalContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitValue_access_expr is called when exiting the value_access_expr production.
	ExitValue_access_expr(c *Value_access_exprContext)

	// ExitPre_incdec_expr is called when exiting the pre_incdec_expr production.
	ExitPre_incdec_expr(c *Pre_incdec_exprContext)

	// ExitPost_incdec_expr is called when exiting the post_incdec_expr production.
	ExitPost_incdec_expr(c *Post_incdec_exprContext)

	// ExitDot_expr is called when exiting the dot_expr production.
	ExitDot_expr(c *Dot_exprContext)

	// ExitSubscript_expr is called when exiting the subscript_expr production.
	ExitSubscript_expr(c *Subscript_exprContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitAssign_expr is called when exiting the assign_expr production.
	ExitAssign_expr(c *Assign_exprContext)

	// ExitConst is called when exiting the const production.
	ExitConst(c *ConstContext)

	// ExitAsd_expr is called when exiting the asd_expr production.
	ExitAsd_expr(c *Asd_exprContext)

	// ExitAsd is called when exiting the asd production.
	ExitAsd(c *AsdContext)

	// ExitCase is called when exiting the case production.
	ExitCase(c *CaseContext)

	// ExitDefault is called when exiting the default production.
	ExitDefault(c *DefaultContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitForCondition is called when exiting the forCondition production.
	ExitForCondition(c *ForConditionContext)

	// ExitForIter is called when exiting the forIter production.
	ExitForIter(c *ForIterContext)

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

	// ExitIncdec is called when exiting the incdec production.
	ExitIncdec(c *IncdecContext)
}
