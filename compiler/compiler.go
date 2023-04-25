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
	t     *Type
	name  string
	array bool
	field bool
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

type StructFunction struct {
	t    string
	name string
	f    *Function
}

type Function struct {
	t    string
	name string

	returns []*FunctionParam

	builtin  *BuiltinFunction
	user     *UserFunction
	compiler *CompilerFunction
}

type StructField struct {
	t    *Type
	name string
	fun  *Function
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

	t    *Type
	name string

	local  *LocalVariable
	param  *ParameterVariable
	const_ *ConstVariable
	fun    *Function

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
	if t.simple != nil && t.complex != nil {
		panic(fmt.Sprintf("type cannot be both simple and complex: '%s'", t.name))
	}

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
	args []*FunctionParam

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

	// TODO: don't really have an idea how to handle the runtime types below
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
		name: "key_t",
		simple: &SimpleType{
			kind: SimpleTypeBytes,
		},
	})

	global.registerType(&Type{
		name: "label",
	})

	global.registerType(&Type{
		name: "any_t",
	})

	global.registerType(&Type{
		name: "bigint_t",
		simple: &SimpleType{
			kind: SimpleTypeInt,
		},
	})

	global.registerType(&Type{
		name: "name_t",
		simple: &SimpleType{
			kind: SimpleTypeBytes,
		},
	})

	global.registerType(&Type{
		name: "addr_t",
		simple: &SimpleType{
			kind: SimpleTypeBytes,
		},
	})

	global.registerType(&Type{
		name: "int8",
		simple: &SimpleType{
			kind: SimpleTypeInt,
		},
	})

	global.registerType(&Type{
		name: "uint64_t",
		simple: &SimpleType{
			kind: SimpleTypeInt,
		},
	})

	global.registerType(&Type{
		name: "bool_t",
		simple: &SimpleType{
			kind: SimpleTypeInt,
		},
	})

	global.registerType(&Type{
		name: "r_byte_t",
		simple: &SimpleType{
			kind: SimpleTypeBytes,
		},
	})

	global.registerType(&Type{
		name: "r32_byte_t",
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

	global.registerType(&Type{
		name:   "method",
		simple: &SimpleType{},
	})

	{
		f := &Function{
			name: "const_string",
			t:    "bytes",
			returns: []*FunctionParam{
				{
					t:    global.resolveType("bytes"),
					name: "r1",
				},
			},
			compiler: &CompilerFunction{
				parameters: []*FunctionParam{
					{
						t:    global.resolveType("bytes"),
						name: "value",
					},
				},
				handler: func(args []CealAst) teal.TealAst {
					return &teal.Teal_literal{V: args[0].TealAst().Teal().String()}
				},
			},
		}

		global.registerFunction(f)
	}

	{
		f := &Function{
			name: "avm_method",
			t:    "bytes",
			returns: []*FunctionParam{
				{
					t:    global.resolveType("bytes"),
					name: "r1",
				},
			},
			compiler: &CompilerFunction{
				parameters: []*FunctionParam{
					{
						t:    global.resolveType("method"),
						name: "name",
					},
				},
				handler: func(args []CealAst) teal.TealAst {
					return &teal.Teal_method{
						Name: args[0].TealAst().Teal().String(),
					}
				},
			},
		}

		global.registerFunction(f)
	}

	{
		f := &Function{
			name: "abi_encode",
			t:    "bytes",
			returns: []*FunctionParam{
				{
					t:    global.resolveType("bytes"),
					name: "r1",
				},
			},
			compiler: &CompilerFunction{
				handler: func(args []CealAst) teal.TealAst {
					v, ok := args[0].(*CealVariable)
					if !ok {
						panic("abi_encode data argument expects a variable")
					}

					vt := v.D.V.t

					if vt.simple != nil {
						// TODO: implement
						panic("abi_encode does not support simple types yet")
					}

					if vt.complex != nil {
						for _, name := range vt.complex.fieldsNames {
							field := vt.complex.fields[name]
							if field.fun != nil {
								panic(fmt.Sprintf("abi_encode does not support built-in fields: '%s'", field.name))
							}

							ft := field.t
							// TODO: implement
							panic(fmt.Sprintf("abi_encode does not support encoding type: '%s'", ft.name))
						}
					}

					panic(fmt.Sprintf("abi_encode does not support the type: '%s'", vt.name))
				},
				parameters: []*FunctionParam{
					{
						t:    global.resolveType("any"),
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

					v1t := v1.D.V.t

					if v1t.simple == nil || v1t.simple.kind != SimpleTypeBytes {
						panic("abi_decode data argument must be of bytes type")
					}

					v2, ok := args[1].(*CealVariable)
					if !ok {
						panic("abi_decode out argument expects a variable")
					}

					v2t := v2.D.V.t

					if v2t.simple != nil {
						// TODO: implement
						panic("abi_decode does not support simple types yet")
					}

					if v2t.complex != nil {
						for _, name := range v2t.complex.fieldsNames {
							field := v2t.complex.fields[name]
							if field.fun != nil {
								panic(fmt.Sprintf("abi_decode does not support built-in fields: '%s'", field.name))
							}

							ft := field.t
							// TODO: implement
							panic(fmt.Sprintf("abi_decode does not support decoding type: '%s'", ft.name))
						}
					}

					panic(fmt.Sprintf("abi_decode does not support the type: '%s'", v2t.name))
				},
				parameters: []*FunctionParam{
					{
						t:    global.resolveType("bytes"),
						name: "data",
					},
					{
						t:    global.resolveType("any"),
						name: "out",
					},
				},
			},
		}

		global.registerFunction(f)
	}

	for _, item := range builtin_functions {
		f := &Function{
			t:    item.t,
			name: item.name,
			builtin: &BuiltinFunction{
				op: item.op,
			},
		}

		for _, item := range item.returns {
			f.returns = append(f.returns, &FunctionParam{
				t:    global.resolveType(item.t),
				name: item.name,
			})
		}

		for _, item := range item.stack {
			f.builtin.stack = append(f.builtin.stack, &FunctionParam{
				t:     global.resolveType(item.t),
				name:  item.name,
				array: item.array,
				field: item.field,
			})
		}

		for _, item := range item.imm {
			f.builtin.imm = append(f.builtin.imm, &FunctionParam{
				t:     global.resolveType(item.t),
				name:  item.name,
				array: item.array,
				field: item.field,
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
			fun := global.functions[item.fun]

			t := global.resolveType(item.t)
			s.fields[item.name] = &StructField{
				t:    t,
				name: item.name,
				fun:  fun,
			}
		}

		for _, item := range item.functions {
			bf := global.functions[item.fun]

			f := &StructFunction{
				t:    item.t,
				name: item.name,
				f:    bf,
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
		t := global.resolveType(item.t)
		v := &Variable{
			t:    t,
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
