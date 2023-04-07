package compiler

import (
	"ceal/parser"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type FunctionParam struct {
	t    string
	name string
}

type BuiltinFunction struct {
	op    string
	stack []*FunctionParam
	imm   []*FunctionParam
}

type Function struct {
	t    string
	name string

	builtin *BuiltinFunction
	user    *UserFunction
}

type StructField struct {
	t    string
	name string
}

type StructFunction struct {
	t    string
	name string

	params []*FunctionParam
}

type BuiltinStruct struct {
	op string
}

type Struct struct {
	name      string
	fields    map[string]*StructField
	functions map[string]*StructFunction

	builtin *BuiltinStruct
}

type LocalVariable struct {
	slot int
}

type ParameterVariable struct {
	index int
}

type Variable struct {
	t    string
	name string

	readonly bool

	local *LocalVariable
	param *ParameterVariable
}

type Scope struct {
	structs   map[string]*Struct
	functions map[string]*Function
	variables map[string]*Variable

	children []*Scope

	slot     int
	function *Function

	read bool
	i    int

	parent *Scope
}

func NewScope(parent *Scope) *Scope {
	s := &Scope{
		parent:    parent,
		structs:   map[string]*Struct{},
		functions: map[string]*Function{},
		variables: map[string]*Variable{},
	}

	if parent != nil {
		s.slot = parent.slot + len(parent.children)
	} else {
		s.slot = 0
	}

	return s
}

func (s *Scope) enter() *Scope {
	if s.read {
		if s.i >= len(s.children) {
			return nil
		}

		scope := s.children[s.i]
		s.i++

		return scope
	}

	child := NewScope(s)
	s.children = append(s.children, child)

	return child
}

func (s *Scope) exit() *Scope {
	return s.parent
}

func (s *Scope) readonly() {
	s.read = true

	for _, f := range s.functions {
		user := f.user
		if user != nil {
			scope := user.scope
			if scope != nil {
				scope.readonly()
			}
		}
	}

	for _, item := range s.children {
		item.readonly()
	}

	s.i = 0
}

type UserFunction struct {
	args    int
	returns int

	sub bool

	scope *Scope
}

type SymbolTableVisitor struct {
	*parser.BaseCVisitor

	global *Scope
	scope  *Scope
}

func (v *SymbolTableVisitor) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	id := ctx.ID().GetText()

	if _, ok := v.scope.variables[id]; ok {
		panic(fmt.Sprintf("variable '%s' is already defined", id))
	}

	local := &LocalVariable{
		slot: v.scope.slot + len(v.scope.variables),
	}

	vr := &Variable{
		t:     ctx.Type_().ID().GetText(),
		name:  id,
		local: local,
	}

	v.scope.variables[vr.name] = vr

	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitDefinition(ctx *parser.DefinitionContext) interface{} {
	id := ctx.ID().GetText()

	if _, ok := v.scope.variables[id]; ok {
		panic(fmt.Sprintf("variable '%s' is already defined", id))
	}

	local := &LocalVariable{
		slot: v.scope.slot + len(v.scope.variables),
	}

	vr := &Variable{
		t:     ctx.Type_().ID().GetText(),
		name:  id,
		local: local,
	}

	v.scope.variables[vr.name] = vr

	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.scope = v.scope.enter()
	v.VisitChildren(ctx)
	v.scope = v.scope.exit()

	return nil
}

func (v *SymbolTableVisitor) VisitStruct(ctx *parser.StructContext) interface{} {
	name := ctx.ID().GetText()

	if _, ok := v.global.structs[name]; ok {
		panic(fmt.Sprintf("struct '%s' is already defined", name))
	}

	s := &Struct{
		name:      name,
		fields:    map[string]*StructField{},
		functions: map[string]*StructFunction{},
	}

	for _, item := range ctx.AllField() {
		t := item.Type_().ID().GetText()
		name := item.ID().GetText()

		f := &StructField{
			t:    t,
			name: name,
		}

		if _, ok := s.fields[name]; ok {
			panic(fmt.Sprintf("field '%s' is already defined", name))
		}

		s.fields[name] = f
	}

	v.global.structs[s.name] = s

	return nil
}

func (v *SymbolTableVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.scope = v.global
	v.VisitChildren(ctx)
	v.scope = nil

	return nil
}

func (v *SymbolTableVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	id := ctx.ID().GetText()

	if _, ok := v.scope.functions[id]; ok {
		panic(fmt.Sprintf("function '%s' already defined", id))
	}

	user := &UserFunction{
		scope: NewScope(v.scope),
		args:  len(ctx.Params().AllParam()),
	}

	ret := ctx.Type_().ID().GetText()
	t := v.global.structs[ret]

	if t != nil {
		user.returns = len(t.fields)
	} else {
		if ret == "void" {
			user.returns = 0
		} else {
			user.returns = 1
		}
	}

	user.sub = id != AvmMainName

	fun := &Function{
		t:    ret,
		name: id,
		user: user,
	}

	v.scope.functions[id] = fun

	v.scope = user.scope
	v.scope.function = fun

	index := -len(ctx.Params().AllParam())

	for _, pctx := range ctx.Params().AllParam() {
		id := pctx.ID().GetText()

		if _, ok := v.scope.variables[id]; ok {
			panic(fmt.Sprintf("param '%s' already defined", id))
		}

		param := &ParameterVariable{
			index: index,
		}

		vr := &Variable{
			t:     pctx.Type_().ID().GetText(),
			name:  id,
			param: param,
		}

		index++

		v.scope.variables[id] = vr
	}

	v.VisitChildren(ctx)

	v.scope = v.scope.exit()

	return nil
}

func (v *SymbolTableVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *SymbolTableVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}

type ASTVisitor struct {
	*parser.BaseCVisitor

	global *Scope
	scope  *Scope

	program *AstProgram

	label int
}

func (v *ASTVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *ASTVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}

type AstIsReturn interface {
	IsReturn()
}

type AstVariable struct {
	v *Variable
}

func (a *AstVariable) String() string {
	if a.v.local != nil {
		return fmt.Sprintf("load %d", a.v.local.slot)
	}

	if a.v.param != nil {
		return fmt.Sprintf("frame_dig %d", a.v.param.index)
	}

	if a.v.t == "uint64" {
		return fmt.Sprintf("int %s", a.v.name)
	}

	return fmt.Sprintf("byte %s", a.v.name)
}

type AstAssignVariable struct {
	v     *Variable
	value AstStatement
}

func (a *AstAssignVariable) String() string {
	if a.v.param != nil {
		// TODO: add param var assignment support
		panic("cannot assign param var")
	}

	return fmt.Sprintf("%s\nstore %d", a.value.String(), a.v.local.slot)
}

type AstBinop struct {
	l  AstStatement
	op string
	r  AstStatement
}

func (a *AstBinop) String() string {
	return fmt.Sprintf("%s\n%s\n%s", a.l.String(), a.r.String(), a.op)
}

type AstMinusOp struct {
	value AstStatement
}

func (a *AstMinusOp) String() string {
	return fmt.Sprintf("int 0\n%s\n-", a.value.String())
}

type AstAssignStructField struct {
	v *Variable
	s *Struct
	f *StructField

	value AstStatement
}

func (a *AstAssignStructField) String() string {
	if a.s.builtin != nil {
		return fmt.Sprintf("%s %s", a.s.builtin.op, a.f.name)
	} else {
		// TODO
		panic("not implemented")
	}
}

type AstStructField struct {
	v *Variable
	s *Struct
	f *StructField
}

func (a *AstStructField) String() string {
	if a.s.builtin != nil {
		return fmt.Sprintf("%s %s", a.s.builtin.op, a.f.name)
	} else {
		panic("not implemented")
	}
}

type AstCall struct {
	fun  *Function
	args []AstStatement
}

func (a *AstCall) String() string {
	s := strings.Builder{}

	if a.fun.builtin != nil {
		i := 0

		for ; i < len(a.fun.builtin.stack); i++ {
			arg := a.args[i]
			s.WriteString(arg.String())
			s.WriteString("\n")
		}

		s.WriteString(a.fun.builtin.op)

		for ; i < len(a.fun.builtin.imm); i++ {
			arg := a.args[i]
			s.WriteString(" ")
			s.WriteString(arg.String())
		}

		s.WriteString("\n")

		if a.fun.user != nil {
			for _, arg := range a.args {
				s.WriteString(arg.String())
				s.WriteString("\n")
			}

			s.WriteString(fmt.Sprintf("callsub %s", a.fun.name))
		}
	}

	return s.String()
}

type AstIntConstant struct {
	value string
}

func (a *AstIntConstant) String() string {
	return fmt.Sprintf("int %s", a.value)
}

type AstByteConstant struct {
	value string
}

func (a *AstByteConstant) String() string {
	return fmt.Sprintf("byte %s", a.value)
}

type AstReturn struct {
	value    AstStatement
	function *Function
}

func (a *AstReturn) IsReturn() {
}

func (a *AstReturn) String() string {
	s := strings.Builder{}

	if a.value != nil {
		s.WriteString(a.value.String())
		s.WriteString("\n")
	}

	if a.function != nil {
		if a.function.user.sub {
			s.WriteString("retsub")
		} else {
			s.WriteString("return")
		}
	} else {
		s.WriteString("return")
	}

	return s.String()
}

type AstBlock struct {
	statements []AstStatement
}

func (a *AstBlock) String() string {
	ops := strings.Builder{}

	for _, stmt := range a.statements {
		ops.WriteString(stmt.String())
		ops.WriteString("\n")
	}

	return ops.String()
}

type AstIf struct {
	condition    AstStatement
	label        int
	statements   []AstStatement
	alternatives []AstStatement
}

func (a *AstIf) String() string {
	ops := strings.Builder{}

	for _, stmt := range a.statements {
		ops.WriteString(stmt.String())
		ops.WriteString("\n")
	}

	alts := strings.Builder{}

	if len(a.alternatives) > 0 {
		alts.WriteString("\n")

		for _, stmt := range a.alternatives {
			alts.WriteString(stmt.String())
		}
	}

	return fmt.Sprintf("%s\nbz skip_%d\n%s\nb end_%d\nskip_%d:\n%s\nend_%d:",
		a.condition.String(), a.label, ops.String(), a.label, a.label, alts.String(), a.label)
}

type AstFunction struct {
	fun        *Function
	statements []AstStatement
}

func (a *AstFunction) String() string {
	res := strings.Builder{}

	res.WriteString(fmt.Sprintf("// function %s\n", a.fun.name))
	res.WriteString(fmt.Sprintf("%s:\n", a.fun.name))

	if a.fun.user.sub {
		res.WriteString(fmt.Sprintf("proto %d %d\n", a.fun.user.args, a.fun.user.returns))
	}

	for _, stmt := range a.statements {
		res.WriteString(stmt.String())
		res.WriteString("\n")
	}

	if a.fun.user.sub {
		if len(a.statements) > 0 {
			last := a.statements[len(a.statements)-1]
			if _, ok := last.(AstIsReturn); !ok {
				res.WriteString("retsub\n")
			}
		}
	}

	return res.String()
}

type AstProgram struct {
	functions map[string]*AstFunction
}

func (a *AstProgram) String() string {
	res := strings.Builder{}

	res.WriteString(fmt.Sprintf("#pragma version %d\n", AvmVersion))

	main := a.functions[AvmMainName]

	res.WriteString(fmt.Sprintf("b %s\n", main.fun.name))

	for name, fun := range a.functions {
		if name == AvmMainName {
			continue
		}

		res.WriteString(fun.String())
		res.WriteString("\n")
	}

	res.WriteString(main.String())
	res.WriteString("\n")

	return res.String()
}

func (v *ASTVisitor) visitStatement(tree antlr.ParseTree) AstStatement {
	res := v.Visit(tree)
	if res == nil {
		return nil
	}

	return res.(AstStatement)
}

func (v *ASTVisitor) VisitMemberExpr(ctx *parser.MemberExprContext) interface{} {
	ids := ctx.AllID()
	id := ids[0].GetText()

	vr := v.resolveVariable(id)
	for i := 1; i < len(ids); i++ {
		s := v.global.structs[vr.t]
		id := ids[i].GetText()

		ast := &AstStructField{
			v: vr,
			s: s,
			f: s.fields[id],
		}

		return ast
	}

	return nil
}

func (v *ASTVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	ids := ctx.AllID()

	vr := v.resolveVariable(ids[0].GetText())
	if vr.readonly {
		panic(fmt.Sprintf("variable '%s' is read only", vr.name))
	}

	if len(ids) == 1 {
		ast := &AstAssignVariable{
			value: v.visitStatement(ctx.Expr()),
			v:     vr,
		}

		return ast
	}

	for i := 1; i < len(ids); i++ {
		s := v.global.structs[vr.t]
		id := ids[i].GetText()

		// TODO: support nested fields
		ast := &AstAssignStructField{
			s:     s,
			f:     s.fields[id],
			v:     vr,
			value: v.visitStatement(ctx.Expr()),
		}

		return ast
	}

	return nil
}

func (v *ASTVisitor) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	t := ctx.Type_().ID().GetText()
	if t != "uint64" && t != "bytes" {
		if _, ok := v.global.structs[t]; !ok {
			panic(fmt.Sprintf("type '%s' not found", t))
		}
	}

	return nil
}

func (v *ASTVisitor) VisitVariableExpr(ctx *parser.VariableExprContext) interface{} {
	id := ctx.ID().GetText()
	vr := v.resolveVariable(id)

	ast := &AstVariable{
		v: vr,
	}

	return ast
}

func (v *ASTVisitor) VisitMinusExpr(ctx *parser.MinusExprContext) interface{} {
	return &AstMinusOp{
		value: v.visitStatement(ctx.Expr()),
	}
}

func (v *ASTVisitor) VisitAddSubExpr(ctx *parser.AddSubExprContext) interface{} {
	exprs := ctx.AllExpr()
	return &AstBinop{
		l:  v.visitStatement(exprs[0]),
		r:  v.visitStatement(exprs[1]),
		op: ctx.Addsub().GetText(),
	}
}
func (v *ASTVisitor) VisitMulDivExpr(ctx *parser.MulDivExprContext) interface{} {
	exprs := ctx.AllExpr()
	return &AstBinop{
		l:  v.visitStatement(exprs[0]),
		r:  v.visitStatement(exprs[1]),
		op: ctx.Muldiv().GetText(),
	}
}

func (v *ASTVisitor) VisitEqNeqExpr(ctx *parser.EqNeqExprContext) interface{} {
	exprs := ctx.AllExpr()
	return &AstBinop{
		l:  v.visitStatement(exprs[0]),
		r:  v.visitStatement(exprs[1]),
		op: ctx.Eqneq().GetText(),
	}
}

func (v *ASTVisitor) VisitConstantExpr(ctx *parser.ConstantExprContext) interface{} {
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

func (v *ASTVisitor) VisitCallExpr(ctx *parser.CallExprContext) interface{} {
	return v.Visit(ctx.Call_expr())
}
func (v *ASTVisitor) VisitCall(ctx *parser.CallContext) interface{} {
	return v.Visit(ctx.Call_expr())
}

func (v *ASTVisitor) VisitCall_expr(ctx *parser.Call_exprContext) interface{} {
	ids := ctx.AllID()

	if len(ids) > 1 {
		panic("calling struct function is not supported yet")
	}

	id := ids[0].GetText()
	fun := v.global.functions[id]

	ast := &AstCall{
		fun: fun,
	}

	for _, arg := range ctx.Args().AllExpr() {
		stmt := v.visitStatement(arg)
		ast.args = append(ast.args, stmt)
	}

	return ast
}

func (v *ASTVisitor) VisitReturn(ctx *parser.ReturnContext) interface{} {
	ast := &AstReturn{
		function: v.scope.function,
	}

	if ctx.Expr() != nil {
		ast.value = v.visitStatement(ctx.Expr())
	}

	return ast
}

func (v *ASTVisitor) VisitIf(ctx *parser.IfContext) interface{} {
	ast := &AstIf{
		condition: v.visitStatement(ctx.Expr()),
		label:     v.label,
	}

	v.label++

	for _, item := range ctx.AllStmt() {
		if stmt := v.visitStatement(item); stmt != nil {
			ast.statements = append(ast.statements, stmt)
		}
	}

	if len(ctx.AllElseif()) > 0 {
		// TODO: handle 'else if'
	}

	if ctx.Else_() != nil {
		for _, item := range ctx.Else_().AllStmt() {
			stmt := v.visitStatement(item)
			ast.alternatives = append(ast.alternatives, stmt)
		}
	}

	return ast
}

func (v *ASTVisitor) resolveVariable(name string) *Variable {
	current := v.scope

	for current != nil {
		if vr, ok := current.variables[name]; ok {
			return vr
		}

		current = current.exit()
	}

	panic(fmt.Sprintf("variable '%s' not found", name))
}

func (v *ASTVisitor) VisitDefinition(ctx *parser.DefinitionContext) interface{} {
	id := ctx.ID().GetText()
	vr := v.scope.variables[id]

	ast := &AstAssignVariable{
		v:     vr,
		value: v.visitStatement(ctx.Expr()),
	}

	return ast
}

func (v *ASTVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	id := ctx.ID().GetText()
	fun := v.global.functions[id]

	ast := &AstFunction{
		fun: fun,
	}

	v.scope = fun.user.scope

	for _, item := range ctx.AllStmt() {
		if stmt := v.visitStatement(item); stmt != nil {
			ast.statements = append(ast.statements, stmt)
		}
	}

	v.scope = v.scope.exit()

	v.program.functions[id] = ast

	return nil
}

func (v *ASTVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
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

func (v *ASTVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.scope = v.global
	v.VisitChildren(ctx)
	v.scope = nil

	return nil
}

func Compile(src string) string {
	input := antlr.NewInputStream(src)
	lexer := parser.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCParser(stream)

	global := NewScope(nil)

	for _, item := range builtin_functions {
		f := &Function{
			t:    item.t,
			name: item.name,
			builtin: &BuiltinFunction{
				op: item.op,
			},
		}

		for _, item := range item.stack {
			f.builtin.stack = append(f.builtin.stack, &FunctionParam{
				t:    item.t,
				name: item.name,
			})
		}

		for _, item := range item.imm {
			f.builtin.imm = append(f.builtin.imm, &FunctionParam{
				t:    item.t,
				name: item.name,
			})
		}

		global.functions[f.name] = f
	}

	for _, item := range builtin_structs {
		s := &Struct{
			name: item.name,
			builtin: &BuiltinStruct{
				op: item.op,
			},
			fields:    map[string]*StructField{},
			functions: map[string]*StructFunction{},
		}

		for _, item := range item.fields {
			s.fields[item.name] = &StructField{
				t:    item.t,
				name: item.name,
			}
		}

		for _, item := range item.functions {
			f := &StructFunction{
				t:    item.t,
				name: item.name,
			}

			for _, item := range item.params {
				f.params = append(f.params, &FunctionParam{
					t:    item.t,
					name: item.name,
				})
			}

			s.functions[f.name] = f
		}

		global.structs[s.name] = s
	}

	for _, item := range builtin_variables {
		v := &Variable{
			t:    item.t,
			name: item.name,
		}

		global.variables[v.name] = v
	}

	sv := &SymbolTableVisitor{
		BaseCVisitor: &parser.BaseCVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
		global: global,
	}

	stream.Seek(0)
	sv.Visit(p.Program())

	program := &AstProgram{
		functions: map[string]*AstFunction{},
	}

	global.readonly()

	av := &ASTVisitor{
		BaseCVisitor: &parser.BaseCVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
		global:  global,
		program: program,
	}

	stream.Seek(0)
	av.Visit(p.Program())

	return program.String()
}
