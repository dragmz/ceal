package compiler

import (
	"ceal/parser"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type AstVisitor struct {
	*parser.BaseCVisitor

	global *Scope
	scope  *Scope

	program *CealProgram

	index int
	slot  int

	loops *LoopScope

	comments []CealAst
}

func (v *AstVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *AstVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}

func (v *AstVisitor) visitAst(tree antlr.ParseTree) CealAst {
	res := v.Visit(tree)
	if res == nil {
		return nil
	}

	return res.(CealAst)
}

func (v *AstVisitor) VisitValueExpr(ctx *parser.ValueExprContext) interface{} {
	ids := []antlr.TerminalNode{}

	for _, e := range ctx.AllValue_expr() {
		ids = append(ids, e.ID())
	}

	var ast CealAst

	vr, f := v.mustResolve(ids)
	if f == nil {
		ast = &CealVariable{
			V: vr,
		}
	} else {
		fun := v.global.functions[f.fun]

		ast = &CealStructField{
			V:   vr,
			T:   v.scope.resolveType(vr.t),
			F:   f,
			Fun: fun,
		}
	}

	return ast
}

func (v *AstVisitor) VisitAssignExpr(ctx *parser.AssignExprContext) interface{} {
	return v.Visit(ctx.Assign_expr())
}

func (v *AstVisitor) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	res := v.Visit(ctx.Assign_expr())
	if e, ok := res.(CealExpr); ok {
		e.ToStmt()
	}
	return res
}

func (v *AstVisitor) VisitAssign_expr(ctx *parser.Assign_exprContext) interface{} {
	ids := ctx.AllID()

	vr, f := v.mustResolve(ids)
	if vr.readonly {
		panic(fmt.Sprintf("variable '%s' is read only", vr.name))
	}

	if vr.readonly {
		panic(fmt.Sprintf("variable '%s' is read only", vr.name))
	}

	ast := &CealAssign{
		V:     vr,
		T:     v.scope.resolveType(vr.t),
		F:     f,
		Value: v.visitAst(ctx.Expr()),
	}

	return ast
}

func (v *AstVisitor) VisitDeclarationStmt(ctx *parser.DeclarationStmtContext) interface{} {
	id := ctx.Declaration().Type_().ID().GetText()
	t := v.scope.resolveType(id)

	if t == nil {
		panic(fmt.Sprintf("type '%s' not found", id))
	}

	return nil
}

func (v *AstVisitor) VisitMinusExpr(ctx *parser.MinusExprContext) interface{} {
	return &CealNegate{
		Value: v.visitAst(ctx.Expr()),
	}
}

func (v *AstVisitor) VisitAddSubExpr(ctx *parser.AddSubExprContext) interface{} {
	return &CealBinop{
		Left:  v.visitAst(ctx.GetL()),
		Right: v.visitAst(ctx.GetR()),
		Op:    ctx.Addsub().GetText(),
	}
}

func (v *AstVisitor) VisitAssignSumDiffStmt(ctx *parser.AssignSumDiffStmtContext) interface{} {
	vr, f := v.mustResolve(ctx.Asdexpr().AllID())
	t := v.scope.resolveType(vr.t)

	return &CealAssignSumDiff{
		V:      vr,
		F:      f,
		T:      t,
		Value:  v.visitAst(ctx.Asdexpr().Expr()),
		Op:     ctx.Asdexpr().Asd().GetText(),
		IsStmt: true,
	}
}
func (v *AstVisitor) VisitAssignSumDiffExpr(ctx *parser.AssignSumDiffExprContext) interface{} {
	vr, f := v.mustResolve(ctx.Asdexpr().AllID())
	t := v.scope.resolveType(vr.t)

	return &CealAssignSumDiff{
		V:     vr,
		F:     f,
		T:     t,
		Value: v.visitAst(ctx.Asdexpr().Expr()),
		Op:    ctx.Asdexpr().Asd().GetText(),
	}
}

func (v *AstVisitor) VisitOrExpr(ctx *parser.OrExprContext) interface{} {
	ast := &CealOr{
		Index: v.index,
		Left:  v.visitAst(ctx.GetL()),
		Right: v.visitAst(ctx.GetR()),
	}

	v.index++

	return ast
}

func (v *AstVisitor) VisitAndExpr(ctx *parser.AndExprContext) interface{} {
	ast := &CealAnd{
		Index: v.index,
		Left:  v.visitAst(ctx.GetL()),
		Right: v.visitAst(ctx.GetR()),
	}

	v.index++

	return ast
}

func (v *AstVisitor) VisitMulDivExpr(ctx *parser.MulDivExprContext) interface{} {
	return &CealBinop{
		Left:  v.visitAst(ctx.GetL()),
		Right: v.visitAst(ctx.GetR()),
		Op:    ctx.Muldiv().GetText(),
	}
}

func (v *AstVisitor) VisitEqNeqExpr(ctx *parser.EqNeqExprContext) interface{} {
	return &CealBinop{
		Left:  v.visitAst(ctx.GetL()),
		Right: v.visitAst(ctx.GetR()),
		Op:    ctx.Eqneq().GetText(),
	}
}

func (v *AstVisitor) VisitBitAndExpr(ctx *parser.BitAndExprContext) interface{} {
	return &CealBinop{
		Left:  v.visitAst(ctx.GetL()),
		Right: v.visitAst(ctx.GetR()),
		Op:    "&",
	}
}

func (v *AstVisitor) VisitBitXorExpr(ctx *parser.BitXorExprContext) interface{} {
	return &CealBinop{
		Left:  v.visitAst(ctx.GetL()),
		Right: v.visitAst(ctx.GetR()),
		Op:    "^",
	}
}

func (v *AstVisitor) VisitBitOrExpr(ctx *parser.BitOrExprContext) interface{} {
	return &CealBinop{
		Left:  v.visitAst(ctx.GetL()),
		Right: v.visitAst(ctx.GetR()),
		Op:    "|",
	}
}

func (v *AstVisitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	return &CealUnaryOp{
		Statement: v.visitAst(ctx.Expr()),
		Op:        "!",
	}
}

func (v *AstVisitor) VisitConstantExpr(ctx *parser.ConstantExprContext) interface{} {
	var res CealAst

	c := ctx.Constant()

	if c.INT() != nil {
		res = &CealIntConstant{
			Value: c.INT().GetText(),
		}
	}

	if c.STRING() != nil {
		res = &CealStringConstant{
			Value: ceal_TrimSTRINGQuotes(c.STRING().GetText()),
		}
	}

	return res
}

func (v *AstVisitor) VisitGroupExpr(ctx *parser.GroupExprContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *AstVisitor) VisitCallExpr(ctx *parser.CallExprContext) interface{} {
	return v.Visit(ctx.Call_expr())
}

func (v *AstVisitor) VisitCallStmt(ctx *parser.CallStmtContext) interface{} {
	ast := v.Visit(ctx.Call_expr())

	if e, ok := ast.(CealExpr); ok {
		e.ToStmt()
	}

	return ast
}

func (v *AstVisitor) VisitCall_expr(ctx *parser.Call_exprContext) interface{} {
	ids := ctx.AllID()

	id := ids[0].GetText()

	imms := []CealAst{}

	if len(ids) > 1 {
		vr := v.mustResolveVariable(id)
		t := v.scope.resolveType(vr.t)

		if t.simple != nil {
			panic("cannot call simple type")
		}

		if t.complex.builtin == nil {
			panic("calling struct function is not supported yet")
		}

		id = t.complex.fields[ids[1].GetText()].fun

		// TODO: currently supports just a single level of fields

		raw := &CealRaw{
			Value: ids[1].GetText(),
		}

		imms = append(imms, raw)
	}

	fun := v.global.functions[id]

	ast := &CealCall{
		Fun: fun,
	}

	for _, arg := range ctx.Args().AllExpr() {
		stmt := v.visitAst(arg)
		ast.Args = append(ast.Args, stmt)
	}

	ast.Args = append(ast.Args, imms...)

	return ast
}

func (v *AstVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	ast := &CealReturn{
		Fun: v.scope.function,
	}

	if ctx.Expr() != nil {
		ast.Value = v.visitAst(ctx.Expr())
	}

	return ast
}

func (v *AstVisitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	alts := []*CealIfAlternative{}

	ast := &CealIf{
		Index: v.index,
	}

	v.index++

	alt := &CealIfAlternative{
		Condition: v.visitAst(ctx.Expr()),
	}

	for _, item := range ctx.AllStmt() {
		if stmt := v.visitAst(item); stmt != nil {
			alt.Statements = append(alt.Statements, stmt)
		}
	}

	alts = append(alts, alt)

	for _, elif := range ctx.AllElseif() {
		alt := &CealIfAlternative{
			Condition: v.visitAst(elif.Expr()),
		}

		for _, item := range elif.AllStmt() {
			if stmt := v.visitAst(item); stmt != nil {
				alt.Statements = append(alt.Statements, stmt)
			}
		}

		alts = append(alts, alt)
	}

	if ctx.Else_() != nil {
		alt := &CealIfAlternative{}

		for _, item := range ctx.Else_().AllStmt() {
			stmt := v.visitAst(item)
			alt.Statements = append(alt.Statements, stmt)
		}

		alts = append(alts, alt)
	}

	ast.Alternatives = alts

	return ast
}

func (v *AstVisitor) mustResolveVariable(name string) *Variable {
	current := v.scope

	for current != nil {
		if vr, ok := current.variables[name]; ok {
			return vr
		}

		current = current.exit()
	}

	panic(fmt.Sprintf("variable '%s' not found", name))
}

func (v *AstVisitor) mustResolve(ids []antlr.TerminalNode) (*Variable, *StructField) {
	vr := v.mustResolveVariable(ids[0].GetText())
	t := v.scope.resolveType(vr.t)

	if len(ids) == 1 {
		return vr, nil
	}

	if t.simple != nil {
		panic("cannot resolve simple type access")
	}

	var f *StructField

	nvr := vr

	for i := 1; i < len(ids); i++ {
		vr = nvr
		t = v.scope.resolveType(vr.t)
		id := ids[i].GetText()
		f = t.complex.fields[id]
		nvr = vr.fields[id]
	}

	return vr, f
}

func (v *AstVisitor) VisitDefinition(ctx *parser.DefinitionContext) interface{} {
	id := ctx.ID().GetText()
	vr := v.scope.variables[id]
	t := v.scope.resolveType(vr.t)
	e := ctx.Expr()

	ast := &CealAssign{
		V:     vr,
		T:     t,
		Value: v.visitAst(e),
	}

	return ast
}

func (v *AstVisitor) VisitDefinitionStmt(ctx *parser.DefinitionStmtContext) interface{} {
	id := ctx.Definition().ID().GetText()
	vr := v.scope.variables[id]
	t := v.scope.resolveType(vr.t)
	e := ctx.Definition().Expr()

	ast := &CealDefine{
		V:     vr,
		T:     t,
		Value: v.visitAst(e),
	}

	return ast
}

func (v *AstVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	var ast *CealFunction

	id := ctx.ID().GetText()
	fun := v.global.functions[id]

	v.scope = fun.user.scope
	{
		ast = &CealFunction{
			Fun: fun,
		}

		for _, item := range ctx.AllStmt() {
			if stmt := v.visitAst(item); stmt != nil {
				ast.Statements = append(ast.Statements, stmt)
			}
		}
		v.program.registerFunction(ast)
	}
	v.scope = v.scope.exit()

	return ast
}

func (v *AstVisitor) VisitBlockStmt(ctx *parser.BlockStmtContext) interface{} {
	ast := &CealBlock{}

	v.scope = v.scope.enter()
	{
		for _, item := range ctx.AllStmt() {
			stmt := v.visitAst(item)
			if stmt != nil {
				ast.Statements = append(ast.Statements, stmt)
			}
		}
	}
	v.scope = v.scope.exit()

	return ast
}

func (v *AstVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.scope = v.global

	var aff CealAffixable

	{
		for _, ch := range ctx.GetChildren() {
			tree := ch.(antlr.ParseTree)
			res := tree.Accept(v)

			if e, ok := res.(CealAffixable); ok {
				aff = e

				e.AddPrefix(v.comments)
				v.comments = []CealAst{}
			}
		}
	}

	if len(v.comments) > 0 {
		if aff != nil {
			aff.AddSuffix(v.comments)
			v.comments = []CealAst{}
		}
	}

	v.scope = nil

	return nil
}

func (v *AstVisitor) VisitGotoStmt(ctx *parser.GotoStmtContext) interface{} {
	ast := &CealGoto{
		Label: ctx.ID().GetText(),
	}

	return ast
}

func (v *AstVisitor) VisitLabelStmt(ctx *parser.LabelStmtContext) interface{} {
	ast := &CealLabel{
		Name: ctx.ID().GetText(),
	}

	return ast
}

func (v *AstVisitor) VisitPostIncDecExpr(ctx *parser.PostIncDecExprContext) interface{} {
	ast := &CealPostfix{
		V:  v.mustResolveVariable(ctx.ID().GetText()),
		Op: ctx.Incdec().GetText(),
	}

	return ast
}

func (v *AstVisitor) VisitPreIncDecExpr(ctx *parser.PreIncDecExprContext) interface{} {
	ast := &CealPrefix{
		V:  v.mustResolveVariable(ctx.ID().GetText()),
		Op: ctx.Incdec().GetText(),
	}

	return ast
}

func (v *AstVisitor) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	v.scope = v.scope.enter()
	loop := v.loops.Push("for")

	init := []CealAst{}

	if ctx.ForInit().Definition() != nil {
		ast := v.visitAst(ctx.ForInit().Definition())
		if e, ok := ast.(CealExpr); ok {
			e.ToStmt()
		}
		init = append(init, ast)
	}

	for _, e := range ctx.ForInit().AllExpr() {
		ast := v.visitAst(e)
		if e, ok := ast.(CealExpr); ok {
			e.ToStmt()
		}

		init = append(init, ast)
	}

	condition := v.visitAst(ctx.ForCondition().Expr())

	iter := []CealAst{}

	for _, e := range ctx.ForIter().AllExpr() {
		ast := v.visitAst(e)
		if e, ok := ast.(CealExpr); ok {
			e.ToStmt()
		}

		iter = append(iter, ast)
	}

	stmt := v.visitAst(ctx.Stmt())

	ast := &CealFor{
		Index:     v.index,
		Loop:      loop,
		Init:      init,
		Condition: condition,
		Iter:      iter,
		Statement: stmt,
	}

	v.index++

	v.loops.Pop()
	v.scope = v.scope.exit()

	return ast
}

func (v *AstVisitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	ast := &CealWhile{
		Index:     v.index,
		Loop:      v.loops.Push("while"),
		Condition: v.visitAst(ctx.Expr()),
		Statement: v.visitAst(ctx.Stmt()),
	}

	v.index++

	v.loops.Pop()

	return ast
}

func (v *AstVisitor) VisitDoWhileStmt(ctx *parser.DoWhileStmtContext) interface{} {
	ast := &CealDoWhile{
		Index:     v.index,
		Loop:      v.loops.Push("do"),
		Condition: v.visitAst(ctx.Expr()),
		Statement: v.visitAst(ctx.Stmt()),
	}

	v.index++

	v.loops.Pop()

	return ast
}

func (v *AstVisitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	return &CealBreak{
		Label: v.loops.Break(),
		Index: v.index,
	}
}

func (v *AstVisitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	return &CealContinue{
		Label: v.loops.Continue(),
		Index: v.index,
	}
}

func (v *AstVisitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	ast := &CealSwitch{
		Index: v.index,
		Loop:  v.loops.Push("switch"),
		Value: v.visitAst(ctx.Expr()),
	}

	for _, c := range ctx.AllCase_() {
		stmts := []CealAst{}

		for _, item := range c.AllStmt() {
			stmts = append(stmts, v.visitAst(item))
		}

		sc := &CealSwitchCase{
			Value:      v.visitAst(c.Expr()),
			Statements: stmts,
		}

		ast.Cases = append(ast.Cases, sc)
	}

	if ctx.Default_() != nil {
		stmts := []CealAst{}

		for _, item := range ctx.Default_().AllStmt() {
			stmts = append(stmts, v.visitAst(item))
		}

		ast.Default = stmts
	}

	v.loops.Pop()
	v.index++

	return ast
}

func (v *AstVisitor) VisitConditionalExpr(ctx *parser.ConditionalExprContext) interface{} {
	ast := &CealConditional{
		Index:     v.index,
		Condition: v.visitAst(ctx.GetCondition()),
		True:      v.visitAst(ctx.GetTrue_()),
		False:     v.visitAst(ctx.GetFalse_()),
	}

	v.index++

	return ast

}

func (v *AstVisitor) VisitCommentStmt(ctx *parser.CommentStmtContext) interface{} {
	comment := ctx.Comment()
	return v.visitComment(comment)
}

func (v *AstVisitor) VisitComment(ctx *parser.CommentContext) interface{} {
	ast := v.visitComment(ctx)
	v.comments = append(v.comments, ast)
	return nil
}

func (v *AstVisitor) visitComment(ctx parser.ICommentContext) CealAst {
	if ctx.SINGLE_COMMENT() != nil {
		return &CealSingleLineComment{
			Line: strings.TrimPrefix(
				strings.Trim(ctx.GetText(), "\r\n"),
				"//",
			),
		}
	}

	if ctx.MULTILINE_COMMENT() != nil {
		return &CealMultiLineComment{
			Text: strings.TrimSuffix(strings.TrimPrefix(ctx.GetText(), "/*"), "*/"),
		}
	}

	return nil
}

func (v *AstVisitor) VisitAsmStmt(ctx *parser.AsmStmtContext) interface{} {
	var ops []string
	for _, s := range ctx.AllSTRING() {
		ops = append(ops, ceal_TrimSTRINGQuotes(s.GetText()))
	}
	return &CealAsm{
		Value: strings.Join(ops, "\n"),
	}
}

func (v *AstVisitor) VisitSubscriptExpr(ctx *parser.SubscriptExprContext) interface{} {
	return v.Visit(ctx.Subscript_expr())
}

func (v *AstVisitor) VisitSubscript_expr(ctx *parser.Subscript_exprContext) interface{} {
	id := ctx.ID().GetText()
	vr := v.mustResolveVariable(id)

	return &CealSubscript{
		V:     vr,
		Index: v.visitAst(ctx.Expr()),
	}
}
