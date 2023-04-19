// Code generated from C.g4 by ANTLR 4.12.0. DO NOT EDIT.

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

// EnterDeclarationStmt is called when production DeclarationStmt is entered.
func (s *BaseCListener) EnterDeclarationStmt(ctx *DeclarationStmtContext) {}

// ExitDeclarationStmt is called when production DeclarationStmt is exited.
func (s *BaseCListener) ExitDeclarationStmt(ctx *DeclarationStmtContext) {}

// EnterDefinitionStmt is called when production DefinitionStmt is entered.
func (s *BaseCListener) EnterDefinitionStmt(ctx *DefinitionStmtContext) {}

// ExitDefinitionStmt is called when production DefinitionStmt is exited.
func (s *BaseCListener) ExitDefinitionStmt(ctx *DefinitionStmtContext) {}

// EnterAssignStmt is called when production AssignStmt is entered.
func (s *BaseCListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production AssignStmt is exited.
func (s *BaseCListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterAssignSumDiffStmt is called when production AssignSumDiffStmt is entered.
func (s *BaseCListener) EnterAssignSumDiffStmt(ctx *AssignSumDiffStmtContext) {}

// ExitAssignSumDiffStmt is called when production AssignSumDiffStmt is exited.
func (s *BaseCListener) ExitAssignSumDiffStmt(ctx *AssignSumDiffStmtContext) {}

// EnterAsmStmt is called when production AsmStmt is entered.
func (s *BaseCListener) EnterAsmStmt(ctx *AsmStmtContext) {}

// ExitAsmStmt is called when production AsmStmt is exited.
func (s *BaseCListener) ExitAsmStmt(ctx *AsmStmtContext) {}

// EnterCallStmt is called when production CallStmt is entered.
func (s *BaseCListener) EnterCallStmt(ctx *CallStmtContext) {}

// ExitCallStmt is called when production CallStmt is exited.
func (s *BaseCListener) ExitCallStmt(ctx *CallStmtContext) {}

// EnterIfStmt is called when production IfStmt is entered.
func (s *BaseCListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production IfStmt is exited.
func (s *BaseCListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterReturnStmt is called when production ReturnStmt is entered.
func (s *BaseCListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production ReturnStmt is exited.
func (s *BaseCListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterBlockStmt is called when production BlockStmt is entered.
func (s *BaseCListener) EnterBlockStmt(ctx *BlockStmtContext) {}

// ExitBlockStmt is called when production BlockStmt is exited.
func (s *BaseCListener) ExitBlockStmt(ctx *BlockStmtContext) {}

// EnterGotoStmt is called when production GotoStmt is entered.
func (s *BaseCListener) EnterGotoStmt(ctx *GotoStmtContext) {}

// ExitGotoStmt is called when production GotoStmt is exited.
func (s *BaseCListener) ExitGotoStmt(ctx *GotoStmtContext) {}

// EnterLabelStmt is called when production LabelStmt is entered.
func (s *BaseCListener) EnterLabelStmt(ctx *LabelStmtContext) {}

// ExitLabelStmt is called when production LabelStmt is exited.
func (s *BaseCListener) ExitLabelStmt(ctx *LabelStmtContext) {}

// EnterForStmt is called when production ForStmt is entered.
func (s *BaseCListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production ForStmt is exited.
func (s *BaseCListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterWhileStmt is called when production WhileStmt is entered.
func (s *BaseCListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production WhileStmt is exited.
func (s *BaseCListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterDoWhileStmt is called when production DoWhileStmt is entered.
func (s *BaseCListener) EnterDoWhileStmt(ctx *DoWhileStmtContext) {}

// ExitDoWhileStmt is called when production DoWhileStmt is exited.
func (s *BaseCListener) ExitDoWhileStmt(ctx *DoWhileStmtContext) {}

// EnterSwitchStmt is called when production SwitchStmt is entered.
func (s *BaseCListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production SwitchStmt is exited.
func (s *BaseCListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterBreakStmt is called when production BreakStmt is entered.
func (s *BaseCListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production BreakStmt is exited.
func (s *BaseCListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production ContinueStmt is entered.
func (s *BaseCListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production ContinueStmt is exited.
func (s *BaseCListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterCommentStmt is called when production CommentStmt is entered.
func (s *BaseCListener) EnterCommentStmt(ctx *CommentStmtContext) {}

// ExitCommentStmt is called when production CommentStmt is exited.
func (s *BaseCListener) ExitCommentStmt(ctx *CommentStmtContext) {}

// EnterAndExpr is called when production AndExpr is entered.
func (s *BaseCListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production AndExpr is exited.
func (s *BaseCListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterSubscriptExpr is called when production SubscriptExpr is entered.
func (s *BaseCListener) EnterSubscriptExpr(ctx *SubscriptExprContext) {}

// ExitSubscriptExpr is called when production SubscriptExpr is exited.
func (s *BaseCListener) ExitSubscriptExpr(ctx *SubscriptExprContext) {}

// EnterConstantExpr is called when production ConstantExpr is entered.
func (s *BaseCListener) EnterConstantExpr(ctx *ConstantExprContext) {}

// ExitConstantExpr is called when production ConstantExpr is exited.
func (s *BaseCListener) ExitConstantExpr(ctx *ConstantExprContext) {}

// EnterBitAndExpr is called when production BitAndExpr is entered.
func (s *BaseCListener) EnterBitAndExpr(ctx *BitAndExprContext) {}

// ExitBitAndExpr is called when production BitAndExpr is exited.
func (s *BaseCListener) ExitBitAndExpr(ctx *BitAndExprContext) {}

// EnterDotExpr is called when production DotExpr is entered.
func (s *BaseCListener) EnterDotExpr(ctx *DotExprContext) {}

// ExitDotExpr is called when production DotExpr is exited.
func (s *BaseCListener) ExitDotExpr(ctx *DotExprContext) {}

// EnterEqNeqExpr is called when production EqNeqExpr is entered.
func (s *BaseCListener) EnterEqNeqExpr(ctx *EqNeqExprContext) {}

// ExitEqNeqExpr is called when production EqNeqExpr is exited.
func (s *BaseCListener) ExitEqNeqExpr(ctx *EqNeqExprContext) {}

// EnterBitOrExpr is called when production BitOrExpr is entered.
func (s *BaseCListener) EnterBitOrExpr(ctx *BitOrExprContext) {}

// ExitBitOrExpr is called when production BitOrExpr is exited.
func (s *BaseCListener) ExitBitOrExpr(ctx *BitOrExprContext) {}

// EnterConditionalExpr is called when production ConditionalExpr is entered.
func (s *BaseCListener) EnterConditionalExpr(ctx *ConditionalExprContext) {}

// ExitConditionalExpr is called when production ConditionalExpr is exited.
func (s *BaseCListener) ExitConditionalExpr(ctx *ConditionalExprContext) {}

// EnterOrExpr is called when production OrExpr is entered.
func (s *BaseCListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production OrExpr is exited.
func (s *BaseCListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterAssignExpr is called when production AssignExpr is entered.
func (s *BaseCListener) EnterAssignExpr(ctx *AssignExprContext) {}

// ExitAssignExpr is called when production AssignExpr is exited.
func (s *BaseCListener) ExitAssignExpr(ctx *AssignExprContext) {}

// EnterGroupExpr is called when production GroupExpr is entered.
func (s *BaseCListener) EnterGroupExpr(ctx *GroupExprContext) {}

// ExitGroupExpr is called when production GroupExpr is exited.
func (s *BaseCListener) ExitGroupExpr(ctx *GroupExprContext) {}

// EnterMulDivExpr is called when production MulDivExpr is entered.
func (s *BaseCListener) EnterMulDivExpr(ctx *MulDivExprContext) {}

// ExitMulDivExpr is called when production MulDivExpr is exited.
func (s *BaseCListener) ExitMulDivExpr(ctx *MulDivExprContext) {}

// EnterAssignSumDiffExpr is called when production AssignSumDiffExpr is entered.
func (s *BaseCListener) EnterAssignSumDiffExpr(ctx *AssignSumDiffExprContext) {}

// ExitAssignSumDiffExpr is called when production AssignSumDiffExpr is exited.
func (s *BaseCListener) ExitAssignSumDiffExpr(ctx *AssignSumDiffExprContext) {}

// EnterPostIncDecExpr is called when production PostIncDecExpr is entered.
func (s *BaseCListener) EnterPostIncDecExpr(ctx *PostIncDecExprContext) {}

// ExitPostIncDecExpr is called when production PostIncDecExpr is exited.
func (s *BaseCListener) ExitPostIncDecExpr(ctx *PostIncDecExprContext) {}

// EnterPreIncDecExpr is called when production PreIncDecExpr is entered.
func (s *BaseCListener) EnterPreIncDecExpr(ctx *PreIncDecExprContext) {}

// ExitPreIncDecExpr is called when production PreIncDecExpr is exited.
func (s *BaseCListener) ExitPreIncDecExpr(ctx *PreIncDecExprContext) {}

// EnterCallExpr is called when production CallExpr is entered.
func (s *BaseCListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production CallExpr is exited.
func (s *BaseCListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseCListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseCListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterAddSubExpr is called when production AddSubExpr is entered.
func (s *BaseCListener) EnterAddSubExpr(ctx *AddSubExprContext) {}

// ExitAddSubExpr is called when production AddSubExpr is exited.
func (s *BaseCListener) ExitAddSubExpr(ctx *AddSubExprContext) {}

// EnterBitXorExpr is called when production BitXorExpr is entered.
func (s *BaseCListener) EnterBitXorExpr(ctx *BitXorExprContext) {}

// ExitBitXorExpr is called when production BitXorExpr is exited.
func (s *BaseCListener) ExitBitXorExpr(ctx *BitXorExprContext) {}

// EnterMinusExpr is called when production MinusExpr is entered.
func (s *BaseCListener) EnterMinusExpr(ctx *MinusExprContext) {}

// ExitMinusExpr is called when production MinusExpr is exited.
func (s *BaseCListener) ExitMinusExpr(ctx *MinusExprContext) {}

// EnterDot_expr is called when production dot_expr is entered.
func (s *BaseCListener) EnterDot_expr(ctx *Dot_exprContext) {}

// ExitDot_expr is called when production dot_expr is exited.
func (s *BaseCListener) ExitDot_expr(ctx *Dot_exprContext) {}

// EnterValue_expr is called when production value_expr is entered.
func (s *BaseCListener) EnterValue_expr(ctx *Value_exprContext) {}

// ExitValue_expr is called when production value_expr is exited.
func (s *BaseCListener) ExitValue_expr(ctx *Value_exprContext) {}

// EnterSubscript_expr is called when production subscript_expr is entered.
func (s *BaseCListener) EnterSubscript_expr(ctx *Subscript_exprContext) {}

// ExitSubscript_expr is called when production subscript_expr is exited.
func (s *BaseCListener) ExitSubscript_expr(ctx *Subscript_exprContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseCListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseCListener) ExitComment(ctx *CommentContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseCListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseCListener) ExitConstant(ctx *ConstantContext) {}

// EnterAssign_expr is called when production assign_expr is entered.
func (s *BaseCListener) EnterAssign_expr(ctx *Assign_exprContext) {}

// ExitAssign_expr is called when production assign_expr is exited.
func (s *BaseCListener) ExitAssign_expr(ctx *Assign_exprContext) {}

// EnterConst is called when production const is entered.
func (s *BaseCListener) EnterConst(ctx *ConstContext) {}

// ExitConst is called when production const is exited.
func (s *BaseCListener) ExitConst(ctx *ConstContext) {}

// EnterAsd_expr is called when production asd_expr is entered.
func (s *BaseCListener) EnterAsd_expr(ctx *Asd_exprContext) {}

// ExitAsd_expr is called when production asd_expr is exited.
func (s *BaseCListener) ExitAsd_expr(ctx *Asd_exprContext) {}

// EnterAsd is called when production asd is entered.
func (s *BaseCListener) EnterAsd(ctx *AsdContext) {}

// ExitAsd is called when production asd is exited.
func (s *BaseCListener) ExitAsd(ctx *AsdContext) {}

// EnterCase is called when production case is entered.
func (s *BaseCListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production case is exited.
func (s *BaseCListener) ExitCase(ctx *CaseContext) {}

// EnterDefault is called when production default is entered.
func (s *BaseCListener) EnterDefault(ctx *DefaultContext) {}

// ExitDefault is called when production default is exited.
func (s *BaseCListener) ExitDefault(ctx *DefaultContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseCListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseCListener) ExitForInit(ctx *ForInitContext) {}

// EnterForCondition is called when production forCondition is entered.
func (s *BaseCListener) EnterForCondition(ctx *ForConditionContext) {}

// ExitForCondition is called when production forCondition is exited.
func (s *BaseCListener) ExitForCondition(ctx *ForConditionContext) {}

// EnterForIter is called when production forIter is entered.
func (s *BaseCListener) EnterForIter(ctx *ForIterContext) {}

// ExitForIter is called when production forIter is exited.
func (s *BaseCListener) ExitForIter(ctx *ForIterContext) {}

// EnterGlobal is called when production global is entered.
func (s *BaseCListener) EnterGlobal(ctx *GlobalContext) {}

// ExitGlobal is called when production global is exited.
func (s *BaseCListener) ExitGlobal(ctx *GlobalContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseCListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseCListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseCListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseCListener) ExitDefinition(ctx *DefinitionContext) {}

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

// EnterIncdec is called when production incdec is entered.
func (s *BaseCListener) EnterIncdec(ctx *IncdecContext) {}

// ExitIncdec is called when production incdec is exited.
func (s *BaseCListener) ExitIncdec(ctx *IncdecContext) {}
