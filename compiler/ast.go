package compiler

import (
	"ceal/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type AstVisitor struct {
	*parser.BaseCVisitor

	global *Scope
	scope  *Scope

	program *AstProgram

	index int
	slot  int

	labels *LabelScope
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

func (v *AstVisitor) visitStatement(tree antlr.ParseTree) AstStatement {
	res := v.Visit(tree)
	if res == nil {
		return nil
	}

	return res.(AstStatement)
}

func (v *AstVisitor) VisitMemberExpr(ctx *parser.MemberExprContext) interface{} {
	vr, f := v.mustResolve(ctx.AllID())
	fun := v.global.functions[f.fun]

	ast := &AstStructField{
		v:   vr,
		t:   v.scope.resolveType(vr.t),
		f:   f,
		fun: fun,
	}

	return ast
}

func (v *AstVisitor) VisitAssignmentStmt(ctx *parser.AssignmentStmtContext) interface{} {
	ids := ctx.AllID()

	vr, f := v.mustResolve(ids)
	if vr.readonly {
		panic(fmt.Sprintf("variable '%s' is read only", vr.name))
	}

	if vr.readonly {
		panic(fmt.Sprintf("variable '%s' is read only", vr.name))
	}

	ast := &AstAssign{
		v:     vr,
		t:     v.scope.resolveType(vr.t),
		f:     f,
		value: v.visitStatement(ctx.Expr()),
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

func (v *AstVisitor) VisitVariableExpr(ctx *parser.VariableExprContext) interface{} {
	id := ctx.ID().GetText()
	vr := v.mustResolveVariable(id)

	ast := &AstVariable{
		v: vr,
	}

	return ast
}

func (v *AstVisitor) VisitMinusExpr(ctx *parser.MinusExprContext) interface{} {
	return &AstMinusOp{
		value: v.visitStatement(ctx.Expr()),
	}
}

func (v *AstVisitor) VisitAddSubExpr(ctx *parser.AddSubExprContext) interface{} {
	exprs := ctx.AllExpr()
	return &AstBinop{
		l:  v.visitStatement(exprs[0]),
		r:  v.visitStatement(exprs[1]),
		op: ctx.Addsub().GetText(),
	}
}

func (v *AstVisitor) VisitOrExpr(ctx *parser.OrExprContext) interface{} {
	exprs := ctx.AllExpr()
	return &AstBinop{
		l:  v.visitStatement(exprs[0]),
		r:  v.visitStatement(exprs[1]),
		op: "||",
	}
}

func (v *AstVisitor) VisitAndExpr(ctx *parser.AndExprContext) interface{} {
	exprs := ctx.AllExpr()
	return &AstBinop{
		l:  v.visitStatement(exprs[0]),
		r:  v.visitStatement(exprs[1]),
		op: "&&",
	}
}

func (v *AstVisitor) VisitMulDivExpr(ctx *parser.MulDivExprContext) interface{} {
	exprs := ctx.AllExpr()
	return &AstBinop{
		l:  v.visitStatement(exprs[0]),
		r:  v.visitStatement(exprs[1]),
		op: ctx.Muldiv().GetText(),
	}
}

func (v *AstVisitor) VisitEqNeqExpr(ctx *parser.EqNeqExprContext) interface{} {
	exprs := ctx.AllExpr()
	return &AstBinop{
		l:  v.visitStatement(exprs[0]),
		r:  v.visitStatement(exprs[1]),
		op: ctx.Eqneq().GetText(),
	}
}

func (v *AstVisitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	return &AstUnaryOp{
		s:  v.visitStatement(ctx.Expr()),
		op: "!",
	}
}

func (v *AstVisitor) VisitConstantExpr(ctx *parser.ConstantExprContext) interface{} {
	var res AstStatement

	if ctx.INT() != nil {
		res = &AstIntConstant{
			value: ctx.INT().GetText(),
		}
	}

	if ctx.STRING() != nil {
		res = &AstByteConstant{
			value: ctx.STRING().GetText(),
		}
	}

	return res
}

func (v *AstVisitor) VisitCallExpr(ctx *parser.CallExprContext) interface{} {
	return v.Visit(ctx.Call_expr())
}
func (v *AstVisitor) VisitCallStmt(ctx *parser.CallStmtContext) interface{} {
	return v.Visit(ctx.Call_expr())
}

func (v *AstVisitor) VisitCall_expr(ctx *parser.Call_exprContext) interface{} {
	ids := ctx.AllID()

	id := ids[0].GetText()

	imms := []AstStatement{}

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

		imms = append(imms, &AstRaw{
			value: ids[1].GetText(),
		})
	}

	fun := v.global.functions[id]

	ast := &AstCall{
		fun: fun,
	}

	for _, arg := range ctx.Args().AllExpr() {
		stmt := v.visitStatement(arg)
		ast.args = append(ast.args, stmt)
	}

	ast.args = append(ast.args, imms...)

	return ast
}

func (v *AstVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	ast := &AstReturn{
		function: v.scope.function,
	}

	if ctx.Expr() != nil {
		ast.value = v.visitStatement(ctx.Expr())
	}

	return ast
}

func (v *AstVisitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	alts := []*AstIfAlternative{}

	ast := &AstIf{
		index: v.index,
	}

	v.index++

	alt := &AstIfAlternative{
		condition: v.visitStatement(ctx.Expr()),
	}

	for _, item := range ctx.AllStmt() {
		if stmt := v.visitStatement(item); stmt != nil {
			alt.statements = append(alt.statements, stmt)
		}
	}

	alts = append(alts, alt)

	for _, elif := range ctx.AllElseif() {
		alt := &AstIfAlternative{
			condition: v.visitStatement(elif.Expr()),
		}

		for _, item := range elif.AllStmt() {
			if stmt := v.visitStatement(item); stmt != nil {
				alt.statements = append(alt.statements, stmt)
			}
		}

		alts = append(alts, alt)
	}

	if ctx.Else_() != nil {
		alt := &AstIfAlternative{}

		for _, item := range ctx.Else_().AllStmt() {
			stmt := v.visitStatement(item)
			alt.statements = append(alt.statements, stmt)
		}

		alts = append(alts, alt)
	}

	ast.alternatives = alts

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

	ast := &AstAssign{
		v:     vr,
		t:     t,
		value: v.visitStatement(e),
	}

	return ast
}

func (v *AstVisitor) VisitDefinitionStmt(ctx *parser.DefinitionStmtContext) interface{} {
	id := ctx.Definition().ID().GetText()
	vr := v.scope.variables[id]
	t := v.scope.resolveType(vr.t)
	e := ctx.Definition().Expr()

	ast := &AstAssign{
		v:     vr,
		t:     t,
		value: v.visitStatement(e),
	}

	return ast
}

func (v *AstVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	id := ctx.ID().GetText()
	fun := v.global.functions[id]

	v.scope = fun.user.scope

	ast := &AstFunction{
		fun: fun,
	}

	for _, item := range ctx.AllStmt() {
		if stmt := v.visitStatement(item); stmt != nil {
			ast.statements = append(ast.statements, stmt)
		}
	}

	v.scope = v.scope.exit()

	v.program.functions[id] = ast
	v.program.functionsNames = append(v.program.functionsNames, ast.fun.name)

	return nil
}

func (v *AstVisitor) VisitBlockStmt(ctx *parser.BlockStmtContext) interface{} {
	v.scope = v.scope.enter()

	ast := &AstBlock{}

	for _, item := range ctx.AllStmt() {
		stmt := v.visitStatement(item)
		if stmt != nil {
			ast.statements = append(ast.statements, stmt)
		}
	}

	v.scope = v.scope.exit()

	return ast
}

func (v *AstVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.scope = v.global
	v.VisitChildren(ctx)
	v.scope = nil

	return nil
}

func (v *AstVisitor) VisitGotoStmt(ctx *parser.GotoStmtContext) interface{} {
	ast := &AstGoto{
		label: ctx.ID().GetText(),
	}

	return ast
}

func (v *AstVisitor) VisitLabelStmt(ctx *parser.LabelStmtContext) interface{} {
	ast := &AstLabel{
		name: ctx.ID().GetText(),
	}

	return ast
}

func (v *AstVisitor) VisitPostIncDecExpr(ctx *parser.PostIncDecExprContext) interface{} {
	ast := &AstPostfix{
		v:  v.mustResolveVariable(ctx.ID().GetText()),
		op: ctx.Incdec().GetText(),
	}

	return ast
}

func (v *AstVisitor) VisitPreIncDecExpr(ctx *parser.PreIncDecExprContext) interface{} {
	ast := &AstPrefix{
		v:  v.mustResolveVariable(ctx.ID().GetText()),
		op: ctx.Incdec().GetText(),
	}

	return ast
}

func (v *AstVisitor) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	v.scope = v.scope.enter()
	v.labels.Push("for")

	init := []AstStatement{}

	if ctx.ForInit().Definition() != nil {
		ast := v.visitStatement(ctx.ForInit().Definition())
		init = append(init, ast)
	}

	for _, e := range ctx.ForInit().AllExpr() {
		ast := v.visitStatement(e)
		init = append(init, ast)
	}

	condition := v.visitStatement(ctx.ForCondition().Expr())

	iter := []AstStatement{}

	for _, e := range ctx.ForIter().AllExpr() {
		// TODO: pop is just a temporary solution; should not push onto the stack at all
		ast := &AstPop{
			s: v.visitStatement(e),
		}

		iter = append(iter, ast)
	}

	stmt := v.visitStatement(ctx.Stmt())

	ast := &AstFor{
		index:     v.index,
		init:      init,
		condition: condition,
		iter:      iter,
		s:         stmt,
	}

	v.index++

	v.labels.Pop()
	v.scope = v.scope.exit()

	return ast
}

func (v *AstVisitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	ast := &AstWhile{
		index:     v.index,
		condition: v.visitStatement(ctx.Expr()),
		s:         v.visitStatement(ctx.Stmt()),
	}

	v.index++

	return ast
}

func (v *AstVisitor) VisitDoWhileStmt(ctx *parser.DoWhileStmtContext) interface{} {
	ast := &AstDoWhile{
		index:     v.index,
		condition: v.visitStatement(ctx.Expr()),
		s:         v.visitStatement(ctx.Stmt()),
	}

	v.index++

	return ast
}

func (v *AstVisitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	return &AstBreak{
		label: v.labels.Get(),
		index: v.index,
	}
}

func (v *AstVisitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	ast := &AstSwitch{
		index: v.index,
		value: v.visitStatement(ctx.Expr()),
	}

	v.labels.Push("switch")

	for _, c := range ctx.AllCase_() {
		stmts := []AstStatement{}

		for _, item := range c.AllStmt() {
			stmts = append(stmts, v.visitStatement(item))
		}

		sc := &AstSwitchCase{
			value:      v.visitStatement(c.Expr()),
			statements: stmts,
		}

		ast.cases = append(ast.cases, sc)
	}

	if ctx.Default_() != nil {
		stmts := []AstStatement{}

		for _, item := range ctx.Default_().AllStmt() {
			stmts = append(stmts, v.visitStatement(item))
		}

		ast.default_ = stmts
	}

	v.labels.Pop()
	v.index++

	return ast
}
