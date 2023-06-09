package compiler

import (
	"ceal/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func (v *SymbolTableVisitor) initVariable(vr *Variable) {
	t := vr.t

	if t.avm != nil {
		if vr.local != nil {
			vr.local.slot = v.slot
			v.slot++
		}
		if vr.param != nil {
			vr.param.index = v.index
			v.index--
		}
		return
	}

	if t.complex.builtin != nil {
		return
	}

	fields := map[string]*Variable{}

	for _, name := range t.complex.fieldsNames {
		f := t.complex.fields[name]

		fv := &Variable{
			name: f.name,
			t:    f.t,
		}

		if vr.local != nil {
			fv.local = &LocalVariable{}
		}

		if vr.param != nil {
			fv.param = &ParameterVariable{
				index: v.index,
			}
		}

		v.initVariable(fv)

		fields[f.name] = fv
	}

	vr.fields = fields
}

type SymbolTableVisitor struct {
	*parser.BaseCVisitor

	program *CealProgram

	global *Scope
	scope  *Scope

	slot  int
	index int
}

func (v *SymbolTableVisitor) VisitDeclarationStmt(ctx *parser.DeclarationStmtContext) interface{} {
	id := ctx.Declaration().ID().GetText()

	if _, ok := v.scope.variables[id]; ok {
		panic(fmt.Sprintf("variable '%s' is already defined", id))
	}

	local := &LocalVariable{}

	t := v.scope.resolveType(ctx.Declaration().Type_().ID().GetText())

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

	local := &LocalVariable{}

	vr := &Variable{
		constant: ctx.Definition().Type_().Const_() != nil,
		t:        v.scope.resolveType(ctx.Definition().Type_().ID().GetText()),
		name:     id,
		local:    local,
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

	s := &Struct{
		fields:    map[string]*StructField{},
		functions: map[string]*StructFunction{},
	}

	for _, item := range ctx.AllField() {
		t := v.scope.resolveType(item.Type_().ID().GetText())
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

	v.global.registerType(t)

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

	var args []*FunctionParam

	for _, p := range ctx.Params().AllParam() {
		t := v.scope.resolveType(p.Type_().ID().GetText())
		arg := &FunctionParam{
			t:    t,
			name: p.ID().GetText(),
		}

		args = append(args, arg)
	}
	user := &UserFunction{
		scope: NewScope(v.scope),
		args:  args,
	}

	ret := ctx.Type_().ID().GetText()
	t := v.scope.resolveType(ret)

	fun := &Function{
		t:    ret,
		name: id,
		user: user,
	}

	if t.complex != nil {
		for _, name := range t.complex.fieldsNames {
			f := t.complex.fields[name]
			fun.returns = append(fun.returns, &FunctionParam{
				t:    f.t,
				name: name,
			})
		}
	} else if t != noneType {
		fun.returns = []*FunctionParam{
			{
				t:    t,
				name: "r1",
			},
		}
	}

	user.sub = id != AvmMainName

	v.scope.registerFunction(fun)

	v.scope.variables[fun.name] = &Variable{
		t:    v.scope.resolveType("method_t"),
		name: fun.name,
		fun:  fun,
	}

	v.scope = user.scope
	v.scope.function = fun

	v.index = -1

	allparam := ctx.Params().AllParam()

	for i := len(allparam) - 1; i >= 0; i-- {
		pctx := allparam[i]

		id := pctx.ID().GetText()

		if _, ok := v.scope.variables[id]; ok {
			panic(fmt.Sprintf("param '%s' already defined", id))
		}

		vr := &Variable{
			constant: pctx.Type_().Const_() != nil,
			t:        v.scope.resolveType(pctx.Type_().ID().GetText()),
			name:     id,
			param: &ParameterVariable{
				index: v.index,
			},
		}

		v.initVariable(vr)
		v.scope.variables[id] = vr
	}

	v.VisitChildren(ctx)

	v.scope = v.scope.exit()

	return nil
}

func (v *SymbolTableVisitor) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	v.scope = v.scope.enter()

	init := ctx.ForInit()

	if init.Definition() != nil {
		id := init.Definition().ID().GetText()

		if _, ok := v.scope.variables[id]; ok {
			panic(fmt.Sprintf("variable '%s' is already defined", id))
		}

		local := &LocalVariable{}

		vr := &Variable{
			t:     v.scope.resolveType(init.Definition().Type_().ID().GetText()),
			name:  id,
			local: local,
		}

		v.initVariable(vr)
		v.scope.variables[vr.name] = vr
	}

	v.Visit(ctx.Stmt())

	v.scope = v.scope.exit()

	return nil
}

func (v *SymbolTableVisitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitDoWhileStmt(ctx *parser.DoWhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
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

func (v *SymbolTableVisitor) VisitGlobal(ctx *parser.GlobalContext) interface{} {
	id := ctx.ID().GetText()
	if _, ok := v.scope.variables[id]; ok {
		panic(fmt.Sprintf("variable '%s' is already defined", id))
	}

	tn := ctx.Type_().ID().GetText()

	t := v.scope.resolveType(tn)

	if t.avm == nil {
		panic(fmt.Sprintf("global variables of non built-in type '%s' are not supported yet", t.name))
	}

	vr := &Variable{
		constant: ctx.Type_().Const_() != nil,
		name:     id,
		t:        t,
		const_: &ConstVariable{
			kind: t.avm.kind,
		},
	}

	if !vr.constant {
		panic("non-const global variables are not supported yet")
	}

	v.scope.variables[id] = vr

	switch t.avm.kind {
	case AvmTypeInt:
		vr.const_.index = len(v.program.ConstInts)
		v.program.ConstInts = append(v.program.ConstInts, atoi(ctx.Constant().GetText()))
	case AvmTypeBytes:
		vr.const_.index = len(v.program.ConstBytes)
		v.program.ConstBytes = append(v.program.ConstBytes, []byte(ctx.Constant().GetText()))
	default:
		panic("global variable of this simple type kind isn't supported yet")
	}

	return nil
}
