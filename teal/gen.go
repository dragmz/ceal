package teal

import (
	"fmt"
	"strings"
)

type TealAst interface {
	Teal() Teal
}

type Teal_err struct {
}

func (a *Teal_err) Teal() Teal {
	return Teal{a}
}
func (a *Teal_err) String() string {
	res := strings.Builder{}
	res.WriteString("err")
	return res.String()
}

type Teal_err_op struct {
}

func Parse_Teal_err_op(ctx ParserContext) *Teal_err_op {
	t := &Teal_err_op{}
	return t
}

func (a *Teal_err_op) String() string {
	res := strings.Builder{}
	res.WriteString("err")
	return res.String()
}

type Teal_sha256 struct {
	STACK_1 TealAst
}

func (a *Teal_sha256) Teal() Teal {
	return Teal{a}
}
func (a *Teal_sha256) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("sha256")
	return res.String()
}

type Teal_sha256_op struct {
}

func Parse_Teal_sha256_op(ctx ParserContext) *Teal_sha256_op {
	t := &Teal_sha256_op{}
	return t
}

func (a *Teal_sha256_op) String() string {
	res := strings.Builder{}
	res.WriteString("sha256")
	return res.String()
}

type Teal_keccak256 struct {
	STACK_1 TealAst
}

func (a *Teal_keccak256) Teal() Teal {
	return Teal{a}
}
func (a *Teal_keccak256) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("keccak256")
	return res.String()
}

type Teal_keccak256_op struct {
}

func Parse_Teal_keccak256_op(ctx ParserContext) *Teal_keccak256_op {
	t := &Teal_keccak256_op{}
	return t
}

func (a *Teal_keccak256_op) String() string {
	res := strings.Builder{}
	res.WriteString("keccak256")
	return res.String()
}

type Teal_sha512_256 struct {
	STACK_1 TealAst
}

func (a *Teal_sha512_256) Teal() Teal {
	return Teal{a}
}
func (a *Teal_sha512_256) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("sha512_256")
	return res.String()
}

type Teal_sha512_256_op struct {
}

func Parse_Teal_sha512_256_op(ctx ParserContext) *Teal_sha512_256_op {
	t := &Teal_sha512_256_op{}
	return t
}

func (a *Teal_sha512_256_op) String() string {
	res := strings.Builder{}
	res.WriteString("sha512_256")
	return res.String()
}

type Teal_ed25519verify struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_ed25519verify) Teal() Teal {
	return Teal{a}
}
func (a *Teal_ed25519verify) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("ed25519verify")
	return res.String()
}

type Teal_ed25519verify_op struct {
}

func Parse_Teal_ed25519verify_op(ctx ParserContext) *Teal_ed25519verify_op {
	t := &Teal_ed25519verify_op{}
	return t
}

func (a *Teal_ed25519verify_op) String() string {
	res := strings.Builder{}
	res.WriteString("ed25519verify")
	return res.String()
}

type Teal_ecdsa_verify struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
	STACK_4 TealAst
	STACK_5 TealAst
	V1      uint8
}

func (a *Teal_ecdsa_verify) Teal() Teal {
	return Teal{a}
}
func (a *Teal_ecdsa_verify) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_4.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_5.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("ecdsa_verify")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.V1))
	return res.String()
}

type Teal_ecdsa_verify_op struct {
	V1 uint8
}

func Parse_Teal_ecdsa_verify_op(ctx ParserContext) *Teal_ecdsa_verify_op {
	t := &Teal_ecdsa_verify_op{
		V1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_ecdsa_verify_op) String() string {
	res := strings.Builder{}
	res.WriteString("ecdsa_verify")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.V1))
	return res.String()
}

type Teal_ecdsa_pk_decompress struct {
	STACK_1 TealAst
	V1      uint8
}

func (a *Teal_ecdsa_pk_decompress) Teal() Teal {
	return Teal{a}
}
func (a *Teal_ecdsa_pk_decompress) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("ecdsa_pk_decompress")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.V1))
	return res.String()
}

type Teal_ecdsa_pk_decompress_op struct {
	V1 uint8
}

func Parse_Teal_ecdsa_pk_decompress_op(ctx ParserContext) *Teal_ecdsa_pk_decompress_op {
	t := &Teal_ecdsa_pk_decompress_op{
		V1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_ecdsa_pk_decompress_op) String() string {
	res := strings.Builder{}
	res.WriteString("ecdsa_pk_decompress")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.V1))
	return res.String()
}

type Teal_ecdsa_pk_recover struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
	STACK_4 TealAst
	V1      uint8
}

func (a *Teal_ecdsa_pk_recover) Teal() Teal {
	return Teal{a}
}
func (a *Teal_ecdsa_pk_recover) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_4.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("ecdsa_pk_recover")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.V1))
	return res.String()
}

type Teal_ecdsa_pk_recover_op struct {
	V1 uint8
}

func Parse_Teal_ecdsa_pk_recover_op(ctx ParserContext) *Teal_ecdsa_pk_recover_op {
	t := &Teal_ecdsa_pk_recover_op{
		V1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_ecdsa_pk_recover_op) String() string {
	res := strings.Builder{}
	res.WriteString("ecdsa_pk_recover")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.V1))
	return res.String()
}

type Teal_plus struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_plus) Teal() Teal {
	return Teal{a}
}
func (a *Teal_plus) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("+")
	return res.String()
}

type Teal_plus_op struct {
}

func Parse_Teal_plus_op(ctx ParserContext) *Teal_plus_op {
	t := &Teal_plus_op{}
	return t
}

func (a *Teal_plus_op) String() string {
	res := strings.Builder{}
	res.WriteString("+")
	return res.String()
}

type Teal_minus struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_minus) Teal() Teal {
	return Teal{a}
}
func (a *Teal_minus) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("-")
	return res.String()
}

type Teal_minus_op struct {
}

func Parse_Teal_minus_op(ctx ParserContext) *Teal_minus_op {
	t := &Teal_minus_op{}
	return t
}

func (a *Teal_minus_op) String() string {
	res := strings.Builder{}
	res.WriteString("-")
	return res.String()
}

type Teal_div struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_div) Teal() Teal {
	return Teal{a}
}
func (a *Teal_div) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("/")
	return res.String()
}

type Teal_div_op struct {
}

func Parse_Teal_div_op(ctx ParserContext) *Teal_div_op {
	t := &Teal_div_op{}
	return t
}

func (a *Teal_div_op) String() string {
	res := strings.Builder{}
	res.WriteString("/")
	return res.String()
}

type Teal_mul struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_mul) Teal() Teal {
	return Teal{a}
}
func (a *Teal_mul) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("*")
	return res.String()
}

type Teal_mul_op struct {
}

func Parse_Teal_mul_op(ctx ParserContext) *Teal_mul_op {
	t := &Teal_mul_op{}
	return t
}

func (a *Teal_mul_op) String() string {
	res := strings.Builder{}
	res.WriteString("*")
	return res.String()
}

type Teal_lt struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_lt) Teal() Teal {
	return Teal{a}
}
func (a *Teal_lt) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("<")
	return res.String()
}

type Teal_lt_op struct {
}

func Parse_Teal_lt_op(ctx ParserContext) *Teal_lt_op {
	t := &Teal_lt_op{}
	return t
}

func (a *Teal_lt_op) String() string {
	res := strings.Builder{}
	res.WriteString("<")
	return res.String()
}

type Teal_gt struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_gt) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gt) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString(">")
	return res.String()
}

type Teal_gt_op struct {
}

func Parse_Teal_gt_op(ctx ParserContext) *Teal_gt_op {
	t := &Teal_gt_op{}
	return t
}

func (a *Teal_gt_op) String() string {
	res := strings.Builder{}
	res.WriteString(">")
	return res.String()
}

type Teal_lteq struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_lteq) Teal() Teal {
	return Teal{a}
}
func (a *Teal_lteq) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("<=")
	return res.String()
}

type Teal_lteq_op struct {
}

func Parse_Teal_lteq_op(ctx ParserContext) *Teal_lteq_op {
	t := &Teal_lteq_op{}
	return t
}

func (a *Teal_lteq_op) String() string {
	res := strings.Builder{}
	res.WriteString("<=")
	return res.String()
}

type Teal_gteq struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_gteq) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gteq) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString(">=")
	return res.String()
}

type Teal_gteq_op struct {
}

func Parse_Teal_gteq_op(ctx ParserContext) *Teal_gteq_op {
	t := &Teal_gteq_op{}
	return t
}

func (a *Teal_gteq_op) String() string {
	res := strings.Builder{}
	res.WriteString(">=")
	return res.String()
}

type Teal_andand struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_andand) Teal() Teal {
	return Teal{a}
}
func (a *Teal_andand) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("&&")
	return res.String()
}

type Teal_andand_op struct {
}

func Parse_Teal_andand_op(ctx ParserContext) *Teal_andand_op {
	t := &Teal_andand_op{}
	return t
}

func (a *Teal_andand_op) String() string {
	res := strings.Builder{}
	res.WriteString("&&")
	return res.String()
}

type Teal_oror struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_oror) Teal() Teal {
	return Teal{a}
}
func (a *Teal_oror) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("||")
	return res.String()
}

type Teal_oror_op struct {
}

func Parse_Teal_oror_op(ctx ParserContext) *Teal_oror_op {
	t := &Teal_oror_op{}
	return t
}

func (a *Teal_oror_op) String() string {
	res := strings.Builder{}
	res.WriteString("||")
	return res.String()
}

type Teal_eqeq struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_eqeq) Teal() Teal {
	return Teal{a}
}
func (a *Teal_eqeq) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("==")
	return res.String()
}

type Teal_eqeq_op struct {
}

func Parse_Teal_eqeq_op(ctx ParserContext) *Teal_eqeq_op {
	t := &Teal_eqeq_op{}
	return t
}

func (a *Teal_eqeq_op) String() string {
	res := strings.Builder{}
	res.WriteString("==")
	return res.String()
}

type Teal_noteq struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_noteq) Teal() Teal {
	return Teal{a}
}
func (a *Teal_noteq) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("!=")
	return res.String()
}

type Teal_noteq_op struct {
}

func Parse_Teal_noteq_op(ctx ParserContext) *Teal_noteq_op {
	t := &Teal_noteq_op{}
	return t
}

func (a *Teal_noteq_op) String() string {
	res := strings.Builder{}
	res.WriteString("!=")
	return res.String()
}

type Teal_not struct {
	STACK_1 TealAst
}

func (a *Teal_not) Teal() Teal {
	return Teal{a}
}
func (a *Teal_not) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("!")
	return res.String()
}

type Teal_not_op struct {
}

func Parse_Teal_not_op(ctx ParserContext) *Teal_not_op {
	t := &Teal_not_op{}
	return t
}

func (a *Teal_not_op) String() string {
	res := strings.Builder{}
	res.WriteString("!")
	return res.String()
}

type Teal_len struct {
	STACK_1 TealAst
}

func (a *Teal_len) Teal() Teal {
	return Teal{a}
}
func (a *Teal_len) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("len")
	return res.String()
}

type Teal_len_op struct {
}

func Parse_Teal_len_op(ctx ParserContext) *Teal_len_op {
	t := &Teal_len_op{}
	return t
}

func (a *Teal_len_op) String() string {
	res := strings.Builder{}
	res.WriteString("len")
	return res.String()
}

type Teal_itob struct {
	STACK_1 TealAst
}

func (a *Teal_itob) Teal() Teal {
	return Teal{a}
}
func (a *Teal_itob) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("itob")
	return res.String()
}

type Teal_itob_op struct {
}

func Parse_Teal_itob_op(ctx ParserContext) *Teal_itob_op {
	t := &Teal_itob_op{}
	return t
}

func (a *Teal_itob_op) String() string {
	res := strings.Builder{}
	res.WriteString("itob")
	return res.String()
}

type Teal_btoi struct {
	STACK_1 TealAst
}

func (a *Teal_btoi) Teal() Teal {
	return Teal{a}
}
func (a *Teal_btoi) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("btoi")
	return res.String()
}

type Teal_btoi_op struct {
}

func Parse_Teal_btoi_op(ctx ParserContext) *Teal_btoi_op {
	t := &Teal_btoi_op{}
	return t
}

func (a *Teal_btoi_op) String() string {
	res := strings.Builder{}
	res.WriteString("btoi")
	return res.String()
}

type Teal_mod struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_mod) Teal() Teal {
	return Teal{a}
}
func (a *Teal_mod) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("%")
	return res.String()
}

type Teal_mod_op struct {
}

func Parse_Teal_mod_op(ctx ParserContext) *Teal_mod_op {
	t := &Teal_mod_op{}
	return t
}

func (a *Teal_mod_op) String() string {
	res := strings.Builder{}
	res.WriteString("%")
	return res.String()
}

type Teal_or struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_or) Teal() Teal {
	return Teal{a}
}
func (a *Teal_or) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("|")
	return res.String()
}

type Teal_or_op struct {
}

func Parse_Teal_or_op(ctx ParserContext) *Teal_or_op {
	t := &Teal_or_op{}
	return t
}

func (a *Teal_or_op) String() string {
	res := strings.Builder{}
	res.WriteString("|")
	return res.String()
}

type Teal_and struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_and) Teal() Teal {
	return Teal{a}
}
func (a *Teal_and) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("&")
	return res.String()
}

type Teal_and_op struct {
}

func Parse_Teal_and_op(ctx ParserContext) *Teal_and_op {
	t := &Teal_and_op{}
	return t
}

func (a *Teal_and_op) String() string {
	res := strings.Builder{}
	res.WriteString("&")
	return res.String()
}

type Teal_xor struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_xor) Teal() Teal {
	return Teal{a}
}
func (a *Teal_xor) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("^")
	return res.String()
}

type Teal_xor_op struct {
}

func Parse_Teal_xor_op(ctx ParserContext) *Teal_xor_op {
	t := &Teal_xor_op{}
	return t
}

func (a *Teal_xor_op) String() string {
	res := strings.Builder{}
	res.WriteString("^")
	return res.String()
}

type Teal_inv struct {
	STACK_1 TealAst
}

func (a *Teal_inv) Teal() Teal {
	return Teal{a}
}
func (a *Teal_inv) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("~")
	return res.String()
}

type Teal_inv_op struct {
}

func Parse_Teal_inv_op(ctx ParserContext) *Teal_inv_op {
	t := &Teal_inv_op{}
	return t
}

func (a *Teal_inv_op) String() string {
	res := strings.Builder{}
	res.WriteString("~")
	return res.String()
}

type Teal_mulw struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_mulw) Teal() Teal {
	return Teal{a}
}
func (a *Teal_mulw) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("mulw")
	return res.String()
}

type Teal_mulw_op struct {
}

func Parse_Teal_mulw_op(ctx ParserContext) *Teal_mulw_op {
	t := &Teal_mulw_op{}
	return t
}

func (a *Teal_mulw_op) String() string {
	res := strings.Builder{}
	res.WriteString("mulw")
	return res.String()
}

type Teal_addw struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_addw) Teal() Teal {
	return Teal{a}
}
func (a *Teal_addw) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("addw")
	return res.String()
}

type Teal_addw_op struct {
}

func Parse_Teal_addw_op(ctx ParserContext) *Teal_addw_op {
	t := &Teal_addw_op{}
	return t
}

func (a *Teal_addw_op) String() string {
	res := strings.Builder{}
	res.WriteString("addw")
	return res.String()
}

type Teal_divmodw struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
	STACK_4 TealAst
}

func (a *Teal_divmodw) Teal() Teal {
	return Teal{a}
}
func (a *Teal_divmodw) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_4.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("divmodw")
	return res.String()
}

type Teal_divmodw_op struct {
}

func Parse_Teal_divmodw_op(ctx ParserContext) *Teal_divmodw_op {
	t := &Teal_divmodw_op{}
	return t
}

func (a *Teal_divmodw_op) String() string {
	res := strings.Builder{}
	res.WriteString("divmodw")
	return res.String()
}

type Teal_intcblock struct {
	UINT1 [][]byte
}

func (a *Teal_intcblock) Teal() Teal {
	return Teal{a}
}
func (a *Teal_intcblock) String() string {
	res := strings.Builder{}
	res.WriteString("intcblock")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.UINT1))
	return res.String()
}

type Teal_intcblock_op struct {
	UINT1 [][]byte
}

func Parse_Teal_intcblock_op(ctx ParserContext) *Teal_intcblock_op {
	t := &Teal_intcblock_op{
		UINT1: ctx.Read_rrbyte(),
	}
	return t
}

func (a *Teal_intcblock_op) String() string {
	res := strings.Builder{}
	res.WriteString("intcblock")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.UINT1))
	return res.String()
}

type Teal_intc struct {
	I1 uint8
}

func (a *Teal_intc) Teal() Teal {
	return Teal{a}
}
func (a *Teal_intc) String() string {
	res := strings.Builder{}
	res.WriteString("intc")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_intc_op struct {
	I1 uint8
}

func Parse_Teal_intc_op(ctx ParserContext) *Teal_intc_op {
	t := &Teal_intc_op{
		I1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_intc_op) String() string {
	res := strings.Builder{}
	res.WriteString("intc")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_intc_0 struct {
}

func (a *Teal_intc_0) Teal() Teal {
	return Teal{a}
}
func (a *Teal_intc_0) String() string {
	res := strings.Builder{}
	res.WriteString("intc_0")
	return res.String()
}

type Teal_intc_0_op struct {
}

func Parse_Teal_intc_0_op(ctx ParserContext) *Teal_intc_0_op {
	t := &Teal_intc_0_op{}
	return t
}

func (a *Teal_intc_0_op) String() string {
	res := strings.Builder{}
	res.WriteString("intc_0")
	return res.String()
}

type Teal_intc_1 struct {
}

func (a *Teal_intc_1) Teal() Teal {
	return Teal{a}
}
func (a *Teal_intc_1) String() string {
	res := strings.Builder{}
	res.WriteString("intc_1")
	return res.String()
}

type Teal_intc_1_op struct {
}

func Parse_Teal_intc_1_op(ctx ParserContext) *Teal_intc_1_op {
	t := &Teal_intc_1_op{}
	return t
}

func (a *Teal_intc_1_op) String() string {
	res := strings.Builder{}
	res.WriteString("intc_1")
	return res.String()
}

type Teal_intc_2 struct {
}

func (a *Teal_intc_2) Teal() Teal {
	return Teal{a}
}
func (a *Teal_intc_2) String() string {
	res := strings.Builder{}
	res.WriteString("intc_2")
	return res.String()
}

type Teal_intc_2_op struct {
}

func Parse_Teal_intc_2_op(ctx ParserContext) *Teal_intc_2_op {
	t := &Teal_intc_2_op{}
	return t
}

func (a *Teal_intc_2_op) String() string {
	res := strings.Builder{}
	res.WriteString("intc_2")
	return res.String()
}

type Teal_intc_3 struct {
}

func (a *Teal_intc_3) Teal() Teal {
	return Teal{a}
}
func (a *Teal_intc_3) String() string {
	res := strings.Builder{}
	res.WriteString("intc_3")
	return res.String()
}

type Teal_intc_3_op struct {
}

func Parse_Teal_intc_3_op(ctx ParserContext) *Teal_intc_3_op {
	t := &Teal_intc_3_op{}
	return t
}

func (a *Teal_intc_3_op) String() string {
	res := strings.Builder{}
	res.WriteString("intc_3")
	return res.String()
}

type Teal_bytecblock struct {
	BYTES1 [][]byte
}

func (a *Teal_bytecblock) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bytecblock) String() string {
	res := strings.Builder{}
	res.WriteString("bytecblock")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.BYTES1))
	return res.String()
}

type Teal_bytecblock_op struct {
	BYTES1 [][]byte
}

func Parse_Teal_bytecblock_op(ctx ParserContext) *Teal_bytecblock_op {
	t := &Teal_bytecblock_op{
		BYTES1: ctx.Read_rrbyte(),
	}
	return t
}

func (a *Teal_bytecblock_op) String() string {
	res := strings.Builder{}
	res.WriteString("bytecblock")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.BYTES1))
	return res.String()
}

type Teal_bytec struct {
	I1 uint8
}

func (a *Teal_bytec) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bytec) String() string {
	res := strings.Builder{}
	res.WriteString("bytec")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_bytec_op struct {
	I1 uint8
}

func Parse_Teal_bytec_op(ctx ParserContext) *Teal_bytec_op {
	t := &Teal_bytec_op{
		I1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_bytec_op) String() string {
	res := strings.Builder{}
	res.WriteString("bytec")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_bytec_0 struct {
}

func (a *Teal_bytec_0) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bytec_0) String() string {
	res := strings.Builder{}
	res.WriteString("bytec_0")
	return res.String()
}

type Teal_bytec_0_op struct {
}

func Parse_Teal_bytec_0_op(ctx ParserContext) *Teal_bytec_0_op {
	t := &Teal_bytec_0_op{}
	return t
}

func (a *Teal_bytec_0_op) String() string {
	res := strings.Builder{}
	res.WriteString("bytec_0")
	return res.String()
}

type Teal_bytec_1 struct {
}

func (a *Teal_bytec_1) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bytec_1) String() string {
	res := strings.Builder{}
	res.WriteString("bytec_1")
	return res.String()
}

type Teal_bytec_1_op struct {
}

func Parse_Teal_bytec_1_op(ctx ParserContext) *Teal_bytec_1_op {
	t := &Teal_bytec_1_op{}
	return t
}

func (a *Teal_bytec_1_op) String() string {
	res := strings.Builder{}
	res.WriteString("bytec_1")
	return res.String()
}

type Teal_bytec_2 struct {
}

func (a *Teal_bytec_2) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bytec_2) String() string {
	res := strings.Builder{}
	res.WriteString("bytec_2")
	return res.String()
}

type Teal_bytec_2_op struct {
}

func Parse_Teal_bytec_2_op(ctx ParserContext) *Teal_bytec_2_op {
	t := &Teal_bytec_2_op{}
	return t
}

func (a *Teal_bytec_2_op) String() string {
	res := strings.Builder{}
	res.WriteString("bytec_2")
	return res.String()
}

type Teal_bytec_3 struct {
}

func (a *Teal_bytec_3) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bytec_3) String() string {
	res := strings.Builder{}
	res.WriteString("bytec_3")
	return res.String()
}

type Teal_bytec_3_op struct {
}

func Parse_Teal_bytec_3_op(ctx ParserContext) *Teal_bytec_3_op {
	t := &Teal_bytec_3_op{}
	return t
}

func (a *Teal_bytec_3_op) String() string {
	res := strings.Builder{}
	res.WriteString("bytec_3")
	return res.String()
}

type Teal_arg struct {
	N1 uint8
}

func (a *Teal_arg) Teal() Teal {
	return Teal{a}
}
func (a *Teal_arg) String() string {
	res := strings.Builder{}
	res.WriteString("arg")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_arg_op struct {
	N1 uint8
}

func Parse_Teal_arg_op(ctx ParserContext) *Teal_arg_op {
	t := &Teal_arg_op{
		N1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_arg_op) String() string {
	res := strings.Builder{}
	res.WriteString("arg")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_arg_0 struct {
}

func (a *Teal_arg_0) Teal() Teal {
	return Teal{a}
}
func (a *Teal_arg_0) String() string {
	res := strings.Builder{}
	res.WriteString("arg_0")
	return res.String()
}

type Teal_arg_0_op struct {
}

func Parse_Teal_arg_0_op(ctx ParserContext) *Teal_arg_0_op {
	t := &Teal_arg_0_op{}
	return t
}

func (a *Teal_arg_0_op) String() string {
	res := strings.Builder{}
	res.WriteString("arg_0")
	return res.String()
}

type Teal_arg_1 struct {
}

func (a *Teal_arg_1) Teal() Teal {
	return Teal{a}
}
func (a *Teal_arg_1) String() string {
	res := strings.Builder{}
	res.WriteString("arg_1")
	return res.String()
}

type Teal_arg_1_op struct {
}

func Parse_Teal_arg_1_op(ctx ParserContext) *Teal_arg_1_op {
	t := &Teal_arg_1_op{}
	return t
}

func (a *Teal_arg_1_op) String() string {
	res := strings.Builder{}
	res.WriteString("arg_1")
	return res.String()
}

type Teal_arg_2 struct {
}

func (a *Teal_arg_2) Teal() Teal {
	return Teal{a}
}
func (a *Teal_arg_2) String() string {
	res := strings.Builder{}
	res.WriteString("arg_2")
	return res.String()
}

type Teal_arg_2_op struct {
}

func Parse_Teal_arg_2_op(ctx ParserContext) *Teal_arg_2_op {
	t := &Teal_arg_2_op{}
	return t
}

func (a *Teal_arg_2_op) String() string {
	res := strings.Builder{}
	res.WriteString("arg_2")
	return res.String()
}

type Teal_arg_3 struct {
}

func (a *Teal_arg_3) Teal() Teal {
	return Teal{a}
}
func (a *Teal_arg_3) String() string {
	res := strings.Builder{}
	res.WriteString("arg_3")
	return res.String()
}

type Teal_arg_3_op struct {
}

func Parse_Teal_arg_3_op(ctx ParserContext) *Teal_arg_3_op {
	t := &Teal_arg_3_op{}
	return t
}

func (a *Teal_arg_3_op) String() string {
	res := strings.Builder{}
	res.WriteString("arg_3")
	return res.String()
}

type Teal_txn struct {
	F1 uint8
}

func (a *Teal_txn) Teal() Teal {
	return Teal{a}
}
func (a *Teal_txn) String() string {
	res := strings.Builder{}
	res.WriteString("txn")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_txn_op struct {
	F1 uint8
}

func Parse_Teal_txn_op(ctx ParserContext) *Teal_txn_op {
	t := &Teal_txn_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_txn_op) String() string {
	res := strings.Builder{}
	res.WriteString("txn")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_global struct {
	F1 uint8
}

func (a *Teal_global) Teal() Teal {
	return Teal{a}
}
func (a *Teal_global) String() string {
	res := strings.Builder{}
	res.WriteString("global")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_global_op struct {
	F1 uint8
}

func Parse_Teal_global_op(ctx ParserContext) *Teal_global_op {
	t := &Teal_global_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_global_op) String() string {
	res := strings.Builder{}
	res.WriteString("global")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_gtxn struct {
	T1 uint8
	F2 uint8
}

func (a *Teal_gtxn) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gtxn) String() string {
	res := strings.Builder{}
	res.WriteString("gtxn")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F2))
	return res.String()
}

type Teal_gtxn_op struct {
	T1 uint8
	F2 uint8
}

func Parse_Teal_gtxn_op(ctx ParserContext) *Teal_gtxn_op {
	t := &Teal_gtxn_op{
		T1: ctx.Read_uint8(),
		F2: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_gtxn_op) String() string {
	res := strings.Builder{}
	res.WriteString("gtxn")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F2))
	return res.String()
}

type Teal_load struct {
	I1 uint8
}

func (a *Teal_load) Teal() Teal {
	return Teal{a}
}
func (a *Teal_load) String() string {
	res := strings.Builder{}
	res.WriteString("load")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_load_op struct {
	I1 uint8
}

func Parse_Teal_load_op(ctx ParserContext) *Teal_load_op {
	t := &Teal_load_op{
		I1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_load_op) String() string {
	res := strings.Builder{}
	res.WriteString("load")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_store struct {
	STACK_1 TealAst
	I1      uint8
}

func (a *Teal_store) Teal() Teal {
	return Teal{a}
}
func (a *Teal_store) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("store")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_store_op struct {
	I1 uint8
}

func Parse_Teal_store_op(ctx ParserContext) *Teal_store_op {
	t := &Teal_store_op{
		I1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_store_op) String() string {
	res := strings.Builder{}
	res.WriteString("store")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_txna struct {
	F1 uint8
	I2 uint8
}

func (a *Teal_txna) Teal() Teal {
	return Teal{a}
}
func (a *Teal_txna) String() string {
	res := strings.Builder{}
	res.WriteString("txna")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I2))
	return res.String()
}

type Teal_txna_op struct {
	F1 uint8
	I2 uint8
}

func Parse_Teal_txna_op(ctx ParserContext) *Teal_txna_op {
	t := &Teal_txna_op{
		F1: ctx.Read_uint8(),
		I2: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_txna_op) String() string {
	res := strings.Builder{}
	res.WriteString("txna")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I2))
	return res.String()
}

type Teal_gtxna struct {
	T1 uint8
	F2 uint8
	I3 uint8
}

func (a *Teal_gtxna) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gtxna) String() string {
	res := strings.Builder{}
	res.WriteString("gtxna")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F2))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I3))
	return res.String()
}

type Teal_gtxna_op struct {
	T1 uint8
	F2 uint8
	I3 uint8
}

func Parse_Teal_gtxna_op(ctx ParserContext) *Teal_gtxna_op {
	t := &Teal_gtxna_op{
		T1: ctx.Read_uint8(),
		F2: ctx.Read_uint8(),
		I3: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_gtxna_op) String() string {
	res := strings.Builder{}
	res.WriteString("gtxna")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F2))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I3))
	return res.String()
}

type Teal_gtxns struct {
	STACK_1 TealAst
	F1      uint8
}

func (a *Teal_gtxns) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gtxns) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("gtxns")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_gtxns_op struct {
	F1 uint8
}

func Parse_Teal_gtxns_op(ctx ParserContext) *Teal_gtxns_op {
	t := &Teal_gtxns_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_gtxns_op) String() string {
	res := strings.Builder{}
	res.WriteString("gtxns")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_gtxnsa struct {
	STACK_1 TealAst
	F1      uint8
	I2      uint8
}

func (a *Teal_gtxnsa) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gtxnsa) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("gtxnsa")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I2))
	return res.String()
}

type Teal_gtxnsa_op struct {
	F1 uint8
	I2 uint8
}

func Parse_Teal_gtxnsa_op(ctx ParserContext) *Teal_gtxnsa_op {
	t := &Teal_gtxnsa_op{
		F1: ctx.Read_uint8(),
		I2: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_gtxnsa_op) String() string {
	res := strings.Builder{}
	res.WriteString("gtxnsa")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I2))
	return res.String()
}

type Teal_gload struct {
	T1 uint8
	I2 uint8
}

func (a *Teal_gload) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gload) String() string {
	res := strings.Builder{}
	res.WriteString("gload")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I2))
	return res.String()
}

type Teal_gload_op struct {
	T1 uint8
	I2 uint8
}

func Parse_Teal_gload_op(ctx ParserContext) *Teal_gload_op {
	t := &Teal_gload_op{
		T1: ctx.Read_uint8(),
		I2: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_gload_op) String() string {
	res := strings.Builder{}
	res.WriteString("gload")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I2))
	return res.String()
}

type Teal_gloads struct {
	STACK_1 TealAst
	I1      uint8
}

func (a *Teal_gloads) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gloads) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("gloads")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_gloads_op struct {
	I1 uint8
}

func Parse_Teal_gloads_op(ctx ParserContext) *Teal_gloads_op {
	t := &Teal_gloads_op{
		I1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_gloads_op) String() string {
	res := strings.Builder{}
	res.WriteString("gloads")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_gaid struct {
	T1 uint8
}

func (a *Teal_gaid) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gaid) String() string {
	res := strings.Builder{}
	res.WriteString("gaid")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	return res.String()
}

type Teal_gaid_op struct {
	T1 uint8
}

func Parse_Teal_gaid_op(ctx ParserContext) *Teal_gaid_op {
	t := &Teal_gaid_op{
		T1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_gaid_op) String() string {
	res := strings.Builder{}
	res.WriteString("gaid")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	return res.String()
}

type Teal_gaids struct {
	STACK_1 TealAst
}

func (a *Teal_gaids) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gaids) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("gaids")
	return res.String()
}

type Teal_gaids_op struct {
}

func Parse_Teal_gaids_op(ctx ParserContext) *Teal_gaids_op {
	t := &Teal_gaids_op{}
	return t
}

func (a *Teal_gaids_op) String() string {
	res := strings.Builder{}
	res.WriteString("gaids")
	return res.String()
}

type Teal_loads struct {
	STACK_1 TealAst
}

func (a *Teal_loads) Teal() Teal {
	return Teal{a}
}
func (a *Teal_loads) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("loads")
	return res.String()
}

type Teal_loads_op struct {
}

func Parse_Teal_loads_op(ctx ParserContext) *Teal_loads_op {
	t := &Teal_loads_op{}
	return t
}

func (a *Teal_loads_op) String() string {
	res := strings.Builder{}
	res.WriteString("loads")
	return res.String()
}

type Teal_stores struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_stores) Teal() Teal {
	return Teal{a}
}
func (a *Teal_stores) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("stores")
	return res.String()
}

type Teal_stores_op struct {
}

func Parse_Teal_stores_op(ctx ParserContext) *Teal_stores_op {
	t := &Teal_stores_op{}
	return t
}

func (a *Teal_stores_op) String() string {
	res := strings.Builder{}
	res.WriteString("stores")
	return res.String()
}

type Teal_bnz struct {
	STACK_1 TealAst
	TARGET1 int16
}

func (a *Teal_bnz) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bnz) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("bnz")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.TARGET1))
	return res.String()
}

type Teal_bnz_op struct {
	TARGET1 int16
}

func Parse_Teal_bnz_op(ctx ParserContext) *Teal_bnz_op {
	t := &Teal_bnz_op{
		TARGET1: ctx.Read_int16(),
	}
	return t
}

func (a *Teal_bnz_op) String() string {
	res := strings.Builder{}
	res.WriteString("bnz")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.TARGET1))
	return res.String()
}

type Teal_bz struct {
	STACK_1 TealAst
	TARGET1 int16
}

func (a *Teal_bz) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bz) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("bz")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.TARGET1))
	return res.String()
}

type Teal_bz_op struct {
	TARGET1 int16
}

func Parse_Teal_bz_op(ctx ParserContext) *Teal_bz_op {
	t := &Teal_bz_op{
		TARGET1: ctx.Read_int16(),
	}
	return t
}

func (a *Teal_bz_op) String() string {
	res := strings.Builder{}
	res.WriteString("bz")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.TARGET1))
	return res.String()
}

type Teal_b struct {
	TARGET1 int16
}

func (a *Teal_b) Teal() Teal {
	return Teal{a}
}
func (a *Teal_b) String() string {
	res := strings.Builder{}
	res.WriteString("b")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.TARGET1))
	return res.String()
}

type Teal_b_op struct {
	TARGET1 int16
}

func Parse_Teal_b_op(ctx ParserContext) *Teal_b_op {
	t := &Teal_b_op{
		TARGET1: ctx.Read_int16(),
	}
	return t
}

func (a *Teal_b_op) String() string {
	res := strings.Builder{}
	res.WriteString("b")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.TARGET1))
	return res.String()
}

type Teal_return_ struct {
	STACK_1 TealAst
}

func (a *Teal_return_) Teal() Teal {
	return Teal{a}
}
func (a *Teal_return_) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("return")
	return res.String()
}

type Teal_return__op struct {
}

func Parse_Teal_return__op(ctx ParserContext) *Teal_return__op {
	t := &Teal_return__op{}
	return t
}

func (a *Teal_return__op) String() string {
	res := strings.Builder{}
	res.WriteString("return")
	return res.String()
}

type Teal_assert struct {
	STACK_1 TealAst
}

func (a *Teal_assert) Teal() Teal {
	return Teal{a}
}
func (a *Teal_assert) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("assert")
	return res.String()
}

type Teal_assert_op struct {
}

func Parse_Teal_assert_op(ctx ParserContext) *Teal_assert_op {
	t := &Teal_assert_op{}
	return t
}

func (a *Teal_assert_op) String() string {
	res := strings.Builder{}
	res.WriteString("assert")
	return res.String()
}

type Teal_bury struct {
	STACK_1 TealAst
	N1      uint8
}

func (a *Teal_bury) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bury) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("bury")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_bury_op struct {
	N1 uint8
}

func Parse_Teal_bury_op(ctx ParserContext) *Teal_bury_op {
	t := &Teal_bury_op{
		N1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_bury_op) String() string {
	res := strings.Builder{}
	res.WriteString("bury")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_popn struct {
	N1 uint8
}

func (a *Teal_popn) Teal() Teal {
	return Teal{a}
}
func (a *Teal_popn) String() string {
	res := strings.Builder{}
	res.WriteString("popn")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_popn_op struct {
	N1 uint8
}

func Parse_Teal_popn_op(ctx ParserContext) *Teal_popn_op {
	t := &Teal_popn_op{
		N1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_popn_op) String() string {
	res := strings.Builder{}
	res.WriteString("popn")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_dupn struct {
	STACK_1 TealAst
	N1      uint8
}

func (a *Teal_dupn) Teal() Teal {
	return Teal{a}
}
func (a *Teal_dupn) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("dupn")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_dupn_op struct {
	N1 uint8
}

func Parse_Teal_dupn_op(ctx ParserContext) *Teal_dupn_op {
	t := &Teal_dupn_op{
		N1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_dupn_op) String() string {
	res := strings.Builder{}
	res.WriteString("dupn")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_pop struct {
	STACK_1 TealAst
}

func (a *Teal_pop) Teal() Teal {
	return Teal{a}
}
func (a *Teal_pop) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("pop")
	return res.String()
}

type Teal_pop_op struct {
}

func Parse_Teal_pop_op(ctx ParserContext) *Teal_pop_op {
	t := &Teal_pop_op{}
	return t
}

func (a *Teal_pop_op) String() string {
	res := strings.Builder{}
	res.WriteString("pop")
	return res.String()
}

type Teal_dup struct {
	STACK_1 TealAst
}

func (a *Teal_dup) Teal() Teal {
	return Teal{a}
}
func (a *Teal_dup) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("dup")
	return res.String()
}

type Teal_dup_op struct {
}

func Parse_Teal_dup_op(ctx ParserContext) *Teal_dup_op {
	t := &Teal_dup_op{}
	return t
}

func (a *Teal_dup_op) String() string {
	res := strings.Builder{}
	res.WriteString("dup")
	return res.String()
}

type Teal_dup2 struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_dup2) Teal() Teal {
	return Teal{a}
}
func (a *Teal_dup2) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("dup2")
	return res.String()
}

type Teal_dup2_op struct {
}

func Parse_Teal_dup2_op(ctx ParserContext) *Teal_dup2_op {
	t := &Teal_dup2_op{}
	return t
}

func (a *Teal_dup2_op) String() string {
	res := strings.Builder{}
	res.WriteString("dup2")
	return res.String()
}

type Teal_dig struct {
	STACK_1 TealAst
	N1      uint8
}

func (a *Teal_dig) Teal() Teal {
	return Teal{a}
}
func (a *Teal_dig) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("dig")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_dig_op struct {
	N1 uint8
}

func Parse_Teal_dig_op(ctx ParserContext) *Teal_dig_op {
	t := &Teal_dig_op{
		N1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_dig_op) String() string {
	res := strings.Builder{}
	res.WriteString("dig")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_swap struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_swap) Teal() Teal {
	return Teal{a}
}
func (a *Teal_swap) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("swap")
	return res.String()
}

type Teal_swap_op struct {
}

func Parse_Teal_swap_op(ctx ParserContext) *Teal_swap_op {
	t := &Teal_swap_op{}
	return t
}

func (a *Teal_swap_op) String() string {
	res := strings.Builder{}
	res.WriteString("swap")
	return res.String()
}

type Teal_select struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_select) Teal() Teal {
	return Teal{a}
}
func (a *Teal_select) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("select")
	return res.String()
}

type Teal_select_op struct {
}

func Parse_Teal_select_op(ctx ParserContext) *Teal_select_op {
	t := &Teal_select_op{}
	return t
}

func (a *Teal_select_op) String() string {
	res := strings.Builder{}
	res.WriteString("select")
	return res.String()
}

type Teal_cover struct {
	STACK_1 TealAst
	N1      uint8
}

func (a *Teal_cover) Teal() Teal {
	return Teal{a}
}
func (a *Teal_cover) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("cover")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_cover_op struct {
	N1 uint8
}

func Parse_Teal_cover_op(ctx ParserContext) *Teal_cover_op {
	t := &Teal_cover_op{
		N1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_cover_op) String() string {
	res := strings.Builder{}
	res.WriteString("cover")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_uncover struct {
	STACK_1 TealAst
	N1      uint8
}

func (a *Teal_uncover) Teal() Teal {
	return Teal{a}
}
func (a *Teal_uncover) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("uncover")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_uncover_op struct {
	N1 uint8
}

func Parse_Teal_uncover_op(ctx ParserContext) *Teal_uncover_op {
	t := &Teal_uncover_op{
		N1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_uncover_op) String() string {
	res := strings.Builder{}
	res.WriteString("uncover")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.N1))
	return res.String()
}

type Teal_concat struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_concat) Teal() Teal {
	return Teal{a}
}
func (a *Teal_concat) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("concat")
	return res.String()
}

type Teal_concat_op struct {
}

func Parse_Teal_concat_op(ctx ParserContext) *Teal_concat_op {
	t := &Teal_concat_op{}
	return t
}

func (a *Teal_concat_op) String() string {
	res := strings.Builder{}
	res.WriteString("concat")
	return res.String()
}

type Teal_substring struct {
	STACK_1 TealAst
	S1      uint8
	E2      uint8
}

func (a *Teal_substring) Teal() Teal {
	return Teal{a}
}
func (a *Teal_substring) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("substring")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.S1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.E2))
	return res.String()
}

type Teal_substring_op struct {
	S1 uint8
	E2 uint8
}

func Parse_Teal_substring_op(ctx ParserContext) *Teal_substring_op {
	t := &Teal_substring_op{
		S1: ctx.Read_uint8(),
		E2: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_substring_op) String() string {
	res := strings.Builder{}
	res.WriteString("substring")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.S1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.E2))
	return res.String()
}

type Teal_substring3 struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_substring3) Teal() Teal {
	return Teal{a}
}
func (a *Teal_substring3) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("substring3")
	return res.String()
}

type Teal_substring3_op struct {
}

func Parse_Teal_substring3_op(ctx ParserContext) *Teal_substring3_op {
	t := &Teal_substring3_op{}
	return t
}

func (a *Teal_substring3_op) String() string {
	res := strings.Builder{}
	res.WriteString("substring3")
	return res.String()
}

type Teal_getbit struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_getbit) Teal() Teal {
	return Teal{a}
}
func (a *Teal_getbit) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("getbit")
	return res.String()
}

type Teal_getbit_op struct {
}

func Parse_Teal_getbit_op(ctx ParserContext) *Teal_getbit_op {
	t := &Teal_getbit_op{}
	return t
}

func (a *Teal_getbit_op) String() string {
	res := strings.Builder{}
	res.WriteString("getbit")
	return res.String()
}

type Teal_setbit struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_setbit) Teal() Teal {
	return Teal{a}
}
func (a *Teal_setbit) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("setbit")
	return res.String()
}

type Teal_setbit_op struct {
}

func Parse_Teal_setbit_op(ctx ParserContext) *Teal_setbit_op {
	t := &Teal_setbit_op{}
	return t
}

func (a *Teal_setbit_op) String() string {
	res := strings.Builder{}
	res.WriteString("setbit")
	return res.String()
}

type Teal_getbyte struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_getbyte) Teal() Teal {
	return Teal{a}
}
func (a *Teal_getbyte) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("getbyte")
	return res.String()
}

type Teal_getbyte_op struct {
}

func Parse_Teal_getbyte_op(ctx ParserContext) *Teal_getbyte_op {
	t := &Teal_getbyte_op{}
	return t
}

func (a *Teal_getbyte_op) String() string {
	res := strings.Builder{}
	res.WriteString("getbyte")
	return res.String()
}

type Teal_setbyte struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_setbyte) Teal() Teal {
	return Teal{a}
}
func (a *Teal_setbyte) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("setbyte")
	return res.String()
}

type Teal_setbyte_op struct {
}

func Parse_Teal_setbyte_op(ctx ParserContext) *Teal_setbyte_op {
	t := &Teal_setbyte_op{}
	return t
}

func (a *Teal_setbyte_op) String() string {
	res := strings.Builder{}
	res.WriteString("setbyte")
	return res.String()
}

type Teal_extract struct {
	STACK_1 TealAst
	S1      uint8
	L2      uint8
}

func (a *Teal_extract) Teal() Teal {
	return Teal{a}
}
func (a *Teal_extract) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("extract")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.S1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.L2))
	return res.String()
}

type Teal_extract_op struct {
	S1 uint8
	L2 uint8
}

func Parse_Teal_extract_op(ctx ParserContext) *Teal_extract_op {
	t := &Teal_extract_op{
		S1: ctx.Read_uint8(),
		L2: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_extract_op) String() string {
	res := strings.Builder{}
	res.WriteString("extract")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.S1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.L2))
	return res.String()
}

type Teal_extract3 struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_extract3) Teal() Teal {
	return Teal{a}
}
func (a *Teal_extract3) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("extract3")
	return res.String()
}

type Teal_extract3_op struct {
}

func Parse_Teal_extract3_op(ctx ParserContext) *Teal_extract3_op {
	t := &Teal_extract3_op{}
	return t
}

func (a *Teal_extract3_op) String() string {
	res := strings.Builder{}
	res.WriteString("extract3")
	return res.String()
}

type Teal_extract_uint16 struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_extract_uint16) Teal() Teal {
	return Teal{a}
}
func (a *Teal_extract_uint16) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("extract_uint16")
	return res.String()
}

type Teal_extract_uint16_op struct {
}

func Parse_Teal_extract_uint16_op(ctx ParserContext) *Teal_extract_uint16_op {
	t := &Teal_extract_uint16_op{}
	return t
}

func (a *Teal_extract_uint16_op) String() string {
	res := strings.Builder{}
	res.WriteString("extract_uint16")
	return res.String()
}

type Teal_extract_uint32 struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_extract_uint32) Teal() Teal {
	return Teal{a}
}
func (a *Teal_extract_uint32) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("extract_uint32")
	return res.String()
}

type Teal_extract_uint32_op struct {
}

func Parse_Teal_extract_uint32_op(ctx ParserContext) *Teal_extract_uint32_op {
	t := &Teal_extract_uint32_op{}
	return t
}

func (a *Teal_extract_uint32_op) String() string {
	res := strings.Builder{}
	res.WriteString("extract_uint32")
	return res.String()
}

type Teal_extract_uint64 struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_extract_uint64) Teal() Teal {
	return Teal{a}
}
func (a *Teal_extract_uint64) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("extract_uint64")
	return res.String()
}

type Teal_extract_uint64_op struct {
}

func Parse_Teal_extract_uint64_op(ctx ParserContext) *Teal_extract_uint64_op {
	t := &Teal_extract_uint64_op{}
	return t
}

func (a *Teal_extract_uint64_op) String() string {
	res := strings.Builder{}
	res.WriteString("extract_uint64")
	return res.String()
}

type Teal_replace2 struct {
	STACK_1 TealAst
	STACK_2 TealAst
	S1      uint8
}

func (a *Teal_replace2) Teal() Teal {
	return Teal{a}
}
func (a *Teal_replace2) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("replace2")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.S1))
	return res.String()
}

type Teal_replace2_op struct {
	S1 uint8
}

func Parse_Teal_replace2_op(ctx ParserContext) *Teal_replace2_op {
	t := &Teal_replace2_op{
		S1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_replace2_op) String() string {
	res := strings.Builder{}
	res.WriteString("replace2")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.S1))
	return res.String()
}

type Teal_replace3 struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_replace3) Teal() Teal {
	return Teal{a}
}
func (a *Teal_replace3) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("replace3")
	return res.String()
}

type Teal_replace3_op struct {
}

func Parse_Teal_replace3_op(ctx ParserContext) *Teal_replace3_op {
	t := &Teal_replace3_op{}
	return t
}

func (a *Teal_replace3_op) String() string {
	res := strings.Builder{}
	res.WriteString("replace3")
	return res.String()
}

type Teal_base64_decode struct {
	STACK_1 TealAst
	E1      uint8
}

func (a *Teal_base64_decode) Teal() Teal {
	return Teal{a}
}
func (a *Teal_base64_decode) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("base64_decode")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.E1))
	return res.String()
}

type Teal_base64_decode_op struct {
	E1 uint8
}

func Parse_Teal_base64_decode_op(ctx ParserContext) *Teal_base64_decode_op {
	t := &Teal_base64_decode_op{
		E1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_base64_decode_op) String() string {
	res := strings.Builder{}
	res.WriteString("base64_decode")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.E1))
	return res.String()
}

type Teal_json_ref struct {
	STACK_1 TealAst
	STACK_2 TealAst
	R1      uint8
}

func (a *Teal_json_ref) Teal() Teal {
	return Teal{a}
}
func (a *Teal_json_ref) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("json_ref")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.R1))
	return res.String()
}

type Teal_json_ref_op struct {
	R1 uint8
}

func Parse_Teal_json_ref_op(ctx ParserContext) *Teal_json_ref_op {
	t := &Teal_json_ref_op{
		R1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_json_ref_op) String() string {
	res := strings.Builder{}
	res.WriteString("json_ref")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.R1))
	return res.String()
}

type Teal_balance struct {
	STACK_1 TealAst
}

func (a *Teal_balance) Teal() Teal {
	return Teal{a}
}
func (a *Teal_balance) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("balance")
	return res.String()
}

type Teal_balance_op struct {
}

func Parse_Teal_balance_op(ctx ParserContext) *Teal_balance_op {
	t := &Teal_balance_op{}
	return t
}

func (a *Teal_balance_op) String() string {
	res := strings.Builder{}
	res.WriteString("balance")
	return res.String()
}

type Teal_app_opted_in struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_app_opted_in) Teal() Teal {
	return Teal{a}
}
func (a *Teal_app_opted_in) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("app_opted_in")
	return res.String()
}

type Teal_app_opted_in_op struct {
}

func Parse_Teal_app_opted_in_op(ctx ParserContext) *Teal_app_opted_in_op {
	t := &Teal_app_opted_in_op{}
	return t
}

func (a *Teal_app_opted_in_op) String() string {
	res := strings.Builder{}
	res.WriteString("app_opted_in")
	return res.String()
}

type Teal_app_local_get struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_app_local_get) Teal() Teal {
	return Teal{a}
}
func (a *Teal_app_local_get) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("app_local_get")
	return res.String()
}

type Teal_app_local_get_op struct {
}

func Parse_Teal_app_local_get_op(ctx ParserContext) *Teal_app_local_get_op {
	t := &Teal_app_local_get_op{}
	return t
}

func (a *Teal_app_local_get_op) String() string {
	res := strings.Builder{}
	res.WriteString("app_local_get")
	return res.String()
}

type Teal_app_local_get_ex struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_app_local_get_ex) Teal() Teal {
	return Teal{a}
}
func (a *Teal_app_local_get_ex) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("app_local_get_ex")
	return res.String()
}

type Teal_app_local_get_ex_op struct {
}

func Parse_Teal_app_local_get_ex_op(ctx ParserContext) *Teal_app_local_get_ex_op {
	t := &Teal_app_local_get_ex_op{}
	return t
}

func (a *Teal_app_local_get_ex_op) String() string {
	res := strings.Builder{}
	res.WriteString("app_local_get_ex")
	return res.String()
}

type Teal_app_global_get struct {
	STACK_1 TealAst
}

func (a *Teal_app_global_get) Teal() Teal {
	return Teal{a}
}
func (a *Teal_app_global_get) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("app_global_get")
	return res.String()
}

type Teal_app_global_get_op struct {
}

func Parse_Teal_app_global_get_op(ctx ParserContext) *Teal_app_global_get_op {
	t := &Teal_app_global_get_op{}
	return t
}

func (a *Teal_app_global_get_op) String() string {
	res := strings.Builder{}
	res.WriteString("app_global_get")
	return res.String()
}

type Teal_app_global_get_ex struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_app_global_get_ex) Teal() Teal {
	return Teal{a}
}
func (a *Teal_app_global_get_ex) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("app_global_get_ex")
	return res.String()
}

type Teal_app_global_get_ex_op struct {
}

func Parse_Teal_app_global_get_ex_op(ctx ParserContext) *Teal_app_global_get_ex_op {
	t := &Teal_app_global_get_ex_op{}
	return t
}

func (a *Teal_app_global_get_ex_op) String() string {
	res := strings.Builder{}
	res.WriteString("app_global_get_ex")
	return res.String()
}

type Teal_app_local_put struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_app_local_put) Teal() Teal {
	return Teal{a}
}
func (a *Teal_app_local_put) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("app_local_put")
	return res.String()
}

type Teal_app_local_put_op struct {
}

func Parse_Teal_app_local_put_op(ctx ParserContext) *Teal_app_local_put_op {
	t := &Teal_app_local_put_op{}
	return t
}

func (a *Teal_app_local_put_op) String() string {
	res := strings.Builder{}
	res.WriteString("app_local_put")
	return res.String()
}

type Teal_app_global_put struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_app_global_put) Teal() Teal {
	return Teal{a}
}
func (a *Teal_app_global_put) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("app_global_put")
	return res.String()
}

type Teal_app_global_put_op struct {
}

func Parse_Teal_app_global_put_op(ctx ParserContext) *Teal_app_global_put_op {
	t := &Teal_app_global_put_op{}
	return t
}

func (a *Teal_app_global_put_op) String() string {
	res := strings.Builder{}
	res.WriteString("app_global_put")
	return res.String()
}

type Teal_app_local_del struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_app_local_del) Teal() Teal {
	return Teal{a}
}
func (a *Teal_app_local_del) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("app_local_del")
	return res.String()
}

type Teal_app_local_del_op struct {
}

func Parse_Teal_app_local_del_op(ctx ParserContext) *Teal_app_local_del_op {
	t := &Teal_app_local_del_op{}
	return t
}

func (a *Teal_app_local_del_op) String() string {
	res := strings.Builder{}
	res.WriteString("app_local_del")
	return res.String()
}

type Teal_app_global_del struct {
	STACK_1 TealAst
}

func (a *Teal_app_global_del) Teal() Teal {
	return Teal{a}
}
func (a *Teal_app_global_del) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("app_global_del")
	return res.String()
}

type Teal_app_global_del_op struct {
}

func Parse_Teal_app_global_del_op(ctx ParserContext) *Teal_app_global_del_op {
	t := &Teal_app_global_del_op{}
	return t
}

func (a *Teal_app_global_del_op) String() string {
	res := strings.Builder{}
	res.WriteString("app_global_del")
	return res.String()
}

type Teal_asset_holding_get struct {
	STACK_1 TealAst
	STACK_2 TealAst
	F1      uint8
}

func (a *Teal_asset_holding_get) Teal() Teal {
	return Teal{a}
}
func (a *Teal_asset_holding_get) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("asset_holding_get")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_asset_holding_get_op struct {
	F1 uint8
}

func Parse_Teal_asset_holding_get_op(ctx ParserContext) *Teal_asset_holding_get_op {
	t := &Teal_asset_holding_get_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_asset_holding_get_op) String() string {
	res := strings.Builder{}
	res.WriteString("asset_holding_get")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_asset_params_get struct {
	STACK_1 TealAst
	F1      uint8
}

func (a *Teal_asset_params_get) Teal() Teal {
	return Teal{a}
}
func (a *Teal_asset_params_get) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("asset_params_get")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_asset_params_get_op struct {
	F1 uint8
}

func Parse_Teal_asset_params_get_op(ctx ParserContext) *Teal_asset_params_get_op {
	t := &Teal_asset_params_get_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_asset_params_get_op) String() string {
	res := strings.Builder{}
	res.WriteString("asset_params_get")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_app_params_get struct {
	STACK_1 TealAst
	F1      uint8
}

func (a *Teal_app_params_get) Teal() Teal {
	return Teal{a}
}
func (a *Teal_app_params_get) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("app_params_get")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_app_params_get_op struct {
	F1 uint8
}

func Parse_Teal_app_params_get_op(ctx ParserContext) *Teal_app_params_get_op {
	t := &Teal_app_params_get_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_app_params_get_op) String() string {
	res := strings.Builder{}
	res.WriteString("app_params_get")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_acct_params_get struct {
	STACK_1 TealAst
	F1      uint8
}

func (a *Teal_acct_params_get) Teal() Teal {
	return Teal{a}
}
func (a *Teal_acct_params_get) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("acct_params_get")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_acct_params_get_op struct {
	F1 uint8
}

func Parse_Teal_acct_params_get_op(ctx ParserContext) *Teal_acct_params_get_op {
	t := &Teal_acct_params_get_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_acct_params_get_op) String() string {
	res := strings.Builder{}
	res.WriteString("acct_params_get")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_min_balance struct {
	STACK_1 TealAst
}

func (a *Teal_min_balance) Teal() Teal {
	return Teal{a}
}
func (a *Teal_min_balance) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("min_balance")
	return res.String()
}

type Teal_min_balance_op struct {
}

func Parse_Teal_min_balance_op(ctx ParserContext) *Teal_min_balance_op {
	t := &Teal_min_balance_op{}
	return t
}

func (a *Teal_min_balance_op) String() string {
	res := strings.Builder{}
	res.WriteString("min_balance")
	return res.String()
}

type Teal_pushbytes struct {
	BYTES1 []byte
}

func (a *Teal_pushbytes) Teal() Teal {
	return Teal{a}
}
func (a *Teal_pushbytes) String() string {
	res := strings.Builder{}
	res.WriteString("pushbytes")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.BYTES1))
	return res.String()
}

type Teal_pushbytes_op struct {
	BYTES1 []byte
}

func Parse_Teal_pushbytes_op(ctx ParserContext) *Teal_pushbytes_op {
	t := &Teal_pushbytes_op{
		BYTES1: ctx.Read_rbyte(),
	}
	return t
}

func (a *Teal_pushbytes_op) String() string {
	res := strings.Builder{}
	res.WriteString("pushbytes")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.BYTES1))
	return res.String()
}

type Teal_pushint struct {
	UINT1 []byte
}

func (a *Teal_pushint) Teal() Teal {
	return Teal{a}
}
func (a *Teal_pushint) String() string {
	res := strings.Builder{}
	res.WriteString("pushint")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.UINT1))
	return res.String()
}

type Teal_pushint_op struct {
	UINT1 []byte
}

func Parse_Teal_pushint_op(ctx ParserContext) *Teal_pushint_op {
	t := &Teal_pushint_op{
		UINT1: ctx.Read_rbyte(),
	}
	return t
}

func (a *Teal_pushint_op) String() string {
	res := strings.Builder{}
	res.WriteString("pushint")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.UINT1))
	return res.String()
}

type Teal_pushbytess struct {
	BYTES1 [][]byte
}

func (a *Teal_pushbytess) Teal() Teal {
	return Teal{a}
}
func (a *Teal_pushbytess) String() string {
	res := strings.Builder{}
	res.WriteString("pushbytess")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.BYTES1))
	return res.String()
}

type Teal_pushbytess_op struct {
	BYTES1 [][]byte
}

func Parse_Teal_pushbytess_op(ctx ParserContext) *Teal_pushbytess_op {
	t := &Teal_pushbytess_op{
		BYTES1: ctx.Read_rrbyte(),
	}
	return t
}

func (a *Teal_pushbytess_op) String() string {
	res := strings.Builder{}
	res.WriteString("pushbytess")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.BYTES1))
	return res.String()
}

type Teal_pushints struct {
	UINT1 [][]byte
}

func (a *Teal_pushints) Teal() Teal {
	return Teal{a}
}
func (a *Teal_pushints) String() string {
	res := strings.Builder{}
	res.WriteString("pushints")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.UINT1))
	return res.String()
}

type Teal_pushints_op struct {
	UINT1 [][]byte
}

func Parse_Teal_pushints_op(ctx ParserContext) *Teal_pushints_op {
	t := &Teal_pushints_op{
		UINT1: ctx.Read_rrbyte(),
	}
	return t
}

func (a *Teal_pushints_op) String() string {
	res := strings.Builder{}
	res.WriteString("pushints")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.UINT1))
	return res.String()
}

type Teal_ed25519verify_bare struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_ed25519verify_bare) Teal() Teal {
	return Teal{a}
}
func (a *Teal_ed25519verify_bare) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("ed25519verify_bare")
	return res.String()
}

type Teal_ed25519verify_bare_op struct {
}

func Parse_Teal_ed25519verify_bare_op(ctx ParserContext) *Teal_ed25519verify_bare_op {
	t := &Teal_ed25519verify_bare_op{}
	return t
}

func (a *Teal_ed25519verify_bare_op) String() string {
	res := strings.Builder{}
	res.WriteString("ed25519verify_bare")
	return res.String()
}

type Teal_callsub struct {
	TARGET1 int16
}

func (a *Teal_callsub) Teal() Teal {
	return Teal{a}
}
func (a *Teal_callsub) String() string {
	res := strings.Builder{}
	res.WriteString("callsub")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.TARGET1))
	return res.String()
}

type Teal_callsub_op struct {
	TARGET1 int16
}

func Parse_Teal_callsub_op(ctx ParserContext) *Teal_callsub_op {
	t := &Teal_callsub_op{
		TARGET1: ctx.Read_int16(),
	}
	return t
}

func (a *Teal_callsub_op) String() string {
	res := strings.Builder{}
	res.WriteString("callsub")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.TARGET1))
	return res.String()
}

type Teal_retsub struct {
}

func (a *Teal_retsub) Teal() Teal {
	return Teal{a}
}
func (a *Teal_retsub) String() string {
	res := strings.Builder{}
	res.WriteString("retsub")
	return res.String()
}

type Teal_retsub_op struct {
}

func Parse_Teal_retsub_op(ctx ParserContext) *Teal_retsub_op {
	t := &Teal_retsub_op{}
	return t
}

func (a *Teal_retsub_op) String() string {
	res := strings.Builder{}
	res.WriteString("retsub")
	return res.String()
}

type Teal_proto struct {
	A1 uint8
	R2 uint8
}

func (a *Teal_proto) Teal() Teal {
	return Teal{a}
}
func (a *Teal_proto) String() string {
	res := strings.Builder{}
	res.WriteString("proto")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.A1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.R2))
	return res.String()
}

type Teal_proto_op struct {
	A1 uint8
	R2 uint8
}

func Parse_Teal_proto_op(ctx ParserContext) *Teal_proto_op {
	t := &Teal_proto_op{
		A1: ctx.Read_uint8(),
		R2: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_proto_op) String() string {
	res := strings.Builder{}
	res.WriteString("proto")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.A1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.R2))
	return res.String()
}

type Teal_frame_dig struct {
	I1 int8
}

func (a *Teal_frame_dig) Teal() Teal {
	return Teal{a}
}
func (a *Teal_frame_dig) String() string {
	res := strings.Builder{}
	res.WriteString("frame_dig")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_frame_dig_op struct {
	I1 int8
}

func Parse_Teal_frame_dig_op(ctx ParserContext) *Teal_frame_dig_op {
	t := &Teal_frame_dig_op{
		I1: ctx.Read_int8(),
	}
	return t
}

func (a *Teal_frame_dig_op) String() string {
	res := strings.Builder{}
	res.WriteString("frame_dig")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_frame_bury struct {
	STACK_1 TealAst
	I1      int8
}

func (a *Teal_frame_bury) Teal() Teal {
	return Teal{a}
}
func (a *Teal_frame_bury) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("frame_bury")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_frame_bury_op struct {
	I1 int8
}

func Parse_Teal_frame_bury_op(ctx ParserContext) *Teal_frame_bury_op {
	t := &Teal_frame_bury_op{
		I1: ctx.Read_int8(),
	}
	return t
}

func (a *Teal_frame_bury_op) String() string {
	res := strings.Builder{}
	res.WriteString("frame_bury")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I1))
	return res.String()
}

type Teal_switch_ struct {
	STACK_1 TealAst
	TARGET1 [][]byte
}

func (a *Teal_switch_) Teal() Teal {
	return Teal{a}
}
func (a *Teal_switch_) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("switch")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.TARGET1))
	return res.String()
}

type Teal_switch__op struct {
	TARGET1 [][]byte
}

func Parse_Teal_switch__op(ctx ParserContext) *Teal_switch__op {
	t := &Teal_switch__op{
		TARGET1: ctx.Read_rrbyte(),
	}
	return t
}

func (a *Teal_switch__op) String() string {
	res := strings.Builder{}
	res.WriteString("switch")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.TARGET1))
	return res.String()
}

type Teal_match struct {
	TARGET1 [][]byte
}

func (a *Teal_match) Teal() Teal {
	return Teal{a}
}
func (a *Teal_match) String() string {
	res := strings.Builder{}
	res.WriteString("match")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.TARGET1))
	return res.String()
}

type Teal_match_op struct {
	TARGET1 [][]byte
}

func Parse_Teal_match_op(ctx ParserContext) *Teal_match_op {
	t := &Teal_match_op{
		TARGET1: ctx.Read_rrbyte(),
	}
	return t
}

func (a *Teal_match_op) String() string {
	res := strings.Builder{}
	res.WriteString("match")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.TARGET1))
	return res.String()
}

type Teal_shl struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_shl) Teal() Teal {
	return Teal{a}
}
func (a *Teal_shl) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("shl")
	return res.String()
}

type Teal_shl_op struct {
}

func Parse_Teal_shl_op(ctx ParserContext) *Teal_shl_op {
	t := &Teal_shl_op{}
	return t
}

func (a *Teal_shl_op) String() string {
	res := strings.Builder{}
	res.WriteString("shl")
	return res.String()
}

type Teal_shr struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_shr) Teal() Teal {
	return Teal{a}
}
func (a *Teal_shr) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("shr")
	return res.String()
}

type Teal_shr_op struct {
}

func Parse_Teal_shr_op(ctx ParserContext) *Teal_shr_op {
	t := &Teal_shr_op{}
	return t
}

func (a *Teal_shr_op) String() string {
	res := strings.Builder{}
	res.WriteString("shr")
	return res.String()
}

type Teal_sqrt struct {
	STACK_1 TealAst
}

func (a *Teal_sqrt) Teal() Teal {
	return Teal{a}
}
func (a *Teal_sqrt) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("sqrt")
	return res.String()
}

type Teal_sqrt_op struct {
}

func Parse_Teal_sqrt_op(ctx ParserContext) *Teal_sqrt_op {
	t := &Teal_sqrt_op{}
	return t
}

func (a *Teal_sqrt_op) String() string {
	res := strings.Builder{}
	res.WriteString("sqrt")
	return res.String()
}

type Teal_bitlen struct {
	STACK_1 TealAst
}

func (a *Teal_bitlen) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bitlen) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("bitlen")
	return res.String()
}

type Teal_bitlen_op struct {
}

func Parse_Teal_bitlen_op(ctx ParserContext) *Teal_bitlen_op {
	t := &Teal_bitlen_op{}
	return t
}

func (a *Teal_bitlen_op) String() string {
	res := strings.Builder{}
	res.WriteString("bitlen")
	return res.String()
}

type Teal_exp struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_exp) Teal() Teal {
	return Teal{a}
}
func (a *Teal_exp) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("exp")
	return res.String()
}

type Teal_exp_op struct {
}

func Parse_Teal_exp_op(ctx ParserContext) *Teal_exp_op {
	t := &Teal_exp_op{}
	return t
}

func (a *Teal_exp_op) String() string {
	res := strings.Builder{}
	res.WriteString("exp")
	return res.String()
}

type Teal_expw struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_expw) Teal() Teal {
	return Teal{a}
}
func (a *Teal_expw) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("expw")
	return res.String()
}

type Teal_expw_op struct {
}

func Parse_Teal_expw_op(ctx ParserContext) *Teal_expw_op {
	t := &Teal_expw_op{}
	return t
}

func (a *Teal_expw_op) String() string {
	res := strings.Builder{}
	res.WriteString("expw")
	return res.String()
}

type Teal_bsqrt struct {
	STACK_1 TealAst
}

func (a *Teal_bsqrt) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bsqrt) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("bsqrt")
	return res.String()
}

type Teal_bsqrt_op struct {
}

func Parse_Teal_bsqrt_op(ctx ParserContext) *Teal_bsqrt_op {
	t := &Teal_bsqrt_op{}
	return t
}

func (a *Teal_bsqrt_op) String() string {
	res := strings.Builder{}
	res.WriteString("bsqrt")
	return res.String()
}

type Teal_divw struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_divw) Teal() Teal {
	return Teal{a}
}
func (a *Teal_divw) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("divw")
	return res.String()
}

type Teal_divw_op struct {
}

func Parse_Teal_divw_op(ctx ParserContext) *Teal_divw_op {
	t := &Teal_divw_op{}
	return t
}

func (a *Teal_divw_op) String() string {
	res := strings.Builder{}
	res.WriteString("divw")
	return res.String()
}

type Teal_sha3_256 struct {
	STACK_1 TealAst
}

func (a *Teal_sha3_256) Teal() Teal {
	return Teal{a}
}
func (a *Teal_sha3_256) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("sha3_256")
	return res.String()
}

type Teal_sha3_256_op struct {
}

func Parse_Teal_sha3_256_op(ctx ParserContext) *Teal_sha3_256_op {
	t := &Teal_sha3_256_op{}
	return t
}

func (a *Teal_sha3_256_op) String() string {
	res := strings.Builder{}
	res.WriteString("sha3_256")
	return res.String()
}

type Teal_bplus struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_bplus) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bplus) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b+")
	return res.String()
}

type Teal_bplus_op struct {
}

func Parse_Teal_bplus_op(ctx ParserContext) *Teal_bplus_op {
	t := &Teal_bplus_op{}
	return t
}

func (a *Teal_bplus_op) String() string {
	res := strings.Builder{}
	res.WriteString("b+")
	return res.String()
}

type Teal_bminus struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_bminus) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bminus) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b-")
	return res.String()
}

type Teal_bminus_op struct {
}

func Parse_Teal_bminus_op(ctx ParserContext) *Teal_bminus_op {
	t := &Teal_bminus_op{}
	return t
}

func (a *Teal_bminus_op) String() string {
	res := strings.Builder{}
	res.WriteString("b-")
	return res.String()
}

type Teal_bdiv struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_bdiv) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bdiv) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b/")
	return res.String()
}

type Teal_bdiv_op struct {
}

func Parse_Teal_bdiv_op(ctx ParserContext) *Teal_bdiv_op {
	t := &Teal_bdiv_op{}
	return t
}

func (a *Teal_bdiv_op) String() string {
	res := strings.Builder{}
	res.WriteString("b/")
	return res.String()
}

type Teal_bmul struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_bmul) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bmul) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b*")
	return res.String()
}

type Teal_bmul_op struct {
}

func Parse_Teal_bmul_op(ctx ParserContext) *Teal_bmul_op {
	t := &Teal_bmul_op{}
	return t
}

func (a *Teal_bmul_op) String() string {
	res := strings.Builder{}
	res.WriteString("b*")
	return res.String()
}

type Teal_blt struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_blt) Teal() Teal {
	return Teal{a}
}
func (a *Teal_blt) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b<")
	return res.String()
}

type Teal_blt_op struct {
}

func Parse_Teal_blt_op(ctx ParserContext) *Teal_blt_op {
	t := &Teal_blt_op{}
	return t
}

func (a *Teal_blt_op) String() string {
	res := strings.Builder{}
	res.WriteString("b<")
	return res.String()
}

type Teal_bgt struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_bgt) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bgt) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b>")
	return res.String()
}

type Teal_bgt_op struct {
}

func Parse_Teal_bgt_op(ctx ParserContext) *Teal_bgt_op {
	t := &Teal_bgt_op{}
	return t
}

func (a *Teal_bgt_op) String() string {
	res := strings.Builder{}
	res.WriteString("b>")
	return res.String()
}

type Teal_blteq struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_blteq) Teal() Teal {
	return Teal{a}
}
func (a *Teal_blteq) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b<=")
	return res.String()
}

type Teal_blteq_op struct {
}

func Parse_Teal_blteq_op(ctx ParserContext) *Teal_blteq_op {
	t := &Teal_blteq_op{}
	return t
}

func (a *Teal_blteq_op) String() string {
	res := strings.Builder{}
	res.WriteString("b<=")
	return res.String()
}

type Teal_bgteq struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_bgteq) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bgteq) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b>=")
	return res.String()
}

type Teal_bgteq_op struct {
}

func Parse_Teal_bgteq_op(ctx ParserContext) *Teal_bgteq_op {
	t := &Teal_bgteq_op{}
	return t
}

func (a *Teal_bgteq_op) String() string {
	res := strings.Builder{}
	res.WriteString("b>=")
	return res.String()
}

type Teal_beqeq struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_beqeq) Teal() Teal {
	return Teal{a}
}
func (a *Teal_beqeq) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b==")
	return res.String()
}

type Teal_beqeq_op struct {
}

func Parse_Teal_beqeq_op(ctx ParserContext) *Teal_beqeq_op {
	t := &Teal_beqeq_op{}
	return t
}

func (a *Teal_beqeq_op) String() string {
	res := strings.Builder{}
	res.WriteString("b==")
	return res.String()
}

type Teal_bnoteq struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_bnoteq) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bnoteq) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b!=")
	return res.String()
}

type Teal_bnoteq_op struct {
}

func Parse_Teal_bnoteq_op(ctx ParserContext) *Teal_bnoteq_op {
	t := &Teal_bnoteq_op{}
	return t
}

func (a *Teal_bnoteq_op) String() string {
	res := strings.Builder{}
	res.WriteString("b!=")
	return res.String()
}

type Teal_bmod struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_bmod) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bmod) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b%")
	return res.String()
}

type Teal_bmod_op struct {
}

func Parse_Teal_bmod_op(ctx ParserContext) *Teal_bmod_op {
	t := &Teal_bmod_op{}
	return t
}

func (a *Teal_bmod_op) String() string {
	res := strings.Builder{}
	res.WriteString("b%")
	return res.String()
}

type Teal_bor struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_bor) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bor) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b|")
	return res.String()
}

type Teal_bor_op struct {
}

func Parse_Teal_bor_op(ctx ParserContext) *Teal_bor_op {
	t := &Teal_bor_op{}
	return t
}

func (a *Teal_bor_op) String() string {
	res := strings.Builder{}
	res.WriteString("b|")
	return res.String()
}

type Teal_band struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_band) Teal() Teal {
	return Teal{a}
}
func (a *Teal_band) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b&")
	return res.String()
}

type Teal_band_op struct {
}

func Parse_Teal_band_op(ctx ParserContext) *Teal_band_op {
	t := &Teal_band_op{}
	return t
}

func (a *Teal_band_op) String() string {
	res := strings.Builder{}
	res.WriteString("b&")
	return res.String()
}

type Teal_bxor struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_bxor) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bxor) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b^")
	return res.String()
}

type Teal_bxor_op struct {
}

func Parse_Teal_bxor_op(ctx ParserContext) *Teal_bxor_op {
	t := &Teal_bxor_op{}
	return t
}

func (a *Teal_bxor_op) String() string {
	res := strings.Builder{}
	res.WriteString("b^")
	return res.String()
}

type Teal_binv struct {
	STACK_1 TealAst
}

func (a *Teal_binv) Teal() Teal {
	return Teal{a}
}
func (a *Teal_binv) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("b~")
	return res.String()
}

type Teal_binv_op struct {
}

func Parse_Teal_binv_op(ctx ParserContext) *Teal_binv_op {
	t := &Teal_binv_op{}
	return t
}

func (a *Teal_binv_op) String() string {
	res := strings.Builder{}
	res.WriteString("b~")
	return res.String()
}

type Teal_bzero struct {
	STACK_1 TealAst
}

func (a *Teal_bzero) Teal() Teal {
	return Teal{a}
}
func (a *Teal_bzero) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("bzero")
	return res.String()
}

type Teal_bzero_op struct {
}

func Parse_Teal_bzero_op(ctx ParserContext) *Teal_bzero_op {
	t := &Teal_bzero_op{}
	return t
}

func (a *Teal_bzero_op) String() string {
	res := strings.Builder{}
	res.WriteString("bzero")
	return res.String()
}

type Teal_log struct {
	STACK_1 TealAst
}

func (a *Teal_log) Teal() Teal {
	return Teal{a}
}
func (a *Teal_log) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("log")
	return res.String()
}

type Teal_log_op struct {
}

func Parse_Teal_log_op(ctx ParserContext) *Teal_log_op {
	t := &Teal_log_op{}
	return t
}

func (a *Teal_log_op) String() string {
	res := strings.Builder{}
	res.WriteString("log")
	return res.String()
}

type Teal_itxn_begin struct {
}

func (a *Teal_itxn_begin) Teal() Teal {
	return Teal{a}
}
func (a *Teal_itxn_begin) String() string {
	res := strings.Builder{}
	res.WriteString("itxn_begin")
	return res.String()
}

type Teal_itxn_begin_op struct {
}

func Parse_Teal_itxn_begin_op(ctx ParserContext) *Teal_itxn_begin_op {
	t := &Teal_itxn_begin_op{}
	return t
}

func (a *Teal_itxn_begin_op) String() string {
	res := strings.Builder{}
	res.WriteString("itxn_begin")
	return res.String()
}

type Teal_itxn_field struct {
	STACK_1 TealAst
	F1      uint8
}

func (a *Teal_itxn_field) Teal() Teal {
	return Teal{a}
}
func (a *Teal_itxn_field) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("itxn_field")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_itxn_field_op struct {
	F1 uint8
}

func Parse_Teal_itxn_field_op(ctx ParserContext) *Teal_itxn_field_op {
	t := &Teal_itxn_field_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_itxn_field_op) String() string {
	res := strings.Builder{}
	res.WriteString("itxn_field")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_itxn_submit struct {
}

func (a *Teal_itxn_submit) Teal() Teal {
	return Teal{a}
}
func (a *Teal_itxn_submit) String() string {
	res := strings.Builder{}
	res.WriteString("itxn_submit")
	return res.String()
}

type Teal_itxn_submit_op struct {
}

func Parse_Teal_itxn_submit_op(ctx ParserContext) *Teal_itxn_submit_op {
	t := &Teal_itxn_submit_op{}
	return t
}

func (a *Teal_itxn_submit_op) String() string {
	res := strings.Builder{}
	res.WriteString("itxn_submit")
	return res.String()
}

type Teal_itxn struct {
	F1 uint8
}

func (a *Teal_itxn) Teal() Teal {
	return Teal{a}
}
func (a *Teal_itxn) String() string {
	res := strings.Builder{}
	res.WriteString("itxn")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_itxn_op struct {
	F1 uint8
}

func Parse_Teal_itxn_op(ctx ParserContext) *Teal_itxn_op {
	t := &Teal_itxn_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_itxn_op) String() string {
	res := strings.Builder{}
	res.WriteString("itxn")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_itxna struct {
	F1 uint8
	I2 uint8
}

func (a *Teal_itxna) Teal() Teal {
	return Teal{a}
}
func (a *Teal_itxna) String() string {
	res := strings.Builder{}
	res.WriteString("itxna")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I2))
	return res.String()
}

type Teal_itxna_op struct {
	F1 uint8
	I2 uint8
}

func Parse_Teal_itxna_op(ctx ParserContext) *Teal_itxna_op {
	t := &Teal_itxna_op{
		F1: ctx.Read_uint8(),
		I2: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_itxna_op) String() string {
	res := strings.Builder{}
	res.WriteString("itxna")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I2))
	return res.String()
}

type Teal_itxn_next struct {
}

func (a *Teal_itxn_next) Teal() Teal {
	return Teal{a}
}
func (a *Teal_itxn_next) String() string {
	res := strings.Builder{}
	res.WriteString("itxn_next")
	return res.String()
}

type Teal_itxn_next_op struct {
}

func Parse_Teal_itxn_next_op(ctx ParserContext) *Teal_itxn_next_op {
	t := &Teal_itxn_next_op{}
	return t
}

func (a *Teal_itxn_next_op) String() string {
	res := strings.Builder{}
	res.WriteString("itxn_next")
	return res.String()
}

type Teal_gitxn struct {
	T1 uint8
	F2 uint8
}

func (a *Teal_gitxn) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gitxn) String() string {
	res := strings.Builder{}
	res.WriteString("gitxn")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F2))
	return res.String()
}

type Teal_gitxn_op struct {
	T1 uint8
	F2 uint8
}

func Parse_Teal_gitxn_op(ctx ParserContext) *Teal_gitxn_op {
	t := &Teal_gitxn_op{
		T1: ctx.Read_uint8(),
		F2: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_gitxn_op) String() string {
	res := strings.Builder{}
	res.WriteString("gitxn")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F2))
	return res.String()
}

type Teal_gitxna struct {
	T1 uint8
	F2 uint8
	I3 uint8
}

func (a *Teal_gitxna) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gitxna) String() string {
	res := strings.Builder{}
	res.WriteString("gitxna")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F2))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I3))
	return res.String()
}

type Teal_gitxna_op struct {
	T1 uint8
	F2 uint8
	I3 uint8
}

func Parse_Teal_gitxna_op(ctx ParserContext) *Teal_gitxna_op {
	t := &Teal_gitxna_op{
		T1: ctx.Read_uint8(),
		F2: ctx.Read_uint8(),
		I3: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_gitxna_op) String() string {
	res := strings.Builder{}
	res.WriteString("gitxna")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F2))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.I3))
	return res.String()
}

type Teal_box_create struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_box_create) Teal() Teal {
	return Teal{a}
}
func (a *Teal_box_create) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("box_create")
	return res.String()
}

type Teal_box_create_op struct {
}

func Parse_Teal_box_create_op(ctx ParserContext) *Teal_box_create_op {
	t := &Teal_box_create_op{}
	return t
}

func (a *Teal_box_create_op) String() string {
	res := strings.Builder{}
	res.WriteString("box_create")
	return res.String()
}

type Teal_box_extract struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_box_extract) Teal() Teal {
	return Teal{a}
}
func (a *Teal_box_extract) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("box_extract")
	return res.String()
}

type Teal_box_extract_op struct {
}

func Parse_Teal_box_extract_op(ctx ParserContext) *Teal_box_extract_op {
	t := &Teal_box_extract_op{}
	return t
}

func (a *Teal_box_extract_op) String() string {
	res := strings.Builder{}
	res.WriteString("box_extract")
	return res.String()
}

type Teal_box_replace struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
}

func (a *Teal_box_replace) Teal() Teal {
	return Teal{a}
}
func (a *Teal_box_replace) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("box_replace")
	return res.String()
}

type Teal_box_replace_op struct {
}

func Parse_Teal_box_replace_op(ctx ParserContext) *Teal_box_replace_op {
	t := &Teal_box_replace_op{}
	return t
}

func (a *Teal_box_replace_op) String() string {
	res := strings.Builder{}
	res.WriteString("box_replace")
	return res.String()
}

type Teal_box_del struct {
	STACK_1 TealAst
}

func (a *Teal_box_del) Teal() Teal {
	return Teal{a}
}
func (a *Teal_box_del) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("box_del")
	return res.String()
}

type Teal_box_del_op struct {
}

func Parse_Teal_box_del_op(ctx ParserContext) *Teal_box_del_op {
	t := &Teal_box_del_op{}
	return t
}

func (a *Teal_box_del_op) String() string {
	res := strings.Builder{}
	res.WriteString("box_del")
	return res.String()
}

type Teal_box_len struct {
	STACK_1 TealAst
}

func (a *Teal_box_len) Teal() Teal {
	return Teal{a}
}
func (a *Teal_box_len) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("box_len")
	return res.String()
}

type Teal_box_len_op struct {
}

func Parse_Teal_box_len_op(ctx ParserContext) *Teal_box_len_op {
	t := &Teal_box_len_op{}
	return t
}

func (a *Teal_box_len_op) String() string {
	res := strings.Builder{}
	res.WriteString("box_len")
	return res.String()
}

type Teal_box_get struct {
	STACK_1 TealAst
}

func (a *Teal_box_get) Teal() Teal {
	return Teal{a}
}
func (a *Teal_box_get) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("box_get")
	return res.String()
}

type Teal_box_get_op struct {
}

func Parse_Teal_box_get_op(ctx ParserContext) *Teal_box_get_op {
	t := &Teal_box_get_op{}
	return t
}

func (a *Teal_box_get_op) String() string {
	res := strings.Builder{}
	res.WriteString("box_get")
	return res.String()
}

type Teal_box_put struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_box_put) Teal() Teal {
	return Teal{a}
}
func (a *Teal_box_put) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("box_put")
	return res.String()
}

type Teal_box_put_op struct {
}

func Parse_Teal_box_put_op(ctx ParserContext) *Teal_box_put_op {
	t := &Teal_box_put_op{}
	return t
}

func (a *Teal_box_put_op) String() string {
	res := strings.Builder{}
	res.WriteString("box_put")
	return res.String()
}

type Teal_txnas struct {
	STACK_1 TealAst
	F1      uint8
}

func (a *Teal_txnas) Teal() Teal {
	return Teal{a}
}
func (a *Teal_txnas) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("txnas")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_txnas_op struct {
	F1 uint8
}

func Parse_Teal_txnas_op(ctx ParserContext) *Teal_txnas_op {
	t := &Teal_txnas_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_txnas_op) String() string {
	res := strings.Builder{}
	res.WriteString("txnas")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_gtxnas struct {
	STACK_1 TealAst
	T1      uint8
	F2      uint8
}

func (a *Teal_gtxnas) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gtxnas) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("gtxnas")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F2))
	return res.String()
}

type Teal_gtxnas_op struct {
	T1 uint8
	F2 uint8
}

func Parse_Teal_gtxnas_op(ctx ParserContext) *Teal_gtxnas_op {
	t := &Teal_gtxnas_op{
		T1: ctx.Read_uint8(),
		F2: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_gtxnas_op) String() string {
	res := strings.Builder{}
	res.WriteString("gtxnas")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F2))
	return res.String()
}

type Teal_gtxnsas struct {
	STACK_1 TealAst
	STACK_2 TealAst
	F1      uint8
}

func (a *Teal_gtxnsas) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gtxnsas) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("gtxnsas")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_gtxnsas_op struct {
	F1 uint8
}

func Parse_Teal_gtxnsas_op(ctx ParserContext) *Teal_gtxnsas_op {
	t := &Teal_gtxnsas_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_gtxnsas_op) String() string {
	res := strings.Builder{}
	res.WriteString("gtxnsas")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_args struct {
	STACK_1 TealAst
}

func (a *Teal_args) Teal() Teal {
	return Teal{a}
}
func (a *Teal_args) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("args")
	return res.String()
}

type Teal_args_op struct {
}

func Parse_Teal_args_op(ctx ParserContext) *Teal_args_op {
	t := &Teal_args_op{}
	return t
}

func (a *Teal_args_op) String() string {
	res := strings.Builder{}
	res.WriteString("args")
	return res.String()
}

type Teal_gloadss struct {
	STACK_1 TealAst
	STACK_2 TealAst
}

func (a *Teal_gloadss) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gloadss) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("gloadss")
	return res.String()
}

type Teal_gloadss_op struct {
}

func Parse_Teal_gloadss_op(ctx ParserContext) *Teal_gloadss_op {
	t := &Teal_gloadss_op{}
	return t
}

func (a *Teal_gloadss_op) String() string {
	res := strings.Builder{}
	res.WriteString("gloadss")
	return res.String()
}

type Teal_itxnas struct {
	STACK_1 TealAst
	F1      uint8
}

func (a *Teal_itxnas) Teal() Teal {
	return Teal{a}
}
func (a *Teal_itxnas) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("itxnas")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_itxnas_op struct {
	F1 uint8
}

func Parse_Teal_itxnas_op(ctx ParserContext) *Teal_itxnas_op {
	t := &Teal_itxnas_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_itxnas_op) String() string {
	res := strings.Builder{}
	res.WriteString("itxnas")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_gitxnas struct {
	STACK_1 TealAst
	T1      uint8
	F2      uint8
}

func (a *Teal_gitxnas) Teal() Teal {
	return Teal{a}
}
func (a *Teal_gitxnas) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("gitxnas")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F2))
	return res.String()
}

type Teal_gitxnas_op struct {
	T1 uint8
	F2 uint8
}

func Parse_Teal_gitxnas_op(ctx ParserContext) *Teal_gitxnas_op {
	t := &Teal_gitxnas_op{
		T1: ctx.Read_uint8(),
		F2: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_gitxnas_op) String() string {
	res := strings.Builder{}
	res.WriteString("gitxnas")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.T1))
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F2))
	return res.String()
}

type Teal_vrf_verify struct {
	STACK_1 TealAst
	STACK_2 TealAst
	STACK_3 TealAst
	S1      uint8
}

func (a *Teal_vrf_verify) Teal() Teal {
	return Teal{a}
}
func (a *Teal_vrf_verify) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_2.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	for _, op := range a.STACK_3.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("vrf_verify")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.S1))
	return res.String()
}

type Teal_vrf_verify_op struct {
	S1 uint8
}

func Parse_Teal_vrf_verify_op(ctx ParserContext) *Teal_vrf_verify_op {
	t := &Teal_vrf_verify_op{
		S1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_vrf_verify_op) String() string {
	res := strings.Builder{}
	res.WriteString("vrf_verify")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.S1))
	return res.String()
}

type Teal_block struct {
	STACK_1 TealAst
	F1      uint8
}

func (a *Teal_block) Teal() Teal {
	return Teal{a}
}
func (a *Teal_block) String() string {
	res := strings.Builder{}
	for _, op := range a.STACK_1.Teal() {
		res.WriteString(op.String())
		res.WriteString("\n")
	}
	res.WriteString("block")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

type Teal_block_op struct {
	F1 uint8
}

func Parse_Teal_block_op(ctx ParserContext) *Teal_block_op {
	t := &Teal_block_op{
		F1: ctx.Read_uint8(),
	}
	return t
}

func (a *Teal_block_op) String() string {
	res := strings.Builder{}
	res.WriteString("block")
	res.WriteString(" ")
	res.WriteString(fmt.Sprintf("%d", a.F1))
	return res.String()
}

func parseTeal(ctx ParserContext) TealOp {
	first := ctx.Read()
	value := first.String()
	switch value {
	case "err":
		return Parse_Teal_err_op(ctx)
	case "sha256":
		return Parse_Teal_sha256_op(ctx)
	case "keccak256":
		return Parse_Teal_keccak256_op(ctx)
	case "sha512_256":
		return Parse_Teal_sha512_256_op(ctx)
	case "ed25519verify":
		return Parse_Teal_ed25519verify_op(ctx)
	case "ecdsa_verify":
		return Parse_Teal_ecdsa_verify_op(ctx)
	case "ecdsa_pk_decompress":
		return Parse_Teal_ecdsa_pk_decompress_op(ctx)
	case "ecdsa_pk_recover":
		return Parse_Teal_ecdsa_pk_recover_op(ctx)
	case "+":
		return Parse_Teal_plus_op(ctx)
	case "-":
		return Parse_Teal_minus_op(ctx)
	case "/":
		return Parse_Teal_div_op(ctx)
	case "*":
		return Parse_Teal_mul_op(ctx)
	case "<":
		return Parse_Teal_lt_op(ctx)
	case ">":
		return Parse_Teal_gt_op(ctx)
	case "<=":
		return Parse_Teal_lteq_op(ctx)
	case ">=":
		return Parse_Teal_gteq_op(ctx)
	case "&&":
		return Parse_Teal_andand_op(ctx)
	case "||":
		return Parse_Teal_oror_op(ctx)
	case "==":
		return Parse_Teal_eqeq_op(ctx)
	case "!=":
		return Parse_Teal_noteq_op(ctx)
	case "!":
		return Parse_Teal_not_op(ctx)
	case "len":
		return Parse_Teal_len_op(ctx)
	case "itob":
		return Parse_Teal_itob_op(ctx)
	case "btoi":
		return Parse_Teal_btoi_op(ctx)
	case "%":
		return Parse_Teal_mod_op(ctx)
	case "|":
		return Parse_Teal_or_op(ctx)
	case "&":
		return Parse_Teal_and_op(ctx)
	case "^":
		return Parse_Teal_xor_op(ctx)
	case "~":
		return Parse_Teal_inv_op(ctx)
	case "mulw":
		return Parse_Teal_mulw_op(ctx)
	case "addw":
		return Parse_Teal_addw_op(ctx)
	case "divmodw":
		return Parse_Teal_divmodw_op(ctx)
	case "intcblock":
		return Parse_Teal_intcblock_op(ctx)
	case "intc":
		return Parse_Teal_intc_op(ctx)
	case "intc_0":
		return Parse_Teal_intc_0_op(ctx)
	case "intc_1":
		return Parse_Teal_intc_1_op(ctx)
	case "intc_2":
		return Parse_Teal_intc_2_op(ctx)
	case "intc_3":
		return Parse_Teal_intc_3_op(ctx)
	case "bytecblock":
		return Parse_Teal_bytecblock_op(ctx)
	case "bytec":
		return Parse_Teal_bytec_op(ctx)
	case "bytec_0":
		return Parse_Teal_bytec_0_op(ctx)
	case "bytec_1":
		return Parse_Teal_bytec_1_op(ctx)
	case "bytec_2":
		return Parse_Teal_bytec_2_op(ctx)
	case "bytec_3":
		return Parse_Teal_bytec_3_op(ctx)
	case "arg":
		return Parse_Teal_arg_op(ctx)
	case "arg_0":
		return Parse_Teal_arg_0_op(ctx)
	case "arg_1":
		return Parse_Teal_arg_1_op(ctx)
	case "arg_2":
		return Parse_Teal_arg_2_op(ctx)
	case "arg_3":
		return Parse_Teal_arg_3_op(ctx)
	case "txn":
		return Parse_Teal_txn_op(ctx)
	case "global":
		return Parse_Teal_global_op(ctx)
	case "gtxn":
		return Parse_Teal_gtxn_op(ctx)
	case "load":
		return Parse_Teal_load_op(ctx)
	case "store":
		return Parse_Teal_store_op(ctx)
	case "txna":
		return Parse_Teal_txna_op(ctx)
	case "gtxna":
		return Parse_Teal_gtxna_op(ctx)
	case "gtxns":
		return Parse_Teal_gtxns_op(ctx)
	case "gtxnsa":
		return Parse_Teal_gtxnsa_op(ctx)
	case "gload":
		return Parse_Teal_gload_op(ctx)
	case "gloads":
		return Parse_Teal_gloads_op(ctx)
	case "gaid":
		return Parse_Teal_gaid_op(ctx)
	case "gaids":
		return Parse_Teal_gaids_op(ctx)
	case "loads":
		return Parse_Teal_loads_op(ctx)
	case "stores":
		return Parse_Teal_stores_op(ctx)
	case "bnz":
		return Parse_Teal_bnz_op(ctx)
	case "bz":
		return Parse_Teal_bz_op(ctx)
	case "b":
		return Parse_Teal_b_op(ctx)
	case "return":
		return Parse_Teal_return__op(ctx)
	case "assert":
		return Parse_Teal_assert_op(ctx)
	case "bury":
		return Parse_Teal_bury_op(ctx)
	case "popn":
		return Parse_Teal_popn_op(ctx)
	case "dupn":
		return Parse_Teal_dupn_op(ctx)
	case "pop":
		return Parse_Teal_pop_op(ctx)
	case "dup":
		return Parse_Teal_dup_op(ctx)
	case "dup2":
		return Parse_Teal_dup2_op(ctx)
	case "dig":
		return Parse_Teal_dig_op(ctx)
	case "swap":
		return Parse_Teal_swap_op(ctx)
	case "select":
		return Parse_Teal_select_op(ctx)
	case "cover":
		return Parse_Teal_cover_op(ctx)
	case "uncover":
		return Parse_Teal_uncover_op(ctx)
	case "concat":
		return Parse_Teal_concat_op(ctx)
	case "substring":
		return Parse_Teal_substring_op(ctx)
	case "substring3":
		return Parse_Teal_substring3_op(ctx)
	case "getbit":
		return Parse_Teal_getbit_op(ctx)
	case "setbit":
		return Parse_Teal_setbit_op(ctx)
	case "getbyte":
		return Parse_Teal_getbyte_op(ctx)
	case "setbyte":
		return Parse_Teal_setbyte_op(ctx)
	case "extract":
		return Parse_Teal_extract_op(ctx)
	case "extract3":
		return Parse_Teal_extract3_op(ctx)
	case "extract_uint16":
		return Parse_Teal_extract_uint16_op(ctx)
	case "extract_uint32":
		return Parse_Teal_extract_uint32_op(ctx)
	case "extract_uint64":
		return Parse_Teal_extract_uint64_op(ctx)
	case "replace2":
		return Parse_Teal_replace2_op(ctx)
	case "replace3":
		return Parse_Teal_replace3_op(ctx)
	case "base64_decode":
		return Parse_Teal_base64_decode_op(ctx)
	case "json_ref":
		return Parse_Teal_json_ref_op(ctx)
	case "balance":
		return Parse_Teal_balance_op(ctx)
	case "app_opted_in":
		return Parse_Teal_app_opted_in_op(ctx)
	case "app_local_get":
		return Parse_Teal_app_local_get_op(ctx)
	case "app_local_get_ex":
		return Parse_Teal_app_local_get_ex_op(ctx)
	case "app_global_get":
		return Parse_Teal_app_global_get_op(ctx)
	case "app_global_get_ex":
		return Parse_Teal_app_global_get_ex_op(ctx)
	case "app_local_put":
		return Parse_Teal_app_local_put_op(ctx)
	case "app_global_put":
		return Parse_Teal_app_global_put_op(ctx)
	case "app_local_del":
		return Parse_Teal_app_local_del_op(ctx)
	case "app_global_del":
		return Parse_Teal_app_global_del_op(ctx)
	case "asset_holding_get":
		return Parse_Teal_asset_holding_get_op(ctx)
	case "asset_params_get":
		return Parse_Teal_asset_params_get_op(ctx)
	case "app_params_get":
		return Parse_Teal_app_params_get_op(ctx)
	case "acct_params_get":
		return Parse_Teal_acct_params_get_op(ctx)
	case "min_balance":
		return Parse_Teal_min_balance_op(ctx)
	case "pushbytes":
		return Parse_Teal_pushbytes_op(ctx)
	case "pushint":
		return Parse_Teal_pushint_op(ctx)
	case "pushbytess":
		return Parse_Teal_pushbytess_op(ctx)
	case "pushints":
		return Parse_Teal_pushints_op(ctx)
	case "ed25519verify_bare":
		return Parse_Teal_ed25519verify_bare_op(ctx)
	case "callsub":
		return Parse_Teal_callsub_op(ctx)
	case "retsub":
		return Parse_Teal_retsub_op(ctx)
	case "proto":
		return Parse_Teal_proto_op(ctx)
	case "frame_dig":
		return Parse_Teal_frame_dig_op(ctx)
	case "frame_bury":
		return Parse_Teal_frame_bury_op(ctx)
	case "switch":
		return Parse_Teal_switch__op(ctx)
	case "match":
		return Parse_Teal_match_op(ctx)
	case "shl":
		return Parse_Teal_shl_op(ctx)
	case "shr":
		return Parse_Teal_shr_op(ctx)
	case "sqrt":
		return Parse_Teal_sqrt_op(ctx)
	case "bitlen":
		return Parse_Teal_bitlen_op(ctx)
	case "exp":
		return Parse_Teal_exp_op(ctx)
	case "expw":
		return Parse_Teal_expw_op(ctx)
	case "bsqrt":
		return Parse_Teal_bsqrt_op(ctx)
	case "divw":
		return Parse_Teal_divw_op(ctx)
	case "sha3_256":
		return Parse_Teal_sha3_256_op(ctx)
	case "b+":
		return Parse_Teal_bplus_op(ctx)
	case "b-":
		return Parse_Teal_bminus_op(ctx)
	case "b/":
		return Parse_Teal_bdiv_op(ctx)
	case "b*":
		return Parse_Teal_bmul_op(ctx)
	case "b<":
		return Parse_Teal_blt_op(ctx)
	case "b>":
		return Parse_Teal_bgt_op(ctx)
	case "b<=":
		return Parse_Teal_blteq_op(ctx)
	case "b>=":
		return Parse_Teal_bgteq_op(ctx)
	case "b==":
		return Parse_Teal_beqeq_op(ctx)
	case "b!=":
		return Parse_Teal_bnoteq_op(ctx)
	case "b%":
		return Parse_Teal_bmod_op(ctx)
	case "b|":
		return Parse_Teal_bor_op(ctx)
	case "b&":
		return Parse_Teal_band_op(ctx)
	case "b^":
		return Parse_Teal_bxor_op(ctx)
	case "b~":
		return Parse_Teal_binv_op(ctx)
	case "bzero":
		return Parse_Teal_bzero_op(ctx)
	case "log":
		return Parse_Teal_log_op(ctx)
	case "itxn_begin":
		return Parse_Teal_itxn_begin_op(ctx)
	case "itxn_field":
		return Parse_Teal_itxn_field_op(ctx)
	case "itxn_submit":
		return Parse_Teal_itxn_submit_op(ctx)
	case "itxn":
		return Parse_Teal_itxn_op(ctx)
	case "itxna":
		return Parse_Teal_itxna_op(ctx)
	case "itxn_next":
		return Parse_Teal_itxn_next_op(ctx)
	case "gitxn":
		return Parse_Teal_gitxn_op(ctx)
	case "gitxna":
		return Parse_Teal_gitxna_op(ctx)
	case "box_create":
		return Parse_Teal_box_create_op(ctx)
	case "box_extract":
		return Parse_Teal_box_extract_op(ctx)
	case "box_replace":
		return Parse_Teal_box_replace_op(ctx)
	case "box_del":
		return Parse_Teal_box_del_op(ctx)
	case "box_len":
		return Parse_Teal_box_len_op(ctx)
	case "box_get":
		return Parse_Teal_box_get_op(ctx)
	case "box_put":
		return Parse_Teal_box_put_op(ctx)
	case "txnas":
		return Parse_Teal_txnas_op(ctx)
	case "gtxnas":
		return Parse_Teal_gtxnas_op(ctx)
	case "gtxnsas":
		return Parse_Teal_gtxnsas_op(ctx)
	case "args":
		return Parse_Teal_args_op(ctx)
	case "gloadss":
		return Parse_Teal_gloadss_op(ctx)
	case "itxnas":
		return Parse_Teal_itxnas_op(ctx)
	case "gitxnas":
		return Parse_Teal_gitxnas_op(ctx)
	case "vrf_verify":
		return Parse_Teal_vrf_verify_op(ctx)
	case "block":
		return Parse_Teal_block_op(ctx)
	default:
		panic(fmt.Sprintf("unexpected op: %s", value))
	}
}
