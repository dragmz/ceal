package compiler

import (
	"ceal/teal"
	"fmt"
	"strings"
)

type CealAffixable interface {
	AddPrefix([]CealAst)
	AddSuffix([]CealAst)
}

type CealAst interface {
	TealAst() teal.TealAst
}

type CealProgram struct {
	ConstInts  []int
	ConstBytes [][]byte

	Functions     map[string]*CealFunction
	FunctionNames []string
}

func (a *CealProgram) registerFunction(f *CealFunction) {
	if _, ok := a.Functions[f.Fun.name]; ok {
		panic(fmt.Sprintf("function '%s' is already defined", f.Fun.name))
	}

	a.Functions[f.Fun.name] = f
	a.FunctionNames = append(a.FunctionNames, f.Fun.name)

}

func (a *CealProgram) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	res.Write(&teal.Teal_pragma_version{Version: AvmVersion})

	if len(a.ConstInts) > 0 {
		items := []uint64{}
		for _, v := range a.ConstInts {
			items = append(items, uint64(v))
		}

		res.Write(&teal.Teal_intcblock{Teal_intcblock_op: teal.Teal_intcblock_op{UINT1: items}})
	}

	if len(a.ConstBytes) > 0 {
		res.Write(&teal.Teal_bytecblock{Teal_bytecblock_op: teal.Teal_bytecblock_op{BYTES1: a.ConstBytes}})
	}

	main := a.Functions[AvmMainName]

	if len(a.Functions) > 1 {
		res.Write(&teal.Teal_b{
			Teal_b_op: teal.Teal_b_op{
				TARGET1: main.Fun.name,
			},
		})
	}

	writeAffix := func(items []CealAst) {
		for _, item := range items {
			res.Write(item.TealAst())
		}
	}

	for _, name := range a.FunctionNames {
		ast := a.Functions[name]

		if name == AvmMainName {
			continue
		}

		writeAffix(ast.pre)
		res.Write(&teal.Teal_label{Name: ast.Fun.name})
		res.Write(ast.TealAst())
		writeAffix(ast.post)
	}

	writeAffix(main.pre)

	if len(a.Functions) > 1 {
		res.Write(&teal.Teal_label{Name: main.Fun.name})
	}

	res.Write(main.TealAst())
	writeAffix(main.post)

	return res.Build()
}

type CealContinue struct {
	Label string
	Index int
}

func (a *CealContinue) TealAst() teal.TealAst {
	return &teal.Teal_b{
		Teal_b_op: teal.Teal_b_op{
			TARGET1: fmt.Sprintf("%s_%d_continue", a.Label, a.Index)},
	}
}

type CealBreak struct {
	Label string
	Index int
}

func (a *CealBreak) TealAst() teal.TealAst {
	return &teal.Teal_b{
		Teal_b_op: teal.Teal_b_op{
			TARGET1: fmt.Sprintf("%s_%d_end", a.Label, a.Index),
		},
	}
}

type CealSwitchCase struct {
	Value      CealAst
	Statements []CealAst
}

type CealSwitch struct {
	Index int
	Loop  *LoopScopeItem

	Value CealAst
	Cases []*CealSwitchCase

	Default []CealAst
}

func (a *CealSwitch) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	labels := []string{}

	for i, c := range a.Cases {
		label := fmt.Sprintf("switch_%d_%d", a.Index, i)
		labels = append(labels, label)

		res.Write(c.Value.TealAst())
	}

	res.Write(a.Value.TealAst())
	res.Write(&teal.Teal_match{
		Teal_match_op: teal.Teal_match_op{
			TARGET1: labels,
		},
	})

	if len(a.Default) > 0 {
		for _, stmt := range a.Default {
			res.Write(stmt.TealAst())
		}
	}

	for i, c := range a.Cases {
		label := labels[i]

		res.Write(&teal.Teal_label{Name: label})

		for _, stmt := range c.Statements {
			res.Write(stmt.TealAst())
		}
	}

	if a.Loop.breaks {
		res.Write(&teal.Teal_label{Name: fmt.Sprintf("switch_%d_end", a.Index)})
	}

	return res.Build()
}

type CealDoWhile struct {
	Index     int
	Loop      *LoopScopeItem
	Condition CealAst
	Statement CealAst
}

func (a *CealDoWhile) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}
	res.Write(&teal.Teal_label{Name: fmt.Sprintf("do_%d", a.Index)})
	res.Write(a.Statement.TealAst())

	if a.Loop.continues {
		res.Write(&teal.Teal_label{Name: fmt.Sprintf("do_%d_continue", a.Index)})
	}

	res.Write(&teal.Teal_bnz{
		Teal_bnz_op: teal.Teal_bnz_op{
			TARGET1: fmt.Sprintf("do_%d", a.Index),
		},
		STACK_1: a.Condition.TealAst(),
	})

	if a.Loop.breaks {
		res.Write(&teal.Teal_label{Name: fmt.Sprintf("do_%d_end", a.Index)})
	}

	return res.Build()
}

type CealWhile struct {
	Index     int
	Loop      *LoopScopeItem
	Condition CealAst
	Statement CealAst
}

func (a *CealWhile) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	res.Write(&teal.Teal_label{Name: fmt.Sprintf("while_%d", a.Index)})
	res.Write(&teal.Teal_bz{
		Teal_bz_op: teal.Teal_bz_op{
			TARGET1: fmt.Sprintf("while_%d_end", a.Index),
		},
		STACK_1: a.Condition.TealAst(),
	})

	res.Write(a.Statement.TealAst())

	if a.Loop.continues {
		res.Write(&teal.Teal_label{Name: fmt.Sprintf("while_%d_continue", a.Index)})
	}

	res.Write(&teal.Teal_b{
		Teal_b_op: teal.Teal_b_op{
			TARGET1: fmt.Sprintf("while_%d", a.Index),
		},
	})

	res.Write(&teal.Teal_label{Name: fmt.Sprintf("while_%d_end", a.Index)})

	return res.Build()
}

type CealFor struct {
	Index     int
	Loop      *LoopScopeItem
	Init      []CealAst
	Condition CealAst
	Statement CealAst
	Iter      []CealAst
}

func (a *CealFor) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	for _, stmt := range a.Init {
		res.Write(stmt.TealAst())
	}

	res.Write(&teal.Teal_label{Name: fmt.Sprintf("for_%d", a.Index)})
	res.Write(&teal.Teal_bz{
		Teal_bz_op: teal.Teal_bz_op{
			TARGET1: fmt.Sprintf("for_%d_end", a.Index),
		},
		STACK_1: a.Condition.TealAst(),
	})
	res.Write(a.Statement.TealAst())

	if a.Loop.continues {
		res.Write(&teal.Teal_label{Name: fmt.Sprintf("for_%d_continue", a.Index)})
	}

	for _, stmt := range a.Iter {
		res.Write(stmt.TealAst())
	}

	res.Write(&teal.Teal_b{
		Teal_b_op: teal.Teal_b_op{
			TARGET1: fmt.Sprintf("for_%d", a.Index),
		},
	})
	res.Write(&teal.Teal_label{Name: fmt.Sprintf("for_%d_end", a.Index)})

	return res.Build()
}

type CealExpr interface {
	// ToStmt converts the expression to a statement so it does not push a value onto the stack
	ToStmt()
}

type CealValue interface {
	ToValue()
}

type CealPrefix struct {
	D valueData

	Op string

	IsStmt bool
}

func (a *CealPrefix) ToStmt() {
	a.IsStmt = true
}

func (a *CealPrefix) TealAst() teal.TealAst {
	if a.D.V.constant {
		panic("cannot modify const var")
	}

	var op teal.TealAst

	switch a.Op {
	case "++":
		op = &teal.Teal_plus{
			STACK_1: a.D.Load(),
			STACK_2: &teal.Teal_int{V: 1},
		}
	case "--":
		op = &teal.Teal_minus{
			STACK_1: a.D.Load(),
			STACK_2: &teal.Teal_int{V: 1},
		}
	default:
		panic(fmt.Sprintf("prefix operator not supported: '%s'", a.Op))
	}

	res := &teal.TealAstBuilder{}

	res.Write(a.D.Store(op))

	if !a.IsStmt {
		res.Write(a.D.Load())
	}

	return res.Build()
}

type CealPostfix struct {
	D valueData

	Op string

	IsStmt bool
}

func (a *CealPostfix) ToStmt() {
	a.IsStmt = true
}

func (a *CealPostfix) TealAst() teal.TealAst {
	if a.D.V.constant {
		panic("cannot modify const var")
	}

	var s1 teal.TealAst
	s1 = a.D.Load()

	if !a.IsStmt {
		s1 = &teal.Teal_dup{STACK_1: s1}
	}

	s2 := &teal.Teal_int{V: 1}

	var op teal.TealAst
	switch a.Op {
	case "++":
		op = &teal.Teal_plus{STACK_1: s1, STACK_2: s2}
	case "--":
		op = &teal.Teal_minus{STACK_1: s1, STACK_2: s2}
	default:
		panic(fmt.Sprintf("postfix operator not supported: '%s'", a.Op))
	}

	return a.D.Store(op)
}

type CealLabel struct {
	Name string
}

func (a *CealLabel) TealAst() teal.TealAst {
	return &teal.Teal_label{Name: fmt.Sprintf("label_%s", a.Name)}
}

type CealGoto struct {
	Label string
}

func (a *CealGoto) TealAst() teal.TealAst {
	return &teal.Teal_b{
		Teal_b_op: teal.Teal_b_op{
			TARGET1: fmt.Sprintf("label_%s", a.Label),
		},
	}
}

type CealVariable struct {
	D valueData
}

func (a *CealVariable) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	if a.D.V.local != nil {
		ast := a.D.Load()
		res.Write(ast)

		return res.Build()
	}

	if a.D.V.param != nil {
		ast := a.D.Load()
		res.Write(ast)
		return res.Build()
	}

	if a.D.V.const_ != nil {
		switch a.D.V.const_.kind {
		case SimpleTypeInt:
			ast := &teal.Teal_intc{Teal_intc_op: teal.Teal_intc_op{I1: uint8(a.D.V.const_.index)}}
			res.Write(ast)
			return res.Build()
		case SimpleTypeBytes:
			ast := &teal.Teal_bytec{Teal_bytec_op: teal.Teal_bytec_op{I1: uint8(a.D.V.const_.index)}}
			res.Write(ast)
			return res.Build()
		}
	}

	switch a.D.V.t {
	case "uint64":
		res.Write(&teal.Teal_named_int{V: &teal.Teal_named_int_value{V: a.D.V.name}})
		return res.Build()
	case "bytes":
		res.Write(&teal.Teal_byte{S: &teal.Teal_byte_string_value{V: a.D.V.name}})
	default:
		panic(fmt.Sprintf("type '%s' is not supported", a.D.V.t))
	}

	return res.Build()
}

type CealUnaryOp struct {
	Op        string
	Statement CealAst
}

func (a *CealUnaryOp) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}
	switch a.Op {
	case "!":
		res.Write(&teal.Teal_not{STACK_1: a.Statement.TealAst()})
	default:
		panic(fmt.Sprintf("unary op '%s' not supported", a.Op))
	}
	return res.Build()
}

type CealAssignSumDiff struct {
	D     valueData
	Value CealAst
	Op    string

	IsStmt bool
}

func (a *CealAssignSumDiff) TealAst() teal.TealAst {
	s1 := a.D.Load()

	var op teal.TealAst

	switch a.Op {
	case "+=":
		op = &teal.Teal_plus{STACK_1: s1, STACK_2: a.Value.TealAst()}
	case "-=":
		op = &teal.Teal_minus{STACK_1: s1, STACK_2: a.Value.TealAst()}
	}

	if !a.IsStmt {
		op = &teal.Teal_dup{STACK_1: op}
	}

	ast := a.D.Store(op)

	return ast
}

type CealAnd struct {
	Index int

	Left  CealAst
	Right CealAst
}

func (a *CealAnd) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	res.Write(&teal.Teal_andand{
		STACK_1: &teal.Teal_bz{
			Teal_bz_op: teal.Teal_bz_op{
				TARGET1: fmt.Sprintf("and_%d_end", a.Index),
			},
			STACK_1: &teal.Teal_dup{STACK_1: a.Left.TealAst()},
		},
		STACK_2: a.Right.TealAst(),
	})

	res.Write(&teal.Teal_label{Name: fmt.Sprintf("and_%d_end", a.Index)})

	return res.Build()
}

type CealOr struct {
	Index int
	Left  CealAst
	Right CealAst
}

func (a *CealOr) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	res.Write(
		&teal.Teal_oror{
			STACK_1: &teal.Teal_bnz{
				Teal_bnz_op: teal.Teal_bnz_op{
					TARGET1: fmt.Sprintf("or_%d_end", a.Index),
				},
				STACK_1: &teal.Teal_dup{
					STACK_1: a.Left.TealAst(),
				},
			},
			STACK_2: a.Right.TealAst(),
		})

	res.Write(&teal.Teal_label{Name: fmt.Sprintf("or_%d_end", a.Index)})

	return res.Build()
}

type CealBinop struct {
	Left  CealAst
	Op    string
	Right CealAst
}

func (a *CealBinop) TealAst() teal.TealAst {
	var op teal.TealAst

	switch a.Op {
	case "+":
		op = &teal.Teal_plus{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "-":
		op = &teal.Teal_minus{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "*":
		op = &teal.Teal_mul{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "/":
		op = &teal.Teal_div{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "==":
		op = &teal.Teal_eqeq{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "!=":
		op = &teal.Teal_noteq{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "&":
		op = &teal.Teal_and{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "^":
		op = &teal.Teal_xor{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "|":
		op = &teal.Teal_or{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "<":
		op = &teal.Teal_lt{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case ">":
		op = &teal.Teal_gt{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "<=":
		op = &teal.Teal_lteq{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case ">=":
		op = &teal.Teal_gteq{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "%":
		op = &teal.Teal_mod{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	default:
		panic(fmt.Sprintf("binary op '%s' is not supported yet", a.Op))
	}

	return op
}

type CealNegate struct {
	Value CealAst
}

func (a *CealNegate) TealAst() teal.TealAst {
	return &teal.Teal_minus{
		STACK_1: &teal.Teal_int{V: 0},
		STACK_2: a.Value.TealAst(),
	}
}

type CealDefine struct {
	D valueData

	Value CealAst
}

func (a *CealDefine) TealAst() teal.TealAst {
	if a.D.T.complex != nil {
		panic("defining complex variable is not supported yet")
	}

	ast := a.D.Store(a.Value.TealAst())

	return ast
}

type CealAssign struct {
	D valueData

	Value CealAst

	IsStmt bool
}

func (a *CealAssign) ToStmt() {
	a.IsStmt = true
}

func (a *CealAssign) TealAst() teal.TealAst {
	if a.D.V.constant {
		panic("cannot assign to a const var")
	}

	if a.D.V.param != nil {
		// TODO: add param var assignment support
		panic("cannot assign param var")
	}

	res := &teal.TealAstBuilder{}

	if a.D.T.complex != nil && a.D.T.complex.builtin != nil {
		panic("cannot assign to built-in op")
	}

	ast := a.D.Store(a.Value.TealAst())

	res.Write(ast)

	if !a.IsStmt {
		load := a.D.Load()
		res.Write(load)
	}

	return res.Build()
}

type CealStructField struct {
	D valueData
}

func (a *CealStructField) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	if a.D.T.complex.builtin != nil {
		res.Write(&teal.Teal_call_builtin{
			Name: a.D.Fun.builtin.op,
			Imms: []teal.TealAst{&teal.Teal_named_int_value{
				V: a.D.F.name,
			}},
		})
		return res.Build()
	}

	if a.D.V.param != nil {
		panic("accessing struct param fields is not supported yet")
	}

	ast := a.D.Load()
	res.Write(ast)

	return res.Build()
}

type CealCall struct {
	Fun   *Function
	Field CealAst // TODO: not sure the Field should be here or filled in as an Arg already

	Args []CealAst

	IsStmt bool
}

func (a *CealCall) ToStmt() {
	a.IsStmt = true
}

func (a *CealCall) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	var args []teal.TealAst

	if a.Fun.compiler != nil {
		ast := a.Fun.compiler.handler(a.Args)
		res.Write(ast)
	}

	if a.Fun.builtin != nil {
		sl := len(a.Fun.builtin.stack)

		for i := 0; i < sl; i++ {
			arg := a.Args[i]
			args = append(args, arg.TealAst())
		}

		var imms []teal.TealAst

		imm_index := 0
		arg_index := 0

		l := len(a.Args) - sl
		if a.Field != nil {
			l++
		}

		for i := 0; i < l; i++ {
			var arg CealAst

			imm := a.Fun.builtin.imm[imm_index]
			if imm.field {
				arg = a.Field
			} else {
				arg = a.Args[arg_index+len(a.Fun.builtin.stack)]
				arg_index++
			}

			if e, ok := arg.(CealValue); ok {
				e.ToValue()
			}

			if e, ok := arg.(*CealStringConstant); ok {
				switch imm.t {
				case "label":
					e.kind = CealStringConstantLabel
				}
			}

			ast := arg.TealAst()
			imms = append(imms, ast)

			if !imm.array {
				imm_index++
			}
		}

		res.Write(&teal.Teal_call_builtin{
			Args: args,
			Imms: imms,
			Name: a.Fun.builtin.op,
		})
	}

	if a.Fun.user != nil {
		for _, arg := range a.Args {
			args = append(args, arg.TealAst())
		}

		seq := teal.Teal_seq{}
		seq = append(seq, args...)
		seq = append(seq, &teal.Teal_callsub{
			Teal_callsub_op: teal.Teal_callsub_op{
				TARGET1: a.Fun.name,
			},
		})

		res.Write(seq)
	}

	if a.IsStmt {
		if a.Fun.returns > 0 {
			res.Write(&teal.Teal_popn{Teal_popn_op: teal.Teal_popn_op{N1: uint8(a.Fun.returns)}})
		}
	}

	return res.Build()
}

type CealIsReturn interface {
	IsReturn()
}

type CealIsBreak interface {
	IsBreak()
}

type CealIntConstant struct {
	Value string
	value bool
}

func (a *CealIntConstant) ToValue() {
	a.value = true
}

func (a *CealIntConstant) TealAst() teal.TealAst {
	var v teal.TealAst = &teal.Teal_named_int_value{V: a.Value}
	if !a.value {
		v = &teal.Teal_named_int{V: v}
	}
	return v
}

type CealStringConstantKind uint8

const (
	CealStringConstantLiteral = iota
	CealStringConstantLabel
)

type CealStringConstant struct {
	Value string
	value bool
	kind  CealStringConstantKind
}

func (a *CealStringConstant) ToValue() {
	a.value = true
}

func (a *CealStringConstant) TealAst() teal.TealAst {
	var v teal.TealAst

	switch a.kind {
	case CealStringConstantLabel:
		v = &teal.Teal_literal{V: a.Value}
	case CealStringConstantLiteral:
		v = &teal.Teal_byte_string_value{V: a.Value}
	}

	if !a.value {
		v = &teal.Teal_byte{S: v}
	}

	return v
}

type CealReturn struct {
	Value CealAst
	Fun   *Function
}

func (a *CealReturn) IsReturn() {
}

func (a *CealReturn) TealAst() teal.TealAst {
	var op teal.TealAst

	if a.Fun != nil && a.Fun.user.sub {
		var values []teal.TealAst

		if a.Value != nil {
			values = append(values, a.Value.TealAst())
		}

		seq := teal.Teal_seq{}

		seq = append(seq, values...)
		seq = append(seq, &teal.Teal_retsub{})

		op = seq
	} else {
		op = &teal.Teal_return{
			Teal_return_op: teal.Teal_return_op{},
			STACK_1:        a.Value.TealAst(),
		}
	}

	return op
}

type CealBlock struct {
	Statements []CealAst
}

func (a *CealBlock) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	for _, stmt := range a.Statements {
		res.Write(stmt.TealAst())
	}

	return res.Build()
}

type CealConditional struct {
	Index int

	Condition CealAst

	True  CealAst
	False CealAst
}

func (a *CealConditional) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	false_label := fmt.Sprintf("conditional_%d_false", a.Index)
	end_label := fmt.Sprintf("conditional_%d_end", a.Index)

	res.Write(&teal.Teal_bz{
		STACK_1: a.Condition.TealAst(),
		Teal_bz_op: teal.Teal_bz_op{
			TARGET1: false_label,
		},
	})
	res.Write(a.True.TealAst())
	res.Write(&teal.Teal_b{Teal_b_op: teal.Teal_b_op{TARGET1: end_label}})
	res.Write(&teal.Teal_label{Name: false_label})
	res.Write(a.False.TealAst())
	res.Write(&teal.Teal_label{Name: end_label})

	return res.Build()
}

type CealIfAlternative struct {
	Condition  CealAst
	Statements []CealAst
}

type CealIf struct {
	Index        int
	Alternatives []*CealIfAlternative
}

func (a *CealIf) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	end_label := fmt.Sprintf("if_end_%d", a.Index)

	for i, alt := range a.Alternatives {
		if alt.Condition != nil {
			if i < len(a.Alternatives)-1 {
				res.Write(&teal.Teal_bz{
					Teal_bz_op: teal.Teal_bz_op{
						TARGET1: fmt.Sprintf("if_skip_%d_%d", a.Index, i),
					},
					STACK_1: alt.Condition.TealAst(),
				})
			} else {
				res.Write(&teal.Teal_bz{
					Teal_bz_op: teal.Teal_bz_op{
						TARGET1: end_label,
					},
					STACK_1: alt.Condition.TealAst(),
				})
			}
		}

		for _, stmt := range alt.Statements {
			res.Write(stmt.TealAst())
		}

		if i < len(a.Alternatives)-1 {
			res.Write(&teal.Teal_b{
				Teal_b_op: teal.Teal_b_op{
					TARGET1: end_label,
				},
			})
		}

		if alt.Condition != nil {
			if i < len(a.Alternatives)-1 {
				res.Write(&teal.Teal_label{Name: fmt.Sprintf("if_skip_%d_%d", a.Index, i)})
			}
		}
	}

	res.Write(&teal.Teal_label{Name: end_label})

	return res.Build()
}

type cealAffixableImpl struct {
	pre  []CealAst
	post []CealAst
}

func (a *cealAffixableImpl) AddPrefix(items []CealAst) {
	a.pre = append(a.pre, items...)
}

func (a *cealAffixableImpl) AddSuffix(items []CealAst) {
	a.post = append(a.post, items...)
}

type CealFunction struct {
	cealAffixableImpl

	Fun        *Function
	Statements []CealAst
}

func (a *CealFunction) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	if a.Fun.user.sub {
		if a.Fun.user.args != 0 || a.Fun.returns != 0 {
			ast := &teal.Teal_proto{
				Teal_proto_op: teal.Teal_proto_op{
					A1: uint8(a.Fun.user.args),
					R2: uint8(a.Fun.returns),
				},
			}

			res.Write(ast)
		}
	}

	for _, stmt := range a.Statements {
		res.Write(stmt.TealAst())
	}

	if a.Fun.user.sub {
		if len(a.Statements) > 0 {
			last := a.Statements[len(a.Statements)-1]
			if _, ok := last.(CealIsReturn); !ok {
				res.Write(&teal.Teal_retsub{})
			}
		}
	}

	return res.Build()
}

type CealRaw struct {
	Value string
}

func (a *CealRaw) String() string {
	return a.Value
}

func (a *CealRaw) Teal() teal.Teal {
	return teal.Teal{a}
}

func (a *CealRaw) TealAst() teal.TealAst {
	return a
}

type CealSubscript struct {
	D valueData
}

func (a *CealSubscript) TealAst() teal.TealAst {
	var value teal.TealAst

	if a.D.V.param != nil {
		value = a.D.Load()
	} else if a.D.V.local != nil {
		value = a.D.Load()
	} else {
		panic("unsupported variable type")
	}

	return &teal.Teal_extract3{
		Teal_extract3_op: teal.Teal_extract3_op{},
		STACK_1:          value,
		STACK_2:          a.D.Index.TealAst(),
		STACK_3:          &teal.Teal_int{V: 1},
	}
}

type CealAsm struct {
	Value string
}

func (a *CealAsm) String() string {
	return a.Value
}

func (a *CealAsm) TealAst() teal.TealAst {
	return &teal.Teal_raw{V: a.Value}
}

type CealSingleLineComment struct {
	Line string
}

func (a *CealSingleLineComment) TealAst() teal.TealAst {
	return &teal.Teal_comment{Lines: []string{a.Line}}
}

type CealMultiLineComment struct {
	Text string
}

func (a *CealMultiLineComment) TealAst() teal.TealAst {
	lines := strings.Split(a.Text, "\n")
	return &teal.Teal_comment{Lines: lines}
}
