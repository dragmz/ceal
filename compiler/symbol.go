package compiler

import (
	"ceal/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func (v *SymbolTableVisitor) initVariable(vr *Variable) {
	t := v.scope.resolveType(vr.t)

	if t.simple != nil {
		return
	}

	if t.complex.builtin != nil {
		return
	}

	if vr.param != nil {
		panic("struct types are not supported for parameters yet")
	}

	fields := map[string]*Variable{}

	for _, name := range t.complex.fieldsNames {
		f := t.complex.fields[name]
		fv := &Variable{
			name: f.name,
			t:    f.t,
			local: &LocalVariable{
				slot: v.slot,
			},
		}

		v.slot++

		v.initVariable(fv)

		fields[f.name] = fv
	}

	vr.fields = fields
}

type SymbolTableVisitor struct {
	*parser.BaseCVisitor

	global *Scope
	scope  *Scope

	slot int
}

func (v *SymbolTableVisitor) VisitDeclarationStmt(ctx *parser.DeclarationStmtContext) interface{} {
	id := ctx.Declaration().ID().GetText()

	if _, ok := v.scope.variables[id]; ok {
		panic(fmt.Sprintf("variable '%s' is already defined", id))
	}

	local := &LocalVariable{
		slot: v.slot,
	}

	v.slot++

	t := ctx.Declaration().Type_().ID().GetText()

	vr := &Variable{
		t:     t,
		name:  id,
		local: local,
	}

	v.initVariable(vr)
	v.scope.variables[vr.name] = vr

	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitDefinitionStmt(ctx *parser.DefinitionStmtContext) interface{} {
	id := ctx.Definition().ID().GetText()

	if _, ok := v.scope.variables[id]; ok {
		panic(fmt.Sprintf("variable '%s' is already defined", id))
	}

	local := &LocalVariable{
		slot: v.slot,
	}

	v.slot++

	vr := &Variable{
		t:     ctx.Definition().Type_().ID().GetText(),
		name:  id,
		local: local,
	}

	v.initVariable(vr)
	v.scope.variables[vr.name] = vr

	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitBlockStmt(ctx *parser.BlockStmtContext) interface{} {
	v.scope = v.scope.enter()
	v.VisitChildren(ctx)
	v.scope = v.scope.exit()

	return nil
}

func (v *SymbolTableVisitor) VisitStruct(ctx *parser.StructContext) interface{} {
	name := ctx.ID().GetText()

	if t := v.scope.resolveType(name); t != nil {
		panic(fmt.Sprintf("type '%s' is already defined", name))
	}

	s := &Struct{
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
		s.fieldsNames = append(s.fieldsNames, f.name)
	}

	t := &Type{
		name:    name,
		complex: s,
	}

	v.global.types[t.name] = t

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
	t := v.scope.resolveType(ret)

	if t.complex != nil {
		user.returns = len(t.complex.fields)
	} else {
		if t.simple.empty {
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

		index++

		vr := &Variable{
			t:     pctx.Type_().ID().GetText(),
			name:  id,
			param: param,
		}

		v.initVariable(vr)
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
