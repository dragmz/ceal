package compiler

import (
	"fmt"
	"strconv"
	"strings"
)

func itoa(v int) string {
	return strconv.Itoa(v)
}

func atoi(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return i
}

type Lines struct {
	lines []string
}

func (l *Lines) WriteLine(s string) {
	if s == "" {
		return
	}

	l.lines = append(l.lines, s)
}

func (l *Lines) String() string {
	return strings.Join(l.lines, "\n")
}

type AstContinue struct {
	label string
	index int
}

func (a *AstContinue) String() string {
	return fmt.Sprintf("b %s_%d_continue", a.label, a.index)
}

type AstBreak struct {
	label string
	index int
}

func (a *AstBreak) String() string {
	return fmt.Sprintf("b %s_%d_end", a.label, a.index)
}

type AstSwitchCase struct {
	value      AstStatement
	statements []AstStatement
}

type AstSwitch struct {
	index int

	value AstStatement
	cases []*AstSwitchCase

	default_ []AstStatement
}

func (a *AstSwitch) String() string {
	res := Lines{}

	labels := []string{}

	for i, c := range a.cases {
		label := fmt.Sprintf("switch_%d_%d", a.index, i)
		labels = append(labels, label)

		res.WriteLine(c.value.String())
	}

	res.WriteLine(a.value.String())
	res.WriteLine(fmt.Sprintf("match %s", strings.Join(labels, " ")))

	if len(a.default_) > 0 {
		for _, stmt := range a.default_ {
			res.WriteLine(stmt.String())
		}
	}

	for i, c := range a.cases {
		label := labels[i]

		res.WriteLine(fmt.Sprintf("%s:", label))

		for _, stmt := range c.statements {
			res.WriteLine(stmt.String())
		}
	}

	res.WriteLine(fmt.Sprintf("switch_%d_end:", a.index))

	return res.String()
}

type AstPop struct {
	s AstStatement
}

func (a *AstPop) String() string {
	return fmt.Sprintf("%s\n%s", a.s.String(), "pop")
}

type AstDoWhile struct {
	index     int
	condition AstStatement
	s         AstStatement
}

func (a *AstDoWhile) String() string {
	res := Lines{}
	res.WriteLine(fmt.Sprintf("do_%d:", a.index))
	res.WriteLine(a.s.String())
	res.WriteLine(fmt.Sprintf("do_%d_continue:", a.index))
	res.WriteLine(a.condition.String())
	res.WriteLine(fmt.Sprintf("bnz do_%d", a.index))
	res.WriteLine(fmt.Sprintf("do_%d_end:", a.index))

	return res.String()
}

type AstWhile struct {
	index     int
	condition AstStatement
	s         AstStatement
}

func (a *AstWhile) String() string {
	res := Lines{}

	res.WriteLine(fmt.Sprintf("while_%d:", a.index))
	res.WriteLine(a.condition.String())
	res.WriteLine(fmt.Sprintf("bz while_%d_end", a.index))
	res.WriteLine(a.s.String())
	res.WriteLine(fmt.Sprintf("while_%d_continue:", a.index))
	res.WriteLine(fmt.Sprintf("b while_%d", a.index))
	res.WriteLine(fmt.Sprintf("while_%d_end:", a.index))

	return res.String()
}

type AstFor struct {
	index     int
	init      []AstStatement
	condition AstStatement
	s         AstStatement
	iter      []AstStatement
}

func (a *AstFor) String() string {
	res := Lines{}

	for _, stmt := range a.init {
		res.WriteLine(stmt.String())
	}

	res.WriteLine(fmt.Sprintf("for_%d:", a.index))
	res.WriteLine(a.condition.String())
	res.WriteLine(fmt.Sprintf("bz for_%d_end", a.index))
	res.WriteLine(a.s.String())

	res.WriteLine(fmt.Sprintf("for_%d_continue:", a.index))
	for _, stmt := range a.iter {
		res.WriteLine(stmt.String())
	}

	res.WriteLine(fmt.Sprintf("b for_%d", a.index))
	res.WriteLine(fmt.Sprintf("for_%d_end:", a.index))

	return res.String()
}

type AstExpr interface {
	// ToStmt converts the expression to a statement so it does not push a value onto the stack
	ToStmt()
}

type AstPrefix struct {
	v    *Variable
	op   string
	stmt bool
}

func (a *AstPrefix) ToStmt() {
	a.stmt = true
}

func (a *AstPrefix) String() string {
	if a.v.constant {
		panic("cannot modify const var")
	}

	var op string

	switch a.op {
	case "++":
		op = "+"
	case "--":
		op = "-"
	default:
		panic(fmt.Sprintf("prefix operator not supported: '%s'", a.op))
	}

	res := Lines{}

	res.WriteLine(fmt.Sprintf("load %d", a.v.local.slot))
	res.WriteLine("int 1")
	res.WriteLine(op)
	res.WriteLine(fmt.Sprintf("store %d", a.v.local.slot))

	if !a.stmt {
		res.WriteLine(fmt.Sprintf("load %d", a.v.local.slot))
	}

	return res.String()
}

type AstPostfix struct {
	v    *Variable
	op   string
	stmt bool
}

func (a *AstPostfix) ToStmt() {
	a.stmt = true
}

func (a *AstPostfix) String() string {
	if a.v.constant {
		panic("cannot modify const var")
	}

	var op string

	switch a.op {
	case "++":
		op = "+"
	case "--":
		op = "-"
	default:
		panic(fmt.Sprintf("postfix operator not supported: '%s'", a.op))
	}

	res := Lines{}
	res.WriteLine(fmt.Sprintf("load %d", a.v.local.slot))
	if !a.stmt {
		res.WriteLine("dup")
	}
	res.WriteLine(fmt.Sprintf("int 1\n%s\nstore %d", op, a.v.local.slot))

	return res.String()
}

type AstLabel struct {
	name string
}

func (a *AstLabel) String() string {
	return fmt.Sprintf("label_%s:", a.name)
}

type AstGoto struct {
	label string
}

func (a *AstGoto) String() string {
	return fmt.Sprintf("b label_%s", a.label)
}

type AstVariable struct {
	v *Variable
}

func (a *AstVariable) String() string {
	if a.v.local != nil {
		ast := avm_load_Ast{I1: itoa(a.v.local.slot)}
		return ast.String()
	}

	if a.v.param != nil {
		ast := avm_frame_dig_Ast{I1: itoa(a.v.param.index)}
		return ast.String()
	}

	if a.v.t == "uint64" {
		return fmt.Sprintf("int %s", a.v.name)
	}

	return fmt.Sprintf("byte %s", a.v.name)
}

type AstUnaryOp struct {
	op string
	s  AstStatement
}

func (a *AstUnaryOp) String() string {
	return fmt.Sprintf("%s\n%s", a.s.String(), a.op)
}

type AstAssignSumDiff struct {
	v     *Variable
	f     *StructField
	t     *Type
	value AstStatement
	op    string

	stmt bool
}

func (a *AstAssignSumDiff) String() string {
	res := Lines{}

	var op string

	switch a.op {
	case "+=":
		op = "+"
	case "-=":
		op = "-"
	}

	var slot int

	if a.t.complex != nil {
		v := a.v.fields[a.f.name]
		slot = v.local.slot
	} else {
		slot = a.v.local.slot
	}

	res.WriteLine(fmt.Sprintf("load %d", slot))
	res.WriteLine(a.value.String())
	res.WriteLine(op)
	if !a.stmt {
		res.WriteLine("dup")
	}
	res.WriteLine(fmt.Sprintf("store %d", slot))

	return res.String()
}

type AstAnd struct {
	index int
	l     AstStatement
	r     AstStatement
}

func (a *AstAnd) String() string {
	res := Lines{}

	res.WriteLine(a.l.String())
	res.WriteLine("dup")
	res.WriteLine(fmt.Sprintf("bz and_%d_end", a.index))

	res.WriteLine(a.r.String())
	res.WriteLine("&&")

	res.WriteLine(fmt.Sprintf("and_%d_end:", a.index))

	return res.String()
}

type AstOr struct {
	index int
	l     AstStatement
	r     AstStatement
}

func (a *AstOr) String() string {
	res := Lines{}

	res.WriteLine(a.l.String())
	res.WriteLine("dup")
	res.WriteLine(fmt.Sprintf("bnz or_%d_end", a.index))

	res.WriteLine(a.r.String())
	res.WriteLine("||")

	res.WriteLine(fmt.Sprintf("or_%d_end:", a.index))

	return res.String()
}

type AstBinop struct {
	l  AstStatement
	op string
	r  AstStatement
}

func (a *AstBinop) String() string {
	return fmt.Sprintf("%s\n%s\n%s", a.l.String(), a.r.String(), a.op)
}

type AstMinusOp struct {
	value AstStatement
}

func (a *AstMinusOp) String() string {
	return fmt.Sprintf("int 0\n%s\n-", a.value.String())
}

type AstDefine struct {
	v *Variable
	t *Type

	value AstStatement
}

func (a *AstDefine) String() string {
	if a.t.complex != nil {
		panic("defining complex variable is not supported yet")
	}

	ast := avm_store_Ast{
		s1: a.value,
		I1: itoa(a.v.local.slot),
	}

	return ast.String()
}

type AstAssign struct {
	v   *Variable
	t   *Type
	f   *StructField
	fun *Function

	value AstStatement

	stmt bool
}

func (a *AstAssign) ToStmt() {
	a.stmt = true
}

func (a *AstAssign) String() string {
	if a.v.constant {
		panic("cannot assign to a const var")
	}

	if a.v.param != nil {
		// TODO: add param var assignment support
		panic("cannot assign param var")
	}

	res := Lines{}

	if a.t.complex != nil {
		if a.t.complex.builtin != nil {
			return fmt.Sprintf("%s %s", a.fun.builtin.op, a.f.name)
		} else {
			if a.v.param != nil {
				panic("accessing struct param fields is not supported yet")
			}

			v := a.v.fields[a.f.name]
			ast := avm_store_Ast{
				s1: a.value,
				I1: itoa(v.local.slot),
			}

			res.WriteLine(ast.String())
		}
	} else {
		ast := avm_store_Ast{
			s1: a.value,
			I1: itoa(a.v.local.slot),
		}

		res.WriteLine(ast.String())
	}

	if !a.stmt {
		load := avm_load_Ast{
			I1: itoa(a.v.local.slot),
		}

		res.WriteLine(load.String())
	}

	return res.String()
}

type AstStructField struct {
	v   *Variable
	t   *Type
	f   *StructField
	fun *Function
}

func (a *AstStructField) String() string {
	if a.t.complex.builtin != nil {
		return fmt.Sprintf("%s %s", a.fun.builtin.op, a.f.name)
	}

	if a.v.param != nil {
		panic("accessing struct param fields is not supported yet")
	}

	v := a.v.fields[a.f.name]

	ast := avm_load_Ast{
		I1: itoa(v.local.slot),
	}

	return ast.String()
}

type AstCall struct {
	fun  *Function
	args []AstStatement
}

func (a *AstCall) String() string {
	s := Lines{}

	if a.fun.builtin != nil {
		i := 0

		for ; i < len(a.fun.builtin.stack); i++ {
			arg := a.args[i]
			s.WriteLine(arg.String())
		}

		args := []string{}
		args = append(args, a.fun.builtin.op)

		for ; i < len(a.fun.builtin.stack)+len(a.fun.builtin.imm); i++ {
			arg := a.args[i]
			args = append(args, arg.String())
		}

		s.WriteLine(strings.Join(args, " "))
	}

	if a.fun.user != nil {
		for _, arg := range a.args {
			s.WriteLine(arg.String())
		}

		s.WriteLine(fmt.Sprintf("callsub %s", a.fun.name))
	}

	return s.String()
}

type AstIsReturn interface {
	IsReturn()
}

type AstIsBreak interface {
	IsBreak()
}

type AstIntConstant struct {
	value string
}

func (a *AstIntConstant) String() string {
	return fmt.Sprintf("int %s", a.value)
}

type AstByteConstant struct {
	value string
}

func (a *AstByteConstant) String() string {
	return fmt.Sprintf("byte %s", a.value)
}

type AstReturn struct {
	value    AstStatement
	function *Function
}

func (a *AstReturn) IsReturn() {
}

func (a *AstReturn) String() string {
	s := Lines{}

	if a.value != nil {
		s.WriteLine(a.value.String())
	}

	if a.function != nil {
		if a.function.user.sub {
			s.WriteLine("retsub")
		} else {
			s.WriteLine("return")
		}
	} else {
		s.WriteLine("return")
	}

	return s.String()
}

type AstBlock struct {
	statements []AstStatement
}

func (a *AstBlock) String() string {
	ops := Lines{}

	for _, stmt := range a.statements {
		ops.WriteLine(stmt.String())
	}

	return ops.String()
}

type AstIfAlternative struct {
	condition  AstStatement
	statements []AstStatement
}

type AstIf struct {
	index        int
	alternatives []*AstIfAlternative
}

func (a *AstIf) String() string {
	res := Lines{}

	for i, alt := range a.alternatives {
		if alt.condition != nil {
			res.WriteLine(alt.condition.String())

			if i < len(a.alternatives)-1 {
				res.WriteLine(fmt.Sprintf("bz if_skip_%d_%d", a.index, i))
			} else {
				res.WriteLine(fmt.Sprintf("bz if_end_%d", a.index))
			}
		}

		for _, stmt := range alt.statements {
			res.WriteLine(stmt.String())
		}

		if i < len(a.alternatives)-1 {
			res.WriteLine(fmt.Sprintf("b if_end_%d", a.index))
		}

		if alt.condition != nil {
			if i < len(a.alternatives)-1 {
				res.WriteLine(fmt.Sprintf("if_skip_%d_%d:", a.index, i))
			}
		}
	}

	res.WriteLine(fmt.Sprintf("if_end_%d:", a.index))

	return res.String()
}

type AstFunction struct {
	fun        *Function
	statements []AstStatement
}

func (a *AstFunction) String() string {
	res := Lines{}

	if a.fun.user.sub {
		if a.fun.user.args != 0 || a.fun.user.returns != 0 {
			ast := avm_proto_Ast{
				A1: itoa(a.fun.user.args),
				R2: itoa(a.fun.user.returns),
			}

			res.WriteLine(ast.String())
		}
	}

	for _, stmt := range a.statements {
		res.WriteLine(stmt.String())
	}

	if a.fun.user.sub {
		if len(a.statements) > 0 {
			last := a.statements[len(a.statements)-1]
			if _, ok := last.(AstIsReturn); !ok {
				res.WriteLine("retsub")
			}
		}
	}

	return res.String()
}

type AstRaw struct {
	value string
}

func (a *AstRaw) String() string {
	return a.value
}
