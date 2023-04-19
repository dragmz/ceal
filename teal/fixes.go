package teal

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type TealOp interface {
	String() string
}

type Teal []TealOp

func (t Teal) String() string {
	res := Source{}
	for _, op := range t {
		res.WriteLine(op.String())
	}
	return res.String()
}

type TealAstBuilder struct {
	items []TealAst
}

func (b *TealAstBuilder) Write(ast TealAst) {
	b.items = append(b.items, ast)
}

func (b *TealAstBuilder) Build() TealAst {
	return Teal_seq(b.items)
}

type TealBuilder struct {
	ops []TealOp
}

func (b *TealBuilder) Write(ops ...TealOp) {
	b.ops = append(b.ops, ops...)
}

type Teal_pragma_version struct {
	Version int
}

func (a *Teal_pragma_version) Teal() Teal {
	return Teal{a}
}

func (t *Teal_pragma_version) String() string {
	return fmt.Sprintf("#pragma version %d", t.Version)
}

type Teal_label struct {
	Name string
}

func (a *Teal_label) Teal() Teal {
	return Teal{a}
}

func (a *Teal_label) String() string {
	return fmt.Sprintf("%s:", a.Name)
}

type Teal_byte struct {
	S TealAst
}

func (a *Teal_byte) Teal() Teal {
	return Teal{a}
}

func (a *Teal_byte) String() string {
	return fmt.Sprintf("byte %s", a.S)
}

type Teal_byte_value struct {
	V []byte
}

func (a *Teal_byte_value) Teal() Teal {
	return Teal{a}
}

func (a *Teal_byte_value) String() string {
	return fmt.Sprintf("b64 %s", base64.StdEncoding.EncodeToString(a.V))
}

type Teal_literal struct {
	V string
}

func (a *Teal_literal) Teal() Teal {
	return Teal{a}
}

func (a *Teal_literal) String() string {
	return a.V
}

type Teal_byte_string_value struct {
	V string
}

func (a *Teal_byte_string_value) Teal() Teal {
	return Teal{a}
}

func (a *Teal_byte_string_value) String() string {
	return fmt.Sprintf("\"%s\"", a.V)
}

type Teal_named_int_value struct {
	V string
}

func (a *Teal_named_int_value) Teal() Teal {
	return Teal{a}
}

func (a *Teal_named_int_value) String() string {
	return a.V
}

type Teal_named_int struct {
	V TealAst
}

func (a *Teal_named_int) Teal() Teal {
	return Teal{a}
}

func (a *Teal_named_int) String() string {
	return fmt.Sprintf("int %s", a.V.Teal().String())

}

type Teal_int struct {
	V uint64
}

func (a *Teal_int) Teal() Teal {
	return Teal{a}
}

func (a *Teal_int) String() string {
	return fmt.Sprintf("int %d", a.V)
}

type Teal_retsub_fixed struct {
	Values []TealAst
}

func (a *Teal_retsub_fixed) Teal() Teal {
	return Teal{a}
}

func (a *Teal_retsub_fixed) String() string {
	res := Source{}

	for _, v := range a.Values {
		for _, op := range v.Teal() {
			res.WriteLine(op.String())
		}
	}

	res.WriteLine("retsub")

	return res.String()
}

type Teal_return_fixed struct {
	Value TealAst
}

func (a *Teal_return_fixed) Teal() Teal {
	return Teal{a}
}

func (a *Teal_return_fixed) String() string {
	res := Source{}

	if a.Value != nil {
		for _, op := range a.Value.Teal() {
			res.WriteLine(op.String())
		}
	}

	res.WriteLine("return")

	return res.String()
}

type Teal_callsub_fixed struct {
	Args   []TealAst
	Target string
}

func (a *Teal_callsub_fixed) Teal() Teal {
	return Teal{a}
}

func (a *Teal_callsub_fixed) String() string {
	res := Source{}

	for _, a := range a.Args {
		for _, op := range a.Teal() {
			res.WriteLine(op.String())
		}
	}

	res.WriteLine(fmt.Sprintf("callsub %s", a.Target))

	return res.String()
}

type Teal_call_builtin struct {
	Name string

	Args []TealAst
	Imms []TealAst
}

func (a *Teal_call_builtin) Teal() Teal {
	res := TealBuilder{}

	for _, arg := range a.Args {
		for _, op := range arg.Teal() {
			res.Write(op)
		}
	}

	res.Write(a)

	return res.ops
}

func (a *Teal_call_builtin) String() string {
	res := Source{}

	call := []string{
		a.Name,
	}

	for _, i := range a.Imms {
		for _, op := range i.Teal() {
			call = append(call, op.String())
		}
	}

	res.WriteLine(strings.Join(call, " "))

	return res.String()
}

type Teal_builtin struct {
	Op    string
	Field string
}

type Teal_seq []TealAst

func (a Teal_seq) Teal() Teal {
	res := TealBuilder{}
	for _, ast := range a {
		res.Write(ast.Teal()...)
	}
	return res.ops
}

type Teal_comment struct {
	Lines []string
}

func (a *Teal_comment) Teal() Teal {
	return Teal{a}
}

func (a *Teal_comment) String() string {
	var lines []string

	for _, line := range a.Lines {
		lines = append(lines, fmt.Sprintf("//%s", line))
	}

	return strings.Join(lines, "\n")
}

type Teal_raw struct {
	V string
}

func (a *Teal_raw) Teal() Teal {
	return ParseTeal(a.V)
}

type Teal_int_op struct {
	V int
}

func Parse_Teal_int_op(a ParserContext) TealOp {
	v := a.Read()
	return &Teal_named_int{V: &Teal_named_int_value{V: v.String()}}
}

func Parse_Teal_byte_op(a ParserContext) TealOp {
	v := a.Read_rbyte()
	return &Teal_byte{S: &Teal_byte_value{V: v}}
}

func Parse_Teal_callsub_op_fixed(a ParserContext) TealOp {
	v := a.Read()
	// TODO: should actually return Teal_callsub_op
	return &Teal_callsub_fixed{Target: v.String()}
}
