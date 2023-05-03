package compiler

import (
	"ceal/parser"
	"ceal/teal"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/pkg/errors"
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

type cealPreprocessor struct {
	includes  []string
	processed map[string]bool
	defines   map[string]string

	stack *defStack
}

type defStack struct {
	parent *defStack
	skip   bool
}

func (d *defStack) push() *defStack {
	return &defStack{
		skip:   d.skip,
		parent: d,
	}
}

func (d *defStack) pop() *defStack {
	if d.parent == nil {
		panic("#endif without #if")
	}
	return d.parent
}

func (p *cealPreprocessor) preprocess(name string, src string) (string, error) {
	src = strings.ReplaceAll(strings.ReplaceAll(src, "\r\n", "\n"), "\r", "\n")

	p.stack = p.stack.push()
	defer func() {
		p.stack = p.stack.pop()
	}()

	lines := strings.Split(src, "\n")

	i := 0

	for i < len(lines) {
		clear := func() {
			lines[i] = ""
		}

		line := strings.TrimSpace(lines[i])

		parts := strings.SplitN(strings.TrimSpace(line), " ", 2)

		if len(parts) == 0 {
			clear()
			i += 1
			continue
		}

		switch parts[0] {
		case "#endif":
			p.stack = p.stack.pop()
			clear()
		case "#ifndef":
			if len(parts) < 2 {
				return "", fmt.Errorf("invalid #ifndef: '%s'", line)
			}

			p.stack = p.stack.push()

			name := strings.TrimSpace(parts[1])

			if _, ok := p.defines[name]; ok {
				p.stack.skip = true
			}

			clear()
		case "#ifdef":
			if len(parts) < 2 {
				return "", fmt.Errorf("invalid #ifdef: '%s'", line)
			}

			p.stack = p.stack.push()

			name := strings.TrimSpace(parts[1])

			if _, ok := p.defines[name]; !ok {
				p.stack.skip = true
			}

			clear()
		default:
			if p.stack.skip {
				clear()
			} else {
				switch parts[0] {
				case "#include":
					if len(parts) != 2 {
						return "", fmt.Errorf("invalid #include: '%s'", line)
					}

					quoted := strings.TrimSpace(parts[1])

					begin := string(quoted[0])
					var end string

					switch begin {
					case "\"":
						end = begin
					case "<":
						end = ">"
					default:
						return "", fmt.Errorf("unexpected #include quote in: '%s'", line)
					}

					if end != string(quoted[len(quoted)-1]) {
						return "", fmt.Errorf("#include quote mismatch: '%s'", line)
					}

					name := quoted[1 : len(quoted)-1]

					inc, ipath, err := func() (string, string, error) {
						for _, dir := range p.includes {
							ip := path.Join(dir, name)

							bs, err := os.ReadFile(ip)
							if os.IsNotExist(err) {
								continue
							}

							if err != nil {
								return "", "", errors.Wrapf(err, "failed to read #include file: '%s'", name)
							}

							inc := string(bs)
							return inc, ip, nil
						}

						return "", "", errors.Errorf("failed to find #include file: '%s'", name)
					}()

					if err != nil {
						return "", err
					}

					inc, err = p.preprocess(ipath, inc)
					if err != nil {
						return "", err
					}

					inclines := strings.Split(inc, "\n")

					// insert inclines at current line removing the current line
					lines = append(lines[:i], append(inclines, lines[i+1:]...)...)
				case "#define":
					if len(parts) != 2 {
						return "", fmt.Errorf("invalid #define: '%s'", line)
					}

					items := strings.SplitN(parts[1], " ", 2)

					switch len(items) {
					case 1:
						p.defines[strings.TrimSpace(items[0])] = ""
					case 2:
						p.defines[strings.TrimSpace(items[0])] = strings.TrimSpace(items[1])
					}

					clear()
				case "#undef":
					if len(parts) != 2 {
						return "", fmt.Errorf("invalid #undef: '%s'", line)
					}

					name := strings.TrimSpace(parts[1])

					delete(p.defines, name)

					clear()
				case "#pragma":
					if len(parts) < 2 {
						return "", fmt.Errorf("invalid #pragma: '%s'", line)
					}

					switch parts[1] {
					case "once":
						clear()
						if p.processed[name] {
							p.stack.skip = true
						} else {
							p.processed[name] = true
						}
					}
				}
			}
		}

		i++
	}

	src = strings.Join(lines, "\n")

	return src, nil
}

func (c *cealCompiler) Compile(src string) *CealProgram {
	p := &cealPreprocessor{
		includes:  c.includes,
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
