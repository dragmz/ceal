package compiler

import (
	"ceal/parser"
	"ceal/teal"
	"fmt"
	"os"
	"path"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type CealCompiler struct {
	Includes []string
}

type FunctionParam struct {
	t     string
	name  string
	array bool
}

type BuiltinFunction struct {
	op    string
	stack []*FunctionParam
	imm   []*FunctionParam
}

type CompilerFunction struct {
	parameters []*FunctionParam
	handler    func(args []CealAst) teal.TealAst
}

type Function struct {
	t    string
	name string

	returns int

	builtin  *BuiltinFunction
	user     *UserFunction
	compiler *CompilerFunction
}

type StructField struct {
	t    string
	name string
	fun  string
}

type BuiltinStruct struct {
}

type Struct struct {
	fields      map[string]*StructField
	fieldsNames []string

	functions map[string]*Function

	builtin *BuiltinStruct
}

func (s *Struct) registerFunction(f *Function) {
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

func NewScope(parent *Scope) *Scope {
	s := &Scope{
		parent:    parent,
		types:     map[string]*Type{},
		functions: map[string]*Function{},
		variables: map[string]*Variable{},
	}

	return s
}

func (s *Scope) resolveType(typeName string) *Type {
	current := s

	for current != nil {
		if t, ok := current.types[typeName]; ok {
			return t
		}

		current = current.parent
	}

	panic(fmt.Sprintf("unknown type: '%s'", typeName))
}

func (s *Scope) registerFunction(f *Function) {
	if _, ok := s.functions[f.name]; ok {
		panic(fmt.Sprintf("function '%s' is already defined", f.name))
	}

	s.functions[f.name] = f
}

func (s *Scope) registerType(t *Type) {
	if _, ok := s.types[t.name]; ok {
		panic(fmt.Sprintf("type '%s' is already defined", t.name))
	}

	s.types[t.name] = t
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

		var incsrc string

		// TODO: is a hardcoded avm.hpp okay?
		if name != "avm.hpp" {
			for _, dir := range c.Includes {
				ip := path.Join(dir, name)

				bs, err := os.ReadFile(ip)
				if os.IsNotExist(err) {
					continue
				}

				if err != nil {
					panic(fmt.Sprintf("failed to read #include file: '%s'", name))
				}

				incsrc = string(bs)

				break
			}
		}

		prefix := src[:include.GetStart().GetStart()]
		suffix := src[include.GetStop().GetStop()+1:]

		src = prefix + incsrc + suffix
	}

	input = antlr.NewInputStream(src)
	lexer = parser.NewCLexer(input)
	stream = antlr.NewCommonTokenStream(lexer, 0)
	p = parser.NewCParser(stream)

	global := NewScope(nil)

	global.registerType(&Type{
		name: "void",
		simple: &SimpleType{
			empty: true,
		},
	})

	global.registerType(&Type{
		name: "bytes",
		simple: &SimpleType{
			kind: SimpleTypeBytes,
		},
	})

	global.registerType(&Type{
		name: "bool",
		simple: &SimpleType{
			kind: SimpleTypeInt,
		},
	})

	global.registerType(&Type{
		name: "uint8",
		simple: &SimpleType{
			kind: SimpleTypeInt,
		},
	})

	global.registerType(&Type{
		name: "uint64",
		simple: &SimpleType{
			kind: SimpleTypeInt,
		},
	})

	global.registerType(&Type{
		name:   "any",
		simple: &SimpleType{},
	})

	{
		f := &Function{
			name:    "abi_encode",
			t:       "bytes",
			returns: 1,
			compiler: &CompilerFunction{
				handler: func(args []CealAst) teal.TealAst {
					v, ok := args[0].(*CealVariable)
					if !ok {
						panic("abi_encode data argument expects a variable")
					}

					vt := global.resolveType(v.D.V.t)

					if vt.simple != nil {
						// TODO: implement
						panic("abi_encode does not support simple types yet")
					}

					if vt.complex != nil {
						for _, name := range vt.complex.fieldsNames {
							field := vt.complex.fields[name]
							if field.fun != "" {
								panic(fmt.Sprintf("abi_encode does not support built-in fields: '%s'", field.name))
							}

							ft := global.resolveType(field.t)
							// TODO: implement
							panic(fmt.Sprintf("abi_encode does not support encoding type: '%s'", ft.name))
						}
					}

					panic(fmt.Sprintf("abi_encode does not support the type: '%s'", vt.name))
				},
				parameters: []*FunctionParam{
					{
						t:    "any",
						name: "in",
					},
				},
			},
		}

		global.registerFunction(f)
	}

	{
		f := &Function{
			name: "abi_decode",
			compiler: &CompilerFunction{
				handler: func(args []CealAst) teal.TealAst {
					v1, ok := args[0].(*CealVariable)
					if !ok {
						panic("abi_decode data argument expects a variable")
					}

					v1t := global.resolveType(v1.D.V.t)

					if v1t.simple == nil || v1t.simple.kind != SimpleTypeBytes {
						panic("abi_decode data argument must be of bytes type")
					}

					v2, ok := args[1].(*CealVariable)
					if !ok {
						panic("abi_decode out argument expects a variable")
					}

					v2t := global.resolveType(v2.D.V.t)

					if v2t.simple != nil {
						// TODO: implement
						panic("abi_decode does not support simple types yet")
					}

					if v2t.complex != nil {
						for _, name := range v2t.complex.fieldsNames {
							field := v2t.complex.fields[name]
							if field.fun != "" {
								panic(fmt.Sprintf("abi_decode does not support built-in fields: '%s'", field.name))
							}

							ft := global.resolveType(field.t)
							// TODO: implement
							panic(fmt.Sprintf("abi_decode does not support decoding type: '%s'", ft.name))
						}
					}

					panic(fmt.Sprintf("abi_decode does not support the type: '%s'", v2t.name))
				},
				parameters: []*FunctionParam{
					{
						t:    "bytes",
						name: "data",
					},
					{
						t:    "any&",
						name: "out",
					},
				},
			},
		}

		global.registerFunction(f)
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
				t:     item.t,
				name:  item.name,
				array: item.array,
			})
		}

		for _, item := range item.imm {
			f.builtin.imm = append(f.builtin.imm, &FunctionParam{
				t:     item.t,
				name:  item.name,
				array: item.array,
			})
		}

		global.registerFunction(f)
	}

	for _, item := range builtin_structs {
		s := &Struct{
			builtin:   &BuiltinStruct{},
			fields:    map[string]*StructField{},
			functions: map[string]*Function{},
		}

		for _, item := range item.fields {
			s.fields[item.name] = &StructField{
				t:    item.t,
				name: item.name,
				fun:  item.fun,
			}
		}

		for _, item := range item.functions {
			f := &Function{
				t:       item.t,
				name:    item.name,
				builtin: &BuiltinFunction{},
			}

			for _, item := range item.imms {
				f.builtin.imm = append(f.builtin.imm, &FunctionParam{
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

		global.registerType(t)
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
