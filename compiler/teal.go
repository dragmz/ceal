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
	res := strings.Builder{}

	labels := []string{}

	for i, c := range a.cases {
		label := fmt.Sprintf("switch_%d_%d", a.index, i)
		labels = append(labels, label)

		res.WriteString(fmt.Sprintf("%s\n", c.value.String()))
	}

	res.WriteString(fmt.Sprintf("%s\n", a.value.String()))
	res.WriteString(fmt.Sprintf("match %s\n", strings.Join(labels, " ")))

	if len(a.default_) > 0 {
		res.WriteString("// default\n")

		for _, stmt := range a.default_ {
			res.WriteString(fmt.Sprintf("%s\n", stmt.String()))
		}
	}

	for i, c := range a.cases {
		label := labels[i]

		res.WriteString(fmt.Sprintf("%s:\n", label))

		for _, stmt := range c.statements {
			res.WriteString(fmt.Sprintf("%s\n", stmt.String()))
		}
	}

	res.WriteString(fmt.Sprintf("switch_%d_end:\n", a.index))

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
	res := strings.Builder{}
	res.WriteString(fmt.Sprintf("do_%d:\n", a.index))
	res.WriteString(fmt.Sprintf("%s\n", a.s.String()))
	res.WriteString(fmt.Sprintf("%s\n", a.condition.String()))
	res.WriteString(fmt.Sprintf("bnz do_%d", a.index))

	return res.String()
}

type AstWhile struct {
	index     int
	condition AstStatement
	s         AstStatement
}

func (a *AstWhile) String() string {
	res := strings.Builder{}

	res.WriteString(fmt.Sprintf("while_%d:\n", a.index))
	res.WriteString(a.condition.String())
	res.WriteString("\n")
	res.WriteString(fmt.Sprintf("bz while_%d_end\n", a.index))
	res.WriteString(a.s.String())
	res.WriteString("\n")
	res.WriteString(fmt.Sprintf("b while_%d\n", a.index))
	res.WriteString(fmt.Sprintf("while_%d_end:", a.index))

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
	res := strings.Builder{}

	for _, stmt := range a.init {
		res.WriteString(stmt.String())
		res.WriteString("\n")
	}

	res.WriteString(fmt.Sprintf("for_%d:\n", a.index))
	res.WriteString(fmt.Sprintf("%s\n", a.condition.String()))
	res.WriteString(fmt.Sprintf("bz for_%d_end\n", a.index))
	res.WriteString(fmt.Sprintf("%s\n", a.s.String()))

	for _, stmt := range a.iter {
		res.WriteString(fmt.Sprintf("%s\n", stmt.String()))
	}

	res.WriteString(fmt.Sprintf("b for_%d\n", a.index))
	res.WriteString(fmt.Sprintf("for_%d_end:", a.index))

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
	var op string

	switch a.op {
	case "++":
		op = "+"
	case "--":
		op = "-"
	default:
		panic(fmt.Sprintf("prefix operator not supported: '%s'", a.op))
	}

	res := strings.Builder{}

	res.WriteString(fmt.Sprintf("load %d\n", a.v.local.slot))
	res.WriteString("int 1\n")
	res.WriteString(fmt.Sprintf("%s\n", op))
	res.WriteString(fmt.Sprintf("store %d", a.v.local.slot))

	if !a.stmt {
		res.WriteString(fmt.Sprintf("\nload %d\n", a.v.local.slot))
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
	var op string

	switch a.op {
	case "++":
		op = "+"
	case "--":
		op = "-"
	default:
		panic(fmt.Sprintf("postfix operator not supported: '%s'", a.op))
	}

	res := strings.Builder{}
	res.WriteString(fmt.Sprintf("load %d\n", a.v.local.slot))
	if !a.stmt {
		res.WriteString("dup\n")
	}
	res.WriteString(fmt.Sprintf("int 1\n%s\nstore %d", op, a.v.local.slot))

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
		ast := avm_load_Ast{i1: itoa(a.v.local.slot)}
		return ast.String()
	}

	if a.v.param != nil {
		ast := avm_frame_dig_Ast{i1: itoa(a.v.param.index)}
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
	res := strings.Builder{}

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

	res.WriteString(fmt.Sprintf("load %d\n", slot))
	res.WriteString(fmt.Sprintf("%s\n", a.value.String()))
	res.WriteString(fmt.Sprintf("%s\n", op))
	if !a.stmt {
		res.WriteString("dup\n")
	}
	res.WriteString(fmt.Sprintf("store %d", slot))

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

type AstAssign struct {
	slot int
	v    *Variable
	t    *Type
	f    *StructField
	fun  *Function

	value AstStatement
}

func (a *AstAssign) String() string {
	if a.v.param != nil {
		// TODO: add param var assignment support
		panic("cannot assign param var")
	}

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
				i1: itoa(v.local.slot),
			}

			return ast.String()
		}
	} else {
		ast := avm_store_Ast{
			s1: a.value,
			i1: itoa(a.v.local.slot),
		}

		return ast.String()
	}
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
		i1: itoa(v.local.slot),
	}

	return ast.String()
}

type AstCall struct {
	fun  *Function
	args []AstStatement
}

func (a *AstCall) String() string {
	s := strings.Builder{}

	if a.fun.builtin != nil {
		i := 0

		for ; i < len(a.fun.builtin.stack); i++ {
			arg := a.args[i]
			s.WriteString(arg.String())
			s.WriteString("\n")
		}

		s.WriteString(a.fun.builtin.op)

		for ; i < len(a.fun.builtin.stack)+len(a.fun.builtin.imm); i++ {
			arg := a.args[i]
			s.WriteString(" ")
			s.WriteString(arg.String())
		}
	}

	if a.fun.user != nil {
		for _, arg := range a.args {
			s.WriteString(arg.String())
			s.WriteString("\n")
		}

		s.WriteString(fmt.Sprintf("callsub %s", a.fun.name))
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
	s := strings.Builder{}

	if a.value != nil {
		s.WriteString(a.value.String())
		s.WriteString("\n")
	}

	if a.function != nil {
		if a.function.user.sub {
			s.WriteString("retsub")
		} else {
			s.WriteString("return")
		}
	} else {
		s.WriteString("return")
	}

	return s.String()
}

type AstBlock struct {
	statements []AstStatement
}

func (a *AstBlock) String() string {
	ops := strings.Builder{}

	for _, stmt := range a.statements {
		ops.WriteString(stmt.String())
		ops.WriteString("\n")
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
	res := strings.Builder{}

	for i, alt := range a.alternatives {
		if alt.condition != nil {
			res.WriteString(alt.condition.String())
			res.WriteString("\n")

			if i < len(a.alternatives)-1 {
				res.WriteString(fmt.Sprintf("bz if_skip_%d_%d\n", a.index, i))
			} else {
				res.WriteString(fmt.Sprintf("bz if_end_%d\n", a.index))
			}
		}

		for _, stmt := range alt.statements {
			res.WriteString(stmt.String())
			res.WriteString("\n")
		}

		if i < len(a.alternatives)-1 {
			res.WriteString(fmt.Sprintf("b if_end_%d\n", a.index))
		}

		if alt.condition != nil {
			if i < len(a.alternatives)-1 {
				res.WriteString(fmt.Sprintf("if_skip_%d_%d:\n", a.index, i))
			}
		}
	}

	res.WriteString(fmt.Sprintf("if_end_%d:", a.index))

	return res.String()
}

type AstFunction struct {
	fun        *Function
	statements []AstStatement
}

func (a *AstFunction) String() string {
	res := strings.Builder{}

	if a.fun.user.sub {
		if a.fun.user.args != 0 || a.fun.user.returns != 0 {
			ast := avm_proto_Ast{
				i1: itoa(a.fun.user.args),
				i2: itoa(a.fun.user.returns),
			}

			res.WriteString(ast.String())
			res.WriteString("\n")
		}
	}

	for _, stmt := range a.statements {
		res.WriteString(stmt.String())
		res.WriteString("\n")
	}

	if a.fun.user.sub {
		if len(a.statements) > 0 {
			last := a.statements[len(a.statements)-1]
			if _, ok := last.(AstIsReturn); !ok {
				res.WriteString("retsub")
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
