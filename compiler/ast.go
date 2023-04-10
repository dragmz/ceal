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

	label int
	slot  int
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
		label: v.label,
	}

	v.label++

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

	if len(ids) > 2 {
		panic("multilevel member access is not supported yet")
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

func (v *AstVisitor) VisitDefinitionStmt(ctx *parser.DefinitionStmtContext) interface{} {
	id := ctx.Definition().ID().GetText()
	vr := v.scope.variables[id]
	t := v.scope.resolveType(vr.t)

	ast := &AstAssign{
		v:     vr,
		t:     t,
		value: v.visitStatement(ctx.Definition().Expr()),
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
