package compiler

import (
	"ceal/parser"
	"fmt"
	"os"
	"path"
	"strings"

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

type SimpleType struct {
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

type Variable struct {
	constant bool

	t    string
	name string

	readonly bool

	local *LocalVariable
	param *ParameterVariable

	fields map[string]*Variable
}

type LoopScope struct {
	labels []string
}

func (l *LoopScope) Get() string {
	if len(l.labels) == 0 {
		return ""
	}

	return l.labels[len(l.labels)-1]
}

func (l *LoopScope) Push(label string) {
	l.labels = append(l.labels, label)
}

func (l *LoopScope) Pop() {
	if len(l.labels) == 0 {
		return
	}

	l.labels = l.labels[:len(l.labels)-1]
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

type AstProgram struct {
	functions      map[string]*AstFunction
	functionsNames []string
}

func (a *AstProgram) String() string {
	res := strings.Builder{}

	res.WriteString(fmt.Sprintf("#pragma version %d\n", AvmVersion))

	main := a.functions[AvmMainName]

	if len(a.functions) > 1 {
		res.WriteString(fmt.Sprintf("b %s\n", main.fun.name))
	}

	for _, name := range a.functionsNames {
		ast := a.functions[name]

		if name == AvmMainName {
			continue
		}

		res.WriteString(fmt.Sprintf("%s:\n", ast.fun.name))

		res.WriteString(ast.String())
		res.WriteString("\n")
	}

	if len(a.functions) > 1 {
		res.WriteString(fmt.Sprintf("%s:\n", main.fun.name))
	}

	res.WriteString(main.String())

	return res.String()
}

func (c *CealCompiler) Compile(src string) string {
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
		name:   "bytes",
		simple: &SimpleType{},
	}

	global.types["uint64"] = &Type{
		name:   "uint64",
		simple: &SimpleType{},
	}

	global.types["any"] = &Type{
		name:   "any",
		simple: &SimpleType{},
	}

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

			s.functions[f.name] = f
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

	return program.String()
}
