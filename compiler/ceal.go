package compiler

import (
	"fmt"
	"strings"
)

type CealProgram struct {
	ConstInts  []int
	ConstBytes [][]byte

	Functions     map[string]*CealFunction
	FunctionNames []string
}

func (a *CealProgram) String() string {
	res := strings.Builder{}

	res.WriteString(fmt.Sprintf("#pragma version %d\n", AvmVersion))

	if len(a.ConstInts) > 0 {
		constints := []string{}
		for _, v := range a.ConstInts {
			constints = append(constints, itoa(v))
		}
		res.WriteString(fmt.Sprintf("intcblock %s\n", strings.Join(constints, " ")))
	}

	main := a.Functions[AvmMainName]

	if len(a.Functions) > 1 {
		res.WriteString(fmt.Sprintf("b %s\n", main.Fun.name))
	}

	for _, name := range a.FunctionNames {
		ast := a.Functions[name]

		if name == AvmMainName {
			continue
		}

		res.WriteString(fmt.Sprintf("%s:\n", ast.Fun.name))

		res.WriteString(ast.String())
		res.WriteString("\n")
	}

	if len(a.Functions) > 1 {
		res.WriteString(fmt.Sprintf("%s:\n", main.Fun.name))
	}

	res.WriteString(main.String())

	return res.String()
}

type CealContinue struct {
	Label string
	Index int
}

func (a *CealContinue) String() string {
	return fmt.Sprintf("b %s_%d_continue", a.Label, a.Index)
}

type CealBreak struct {
	Label string
	Index int
}

func (a *CealBreak) String() string {
	return fmt.Sprintf("b %s_%d_end", a.Label, a.Index)
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
	res := Lines{}

	labels := []string{}

	for i, c := range a.Cases {
		label := fmt.Sprintf("switch_%d_%d", a.Index, i)
		labels = append(labels, label)

		res.WriteLine(c.Value.String())
	}

	res.WriteLine(a.Value.String())
	res.WriteLine(fmt.Sprintf("match %s", strings.Join(labels, " ")))

	if len(a.Default) > 0 {
		for _, stmt := range a.Default {
			res.WriteLine(stmt.String())
		}
	}

	for i, c := range a.Cases {
		label := labels[i]

		res.WriteLine(fmt.Sprintf("%s:", label))

		for _, stmt := range c.Statements {
			res.WriteLine(stmt.String())
		}
	}

	if a.Loop.breaks {
		res.WriteLine(fmt.Sprintf("switch_%d_end:", a.Index))
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
	res := Lines{}
	res.WriteLine(fmt.Sprintf("do_%d:", a.Index))
	res.WriteLine(a.Statement.String())

	if a.Loop.continues {
		res.WriteLine(fmt.Sprintf("do_%d_continue:", a.Index))
	}

	res.WriteLine(a.Condition.String())
	res.WriteLine(fmt.Sprintf("bnz do_%d", a.Index))

	if a.Loop.breaks {
		res.WriteLine(fmt.Sprintf("do_%d_end:", a.Index))
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
	res := Lines{}

	res.WriteLine(fmt.Sprintf("while_%d:", a.Index))
	res.WriteLine(a.Condition.String())
	res.WriteLine(fmt.Sprintf("bz while_%d_end", a.Index))
	res.WriteLine(a.Statement.String())
	if a.Loop.continues {
		res.WriteLine(fmt.Sprintf("while_%d_continue:", a.Index))
	}
	res.WriteLine(fmt.Sprintf("b while_%d", a.Index))
	res.WriteLine(fmt.Sprintf("while_%d_end:", a.Index))

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
	res := Lines{}

	for _, stmt := range a.Init {
		res.WriteLine(stmt.String())
	}

	res.WriteLine(fmt.Sprintf("for_%d:", a.Index))
	res.WriteLine(a.Condition.String())
	res.WriteLine(fmt.Sprintf("bz for_%d_end", a.Index))
	res.WriteLine(a.Statement.String())

	if a.Loop.continues {
		res.WriteLine(fmt.Sprintf("for_%d_continue:", a.Index))
	}

	for _, stmt := range a.Iter {
		res.WriteLine(stmt.String())
	}

	res.WriteLine(fmt.Sprintf("b for_%d", a.Index))
	res.WriteLine(fmt.Sprintf("for_%d_end:", a.Index))

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

	var op string

	switch a.Op {
	case "++":
		op = "+"
	case "--":
		op = "-"
	default:
		panic(fmt.Sprintf("prefix operator not supported: '%s'", a.Op))
	}

	res := Lines{}

	res.WriteLine(fmt.Sprintf("load %d", a.V.local.slot))
	res.WriteLine("int 1")
	res.WriteLine(op)
	res.WriteLine(fmt.Sprintf("store %d", a.V.local.slot))

	if !a.IsStmt {
		res.WriteLine(fmt.Sprintf("load %d", a.V.local.slot))
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

	var op string

	switch a.Op {
	case "++":
		op = "+"
	case "--":
		op = "-"
	default:
		panic(fmt.Sprintf("postfix operator not supported: '%s'", a.Op))
	}

	res := Lines{}
	res.WriteLine(fmt.Sprintf("load %d", a.V.local.slot))
	if !a.IsStmt {
		res.WriteLine("dup")
	}
	res.WriteLine(fmt.Sprintf("int 1\n%s\nstore %d", op, a.V.local.slot))

	return res.String()
}

type CealLabel struct {
	Name string
}

func (a *CealLabel) String() string {
	return fmt.Sprintf("label_%s:", a.Name)
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
		ast := Teal_load{I1: itoa(a.V.local.slot)}
		return ast.String()
	}

	if a.V.param != nil {
		ast := Teal_frame_dig{I1: itoa(a.V.param.index)}
		return ast.String()
	}

	if a.V.const_ != nil {
		switch a.V.const_.kind {
		case SimpleTypeInt:
			ast := Teal_intc{
				I1: itoa(a.V.const_.index),
			}
			return ast.String()
		case SimpleTypeBytes:
			ast := Teal_bytec{
				I1: itoa(a.V.const_.index),
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
	res := Lines{}

	var op string

	switch a.Op {
	case "+=":
		op = "+"
	case "-=":
		op = "-"
	}

	var slot int

	if a.T.complex != nil {
		v := a.V.fields[a.F.name]
		slot = v.local.slot
	} else {
		slot = a.V.local.slot
	}

	res.WriteLine(fmt.Sprintf("load %d", slot))
	res.WriteLine(a.Value.String())
	res.WriteLine(op)
	if !a.IsStmt {
		res.WriteLine("dup")
	}
	res.WriteLine(fmt.Sprintf("store %d", slot))

	return res.String()
}

type CealAnd struct {
	Index int

	Left  AstStatement
	Right AstStatement
}

func (a *CealAnd) String() string {
	res := Lines{}

	res.WriteLine(a.Left.String())
	res.WriteLine("dup")
	res.WriteLine(fmt.Sprintf("bz and_%d_end", a.Index))

	res.WriteLine(a.Right.String())
	res.WriteLine("&&")

	res.WriteLine(fmt.Sprintf("and_%d_end:", a.Index))

	return res.String()
}

type CealOr struct {
	Index int
	Left  AstStatement
	Right AstStatement
}

func (a *CealOr) String() string {
	res := Lines{}

	res.WriteLine(a.Left.String())
	res.WriteLine("dup")
	res.WriteLine(fmt.Sprintf("bnz or_%d_end", a.Index))

	res.WriteLine(a.Right.String())
	res.WriteLine("||")

	res.WriteLine(fmt.Sprintf("or_%d_end:", a.Index))

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
		I1: itoa(a.V.local.slot),
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

	res := Lines{}

	if a.T.complex != nil {
		if a.T.complex.builtin != nil {
			return fmt.Sprintf("%s %s", a.Fun.builtin.op, a.F.name)
		} else {
			if a.V.param != nil {
				panic("accessing struct param fields is not supported yet")
			}

			v := a.V.fields[a.F.name]
			ast := Teal_store{
				s1: a.Value,
				I1: itoa(v.local.slot),
			}

			res.WriteLine(ast.String())
		}
	} else {
		ast := Teal_store{
			s1: a.Value,
			I1: itoa(a.V.local.slot),
		}

		res.WriteLine(ast.String())
	}

	if !a.IsStmt {
		load := Teal_load{
			I1: itoa(a.V.local.slot),
		}

		res.WriteLine(load.String())
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
		I1: itoa(v.local.slot),
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
	s := Lines{}

	if a.Fun.builtin != nil {
		i := 0

		for ; i < len(a.Fun.builtin.stack); i++ {
			arg := a.Args[i]
			s.WriteLine(arg.String())
		}

		args := []string{}
		args = append(args, a.Fun.builtin.op)

		for ; i < len(a.Fun.builtin.stack)+len(a.Fun.builtin.imm); i++ {
			arg := a.Args[i]
			args = append(args, arg.String())
		}

		s.WriteLine(strings.Join(args, " "))
	}

	if a.Fun.user != nil {
		for _, arg := range a.Args {
			s.WriteLine(arg.String())
		}

		s.WriteLine(fmt.Sprintf("callsub %s", a.Fun.name))
	}

	if a.IsStmt {
		if a.Fun.returns > 0 {
			ast := Teal_popn{N1: itoa(a.Fun.returns)}
			s.WriteLine(ast.String())
		}
	}

	return s.String()
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
	s := Lines{}

	if a.Value != nil {
		s.WriteLine(a.Value.String())
	}

	if a.Fun != nil {
		if a.Fun.user.sub {
			s.WriteLine("retsub")
		} else {
			s.WriteLine("return")
		}
	} else {
		s.WriteLine("return")
	}

	return s.String()
}

type CealBlock struct {
	Statements []AstStatement
}

func (a *CealBlock) String() string {
	ops := Lines{}

	for _, stmt := range a.Statements {
		ops.WriteLine(stmt.String())
	}

	return ops.String()
}

type CealConditional struct {
	Index int

	Condition AstStatement

	True  AstStatement
	False AstStatement
}

func (a *CealConditional) String() string {
	res := Lines{}

	res.WriteLine(a.Condition.String())
	res.WriteLine(fmt.Sprintf("bz conditional_%d_false", a.Index))
	res.WriteLine(a.True.String())
	res.WriteLine(fmt.Sprintf("b conditional_%d_end", a.Index))
	res.WriteLine(fmt.Sprintf("conditional_%d_false:", a.Index))
	res.WriteLine(a.False.String())
	res.WriteLine(fmt.Sprintf("conditional_%d_end:", a.Index))

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
	res := Lines{}

	for i, alt := range a.Alternatives {
		if alt.Condition != nil {
			res.WriteLine(alt.Condition.String())

			if i < len(a.Alternatives)-1 {
				res.WriteLine(fmt.Sprintf("bz if_skip_%d_%d", a.Index, i))
			} else {
				res.WriteLine(fmt.Sprintf("bz if_end_%d", a.Index))
			}
		}

		for _, stmt := range alt.Statements {
			res.WriteLine(stmt.String())
		}

		if i < len(a.Alternatives)-1 {
			res.WriteLine(fmt.Sprintf("b if_end_%d", a.Index))
		}

		if alt.Condition != nil {
			if i < len(a.Alternatives)-1 {
				res.WriteLine(fmt.Sprintf("if_skip_%d_%d:", a.Index, i))
			}
		}
	}

	res.WriteLine(fmt.Sprintf("if_end_%d:", a.Index))

	return res.String()
}

type CealFunction struct {
	Fun        *Function
	Statements []AstStatement
}

func (a *CealFunction) String() string {
	res := Lines{}

	if a.Fun.user.sub {
		if a.Fun.user.args != 0 || a.Fun.returns != 0 {
			ast := Teal_proto{
				A1: itoa(a.Fun.user.args),
				R2: itoa(a.Fun.returns),
			}

			res.WriteLine(ast.String())
		}
	}

	for _, stmt := range a.Statements {
		res.WriteLine(stmt.String())
	}

	if a.Fun.user.sub {
		if len(a.Statements) > 0 {
			last := a.Statements[len(a.Statements)-1]
			if _, ok := last.(CealIsReturn); !ok {
				res.WriteLine("retsub")
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
