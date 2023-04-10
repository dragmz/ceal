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

type AstVariable struct {
	s *Stack
	v *Variable
}

func (a *AstVariable) String() string {
	if a.v.local != nil {
		ast := avm_load_Ast{i1: itoa(a.v.local.slot + a.s.current.slot)}
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
	s *Stack

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
				i1: itoa(v.local.slot + a.s.current.slot),
			}

			return ast.String()
		}
	} else {
		ast := avm_store_Ast{
			s1: a.value,
			i1: itoa(a.v.local.slot + a.s.current.slot),
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
	s          *Stack
	statements []AstStatement
}

func (a *AstBlock) String() string {
	ops := strings.Builder{}

	a.s.push()

	for _, stmt := range a.statements {
		ops.WriteString(stmt.String())
		ops.WriteString("\n")
	}

	a.s.pop()

	return ops.String()
}

type AstIfAlternative struct {
	condition  AstStatement
	statements []AstStatement
}

type AstIf struct {
	s            *Stack
	label        int
	alternatives []*AstIfAlternative
}

func (a *AstIf) String() string {
	res := strings.Builder{}

	for i, alt := range a.alternatives {
		if alt.condition != nil {
			res.WriteString(alt.condition.String())
			res.WriteString("\n")

			if i < len(a.alternatives)-1 {
				res.WriteString(fmt.Sprintf("bz skip_%d_%d\n", a.label, i))
			} else {
				res.WriteString(fmt.Sprintf("bz end_%d\n", a.label))
			}
		}

		a.s.push()

		for _, stmt := range alt.statements {
			res.WriteString(stmt.String())
			res.WriteString("\n")
		}

		a.s.pop()

		if i < len(a.alternatives)-1 {
			res.WriteString(fmt.Sprintf("b end_%d\n", a.label))
		}

		if alt.condition != nil {
			if i < len(a.alternatives)-1 {
				res.WriteString(fmt.Sprintf("skip_%d_%d:\n", a.label, i))
			}
		}
	}

	res.WriteString(fmt.Sprintf("end_%d:", a.label))

	return res.String()
}

type AstFunction struct {
	s          *Stack
	fun        *Function
	statements []AstStatement
}

func (a *AstFunction) String() string {
	res := strings.Builder{}

	res.WriteString(fmt.Sprintf("// function %s\n", a.fun.name))
	res.WriteString(fmt.Sprintf("%s:\n", a.fun.name))

	if a.fun.user.sub {
		ast := avm_proto_Ast{
			i1: itoa(a.fun.user.args),
			i2: itoa(a.fun.user.returns),
		}

		res.WriteString(ast.String())
		res.WriteString("\n")
	}

	a.s.push()

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

	a.s.pop()

	return res.String()
}

type AstRaw struct {
	value string
}

func (a *AstRaw) String() string {
	return a.value
}
