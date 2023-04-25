package compiler

import (
	"ceal/parser"
	"ceal/teal"
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

func (v *AstVisitor) VisitDotExpr(ctx *parser.DotExprContext) interface{} {
	var ast CealAst

	vd := valueData{
		dotData: v.mustResolveDot(ctx.Dot_expr()),
	}

	if vd.F == nil {
		ast = &CealVariable{D: vd}
	} else {
		ast = &CealStructField{D: vd}
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
	dd := v.mustResolveValueAccess(ctx.Value_access_expr())

	if dd.V.constant {
		panic(fmt.Sprintf("variable '%s' is read only", dd.V.name))
	}

	ast := &CealAssign{
		D:     dd,
		Value: v.visitAst(ctx.Expr()),
	}

	return ast
}

func (v *AstVisitor) mustResolveValueAccess(ctx parser.IValue_access_exprContext) valueData {
	dot, index := v.mustResolveDotSub(ctx.Dot_expr(), ctx.Subscript_expr())

	return valueData{
		dotData: dot,
		Index:   index,
	}
}

func (v *AstVisitor) mustResolveDotSub(dot parser.IDot_exprContext, sub parser.ISubscript_exprContext) (dotData, CealAst) {
	var de parser.IDot_exprContext

	var index CealAst

	if sub != nil {
		de = sub.Dot_expr()
		index = v.visitAst(sub.Expr())
	} else {
		de = dot
	}

	dd := v.mustResolveDot(de)

	return dd, index
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
	dd := v.mustResolveValueAccess(ctx.Asd_expr().Value_access_expr())

	return &CealAssignSumDiff{
		D:      dd,
		Value:  v.visitAst(ctx.Asd_expr().Expr()),
		Op:     ctx.Asd_expr().Asd().GetText(),
		IsStmt: true,
	}
}
func (v *AstVisitor) VisitAssignSumDiffExpr(ctx *parser.AssignSumDiffExprContext) interface{} {
	dd := v.mustResolveValueAccess(ctx.Asd_expr().Value_access_expr())

	if dd.V.constant {
		panic(fmt.Sprintf("variable '%s' is read only", dd.V.name))
	}

	return &CealAssignSumDiff{
		D:     dd,
		Value: v.visitAst(ctx.Asd_expr().Expr()),
		Op:    ctx.Asd_expr().Asd().GetText(),
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
	sf, f := v.mustResolveFunction(ctx.Value_access_expr().Dot_expr().AllID())

	var field CealAst

	if sf != nil {
		f = sf.f
	}

	if f.builtin != nil {
		// TODO: refactor the ^ and v conditions
		if sf != nil {
			field = &CealRaw{Value: sf.name}
		}
	}

	ast := &CealCall{
		Fun:   f,
		Field: field,
	}

	for _, arg := range ctx.Args().AllExpr() {
		stmt := v.visitAst(arg)
		ast.Args = append(ast.Args, stmt)
	}

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

func (v *AstVisitor) resolveVariable(name string) *Variable {
	current := v.scope

	for current != nil {
		if vr, ok := current.variables[name]; ok {
			return vr
		}

		current = current.exit()
	}

	return nil
}

func (v *AstVisitor) resolveFunction(name string) *Function {
	current := v.scope

	for current != nil {
		if fun, ok := current.functions[name]; ok {
			return fun
		}

		current = current.exit()
	}

	return nil
}

func (v *AstVisitor) mustResolveFunction(ids []antlr.TerminalNode) (*StructFunction, *Function) {
	name := ids[len(ids)-1].GetText()

	if len(ids) > 1 {
		vr := v.resolveVariable(ids[0].GetText())

		nvr := vr
		for i := 1; i < len(ids)-1; i++ {
			vr = nvr
			id := ids[i].GetText()
			nvr = vr.fields[id]
		}

		t := v.scope.resolveType(vr.t)
		sf := t.complex.functions[name]

		if sf == nil {
			panic(fmt.Sprintf("failed to resolve id: '%s'", name))
		}

		return sf, nil
	} else {
		fun := v.resolveFunction(name)

		if fun == nil {
			panic(fmt.Sprintf("failed to resolve id: '%s'", name))
		}

		return nil, fun
	}
}

func (v *AstVisitor) mustResolve(ids []antlr.TerminalNode) (*Variable, *StructField) {
	id := ids[0].GetText()

	vr := v.resolveVariable(id)
	if vr == nil {
		panic(fmt.Sprintf("failed to resolve id: '%s", id))
	}

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
		D: valueData{
			dotData: dotData{
				V: vr,
				T: t,
			},
		},
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
		D: valueData{
			dotData: dotData{
				V: vr,
				T: t,
			},
		},
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

type valueData struct {
	dotData
	Index CealAst
}

type dotData struct {
	V   *Variable
	F   *StructField
	T   *Type
	Fun *Function
}

func (d dotData) Slot() uint8 {
	var slot uint8

	if d.T.complex != nil {
		v := d.V.fields[d.F.name]
		slot = uint8(v.local.slot)
	} else {
		slot = uint8(d.V.local.slot)
	}

	return slot
}

func (d dotData) Load() teal.TealAst {
	if d.V.param != nil {
		return &teal.Teal_frame_dig{Teal_frame_dig_op: teal.Teal_frame_dig_op{I1: int8(d.V.param.index)}}
	}

	return &teal.Teal_load{Teal_load_op: teal.Teal_load_op{I1: d.Slot()}}
}

func (d valueData) Store(op teal.TealAst) teal.TealAst {
	var index teal.TealAst

	if d.Index != nil {
		index = d.Index.TealAst()
	}

	var value teal.TealAst

	if index != nil {
		value = &teal.Teal_setbyte{
			Teal_setbyte_op: teal.Teal_setbyte_op{},
			STACK_1:         d.Load(),
			STACK_2:         index,
			STACK_3:         op,
		}
	} else {
		value = op
	}

	var ast teal.TealAst

	if d.V.param != nil {
		ast = &teal.Teal_frame_bury{
			Teal_frame_bury_op: teal.Teal_frame_bury_op{I1: int8(d.V.param.index)},
			STACK_1:            ast,
		}
	} else {
		ast = &teal.Teal_store{STACK_1: value, Teal_store_op: teal.Teal_store_op{I1: d.Slot()}}
	}

	return ast
}

func (v *AstVisitor) mustResolveDot(ctx parser.IDot_exprContext) dotData {
	var ids []antlr.TerminalNode

	ids = append(ids, ctx.AllID()...)

	vr, f := v.mustResolve(ids)
	t := v.scope.resolveType(vr.t)

	var fun *Function
	if f != nil {
		fun = v.global.functions[f.fun]
	}

	return dotData{
		V:   vr,
		F:   f,
		T:   t,
		Fun: fun,
	}
}

func (v *AstVisitor) VisitPostIncDecStmt(ctx *parser.PostIncDecStmtContext) interface{} {
	ast := v.Visit(ctx.Post_incdec_expr())

	if e, ok := ast.(CealExpr); ok {
		e.ToStmt()
	}

	return ast
}

func (v *AstVisitor) VisitPost_incdec_expr(ctx *parser.Post_incdec_exprContext) interface{} {
	de := v.mustResolveValueAccess(ctx.Value_access_expr())

	ast := &CealPostfix{
		D:  de,
		Op: ctx.Incdec().GetText(),
	}

	return ast
}

func (v *AstVisitor) VisitPostIncDecExpr(ctx *parser.PostIncDecExprContext) interface{} {
	return v.Visit(ctx.Post_incdec_expr())
}

func (v *AstVisitor) VisitPreIncDecExpr(ctx *parser.PreIncDecExprContext) interface{} {
	return v.Visit(ctx.Pre_incdec_expr())
}

func (v *AstVisitor) VisitPreIncDecStmt(ctx *parser.PreIncDecStmtContext) interface{} {
	ast := v.Visit(ctx.Pre_incdec_expr())

	if e, ok := ast.(CealExpr); ok {
		e.ToStmt()
	}

	return ast
}

func (v *AstVisitor) VisitPre_incdec_expr(ctx *parser.Pre_incdec_exprContext) interface{} {
	de := v.mustResolveValueAccess(ctx.Value_access_expr())

	ast := &CealPrefix{
		D:  de,
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

		loop := v.loops.Push("switch")

		for _, item := range c.AllStmt() {
			stmts = append(stmts, v.visitAst(item))
		}

		sc := &CealSwitchCase{
			Loop:       loop,
			Value:      v.visitAst(c.Expr()),
			Statements: stmts,
		}

		ast.Cases = append(ast.Cases, sc)

		v.loops.Pop()
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
			Line: ceal_TrimSINGLE_COMMENTQuotes(ctx.GetText()),
		}
	}

	if ctx.MULTILINE_COMMENT() != nil {
		return &CealMultiLineComment{
			Text: ceal_TrimMULTILINE_COMMENTQuotes(ctx.GetText()),
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
	de := v.mustResolveDot(ctx.Dot_expr())

	return &CealSubscript{
		D: valueData{
			dotData: de,
			Index:   v.visitAst(ctx.Expr()),
		},
	}
}
