package compiler

import (
	"fmt"
)

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

func (a *CealProgram) String() string {
	res := Teal{}

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
		res.Write(ast)
	}

	if len(a.Functions) > 1 {
		res.Write(&Teal_label{Name: main.Fun.name})
	}

	res.Write(main)

	return res.String()
}

type CealContinue struct {
	Label string
	Index int
}

func (a *CealContinue) String() string {
	teal := Teal_b_fixed{TARGET1: fmt.Sprintf("%s_%d_continue", a.Label, a.Index)}
	return teal.String()
}

type CealBreak struct {
	Label string
	Index int
}

func (a *CealBreak) String() string {
	teal := Teal_b_fixed{TARGET1: fmt.Sprintf("%s_%d_end", a.Label, a.Index)}
	return teal.String()
}

type CealSwitchCase struct {
	Value      AstStatement
	Statements []AstStatement
}

type CealSwitch struct {
	Index int
	Loop  *LoopScopeItem

	Value AstStatement
	Cases []*CealSwitchCase

	Default []AstStatement
}

func (a *CealSwitch) String() string {
	res := Teal{}

	labels := []string{}

	for i, c := range a.Cases {
		label := fmt.Sprintf("switch_%d_%d", a.Index, i)
		labels = append(labels, label)

		res.Write(c.Value)
	}

	res.Write(a.Value)
	res.Write(&Teal_match_fixed{TARGET1: labels})

	if len(a.Default) > 0 {
		for _, stmt := range a.Default {
			res.Write(stmt)
		}
	}

	for i, c := range a.Cases {
		label := labels[i]

		res.Write(&Teal_label{Name: label})

		for _, stmt := range c.Statements {
			res.Write(stmt)
		}
	}

	if a.Loop.breaks {
		res.Write(&Teal_label{Name: fmt.Sprintf("switch_%d_end", a.Index)})
	}

	return res.String()
}

type CealDoWhile struct {
	Index     int
	Loop      *LoopScopeItem
	Condition AstStatement
	Statement AstStatement
}

func (a *CealDoWhile) String() string {
	res := Teal{}
	res.Write(&Teal_label{Name: fmt.Sprintf("do_%d", a.Index)})
	res.Write(a.Statement)

	if a.Loop.continues {
		res.Write(&Teal_label{Name: fmt.Sprintf("do_%d_continue", a.Index)})
	}

	res.Write(&Teal_bnz_fixed{
		s1:      a.Condition,
		TARGET1: fmt.Sprintf("do_%d", a.Index),
	})

	if a.Loop.breaks {
		res.Write(&Teal_label{Name: fmt.Sprintf("do_%d_end", a.Index)})
	}

	return res.String()
}

type CealWhile struct {
	Index     int
	Loop      *LoopScopeItem
	Condition AstStatement
	Statement AstStatement
}

func (a *CealWhile) String() string {
	res := Teal{}

	res.Write(&Teal_label{Name: fmt.Sprintf("while_%d", a.Index)})
	res.Write(&Teal_bz_fixed{
		s1:      a.Condition,
		TARGET1: fmt.Sprintf("while_%d_end", a.Index),
	})
	res.Write(a.Statement)
	if a.Loop.continues {
		res.Write(&Teal_label{Name: fmt.Sprintf("while_%d_continue", a.Index)})
	}
	res.Write(&Teal_b_fixed{TARGET1: fmt.Sprintf("while_%d", a.Index)})
	res.Write(&Teal_label{Name: fmt.Sprintf("while_%d_end", a.Index)})

	return res.String()
}

type CealFor struct {
	Index     int
	Loop      *LoopScopeItem
	Init      []AstStatement
	Condition AstStatement
	Statement AstStatement
	Iter      []AstStatement
}

func (a *CealFor) String() string {
	res := Teal{}

	for _, stmt := range a.Init {
		res.Write(stmt)
	}

	res.Write(&Teal_label{Name: fmt.Sprintf("for_%d", a.Index)})
	res.Write(&Teal_bz_fixed{
		s1:      a.Condition,
		TARGET1: fmt.Sprintf("for_%d_end", a.Index),
	})
	res.Write(a.Statement)

	if a.Loop.continues {
		res.Write(&Teal_label{Name: fmt.Sprintf("for_%d_continue", a.Index)})
	}

	for _, stmt := range a.Iter {
		res.Write(stmt)
	}

	res.Write(&Teal_b_fixed{TARGET1: fmt.Sprintf("for_%d", a.Index)})
	res.Write(&Teal_label{Name: fmt.Sprintf("for_%d_end", a.Index)})

	return res.String()
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

func (a *CealPrefix) String() string {
	if a.V.constant {
		panic("cannot modify const var")
	}

	var op TealOp

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

	res := Teal{}
	res.Write(&Teal_store{s1: op, I1: uint8(a.V.local.slot)})

	if !a.IsStmt {
		res.Write(&Teal_load{I1: uint8(a.V.local.slot)})
	}

	return res.String()
}

type CealPostfix struct {
	V  *Variable
	Op string

	IsStmt bool
}

func (a *CealPostfix) ToStmt() {
	a.IsStmt = true
}

func (a *CealPostfix) String() string {
	if a.V.constant {
		panic("cannot modify const var")
	}

	var s1 TealOp
	s1 = &Teal_load{I1: uint8(a.V.local.slot)}

	if !a.IsStmt {
		s1 = &Teal_dup{s1: s1}
	}

	s2 := &Teal_int{V: 1}

	var op TealOp
	switch a.Op {
	case "++":
		op = &Teal_plus{s1: s1, s2: s2}
	case "--":
		op = &Teal_minus{s1: s1, s2: s2}
	default:
		panic(fmt.Sprintf("postfix operator not supported: '%s'", a.Op))
	}

	res := Teal{}
	res.Write(&Teal_store{s1: op, I1: uint8(a.V.local.slot)})

	return res.String()
}

type CealLabel struct {
	Name string
}

func (a *CealLabel) String() string {
	teal := &Teal_label{Name: fmt.Sprintf("label_%s", a.Name)}
	return teal.String()
}

type CealGoto struct {
	Label string
}

func (a *CealGoto) String() string {
	return fmt.Sprintf("b label_%s", a.Label)
}

type CealVariable struct {
	V *Variable
}

func (a *CealVariable) String() string {
	if a.V.local != nil {
		ast := Teal_load{I1: uint8(a.V.local.slot)}
		return ast.String()
	}

	if a.V.param != nil {
		ast := Teal_frame_dig{I1: int8(a.V.param.index)}
		return ast.String()
	}

	if a.V.const_ != nil {
		switch a.V.const_.kind {
		case SimpleTypeInt:
			ast := Teal_intc{
				I1: uint8(a.V.const_.index),
			}
			return ast.String()
		case SimpleTypeBytes:
			ast := Teal_bytec{
				I1: uint8(a.V.const_.index),
			}
			return ast.String()
		}
	}

	if a.V.t == "uint64" {
		return fmt.Sprintf("int %s", a.V.name)
	}

	return fmt.Sprintf("byte %s", a.V.name)
}

type CealUnaryOp struct {
	Op        string
	Statement AstStatement
}

func (a *CealUnaryOp) String() string {
	return fmt.Sprintf("%s\n%s", a.Statement.String(), a.Op)
}

type CealAssignSumDiff struct {
	V     *Variable
	F     *StructField
	T     *Type
	Value AstStatement
	Op    string

	IsStmt bool
}

func (a *CealAssignSumDiff) String() string {
	var slot uint8

	if a.T.complex != nil {
		v := a.V.fields[a.F.name]
		slot = uint8(v.local.slot)
	} else {
		slot = uint8(a.V.local.slot)
	}

	s1 := &Teal_load{I1: slot}

	var op TealOp

	switch a.Op {
	case "+=":
		op = &Teal_plus{s1: s1, s2: a.Value}
	case "-=":
		op = &Teal_minus{s1: s1, s2: a.Value}
	}

	if !a.IsStmt {
		op = &Teal_dup{s1: op}
	}

	teal := &Teal_store{s1: op, I1: slot}

	return teal.String()
}

type CealAnd struct {
	Index int

	Left  AstStatement
	Right AstStatement
}

func (a *CealAnd) String() string {
	res := Teal{}

	res.Write(&Teal_andand{
		s1: &Teal_bz_fixed{
			s1:      &Teal_dup{s1: a.Left},
			TARGET1: fmt.Sprintf("and_%d_end", a.Index),
		},
		s2: a.Right,
	})

	res.Write(&Teal_label{Name: fmt.Sprintf("and_%d_end", a.Index)})

	return res.String()
}

type CealOr struct {
	Index int
	Left  AstStatement
	Right AstStatement
}

func (a *CealOr) String() string {
	res := Teal{}

	res.Write(
		&Teal_oror{
			s1: &Teal_bnz_fixed{
				s1: &Teal_dup{
					s1: a.Left,
				},
				TARGET1: fmt.Sprintf("or_%d_end", a.Index),
			},
			s2: a.Right,
		})

	res.Write(&Teal_label{Name: fmt.Sprintf("or_%d_end", a.Index)})

	return res.String()
}

type CealBinop struct {
	Left  AstStatement
	Op    string
	Right AstStatement
}

func (a *CealBinop) String() string {
	return fmt.Sprintf("%s\n%s\n%s", a.Left.String(), a.Right.String(), a.Op)
}

type CealMinusOp struct {
	Value AstStatement
}

func (a *CealMinusOp) String() string {
	return fmt.Sprintf("int 0\n%s\n-", a.Value.String())
}

type CealDefine struct {
	V *Variable
	T *Type

	Value AstStatement
}

func (a *CealDefine) String() string {
	if a.T.complex != nil {
		panic("defining complex variable is not supported yet")
	}

	ast := Teal_store{
		s1: a.Value,
		I1: uint8(a.V.local.slot),
	}

	return ast.String()
}

type CealAssign struct {
	V   *Variable
	T   *Type
	F   *StructField
	Fun *Function

	Value AstStatement

	IsStmt bool
}

func (a *CealAssign) ToStmt() {
	a.IsStmt = true
}

func (a *CealAssign) String() string {
	if a.V.constant {
		panic("cannot assign to a const var")
	}

	if a.V.param != nil {
		// TODO: add param var assignment support
		panic("cannot assign param var")
	}

	res := Teal{}

	if a.T.complex != nil {
		if a.T.complex.builtin != nil {
			return fmt.Sprintf("%s %s", a.Fun.builtin.op, a.F.name)
		} else {
			if a.V.param != nil {
				panic("accessing struct param fields is not supported yet")
			}

			v := a.V.fields[a.F.name]
			ast := &Teal_store{
				s1: a.Value,
				I1: uint8(v.local.slot),
			}

			res.Write(ast)
		}
	} else {
		ast := &Teal_store{
			s1: a.Value,
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

	return res.String()
}

type CealStructField struct {
	V   *Variable
	T   *Type
	F   *StructField
	Fun *Function
}

func (a *CealStructField) String() string {
	if a.T.complex.builtin != nil {
		return fmt.Sprintf("%s %s", a.Fun.builtin.op, a.F.name)
	}

	if a.V.param != nil {
		panic("accessing struct param fields is not supported yet")
	}

	v := a.V.fields[a.F.name]

	ast := Teal_load{
		I1: uint8(v.local.slot),
	}

	return ast.String()
}

type CealCall struct {
	Fun  *Function
	Args []AstStatement

	IsStmt bool
}

func (a *CealCall) ToStmt() {
	a.IsStmt = true
}

func (a *CealCall) String() string {
	res := Teal{}

	var args []TealOp

	if a.Fun.builtin != nil {
		i := 0

		for ; i < len(a.Fun.builtin.stack); i++ {
			arg := a.Args[i]
			args = append(args, arg)
		}

		var imms []TealOp

		for ; i < len(a.Fun.builtin.stack)+len(a.Fun.builtin.imm); i++ {
			arg := a.Args[i]
			imms = append(imms, arg)
		}

		res.Write(&Teal_call_builtin{
			Args: args,
			Imms: imms,
			Name: a.Fun.builtin.op,
		})
	}

	if a.Fun.user != nil {
		for _, arg := range a.Args {
			args = append(args, arg)
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

	return res.String()
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

func (a *CealIntConstant) String() string {
	return fmt.Sprintf("int %s", a.Value)
}

type CealByteConstant struct {
	Value string
}

func (a *CealByteConstant) String() string {
	return fmt.Sprintf("byte %s", a.Value)
}

type CealReturn struct {
	Value AstStatement
	Fun   *Function
}

func (a *CealReturn) IsReturn() {
}

func (a *CealReturn) String() string {
	var op TealOp

	if a.Fun != nil && a.Fun.user.sub {
		var values []TealOp

		if a.Value != nil {
			values = append(values, a.Value)
		}

		op = &Teal_retsub_fixed{
			Values: values,
		}
	} else {
		op = &Teal_return_fixed{
			Value: a.Value,
		}
	}

	return op.String()
}

type CealBlock struct {
	Statements []AstStatement
}

func (a *CealBlock) String() string {
	res := Teal{}

	for _, stmt := range a.Statements {
		res.Write(stmt)
	}

	return res.String()
}

type CealConditional struct {
	Index int

	Condition AstStatement

	True  AstStatement
	False AstStatement
}

func (a *CealConditional) String() string {
	res := Teal{}

	false_label := fmt.Sprintf("conditional_%d_false", a.Index)
	end_label := fmt.Sprintf("conditional_%d_end", a.Index)

	res.Write(&Teal_bz_fixed{
		s1:      a.Condition,
		TARGET1: false_label,
	})
	res.Write(a.True)
	res.Write(&Teal_b_fixed{TARGET1: end_label})
	res.Write(&Teal_label{Name: false_label})
	res.Write(a.False)
	res.Write(&Teal_label{Name: end_label})

	return res.String()
}

type CealIfAlternative struct {
	Condition  AstStatement
	Statements []AstStatement
}

type CealIf struct {
	Index        int
	Alternatives []*CealIfAlternative
}

func (a *CealIf) String() string {
	res := Teal{}

	end_label := fmt.Sprintf("if_end_%d", a.Index)

	for i, alt := range a.Alternatives {
		if alt.Condition != nil {
			if i < len(a.Alternatives)-1 {
				res.Write(&Teal_bz_fixed{
					s1:      alt.Condition,
					TARGET1: fmt.Sprintf("if_skip_%d_%d", a.Index, i),
				})
			} else {
				res.Write(&Teal_bz_fixed{
					s1:      alt.Condition,
					TARGET1: end_label,
				})
			}
		}

		for _, stmt := range alt.Statements {
			res.Write(stmt)
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

	return res.String()
}

type CealFunction struct {
	Fun        *Function
	Statements []AstStatement
}

func (a *CealFunction) String() string {
	res := Teal{}

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
		res.Write(stmt)
	}

	if a.Fun.user.sub {
		if len(a.Statements) > 0 {
			last := a.Statements[len(a.Statements)-1]
			if _, ok := last.(CealIsReturn); !ok {
				res.Write(&Teal_retsub{})
			}
		}
	}

	return res.String()
}

type CealRaw struct {
	Value string
}

func (a *CealRaw) String() string {
	return a.Value
}
