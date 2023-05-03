package compiler

import (
	"ceal/parser"
	"ceal/teal"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type CealCompilerConfig struct {
	Includes []string
}

type cealCompiler struct {
	includes []string
}

func NewCompiler(config CealCompilerConfig) *cealCompiler {
	return &cealCompiler{
		includes: config.Includes,
	}
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

type AvmTypeKind int

const (
	AvmTypeInt = AvmTypeKind(iota)
	AvmTypeBytes
)

type AvmType struct {
	kind AvmTypeKind
}

type Type struct {
	name string

	complex *Struct // struct
	avm     *AvmType
}

var noneType = &Type{
	name: "none",
}

var anyType = &Type{
	name: "any",
}

var bytesType = &Type{
	name: "bytes",
	avm: &AvmType{
		kind: AvmTypeBytes,
	},
}

var avmTypeBytes = &AvmType{
	kind: AvmTypeBytes,
}

var avmTypeUint64 = &AvmType{
	kind: AvmTypeInt,
}

var uint64Type = &Type{
	name: "uint64",
	avm:  avmTypeUint64,
}

type LocalVariable struct {
	slot int
}

type ParameterVariable struct {
	index int
}

type ConstVariable struct {
	kind  AvmTypeKind
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
	validateType(t)

	if _, ok := s.types[t.name]; ok {
		panic(fmt.Sprintf("type '%s' is already defined", t.name))
	}

	s.types[t.name] = t
}

func validateType(t *Type) {
	if t.avm != nil && t.complex != nil {
		panic(fmt.Sprintf("type cannot be both simple and complex: '%s'", t.name))
	}
}

func (s *Scope) registerTypeAlias(name string, t *Type) {
	validateType(t)

	if _, ok := s.types[name]; ok {
		panic(fmt.Sprintf("type '%s' is already defined", name))
	}

	s.types[name] = t
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

func (c *cealCompiler) Compile(src string) *CealProgram {
	p := &cealPreprocessor{
		includes: &cealIncludes{
			paths: c.includes,
		},
		processed: map[string]bool{},
		defines:   map[string]string{},
		stack:     &defStack{},
	}

	p.defines["CEAL"] = "1"

	src, err := p.preprocess("", src)
	if err != nil {
		panic(err)
	}

	input := antlr.NewInputStream(src)
	lexer := parser.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	cp := parser.NewCParser(stream)

	global := NewScope(nil)

	global.registerType(uint64Type)
	global.registerType(bytesType)
	global.registerType(anyType)
	global.registerType(noneType)
	global.registerTypeAlias("void", noneType)
	global.registerType(&Type{
		name: "uint8",
		avm:  avmTypeUint64,
	})
	global.registerType(&Type{
		name: "int8",
		avm:  avmTypeUint64,
	})
	global.registerType(&Type{
		name: "bool",
		avm:  avmTypeUint64,
	})

	global.registerType(&Type{
		name: "label",
	})

	for _, item := range builtin_types {
		bt := &Type{
			name: item.Name,
		}

		switch item.Type {
		case "uint64":
			bt.avm = avmTypeUint64
		case "bytes":
			bt.avm = avmTypeBytes
		case "none":
		}

		global.registerType(bt)
	}

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
						t:    global.resolveType("method_t"),
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

					if vt.avm != nil {
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
						t:    anyType,
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

					if v1t.avm == nil || v1t.avm.kind != AvmTypeBytes {
						panic("abi_decode data argument must be of bytes type")
					}

					v2, ok := args[1].(*CealVariable)
					if !ok {
						panic("abi_decode out argument expects a variable")
					}

					v2t := v2.D.V.t

					if v2t.avm != nil {
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
						t:    anyType,
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
	sv.Visit(cp.Program())

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
	av.Visit(cp.Program())

	return program
}
