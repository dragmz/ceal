package compiler

import (
	"fmt"
	"strings"
)

type CealStatement interface {
	TealAst() TealAst
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

func (a *CealProgram) TealAst() TealAst {
	res := &TealAstBuilder{}

	res.Write(&Teal_pragma_version{Version: AvmVersion})

	if len(a.ConstInts) > 0 {
		items := []uint64{}
		for _, v := range a.ConstInts {
			items = append(items, uint64(v))
		}

		res.Write(&Teal_intcblock_fixed{UINT1: items})
	}

	if len(a.ConstBytes) > 0 {
		res.Write(&Teal_bytecblock_fixed{BYTES1: a.ConstBytes})
	}

	main := a.Functions[AvmMainName]

	if len(a.Functions) > 1 {
		res.Write(&Teal_b_fixed{TARGET1: main.Fun.name})
	}

	for _, name := range a.FunctionNames {
		ast := a.Functions[name]

		if name == AvmMainName {
			continue
		}

		res.Write(&Teal_label{Name: ast.Fun.name})
		res.Write(ast.TealAst())
	}

	if len(a.Functions) > 1 {
		res.Write(&Teal_label{Name: main.Fun.name})
	}

	res.Write(main.TealAst())

	return res.Build()
}

type CealContinue struct {
	Label string
	Index int
}

func (a *CealContinue) TealAst() TealAst {
	res := &TealAstBuilder{}
	res.Write(&Teal_b_fixed{TARGET1: fmt.Sprintf("%s_%d_continue", a.Label, a.Index)})
	return res.Build()
}

type CealBreak struct {
	Label string
	Index int
}

func (a *CealBreak) TealAst() TealAst {
	res := &TealAstBuilder{}
	res.Write(&Teal_b_fixed{TARGET1: fmt.Sprintf("%s_%d_end", a.Label, a.Index)})
	return res.Build()
}

type CealSwitchCase struct {
	Value      CealStatement
	Statements []CealStatement
}

type CealSwitch struct {
	Index int
	Loop  *LoopScopeItem

	Value CealStatement
	Cases []*CealSwitchCase

	Default []CealStatement
}

func (a *CealSwitch) TealAst() TealAst {
	res := &TealAstBuilder{}

	labels := []string{}

	for i, c := range a.Cases {
		label := fmt.Sprintf("switch_%d_%d", a.Index, i)
		labels = append(labels, label)

		res.Write(c.Value.TealAst())
	}

	res.Write(a.Value.TealAst())
	res.Write(&Teal_match_fixed{TARGET1: labels})

	if len(a.Default) > 0 {
		for _, stmt := range a.Default {
			res.Write(stmt.TealAst())
		}
	}

	for i, c := range a.Cases {
		label := labels[i]

		res.Write(&Teal_label{Name: label})

		for _, stmt := range c.Statements {
			res.Write(stmt.TealAst())
		}
	}

	if a.Loop.breaks {
		res.Write(&Teal_label{Name: fmt.Sprintf("switch_%d_end", a.Index)})
	}

	return res.Build()
}

type CealDoWhile struct {
	Index     int
	Loop      *LoopScopeItem
	Condition CealStatement
	Statement CealStatement
}

func (a *CealDoWhile) TealAst() TealAst {
	res := &TealAstBuilder{}
	res.Write(&Teal_label{Name: fmt.Sprintf("do_%d", a.Index)})
	res.Write(a.Statement.TealAst())

	if a.Loop.continues {
		res.Write(&Teal_label{Name: fmt.Sprintf("do_%d_continue", a.Index)})
	}

	res.Write(&Teal_bnz_fixed{
		s1:      a.Condition.TealAst(),
		TARGET1: fmt.Sprintf("do_%d", a.Index),
	})

	if a.Loop.breaks {
		res.Write(&Teal_label{Name: fmt.Sprintf("do_%d_end", a.Index)})
	}

	return res.Build()
}

type CealWhile struct {
	Index     int
	Loop      *LoopScopeItem
	Condition CealStatement
	Statement CealStatement
}

func (a *CealWhile) TealAst() TealAst {
	res := &TealAstBuilder{}

	res.Write(&Teal_label{Name: fmt.Sprintf("while_%d", a.Index)})
	res.Write(&Teal_bz_fixed{
		s1:      a.Condition.TealAst(),
		TARGET1: fmt.Sprintf("while_%d_end", a.Index),
	})

	res.Write(a.Statement.TealAst())

	if a.Loop.continues {
		res.Write(&Teal_label{Name: fmt.Sprintf("while_%d_continue", a.Index)})
	}

	res.Write(&Teal_b_fixed{TARGET1: fmt.Sprintf("while_%d", a.Index)})
	res.Write(&Teal_label{Name: fmt.Sprintf("while_%d_end", a.Index)})

	return res.Build()
}

type CealFor struct {
	Index     int
	Loop      *LoopScopeItem
	Init      []CealStatement
	Condition CealStatement
	Statement CealStatement
	Iter      []CealStatement
}

func (a *CealFor) TealAst() TealAst {
	res := &TealAstBuilder{}

	for _, stmt := range a.Init {
		res.Write(stmt.TealAst())
	}

	res.Write(&Teal_label{Name: fmt.Sprintf("for_%d", a.Index)})
	res.Write(&Teal_bz_fixed{
		s1:      a.Condition.TealAst(),
		TARGET1: fmt.Sprintf("for_%d_end", a.Index),
	})
	res.Write(a.Statement.TealAst())

	if a.Loop.continues {
		res.Write(&Teal_label{Name: fmt.Sprintf("for_%d_continue", a.Index)})
	}

	for _, stmt := range a.Iter {
		res.Write(stmt.TealAst())
	}

	res.Write(&Teal_b_fixed{TARGET1: fmt.Sprintf("for_%d", a.Index)})
	res.Write(&Teal_label{Name: fmt.Sprintf("for_%d_end", a.Index)})

	return res.Build()
}

type CealExpr interface {
	// ToStmt converts the expression to a statement so it does not push a value onto the stack
	ToStmt()
}

type CealPrefix struct {
	V  *Variable
	Op string

	IsStmt bool
}

func (a *CealPrefix) ToStmt() {
	a.IsStmt = true
}

func (a *CealPrefix) TealAst() TealAst {
	if a.V.constant {
		panic("cannot modify const var")
	}

	var op TealAst

	switch a.Op {
	case "++":
		op = &Teal_plus{
			s1: &Teal_load{I1: uint8(a.V.local.slot)},
			s2: &Teal_int{V: 1},
		}
	case "--":
		op = &Teal_minus{
			s1: &Teal_load{I1: uint8(a.V.local.slot)},
			s2: &Teal_int{V: 1},
		}
	default:
		panic(fmt.Sprintf("prefix operator not supported: '%s'", a.Op))
	}

	res := &TealAstBuilder{}
	res.Write(&Teal_store{s1: op, I1: uint8(a.V.local.slot)})

	if !a.IsStmt {
		res.Write(&Teal_load{I1: uint8(a.V.local.slot)})
	}

	return res.Build()
}

type CealPostfix struct {
	V  *Variable
	Op string

	IsStmt bool
}

func (a *CealPostfix) ToStmt() {
	a.IsStmt = true
}

func (a *CealPostfix) TealAst() TealAst {
	if a.V.constant {
		panic("cannot modify const var")
	}

	var s1 TealAst
	s1 = &Teal_load{I1: uint8(a.V.local.slot)}

	if !a.IsStmt {
		s1 = &Teal_dup{s1: s1}
	}

	s2 := &Teal_int{V: 1}

	var op TealAst
	switch a.Op {
	case "++":
		op = &Teal_plus{s1: s1, s2: s2}
	case "--":
		op = &Teal_minus{s1: s1, s2: s2}
	default:
		panic(fmt.Sprintf("postfix operator not supported: '%s'", a.Op))
	}

	res := &TealAstBuilder{}
	res.Write(&Teal_store{s1: op, I1: uint8(a.V.local.slot)})

	return res.Build()
}

type CealLabel struct {
	Name string
}

func (a *CealLabel) TealAst() TealAst {
	res := &TealAstBuilder{}
	res.Write(&Teal_label{Name: fmt.Sprintf("label_%s", a.Name)})
	return res.Build()
}

type CealGoto struct {
	Label string
}

func (a *CealGoto) TealAst() TealAst {
	res := &TealAstBuilder{}
	res.Write(&Teal_b_fixed{TARGET1: fmt.Sprintf("label_%s", a.Label)})
	return res.Build()
}

type CealVariable struct {
	V *Variable
}

func (a *CealVariable) TealAst() TealAst {
	res := &TealAstBuilder{}

	if a.V.local != nil {
		ast := &Teal_load{I1: uint8(a.V.local.slot)}
		res.Write(ast)
		return res.Build()
	}

	if a.V.param != nil {
		ast := &Teal_frame_dig{I1: int8(a.V.param.index)}
		res.Write(ast)
		return res.Build()
	}

	if a.V.const_ != nil {
		switch a.V.const_.kind {
		case SimpleTypeInt:
			ast := &Teal_intc{
				I1: uint8(a.V.const_.index),
			}
			res.Write(ast)
			return res.Build()
		case SimpleTypeBytes:
			ast := &Teal_bytec{
				I1: uint8(a.V.const_.index),
			}
			res.Write(ast)
			return res.Build()
		}
	}

	if a.V.t == "uint64" {
		res.Write(&Teal_named_int{V: &Teal_named_int_value{V: a.V.name}})
		return res.Build()
	}

	res.Write(&Teal_byte{S: a.V.name})
	return res.Build()
}

type CealUnaryOp struct {
	Op        string
	Statement CealStatement
}

func (a *CealUnaryOp) TealAst() TealAst {
	res := &TealAstBuilder{}
	switch a.Op {
	case "!":
		res.Write(&Teal_not{s1: a.Statement.TealAst()})
	default:
		panic(fmt.Sprintf("unary op '%s' not supported", a.Op))
	}
	return res.Build()
}

type CealAssignSumDiff struct {
	V     *Variable
	F     *StructField
	T     *Type
	Value CealStatement
	Op    string

	IsStmt bool
}

func (a *CealAssignSumDiff) TealAst() TealAst {
	var slot uint8

	if a.T.complex != nil {
		v := a.V.fields[a.F.name]
		slot = uint8(v.local.slot)
	} else {
		slot = uint8(a.V.local.slot)
	}

	s1 := &Teal_load{I1: slot}

	var op TealAst

	switch a.Op {
	case "+=":
		op = &Teal_plus{s1: s1, s2: a.Value.TealAst()}
	case "-=":
		op = &Teal_minus{s1: s1, s2: a.Value.TealAst()}
	}

	if !a.IsStmt {
		op = &Teal_dup{s1: op}
	}

	res := &TealAstBuilder{}
	res.Write(&Teal_store{s1: op, I1: slot})

	return res.Build()
}

type CealAnd struct {
	Index int

	Left  CealStatement
	Right CealStatement
}

func (a *CealAnd) TealAst() TealAst {
	res := &TealAstBuilder{}

	res.Write(&Teal_andand{
		s1: &Teal_bz_fixed{
			s1:      &Teal_dup{s1: a.Left.TealAst()},
			TARGET1: fmt.Sprintf("and_%d_end", a.Index),
		},
		s2: a.Right.TealAst(),
	})

	res.Write(&Teal_label{Name: fmt.Sprintf("and_%d_end", a.Index)})

	return res.Build()
}

type CealOr struct {
	Index int
	Left  CealStatement
	Right CealStatement
}

func (a *CealOr) TealAst() TealAst {
	res := &TealAstBuilder{}

	res.Write(
		&Teal_oror{
			s1: &Teal_bnz_fixed{
				s1: &Teal_dup{
					s1: a.Left.TealAst(),
				},
				TARGET1: fmt.Sprintf("or_%d_end", a.Index),
			},
			s2: a.Right.TealAst(),
		})

	res.Write(&Teal_label{Name: fmt.Sprintf("or_%d_end", a.Index)})

	return res.Build()
}

type CealBinop struct {
	Left  CealStatement
	Op    string
	Right CealStatement
}

func (a *CealBinop) TealAst() TealAst {
	var op TealAst

	switch a.Op {
	case "+":
		op = &Teal_plus{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case "-":
		op = &Teal_minus{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case "*":
		op = &Teal_mul{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case "/":
		op = &Teal_div{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case "==":
		op = &Teal_eqeq{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case "!=":
		op = &Teal_noteq{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case "&":
		op = &Teal_and{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case "^":
		op = &Teal_xor{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case "|":
		op = &Teal_or{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case "<":
		op = &Teal_lt{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case ">":
		op = &Teal_gt{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case "<=":
		op = &Teal_lteq{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case ">=":
		op = &Teal_gteq{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	case "%":
		op = &Teal_mod{s1: a.Left.TealAst(), s2: a.Right.TealAst()}
	default:
		panic(fmt.Sprintf("binary op '%s' is not supported yet", a.Op))
	}

	res := &TealAstBuilder{}
	res.Write(op)
	return res.Build()
}

type CealNegate struct {
	Value CealStatement
}

func (a *CealNegate) TealAst() TealAst {
	res := &TealAstBuilder{}
	res.Write(&Teal_minus{
		s1: &Teal_int{V: 0},
		s2: a.Value.TealAst(),
	})
	return res.Build()
}

type CealDefine struct {
	V *Variable
	T *Type

	Value CealStatement
}

func (a *CealDefine) TealAst() TealAst {
	if a.T.complex != nil {
		panic("defining complex variable is not supported yet")
	}

	ast := &Teal_store{
		s1: a.Value.TealAst(),
		I1: uint8(a.V.local.slot),
	}

	res := &TealAstBuilder{}
	res.Write(ast)
	return res.Build()
}

type CealAssign struct {
	V   *Variable
	T   *Type
	F   *StructField
	Fun *Function

	Value CealStatement

	IsStmt bool
}

func (a *CealAssign) ToStmt() {
	a.IsStmt = true
}

func (a *CealAssign) TealAst() TealAst {
	if a.V.constant {
		panic("cannot assign to a const var")
	}

	if a.V.param != nil {
		// TODO: add param var assignment support
		panic("cannot assign param var")
	}

	res := &TealAstBuilder{}

	if a.T.complex != nil {
		if a.T.complex.builtin != nil {
			res.Write(&Teal_call_builtin{
				Name: a.Fun.builtin.op,
				Imms: []TealAst{&Teal_named_int_value{
					V: a.F.name,
				}},
			})
			return res.Build()
		} else {
			if a.V.param != nil {
				panic("accessing struct param fields is not supported yet")
			}

			v := a.V.fields[a.F.name]
			ast := &Teal_store{
				s1: a.Value.TealAst(),
				I1: uint8(v.local.slot),
			}

			res.Write(ast)
		}
	} else {
		ast := &Teal_store{
			s1: a.Value.TealAst(),
			I1: uint8(a.V.local.slot),
		}

		res.Write(ast)
	}

	if !a.IsStmt {
		load := &Teal_load{
			I1: uint8(a.V.local.slot),
		}

		res.Write(load)
	}

	return res.Build()
}

type CealStructField struct {
	V   *Variable
	T   *Type
	F   *StructField
	Fun *Function
}

func (a *CealStructField) TealAst() TealAst {
	res := &TealAstBuilder{}

	if a.T.complex.builtin != nil {
		res.Write(&Teal_call_builtin{
			Name: a.Fun.builtin.op,
			Imms: []TealAst{&Teal_named_int_value{
				V: a.F.name,
			}},
		})
		return res.Build()
	}

	if a.V.param != nil {
		panic("accessing struct param fields is not supported yet")
	}

	v := a.V.fields[a.F.name]

	ast := &Teal_load{
		I1: uint8(v.local.slot),
	}

	res.Write(ast)

	return res.Build()
}

type CealCall struct {
	Fun  *Function
	Args []CealStatement

	IsStmt bool
}

func (a *CealCall) ToStmt() {
	a.IsStmt = true
}

func (a *CealCall) TealAst() TealAst {
	res := &TealAstBuilder{}

	var args []TealAst

	if a.Fun.builtin != nil {
		i := 0

		for ; i < len(a.Fun.builtin.stack); i++ {
			arg := a.Args[i]
			args = append(args, arg.TealAst())
		}

		var imms []TealAst

		for ; i < len(a.Fun.builtin.stack)+len(a.Fun.builtin.imm); i++ {
			arg := a.Args[i]
			imms = append(imms, arg.TealAst())
		}

		res.Write(&Teal_call_builtin{
			Args: args,
			Imms: imms,
			Name: a.Fun.builtin.op,
		})
	}

	if a.Fun.user != nil {
		for _, arg := range a.Args {
			args = append(args, arg.TealAst())
		}

		res.Write(&Teal_callsub_fixed{
			Args:   args,
			Target: a.Fun.name,
		})
	}

	if a.IsStmt {
		if a.Fun.returns > 0 {
			res.Write(&Teal_popn{N1: uint8(a.Fun.returns)})
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
}

func (a *CealIntConstant) TealAst() TealAst {
	res := &TealAstBuilder{}
	res.Write(&Teal_named_int{V: &Teal_named_int_value{V: a.Value}})
	return res.Build()
}

type CealByteConstant struct {
	Value string
}

func (a *CealByteConstant) TealAst() TealAst {
	res := &TealAstBuilder{}
	res.Write(&Teal_byte{S: a.Value})
	return res.Build()
}

type CealReturn struct {
	Value CealStatement
	Fun   *Function
}

func (a *CealReturn) IsReturn() {
}

func (a *CealReturn) TealAst() TealAst {
	var op TealAst

	if a.Fun != nil && a.Fun.user.sub {
		var values []TealAst

		if a.Value != nil {
			values = append(values, a.Value.TealAst())
		}

		op = &Teal_retsub_fixed{
			Values: values,
		}
	} else {
		op = &Teal_return_fixed{
			Value: a.Value.TealAst(),
		}
	}

	res := &TealAstBuilder{}
	res.Write(op)

	return res.Build()
}

type CealBlock struct {
	Statements []CealStatement
}

func (a *CealBlock) TealAst() TealAst {
	res := &TealAstBuilder{}

	for _, stmt := range a.Statements {
		res.Write(stmt.TealAst())
	}

	return res.Build()
}

type CealConditional struct {
	Index int

	Condition CealStatement

	True  CealStatement
	False CealStatement
}

func (a *CealConditional) TealAst() TealAst {
	res := &TealAstBuilder{}

	false_label := fmt.Sprintf("conditional_%d_false", a.Index)
	end_label := fmt.Sprintf("conditional_%d_end", a.Index)

	res.Write(&Teal_bz_fixed{
		s1:      a.Condition.TealAst(),
		TARGET1: false_label,
	})
	res.Write(a.True.TealAst())
	res.Write(&Teal_b_fixed{TARGET1: end_label})
	res.Write(&Teal_label{Name: false_label})
	res.Write(a.False.TealAst())
	res.Write(&Teal_label{Name: end_label})

	return res.Build()
}

type CealIfAlternative struct {
	Condition  CealStatement
	Statements []CealStatement
}

type CealIf struct {
	Index        int
	Alternatives []*CealIfAlternative
}

func (a *CealIf) TealAst() TealAst {
	res := &TealAstBuilder{}

	end_label := fmt.Sprintf("if_end_%d", a.Index)

	for i, alt := range a.Alternatives {
		if alt.Condition != nil {
			if i < len(a.Alternatives)-1 {
				res.Write(&Teal_bz_fixed{
					s1:      alt.Condition.TealAst(),
					TARGET1: fmt.Sprintf("if_skip_%d_%d", a.Index, i),
				})
			} else {
				res.Write(&Teal_bz_fixed{
					s1:      alt.Condition.TealAst(),
					TARGET1: end_label,
				})
			}
		}

		for _, stmt := range alt.Statements {
			res.Write(stmt.TealAst())
		}

		if i < len(a.Alternatives)-1 {
			res.Write(&Teal_b_fixed{TARGET1: end_label})
		}

		if alt.Condition != nil {
			if i < len(a.Alternatives)-1 {
				res.Write(&Teal_label{Name: fmt.Sprintf("if_skip_%d_%d", a.Index, i)})
			}
		}
	}

	res.Write(&Teal_label{Name: end_label})

	return res.Build()
}

type CealFunction struct {
	Fun        *Function
	Statements []CealStatement
}

func (a *CealFunction) TealAst() TealAst {
	res := &TealAstBuilder{}

	if a.Fun.user.sub {
		if a.Fun.user.args != 0 || a.Fun.returns != 0 {
			ast := &Teal_proto{
				A1: uint8(a.Fun.user.args),
				R2: uint8(a.Fun.returns),
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
				res.Write(&Teal_retsub{})
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

func (a *CealRaw) Teal() Teal {
	return Teal{a}
}

func (a *CealRaw) TealAst() TealAst {
	return a
}

type CealSingleLineComment struct {
	Line string
}

func (a *CealSingleLineComment) TealAst() TealAst {
	return &Teal_comment{Lines: []string{a.Line}}
}

type CealMultiLineComment struct {
	Text string
}

func (a *CealMultiLineComment) TealAst() TealAst {
	lines := strings.Split(a.Text, "\n")
	return &Teal_comment{Lines: lines}
}
