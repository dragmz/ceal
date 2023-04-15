package compiler

import (
	"ceal/parser"
	"fmt"
	"os"
	"path"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type CealCompiler struct {
	Includes []string
}

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

	returns int

	builtin *BuiltinFunction
	user    *UserFunction
}

type StructField struct {
	t    string
	name string
	fun  string
}

type StructFunction struct {
	t    string
	name string

	params []*FunctionParam
}

type BuiltinStruct struct {
}

type Struct struct {
	fields      map[string]*StructField
	fieldsNames []string

	functions map[string]*StructFunction

	builtin *BuiltinStruct
}

func (s *Struct) registerFunction(f *StructFunction) {
	if _, ok := s.functions[f.name]; ok {
		panic(fmt.Sprintf("function '%s' is already defined", f.name))
	}

	s.functions[f.name] = f
}

type SimpleTypeKind int

const (
	SimpleTypeInt = SimpleTypeKind(iota)
	SimpleTypeBytes
)

type SimpleType struct {
	kind  SimpleTypeKind
	empty bool
}

type Type struct {
	name    string
	complex *Struct // struct
	simple  *SimpleType
}

type LocalVariable struct {
	slot int
}

type ParameterVariable struct {
	index int
}

type ConstVariable struct {
	kind  SimpleTypeKind
	index int
}

type Variable struct {
	constant bool

	t    string
	name string

	readonly bool

	local  *LocalVariable
	param  *ParameterVariable
	const_ *ConstVariable

	fields map[string]*Variable
}

type LoopScopeItem struct {
	name string

	breaks    bool
	continues bool
}

type LoopScope struct {
	stack []*LoopScopeItem
}

func (l *LoopScope) Break() string {
	if len(l.stack) == 0 {
		return ""
	}

	item := l.stack[len(l.stack)-1]
	item.breaks = true

	return item.name
}

func (l *LoopScope) Continue() string {
	if len(l.stack) == 0 {
		return ""
	}

	item := l.stack[len(l.stack)-1]
	item.continues = true

	return item.name
}

func (l *LoopScope) Push(name string) *LoopScopeItem {
	item := &LoopScopeItem{
		name: name,
	}

	l.stack = append(l.stack, item)

	return item
}

func (l *LoopScope) Pop() {
	if len(l.stack) == 0 {
		return
	}

	l.stack = l.stack[:len(l.stack)-1]
}

type Scope struct {
	types     map[string]*Type
	functions map[string]*Function
	variables map[string]*Variable

	children []*Scope

	function *Function

	read bool
	i    int

	parent *Scope
}

func (s *Scope) resolveType(typeName string) *Type {
	current := s

	for current != nil {
		if t, ok := current.types[typeName]; ok {
			return t
		}

		current = current.parent
	}

	return nil
}

func NewScope(parent *Scope) *Scope {
	s := &Scope{
		parent:    parent,
		types:     map[string]*Type{},
		functions: map[string]*Function{},
		variables: map[string]*Variable{},
	}

	return s
}

func (s *Scope) registerFunction(f *Function) {
	if _, ok := s.functions[f.name]; ok {
		panic(fmt.Sprintf("function '%s' is already defined", f.name))
	}

	s.functions[f.name] = f
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
	args int

	sub bool

	scope *Scope
}

func (c *CealCompiler) Compile(src string) *CealProgram {
	input := antlr.NewInputStream(src)
	lexer := parser.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCParser(stream)

	includes := p.Program().AllInclude()

	for i := len(includes) - 1; i >= 0; i-- {
		include := includes[i]
		str := includes[i].STRING().GetText()
		name := str[1 : len(str)-1]

		// TODO: is a hardcoded avm.hpp okay?
		if name == "avm.hpp" {
			continue
		}

		ip := path.Join(c.Includes[0], name)

		bs, err := os.ReadFile(ip)
		if err != nil {
			panic(fmt.Sprintf("failed to read #include file: '%s'", name))
		}

		incsrc := string(bs)

		prefix := src[:include.GetStart().GetStart()]
		suffix := src[include.GetStop().GetStop()+1:]

		src = prefix + incsrc + suffix
	}

	input = antlr.NewInputStream(src)
	lexer = parser.NewCLexer(input)
	stream = antlr.NewCommonTokenStream(lexer, 0)
	p = parser.NewCParser(stream)

	global := NewScope(nil)

	global.types["void"] = &Type{
		name: "void",
		simple: &SimpleType{
			empty: true,
		},
	}

	global.types["bytes"] = &Type{
		name: "bytes",
		simple: &SimpleType{
			kind: SimpleTypeBytes,
		},
	}

	global.types["uint64"] = &Type{
		name: "uint64",
		simple: &SimpleType{
			kind: SimpleTypeInt,
		},
	}

	global.types["any"] = &Type{
		name:   "any",
		simple: &SimpleType{},
	}

	for _, item := range builtin_functions {
		f := &Function{
			t:       item.t,
			name:    item.name,
			returns: item.returns,
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

		global.registerFunction(f)
	}

	for _, item := range builtin_structs {
		s := &Struct{
			builtin:   &BuiltinStruct{},
			fields:    map[string]*StructField{},
			functions: map[string]*StructFunction{},
		}

		for _, item := range item.fields {
			s.fields[item.name] = &StructField{
				t:    item.t,
				name: item.name,
				fun:  item.fun,
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

			s.registerFunction(f)
		}

		t := &Type{
			name:    item.name,
			complex: s,
		}

		global.types[t.name] = t
	}

	for _, item := range builtin_variables {
		v := &Variable{
			t:    item.t,
			name: item.name,
		}

		global.variables[v.name] = v
	}

	program := &CealProgram{
		Functions: map[string]*CealFunction{},
	}

	sv := &SymbolTableVisitor{
		BaseCVisitor: &parser.BaseCVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
		program: program,
		global:  global,
	}

	stream.Seek(0)
	sv.Visit(p.Program())

	global.readonly()

	av := &AstVisitor{
		BaseCVisitor: &parser.BaseCVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
		global:  global,
		program: program,
		loops:   &LoopScope{},
	}

	stream.Seek(0)
	av.Visit(p.Program())

	return program
}
