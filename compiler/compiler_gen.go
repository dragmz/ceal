package compiler

import "strings"

const AvmMainName = "avm_main"
const AvmVersion = 8

type AstStatement interface {
	String() string
}

type BuiltinFunctionParamData struct {
	t    string
	name string
}

type BuiltinFunctionData struct {
	t    string
	name string
	op   string

	stack []BuiltinFunctionParamData
	imm   []BuiltinFunctionParamData
}

var builtin_functions = []BuiltinFunctionData{
	{
		t: "void", name: "avm_err", op: "err",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_sha256", op: "sha256",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_keccak256", op: "keccak256",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_sha512_256", op: "sha512_256",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_ed25519verify", op: "ed25519verify",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
			{t: "bytes", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_ecdsa_verify", op: "ecdsa_verify",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
			{t: "bytes", name: "s3"},
			{t: "bytes", name: "s4"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_ecdsa_pk_decompress_result_t", name: "avm_ecdsa_pk_decompress", op: "ecdsa_pk_decompress",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "avm_ecdsa_pk_recover_result_t", name: "avm_ecdsa_pk_recover", op: "ecdsa_pk_recover",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
			{t: "bytes", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_plus", op: "+",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_minus", op: "-",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_div", op: "/",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_mul", op: "*",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_lt", op: "<",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_gt", op: ">",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_lteq", op: "<=",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_gteq", op: ">=",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_andand", op: "&&",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_oror", op: "||",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_eqeq", op: "==",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
			{t: "any", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_noteq", op: "!=",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
			{t: "any", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_not", op: "!",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_len", op: "len",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_itob", op: "itob",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_btoi", op: "btoi",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_mod", op: "%",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_or", op: "|",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_and", op: "&",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_xor", op: "^",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_inv", op: "~",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_mulw_result_t", name: "avm_mulw", op: "mulw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_addw_result_t", name: "avm_addw", op: "addw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_divmodw_result_t", name: "avm_divmodw", op: "divmodw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
			{t: "uint64", name: "s3"},
			{t: "uint64", name: "s4"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_intcblock", op: "intcblock",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_intc", op: "intc",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "uint64", name: "avm_intc_0", op: "intc_0",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_intc_1", op: "intc_1",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_intc_2", op: "intc_2",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_intc_3", op: "intc_3",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_bytecblock", op: "bytecblock",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bytec", op: "bytec",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "bytes", name: "avm_bytec_0", op: "bytec_0",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bytec_1", op: "bytec_1",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bytec_2", op: "bytec_2",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bytec_3", op: "bytec_3",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_arg", op: "arg",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "bytes", name: "avm_arg_0", op: "arg_0",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_arg_1", op: "arg_1",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_arg_2", op: "arg_2",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_arg_3", op: "arg_3",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_txn", op: "txn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "any", name: "avm_global", op: "global",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint64", name: "i1"},
		},
	},
	{
		t: "any", name: "avm_gtxn", op: "gtxn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
			{t: "uint64", name: "i2"},
		},
	},
	{
		t: "any", name: "avm_load", op: "load",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "void", name: "avm_store", op: "store",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "any", name: "avm_txna", op: "txna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
			{t: "bytes", name: "i2"},
		},
	},
	{
		t: "any", name: "avm_gtxna", op: "gtxna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
			{t: "bytes", name: "i2"},
			{t: "uint64", name: "i3"},
		},
	},
	{
		t: "any", name: "avm_gtxns", op: "gtxns",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "any", name: "avm_gtxnsa", op: "gtxnsa",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "any", name: "avm_gload", op: "gload",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
			{t: "bytes", name: "i2"},
		},
	},
	{
		t: "any", name: "avm_gloads", op: "gloads",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "uint64", name: "avm_gaid", op: "gaid",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "uint64", name: "avm_gaids", op: "gaids",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_loads", op: "loads",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_stores", op: "stores",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "any", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_bnz", op: "bnz",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "void", name: "avm_bz", op: "bz",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "void", name: "avm_b", op: "b",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
			{t: "bytes", name: "i2"},
		},
	},
	{
		t: "void", name: "avm_return_", op: "return",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_assert", op: "assert",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_bury", op: "bury",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "void", name: "avm_popn", op: "popn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "void", name: "avm_dupn", op: "dupn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "void", name: "avm_pop", op: "pop",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_dup_result_t", name: "avm_dup", op: "dup",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_dup2_result_t", name: "avm_dup2", op: "dup2",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
			{t: "any", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_dig_result_t", name: "avm_dig", op: "dig",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "avm_swap_result_t", name: "avm_swap", op: "swap",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
			{t: "any", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_select", op: "select",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
			{t: "any", name: "s2"},
			{t: "uint64", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_cover", op: "cover",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "any", name: "avm_uncover", op: "uncover",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "bytes", name: "avm_concat", op: "concat",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_substring", op: "substring",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "bytes", name: "avm_substring3", op: "substring3",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
			{t: "uint64", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_getbit", op: "getbit",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_setbit", op: "setbit",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
			{t: "uint64", name: "s2"},
			{t: "uint64", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_getbyte", op: "getbyte",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_setbyte", op: "setbyte",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
			{t: "uint64", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_extract", op: "extract",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "bytes", name: "avm_extract3", op: "extract3",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
			{t: "uint64", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_extract_uint16", op: "extract_uint16",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_extract_uint32", op: "extract_uint32",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_extract_uint64", op: "extract_uint64",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_replace2", op: "replace2",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_replace3", op: "replace3",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
			{t: "bytes", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_base64_decode", op: "base64_decode",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "any", name: "i1"},
		},
	},
	{
		t: "any", name: "avm_json_ref", op: "json_ref",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_balance", op: "balance",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_app_opted_in", op: "app_opted_in",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_app_local_get", op: "app_local_get",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_app_local_get_ex_result_t", name: "avm_app_local_get_ex", op: "app_local_get_ex",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
			{t: "uint64", name: "s2"},
			{t: "bytes", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_app_global_get", op: "app_global_get",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_app_global_get_ex_result_t", name: "avm_app_global_get_ex", op: "app_global_get_ex",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_app_local_put", op: "app_local_put",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
			{t: "bytes", name: "s2"},
			{t: "any", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_app_global_put", op: "app_global_put",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "any", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_app_local_del", op: "app_local_del",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_app_global_del", op: "app_global_del",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_asset_holding_get_result_t", name: "avm_asset_holding_get", op: "asset_holding_get",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_asset_params_get_result_t", name: "avm_asset_params_get", op: "asset_params_get",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint64", name: "i1"},
		},
	},
	{
		t: "avm_app_params_get_result_t", name: "avm_app_params_get", op: "app_params_get",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "avm_acct_params_get_result_t", name: "avm_acct_params_get", op: "acct_params_get",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint64", name: "i1"},
		},
	},
	{
		t: "uint64", name: "avm_min_balance", op: "min_balance",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_pushbytes", op: "pushbytes",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_pushint", op: "pushint",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_pushbytess", op: "pushbytess",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_pushints", op: "pushints",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_ed25519verify_bare", op: "ed25519verify_bare",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
			{t: "bytes", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_callsub", op: "callsub",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
			{t: "bytes", name: "i2"},
		},
	},
	{
		t: "void", name: "avm_retsub", op: "retsub",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_proto", op: "proto",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
			{t: "bytes", name: "i2"},
		},
	},
	{
		t: "any", name: "avm_frame_dig", op: "frame_dig",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "void", name: "avm_frame_bury", op: "frame_bury",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "void", name: "avm_switch_", op: "switch",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_match", op: "match",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_shl", op: "shl",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_shr", op: "shr",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_sqrt", op: "sqrt",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_bitlen", op: "bitlen",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_exp", op: "exp",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_expw_result_t", name: "avm_expw", op: "expw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bsqrt", op: "bsqrt",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_divw", op: "divw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
			{t: "uint64", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_sha3_256", op: "sha3_256",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bplus", op: "b+",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bminus", op: "b-",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bdiv", op: "b/",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bmul", op: "b*",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_blt", op: "b<",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_bgt", op: "b>",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_blteq", op: "b<=",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_bgteq", op: "b>=",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_beqeq", op: "b==",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_bnoteq", op: "b!=",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bmod", op: "b%",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bor", op: "b|",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_band", op: "b&",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bxor", op: "b^",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_binv", op: "b~",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bzero", op: "bzero",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_log", op: "log",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_itxn_begin", op: "itxn_begin",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_itxn_field", op: "itxn_field",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "void", name: "avm_itxn_submit", op: "itxn_submit",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_itxn", op: "itxn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "any", name: "avm_itxna", op: "itxna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
			{t: "bytes", name: "i2"},
		},
	},
	{
		t: "void", name: "avm_itxn_next", op: "itxn_next",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_gitxn", op: "gitxn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
			{t: "uint64", name: "i2"},
		},
	},
	{
		t: "any", name: "avm_gitxna", op: "gitxna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
			{t: "bytes", name: "i2"},
			{t: "uint64", name: "i3"},
		},
	},
	{
		t: "uint64", name: "avm_box_create", op: "box_create",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_box_extract", op: "box_extract",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
			{t: "uint64", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_box_replace", op: "box_replace",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
			{t: "bytes", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_box_del", op: "box_del",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_box_len_result_t", name: "avm_box_len", op: "box_len",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_box_get_result_t", name: "avm_box_get", op: "box_get",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_box_put", op: "box_put",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_txnas", op: "txnas",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "any", name: "avm_gtxnas", op: "gtxnas",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "any", name: "avm_gtxnsas", op: "gtxnsas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_args", op: "args",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_gloadss", op: "gloadss",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_itxnas", op: "itxnas",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "any", name: "avm_gitxnas", op: "gitxnas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
	{
		t: "avm_vrf_verify_result_t", name: "avm_vrf_verify", op: "vrf_verify",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_block", op: "block",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "i1"},
		},
	},
}

type BuiltinStructFieldData struct {
	t    string
	name string
}

type BuiltinStructFunctionParamData struct {
	t    string
	name string
}

type BuiltinStructFunctionData struct {
	t      string
	name   string
	params []BuiltinStructFunctionParamData
}

type BuiltinStructData struct {
	name      string
	op        string
	fields    []BuiltinStructFieldData
	functions []BuiltinStructFunctionData
}

var builtin_structs = []BuiltinStructData{
	{
		name: "avm_ecdsa_verify_t", op: "ecdsa_verify",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "void", name: "Secp256k1",
				params: []BuiltinStructFunctionParamData{},
			},
			{
				t: "void", name: "Secp256r1",
				params: []BuiltinStructFunctionParamData{},
			},
		},
	},
	{
		name: "avm_ecdsa_pk_decompress_t", op: "ecdsa_pk_decompress",
		fields: []BuiltinStructFieldData{
			{"void", "Secp256k1"},
			{"void", "Secp256r1"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_ecdsa_pk_decompress_result_t", op: "ecdsa_pk_decompress",
		fields: []BuiltinStructFieldData{
			{"bytes", "r1"},
			{"bytes", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_ecdsa_pk_recover_t", op: "ecdsa_pk_recover",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "void", name: "Secp256k1",
				params: []BuiltinStructFunctionParamData{},
			},
			{
				t: "void", name: "Secp256r1",
				params: []BuiltinStructFunctionParamData{},
			},
		},
	},
	{
		name: "avm_ecdsa_pk_recover_result_t", op: "ecdsa_pk_recover",
		fields: []BuiltinStructFieldData{
			{"bytes", "r1"},
			{"bytes", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_mulw_result_t", op: "mulw",
		fields: []BuiltinStructFieldData{
			{"uint64", "r1"},
			{"uint64", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_addw_result_t", op: "addw",
		fields: []BuiltinStructFieldData{
			{"uint64", "r1"},
			{"uint64", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_divmodw_result_t", op: "divmodw",
		fields: []BuiltinStructFieldData{
			{"uint64", "r1"},
			{"uint64", "r2"},
			{"uint64", "r3"},
			{"uint64", "r4"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_txn_t", op: "txn",
		fields: []BuiltinStructFieldData{
			{"bytes", "Sender"},
			{"uint64", "Fee"},
			{"uint64", "FirstValid"},
			{"uint64", "FirstValidTime"},
			{"uint64", "LastValid"},
			{"bytes", "Note"},
			{"bytes", "Lease"},
			{"bytes", "Receiver"},
			{"uint64", "Amount"},
			{"bytes", "CloseRemainderTo"},
			{"bytes", "VotePK"},
			{"bytes", "SelectionPK"},
			{"uint64", "VoteFirst"},
			{"uint64", "VoteLast"},
			{"uint64", "VoteKeyDilution"},
			{"bytes", "Type"},
			{"uint64", "TypeEnum"},
			{"uint64", "XferAsset"},
			{"uint64", "AssetAmount"},
			{"bytes", "AssetSender"},
			{"bytes", "AssetReceiver"},
			{"bytes", "AssetCloseTo"},
			{"uint64", "GroupIndex"},
			{"bytes", "TxID"},
			{"uint64", "ApplicationID"},
			{"uint64", "OnCompletion"},
			{"bytes", "ApplicationArgs"},
			{"uint64", "NumAppArgs"},
			{"bytes", "Accounts"},
			{"uint64", "NumAccounts"},
			{"bytes", "ApprovalProgram"},
			{"bytes", "ClearStateProgram"},
			{"bytes", "RekeyTo"},
			{"uint64", "ConfigAsset"},
			{"uint64", "ConfigAssetTotal"},
			{"uint64", "ConfigAssetDecimals"},
			{"uint64", "ConfigAssetDefaultFrozen"},
			{"bytes", "ConfigAssetUnitName"},
			{"bytes", "ConfigAssetName"},
			{"bytes", "ConfigAssetURL"},
			{"bytes", "ConfigAssetMetadataHash"},
			{"bytes", "ConfigAssetManager"},
			{"bytes", "ConfigAssetReserve"},
			{"bytes", "ConfigAssetFreeze"},
			{"bytes", "ConfigAssetClawback"},
			{"uint64", "FreezeAsset"},
			{"bytes", "FreezeAssetAccount"},
			{"uint64", "FreezeAssetFrozen"},
			{"uint64", "Assets"},
			{"uint64", "NumAssets"},
			{"uint64", "Applications"},
			{"uint64", "NumApplications"},
			{"uint64", "GlobalNumUint"},
			{"uint64", "GlobalNumByteSlice"},
			{"uint64", "LocalNumUint"},
			{"uint64", "LocalNumByteSlice"},
			{"uint64", "ExtraProgramPages"},
			{"uint64", "Nonparticipation"},
			{"bytes", "Logs"},
			{"uint64", "NumLogs"},
			{"uint64", "CreatedAssetID"},
			{"uint64", "CreatedApplicationID"},
			{"bytes", "LastLog"},
			{"bytes", "StateProofPK"},
			{"bytes", "ApprovalProgramPages"},
			{"uint64", "NumApprovalProgramPages"},
			{"bytes", "ClearStateProgramPages"},
			{"uint64", "NumClearStateProgramPages"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_global_t", op: "global",
		fields: []BuiltinStructFieldData{
			{"uint64", "MinTxnFee"},
			{"uint64", "MinBalance"},
			{"uint64", "MaxTxnLife"},
			{"bytes", "ZeroAddress"},
			{"uint64", "GroupSize"},
			{"uint64", "LogicSigVersion"},
			{"uint64", "Round"},
			{"uint64", "LatestTimestamp"},
			{"uint64", "CurrentApplicationID"},
			{"bytes", "CreatorAddress"},
			{"bytes", "CurrentApplicationAddress"},
			{"bytes", "GroupID"},
			{"uint64", "OpcodeBudget"},
			{"uint64", "CallerApplicationID"},
			{"bytes", "CallerApplicationAddress"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_gtxn_t", op: "gtxn",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "Sender",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Fee",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "FirstValid",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "FirstValidTime",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "LastValid",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Note",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Lease",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Receiver",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Amount",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "CloseRemainderTo",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "VotePK",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "SelectionPK",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "VoteFirst",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "VoteLast",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "VoteKeyDilution",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Type",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "TypeEnum",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "XferAsset",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "AssetAmount",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "AssetSender",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "AssetReceiver",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "AssetCloseTo",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "GroupIndex",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "TxID",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "ApplicationID",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "OnCompletion",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ApplicationArgs",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumAppArgs",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Accounts",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumAccounts",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgram",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgram",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "RekeyTo",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "ConfigAsset",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "ConfigAssetTotal",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "ConfigAssetDecimals",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "ConfigAssetDefaultFrozen",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetUnitName",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetName",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetURL",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetMetadataHash",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetManager",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetReserve",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetFreeze",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetClawback",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "FreezeAsset",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "FreezeAssetAccount",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "FreezeAssetFrozen",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Assets",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumAssets",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Applications",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumApplications",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "GlobalNumUint",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "GlobalNumByteSlice",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "LocalNumUint",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "LocalNumByteSlice",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "ExtraProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Nonparticipation",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Logs",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumLogs",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "CreatedAssetID",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "CreatedApplicationID",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "LastLog",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "StateProofPK",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumApprovalProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumClearStateProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
		},
	},
	{
		name: "avm_txna_t", op: "txna",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "ApplicationArgs",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Accounts",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Assets",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Applications",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Logs",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
		},
	},
	{
		name: "avm_gtxna_t", op: "gtxna",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "ApplicationArgs",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
			{
				t: "bytes", name: "Accounts",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
			{
				t: "uint64", name: "Assets",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
			{
				t: "uint64", name: "Applications",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
			{
				t: "bytes", name: "Logs",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
		},
	},
	{
		name: "avm_gtxns_t", op: "gtxns",
		fields: []BuiltinStructFieldData{
			{"bytes", "Sender"},
			{"uint64", "Fee"},
			{"uint64", "FirstValid"},
			{"uint64", "FirstValidTime"},
			{"uint64", "LastValid"},
			{"bytes", "Note"},
			{"bytes", "Lease"},
			{"bytes", "Receiver"},
			{"uint64", "Amount"},
			{"bytes", "CloseRemainderTo"},
			{"bytes", "VotePK"},
			{"bytes", "SelectionPK"},
			{"uint64", "VoteFirst"},
			{"uint64", "VoteLast"},
			{"uint64", "VoteKeyDilution"},
			{"bytes", "Type"},
			{"uint64", "TypeEnum"},
			{"uint64", "XferAsset"},
			{"uint64", "AssetAmount"},
			{"bytes", "AssetSender"},
			{"bytes", "AssetReceiver"},
			{"bytes", "AssetCloseTo"},
			{"uint64", "GroupIndex"},
			{"bytes", "TxID"},
			{"uint64", "ApplicationID"},
			{"uint64", "OnCompletion"},
			{"bytes", "ApplicationArgs"},
			{"uint64", "NumAppArgs"},
			{"bytes", "Accounts"},
			{"uint64", "NumAccounts"},
			{"bytes", "ApprovalProgram"},
			{"bytes", "ClearStateProgram"},
			{"bytes", "RekeyTo"},
			{"uint64", "ConfigAsset"},
			{"uint64", "ConfigAssetTotal"},
			{"uint64", "ConfigAssetDecimals"},
			{"uint64", "ConfigAssetDefaultFrozen"},
			{"bytes", "ConfigAssetUnitName"},
			{"bytes", "ConfigAssetName"},
			{"bytes", "ConfigAssetURL"},
			{"bytes", "ConfigAssetMetadataHash"},
			{"bytes", "ConfigAssetManager"},
			{"bytes", "ConfigAssetReserve"},
			{"bytes", "ConfigAssetFreeze"},
			{"bytes", "ConfigAssetClawback"},
			{"uint64", "FreezeAsset"},
			{"bytes", "FreezeAssetAccount"},
			{"uint64", "FreezeAssetFrozen"},
			{"uint64", "Assets"},
			{"uint64", "NumAssets"},
			{"uint64", "Applications"},
			{"uint64", "NumApplications"},
			{"uint64", "GlobalNumUint"},
			{"uint64", "GlobalNumByteSlice"},
			{"uint64", "LocalNumUint"},
			{"uint64", "LocalNumByteSlice"},
			{"uint64", "ExtraProgramPages"},
			{"uint64", "Nonparticipation"},
			{"bytes", "Logs"},
			{"uint64", "NumLogs"},
			{"uint64", "CreatedAssetID"},
			{"uint64", "CreatedApplicationID"},
			{"bytes", "LastLog"},
			{"bytes", "StateProofPK"},
			{"bytes", "ApprovalProgramPages"},
			{"uint64", "NumApprovalProgramPages"},
			{"bytes", "ClearStateProgramPages"},
			{"uint64", "NumClearStateProgramPages"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_gtxnsa_t", op: "gtxnsa",
		fields: []BuiltinStructFieldData{
			{"bytes", "ApplicationArgs"},
			{"bytes", "Accounts"},
			{"uint64", "Assets"},
			{"uint64", "Applications"},
			{"bytes", "Logs"},
			{"bytes", "ApprovalProgramPages"},
			{"bytes", "ClearStateProgramPages"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_dup_result_t", op: "dup",
		fields: []BuiltinStructFieldData{
			{"any", "r1"},
			{"any", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_dup2_result_t", op: "dup2",
		fields: []BuiltinStructFieldData{
			{"any", "r1"},
			{"any", "r2"},
			{"any", "r3"},
			{"any", "r4"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_dig_result_t", op: "dig",
		fields: []BuiltinStructFieldData{
			{"any", "r1"},
			{"any", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_swap_result_t", op: "swap",
		fields: []BuiltinStructFieldData{
			{"any", "r1"},
			{"any", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_base64_decode_t", op: "base64_decode",
		fields: []BuiltinStructFieldData{
			{"any", "URLEncoding"},
			{"any", "StdEncoding"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_json_ref_t", op: "json_ref",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "JSONString",
				params: []BuiltinStructFunctionParamData{},
			},
			{
				t: "uint64", name: "JSONUint64",
				params: []BuiltinStructFunctionParamData{},
			},
			{
				t: "bytes", name: "JSONObject",
				params: []BuiltinStructFunctionParamData{},
			},
		},
	},
	{
		name: "avm_app_local_get_ex_result_t", op: "app_local_get_ex",
		fields: []BuiltinStructFieldData{
			{"any", "r1"},
			{"uint64", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_app_global_get_ex_result_t", op: "app_global_get_ex",
		fields: []BuiltinStructFieldData{
			{"any", "r1"},
			{"uint64", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_asset_holding_get_t", op: "asset_holding_get",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "uint64", name: "AssetBalance",
				params: []BuiltinStructFunctionParamData{},
			},
			{
				t: "uint64", name: "AssetFrozen",
				params: []BuiltinStructFunctionParamData{},
			},
		},
	},
	{
		name: "avm_asset_holding_get_result_t", op: "asset_holding_get",
		fields: []BuiltinStructFieldData{
			{"any", "r1"},
			{"uint64", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_asset_params_get_t", op: "asset_params_get",
		fields: []BuiltinStructFieldData{
			{"uint64", "AssetTotal"},
			{"uint64", "AssetDecimals"},
			{"uint64", "AssetDefaultFrozen"},
			{"bytes", "AssetUnitName"},
			{"bytes", "AssetName"},
			{"bytes", "AssetURL"},
			{"bytes", "AssetMetadataHash"},
			{"bytes", "AssetManager"},
			{"bytes", "AssetReserve"},
			{"bytes", "AssetFreeze"},
			{"bytes", "AssetClawback"},
			{"bytes", "AssetCreator"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_asset_params_get_result_t", op: "asset_params_get",
		fields: []BuiltinStructFieldData{
			{"any", "r1"},
			{"uint64", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_app_params_get_t", op: "app_params_get",
		fields: []BuiltinStructFieldData{
			{"bytes", "AppApprovalProgram"},
			{"bytes", "AppClearStateProgram"},
			{"uint64", "AppGlobalNumUint"},
			{"uint64", "AppGlobalNumByteSlice"},
			{"uint64", "AppLocalNumUint"},
			{"uint64", "AppLocalNumByteSlice"},
			{"uint64", "AppExtraProgramPages"},
			{"bytes", "AppCreator"},
			{"bytes", "AppAddress"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_app_params_get_result_t", op: "app_params_get",
		fields: []BuiltinStructFieldData{
			{"any", "r1"},
			{"uint64", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_acct_params_get_t", op: "acct_params_get",
		fields: []BuiltinStructFieldData{
			{"uint64", "AcctBalance"},
			{"uint64", "AcctMinBalance"},
			{"bytes", "AcctAuthAddr"},
			{"uint64", "AcctTotalNumUint"},
			{"uint64", "AcctTotalNumByteSlice"},
			{"uint64", "AcctTotalExtraAppPages"},
			{"uint64", "AcctTotalAppsCreated"},
			{"uint64", "AcctTotalAppsOptedIn"},
			{"uint64", "AcctTotalAssetsCreated"},
			{"uint64", "AcctTotalAssets"},
			{"uint64", "AcctTotalBoxes"},
			{"uint64", "AcctTotalBoxBytes"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_acct_params_get_result_t", op: "acct_params_get",
		fields: []BuiltinStructFieldData{
			{"any", "r1"},
			{"uint64", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_expw_result_t", op: "expw",
		fields: []BuiltinStructFieldData{
			{"uint64", "r1"},
			{"uint64", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_itxn_field_t", op: "itxn_field",
		fields: []BuiltinStructFieldData{
			{"bytes", "Sender"},
			{"uint64", "Fee"},
			{"bytes", "Note"},
			{"bytes", "Receiver"},
			{"uint64", "Amount"},
			{"bytes", "CloseRemainderTo"},
			{"bytes", "VotePK"},
			{"bytes", "SelectionPK"},
			{"uint64", "VoteFirst"},
			{"uint64", "VoteLast"},
			{"uint64", "VoteKeyDilution"},
			{"bytes", "Type"},
			{"uint64", "TypeEnum"},
			{"uint64", "XferAsset"},
			{"uint64", "AssetAmount"},
			{"bytes", "AssetSender"},
			{"bytes", "AssetReceiver"},
			{"bytes", "AssetCloseTo"},
			{"uint64", "ApplicationID"},
			{"uint64", "OnCompletion"},
			{"bytes", "ApplicationArgs"},
			{"bytes", "Accounts"},
			{"bytes", "ApprovalProgram"},
			{"bytes", "ClearStateProgram"},
			{"bytes", "RekeyTo"},
			{"uint64", "ConfigAsset"},
			{"uint64", "ConfigAssetTotal"},
			{"uint64", "ConfigAssetDecimals"},
			{"uint64", "ConfigAssetDefaultFrozen"},
			{"bytes", "ConfigAssetUnitName"},
			{"bytes", "ConfigAssetName"},
			{"bytes", "ConfigAssetURL"},
			{"bytes", "ConfigAssetMetadataHash"},
			{"bytes", "ConfigAssetManager"},
			{"bytes", "ConfigAssetReserve"},
			{"bytes", "ConfigAssetFreeze"},
			{"bytes", "ConfigAssetClawback"},
			{"uint64", "FreezeAsset"},
			{"bytes", "FreezeAssetAccount"},
			{"uint64", "FreezeAssetFrozen"},
			{"uint64", "Assets"},
			{"uint64", "Applications"},
			{"uint64", "GlobalNumUint"},
			{"uint64", "GlobalNumByteSlice"},
			{"uint64", "LocalNumUint"},
			{"uint64", "LocalNumByteSlice"},
			{"uint64", "ExtraProgramPages"},
			{"uint64", "Nonparticipation"},
			{"bytes", "StateProofPK"},
			{"bytes", "ApprovalProgramPages"},
			{"bytes", "ClearStateProgramPages"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_itxn_t", op: "itxn",
		fields: []BuiltinStructFieldData{
			{"bytes", "Sender"},
			{"uint64", "Fee"},
			{"uint64", "FirstValid"},
			{"uint64", "FirstValidTime"},
			{"uint64", "LastValid"},
			{"bytes", "Note"},
			{"bytes", "Lease"},
			{"bytes", "Receiver"},
			{"uint64", "Amount"},
			{"bytes", "CloseRemainderTo"},
			{"bytes", "VotePK"},
			{"bytes", "SelectionPK"},
			{"uint64", "VoteFirst"},
			{"uint64", "VoteLast"},
			{"uint64", "VoteKeyDilution"},
			{"bytes", "Type"},
			{"uint64", "TypeEnum"},
			{"uint64", "XferAsset"},
			{"uint64", "AssetAmount"},
			{"bytes", "AssetSender"},
			{"bytes", "AssetReceiver"},
			{"bytes", "AssetCloseTo"},
			{"uint64", "GroupIndex"},
			{"bytes", "TxID"},
			{"uint64", "ApplicationID"},
			{"uint64", "OnCompletion"},
			{"bytes", "ApplicationArgs"},
			{"uint64", "NumAppArgs"},
			{"bytes", "Accounts"},
			{"uint64", "NumAccounts"},
			{"bytes", "ApprovalProgram"},
			{"bytes", "ClearStateProgram"},
			{"bytes", "RekeyTo"},
			{"uint64", "ConfigAsset"},
			{"uint64", "ConfigAssetTotal"},
			{"uint64", "ConfigAssetDecimals"},
			{"uint64", "ConfigAssetDefaultFrozen"},
			{"bytes", "ConfigAssetUnitName"},
			{"bytes", "ConfigAssetName"},
			{"bytes", "ConfigAssetURL"},
			{"bytes", "ConfigAssetMetadataHash"},
			{"bytes", "ConfigAssetManager"},
			{"bytes", "ConfigAssetReserve"},
			{"bytes", "ConfigAssetFreeze"},
			{"bytes", "ConfigAssetClawback"},
			{"uint64", "FreezeAsset"},
			{"bytes", "FreezeAssetAccount"},
			{"uint64", "FreezeAssetFrozen"},
			{"uint64", "Assets"},
			{"uint64", "NumAssets"},
			{"uint64", "Applications"},
			{"uint64", "NumApplications"},
			{"uint64", "GlobalNumUint"},
			{"uint64", "GlobalNumByteSlice"},
			{"uint64", "LocalNumUint"},
			{"uint64", "LocalNumByteSlice"},
			{"uint64", "ExtraProgramPages"},
			{"uint64", "Nonparticipation"},
			{"bytes", "Logs"},
			{"uint64", "NumLogs"},
			{"uint64", "CreatedAssetID"},
			{"uint64", "CreatedApplicationID"},
			{"bytes", "LastLog"},
			{"bytes", "StateProofPK"},
			{"bytes", "ApprovalProgramPages"},
			{"uint64", "NumApprovalProgramPages"},
			{"bytes", "ClearStateProgramPages"},
			{"uint64", "NumClearStateProgramPages"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_itxna_t", op: "itxna",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "ApplicationArgs",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Accounts",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Assets",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Applications",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Logs",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
				},
			},
		},
	},
	{
		name: "avm_gitxn_t", op: "gitxn",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "Sender",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Fee",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "FirstValid",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "FirstValidTime",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "LastValid",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Note",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Lease",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Receiver",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Amount",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "CloseRemainderTo",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "VotePK",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "SelectionPK",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "VoteFirst",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "VoteLast",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "VoteKeyDilution",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Type",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "TypeEnum",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "XferAsset",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "AssetAmount",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "AssetSender",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "AssetReceiver",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "AssetCloseTo",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "GroupIndex",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "TxID",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "ApplicationID",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "OnCompletion",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ApplicationArgs",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumAppArgs",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Accounts",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumAccounts",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgram",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgram",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "RekeyTo",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "ConfigAsset",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "ConfigAssetTotal",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "ConfigAssetDecimals",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "ConfigAssetDefaultFrozen",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetUnitName",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetName",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetURL",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetMetadataHash",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetManager",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetReserve",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetFreeze",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetClawback",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "FreezeAsset",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "FreezeAssetAccount",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "FreezeAssetFrozen",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Assets",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumAssets",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Applications",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumApplications",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "GlobalNumUint",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "GlobalNumByteSlice",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "LocalNumUint",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "LocalNumByteSlice",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "ExtraProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "Nonparticipation",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "Logs",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumLogs",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "CreatedAssetID",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "CreatedApplicationID",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "LastLog",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "StateProofPK",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumApprovalProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
			{
				t: "uint64", name: "NumClearStateProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "uint64", name: "i2"},
				},
			},
		},
	},
	{
		name: "avm_gitxna_t", op: "gitxna",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "ApplicationArgs",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
			{
				t: "bytes", name: "Accounts",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
			{
				t: "uint64", name: "Assets",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
			{
				t: "uint64", name: "Applications",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
			{
				t: "bytes", name: "Logs",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages",
				params: []BuiltinStructFunctionParamData{
					{t: "bytes", name: "i2"},
					{t: "uint64", name: "i3"},
				},
			},
		},
	},
	{
		name: "avm_box_len_result_t", op: "box_len",
		fields: []BuiltinStructFieldData{
			{"uint64", "r1"},
			{"uint64", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_box_get_result_t", op: "box_get",
		fields: []BuiltinStructFieldData{
			{"bytes", "r1"},
			{"uint64", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_txnas_t", op: "txnas",
		fields: []BuiltinStructFieldData{
			{"bytes", "ApplicationArgs"},
			{"bytes", "Accounts"},
			{"uint64", "Assets"},
			{"uint64", "Applications"},
			{"bytes", "Logs"},
			{"bytes", "ApprovalProgramPages"},
			{"bytes", "ClearStateProgramPages"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_gtxnas_t", op: "gtxnas",
		fields: []BuiltinStructFieldData{
			{"bytes", "ApplicationArgs"},
			{"bytes", "Accounts"},
			{"uint64", "Assets"},
			{"uint64", "Applications"},
			{"bytes", "Logs"},
			{"bytes", "ApprovalProgramPages"},
			{"bytes", "ClearStateProgramPages"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_gtxnsas_t", op: "gtxnsas",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "ApplicationArgs",
				params: []BuiltinStructFunctionParamData{},
			},
			{
				t: "bytes", name: "Accounts",
				params: []BuiltinStructFunctionParamData{},
			},
			{
				t: "uint64", name: "Assets",
				params: []BuiltinStructFunctionParamData{},
			},
			{
				t: "uint64", name: "Applications",
				params: []BuiltinStructFunctionParamData{},
			},
			{
				t: "bytes", name: "Logs",
				params: []BuiltinStructFunctionParamData{},
			},
			{
				t: "bytes", name: "ApprovalProgramPages",
				params: []BuiltinStructFunctionParamData{},
			},
			{
				t: "bytes", name: "ClearStateProgramPages",
				params: []BuiltinStructFunctionParamData{},
			},
		},
	},
	{
		name: "avm_vrf_verify_t", op: "vrf_verify",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "void", name: "VrfAlgorand",
				params: []BuiltinStructFunctionParamData{},
			},
		},
	},
	{
		name: "avm_vrf_verify_result_t", op: "vrf_verify",
		fields: []BuiltinStructFieldData{
			{"bytes", "r1"},
			{"uint64", "r2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_block_t", op: "block",
		fields: []BuiltinStructFieldData{
			{"bytes", "BlkSeed"},
			{"uint64", "BlkTimestamp"},
		},
		functions: []BuiltinStructFunctionData{},
	},
}

type BuiltinVariableData struct {
	t    string
	name string
}

var builtin_variables = []BuiltinVariableData{
	{"avm_ecdsa_pk_decompress_t", "avm_ecdsa_pk_decompress"},
	{"avm_txn_t", "avm_txn"},
	{"avm_global_t", "avm_global"},
	{"avm_gtxn_t", "avm_gtxn"},
	{"avm_txna_t", "avm_txna"},
	{"avm_gtxna_t", "avm_gtxna"},
	{"avm_gtxns_t", "avm_gtxns"},
	{"avm_gtxnsa_t", "avm_gtxnsa"},
	{"avm_base64_decode_t", "avm_base64_decode"},
	{"avm_asset_params_get_t", "avm_asset_params_get"},
	{"avm_app_params_get_t", "avm_app_params_get"},
	{"avm_acct_params_get_t", "avm_acct_params_get"},
	{"avm_itxn_field_t", "avm_itxn_field"},
	{"avm_itxn_t", "avm_itxn"},
	{"avm_itxna_t", "avm_itxna"},
	{"avm_gitxn_t", "avm_gitxn"},
	{"avm_gitxna_t", "avm_gitxna"},
	{"avm_txnas_t", "avm_txnas"},
	{"avm_gtxnas_t", "avm_gtxnas"},
	{"avm_block_t", "avm_block"},
	{"uint64", "NoOp"},
	{"uint64", "OptIn"},
	{"uint64", "DeleteApplication"},
	{"uint64", "UpdateApplication"},
	{"uint64", "CloseOut"},
	{"uint64", "ClearState"},
}

type TealExpr interface {
	String() string
}

type avm_err_Ast struct {
}

func (a *avm_err_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("err")
	return res.String()
}

type avm_sha256_Ast struct {
	s1 AstStatement
}

func (a *avm_sha256_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("sha256")
	return res.String()
}

type avm_keccak256_Ast struct {
	s1 AstStatement
}

func (a *avm_keccak256_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("keccak256")
	return res.String()
}

type avm_sha512_256_Ast struct {
	s1 AstStatement
}

func (a *avm_sha512_256_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("sha512_256")
	return res.String()
}

type avm_ed25519verify_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_ed25519verify_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("ed25519verify")
	return res.String()
}

type avm_ecdsa_verify_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
	s4 AstStatement
}

func (a *avm_ecdsa_verify_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString(a.s4.String())
	res.WriteString("\n")
	res.WriteString("ecdsa_verify")
	return res.String()
}

type avm_ecdsa_pk_decompress_Ast struct {
	i1 string
}

func (a *avm_ecdsa_pk_decompress_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("ecdsa_pk_decompress")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_ecdsa_pk_recover_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_ecdsa_pk_recover_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("ecdsa_pk_recover")
	return res.String()
}

type avm_plus_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_plus_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("+")
	return res.String()
}

type avm_minus_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_minus_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("-")
	return res.String()
}

type avm_div_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_div_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("/")
	return res.String()
}

type avm_mul_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_mul_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("*")
	return res.String()
}

type avm_lt_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_lt_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("<")
	return res.String()
}

type avm_gt_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_gt_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(">")
	return res.String()
}

type avm_lteq_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_lteq_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("<=")
	return res.String()
}

type avm_gteq_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_gteq_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(">=")
	return res.String()
}

type avm_andand_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_andand_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("&&")
	return res.String()
}

type avm_oror_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_oror_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("||")
	return res.String()
}

type avm_eqeq_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_eqeq_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("==")
	return res.String()
}

type avm_noteq_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_noteq_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("!=")
	return res.String()
}

type avm_not_Ast struct {
	s1 AstStatement
}

func (a *avm_not_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("!")
	return res.String()
}

type avm_len_Ast struct {
	s1 AstStatement
}

func (a *avm_len_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("len")
	return res.String()
}

type avm_itob_Ast struct {
	s1 AstStatement
}

func (a *avm_itob_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("itob")
	return res.String()
}

type avm_btoi_Ast struct {
	s1 AstStatement
}

func (a *avm_btoi_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("btoi")
	return res.String()
}

type avm_mod_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_mod_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("%")
	return res.String()
}

type avm_or_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_or_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("|")
	return res.String()
}

type avm_and_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_and_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("&")
	return res.String()
}

type avm_xor_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_xor_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("^")
	return res.String()
}

type avm_inv_Ast struct {
	s1 AstStatement
}

func (a *avm_inv_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("~")
	return res.String()
}

type avm_mulw_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_mulw_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("mulw")
	return res.String()
}

type avm_addw_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_addw_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("addw")
	return res.String()
}

type avm_divmodw_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
	s4 AstStatement
}

func (a *avm_divmodw_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString(a.s4.String())
	res.WriteString("\n")
	res.WriteString("divmodw")
	return res.String()
}

type avm_intcblock_Ast struct {
}

func (a *avm_intcblock_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("intcblock")
	return res.String()
}

type avm_intc_Ast struct {
	i1 string
}

func (a *avm_intc_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("intc")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_intc_0_Ast struct {
}

func (a *avm_intc_0_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("intc_0")
	return res.String()
}

type avm_intc_1_Ast struct {
}

func (a *avm_intc_1_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("intc_1")
	return res.String()
}

type avm_intc_2_Ast struct {
}

func (a *avm_intc_2_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("intc_2")
	return res.String()
}

type avm_intc_3_Ast struct {
}

func (a *avm_intc_3_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("intc_3")
	return res.String()
}

type avm_bytecblock_Ast struct {
}

func (a *avm_bytecblock_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("bytecblock")
	return res.String()
}

type avm_bytec_Ast struct {
	i1 string
}

func (a *avm_bytec_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("bytec")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_bytec_0_Ast struct {
}

func (a *avm_bytec_0_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("bytec_0")
	return res.String()
}

type avm_bytec_1_Ast struct {
}

func (a *avm_bytec_1_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("bytec_1")
	return res.String()
}

type avm_bytec_2_Ast struct {
}

func (a *avm_bytec_2_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("bytec_2")
	return res.String()
}

type avm_bytec_3_Ast struct {
}

func (a *avm_bytec_3_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("bytec_3")
	return res.String()
}

type avm_arg_Ast struct {
	i1 string
}

func (a *avm_arg_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("arg")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_arg_0_Ast struct {
}

func (a *avm_arg_0_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("arg_0")
	return res.String()
}

type avm_arg_1_Ast struct {
}

func (a *avm_arg_1_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("arg_1")
	return res.String()
}

type avm_arg_2_Ast struct {
}

func (a *avm_arg_2_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("arg_2")
	return res.String()
}

type avm_arg_3_Ast struct {
}

func (a *avm_arg_3_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("arg_3")
	return res.String()
}

type avm_txn_Ast struct {
	i1 string
}

func (a *avm_txn_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("txn")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_global_Ast struct {
	i1 string
}

func (a *avm_global_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("global")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_gtxn_Ast struct {
	i1 string
	i2 string
}

func (a *avm_gtxn_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gtxn")
	res.WriteString(" ")
	res.WriteString("i1")
	res.WriteString("i2")
	return res.String()
}

type avm_load_Ast struct {
	i1 string
}

func (a *avm_load_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("load")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_store_Ast struct {
	i1 string
}

func (a *avm_store_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("store")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_txna_Ast struct {
	i1 string
	i2 string
}

func (a *avm_txna_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("txna")
	res.WriteString(" ")
	res.WriteString("i1")
	res.WriteString("i2")
	return res.String()
}

type avm_gtxna_Ast struct {
	i1 string
	i2 string
	i3 string
}

func (a *avm_gtxna_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gtxna")
	res.WriteString(" ")
	res.WriteString("i1")
	res.WriteString("i2")
	res.WriteString("i3")
	return res.String()
}

type avm_gtxns_Ast struct {
	i1 string
}

func (a *avm_gtxns_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gtxns")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_gtxnsa_Ast struct {
	i1 string
}

func (a *avm_gtxnsa_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gtxnsa")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_gload_Ast struct {
	i1 string
	i2 string
}

func (a *avm_gload_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gload")
	res.WriteString(" ")
	res.WriteString("i1")
	res.WriteString("i2")
	return res.String()
}

type avm_gloads_Ast struct {
	i1 string
}

func (a *avm_gloads_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gloads")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_gaid_Ast struct {
	i1 string
}

func (a *avm_gaid_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gaid")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_gaids_Ast struct {
	s1 AstStatement
}

func (a *avm_gaids_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("gaids")
	return res.String()
}

type avm_loads_Ast struct {
	s1 AstStatement
}

func (a *avm_loads_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("loads")
	return res.String()
}

type avm_stores_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_stores_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("stores")
	return res.String()
}

type avm_bnz_Ast struct {
	i1 string
	s1 AstStatement
}

func (a *avm_bnz_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("bnz")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_bz_Ast struct {
	i1 string
	s1 AstStatement
}

func (a *avm_bz_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("bz")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_b_Ast struct {
	i1 string
	i2 string
}

func (a *avm_b_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("b")
	res.WriteString(" ")
	res.WriteString("i1")
	res.WriteString("i2")
	return res.String()
}

type avm_return__Ast struct {
	s1 AstStatement
}

func (a *avm_return__Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("return")
	return res.String()
}

type avm_assert_Ast struct {
	s1 AstStatement
}

func (a *avm_assert_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("assert")
	return res.String()
}

type avm_bury_Ast struct {
	i1 string
}

func (a *avm_bury_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("bury")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_popn_Ast struct {
	i1 string
}

func (a *avm_popn_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("popn")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_dupn_Ast struct {
	i1 string
}

func (a *avm_dupn_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("dupn")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_pop_Ast struct {
	s1 AstStatement
}

func (a *avm_pop_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("pop")
	return res.String()
}

type avm_dup_Ast struct {
	s1 AstStatement
}

func (a *avm_dup_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("dup")
	return res.String()
}

type avm_dup2_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_dup2_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("dup2")
	return res.String()
}

type avm_dig_Ast struct {
	i1 string
}

func (a *avm_dig_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("dig")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_swap_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_swap_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("swap")
	return res.String()
}

type avm_select_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_select_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("select")
	return res.String()
}

type avm_cover_Ast struct {
	i1 string
}

func (a *avm_cover_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("cover")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_uncover_Ast struct {
	i1 string
}

func (a *avm_uncover_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("uncover")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_concat_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_concat_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("concat")
	return res.String()
}

type avm_substring_Ast struct {
	i1 string
	s1 AstStatement
}

func (a *avm_substring_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("substring")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_substring3_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_substring3_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("substring3")
	return res.String()
}

type avm_getbit_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_getbit_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("getbit")
	return res.String()
}

type avm_setbit_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_setbit_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("setbit")
	return res.String()
}

type avm_getbyte_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_getbyte_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("getbyte")
	return res.String()
}

type avm_setbyte_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_setbyte_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("setbyte")
	return res.String()
}

type avm_extract_Ast struct {
	i1 string
	s1 AstStatement
}

func (a *avm_extract_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("extract")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_extract3_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_extract3_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("extract3")
	return res.String()
}

type avm_extract_uint16_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_extract_uint16_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("extract_uint16")
	return res.String()
}

type avm_extract_uint32_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_extract_uint32_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("extract_uint32")
	return res.String()
}

type avm_extract_uint64_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_extract_uint64_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("extract_uint64")
	return res.String()
}

type avm_replace2_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_replace2_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("replace2")
	return res.String()
}

type avm_replace3_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_replace3_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("replace3")
	return res.String()
}

type avm_base64_decode_Ast struct {
	i1 string
}

func (a *avm_base64_decode_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("base64_decode")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_json_ref_Ast struct {
	s1 AstStatement
}

func (a *avm_json_ref_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("json_ref")
	return res.String()
}

type avm_balance_Ast struct {
	s1 AstStatement
}

func (a *avm_balance_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("balance")
	return res.String()
}

type avm_app_opted_in_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_app_opted_in_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("app_opted_in")
	return res.String()
}

type avm_app_local_get_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_app_local_get_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("app_local_get")
	return res.String()
}

type avm_app_local_get_ex_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_app_local_get_ex_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("app_local_get_ex")
	return res.String()
}

type avm_app_global_get_Ast struct {
	s1 AstStatement
}

func (a *avm_app_global_get_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("app_global_get")
	return res.String()
}

type avm_app_global_get_ex_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_app_global_get_ex_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("app_global_get_ex")
	return res.String()
}

type avm_app_local_put_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_app_local_put_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("app_local_put")
	return res.String()
}

type avm_app_global_put_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_app_global_put_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("app_global_put")
	return res.String()
}

type avm_app_local_del_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_app_local_del_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("app_local_del")
	return res.String()
}

type avm_app_global_del_Ast struct {
	s1 AstStatement
}

func (a *avm_app_global_del_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("app_global_del")
	return res.String()
}

type avm_asset_holding_get_Ast struct {
	s1 AstStatement
}

func (a *avm_asset_holding_get_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("asset_holding_get")
	return res.String()
}

type avm_asset_params_get_Ast struct {
	i1 string
}

func (a *avm_asset_params_get_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("asset_params_get")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_app_params_get_Ast struct {
	i1 string
}

func (a *avm_app_params_get_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("app_params_get")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_acct_params_get_Ast struct {
	i1 string
}

func (a *avm_acct_params_get_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("acct_params_get")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_min_balance_Ast struct {
	s1 AstStatement
}

func (a *avm_min_balance_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("min_balance")
	return res.String()
}

type avm_pushbytes_Ast struct {
}

func (a *avm_pushbytes_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("pushbytes")
	return res.String()
}

type avm_pushint_Ast struct {
}

func (a *avm_pushint_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("pushint")
	return res.String()
}

type avm_pushbytess_Ast struct {
}

func (a *avm_pushbytess_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("pushbytess")
	return res.String()
}

type avm_pushints_Ast struct {
}

func (a *avm_pushints_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("pushints")
	return res.String()
}

type avm_ed25519verify_bare_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_ed25519verify_bare_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("ed25519verify_bare")
	return res.String()
}

type avm_callsub_Ast struct {
	i1 string
	i2 string
}

func (a *avm_callsub_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("callsub")
	res.WriteString(" ")
	res.WriteString("i1")
	res.WriteString("i2")
	return res.String()
}

type avm_retsub_Ast struct {
}

func (a *avm_retsub_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("retsub")
	return res.String()
}

type avm_proto_Ast struct {
	i1 string
	i2 string
}

func (a *avm_proto_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("proto")
	res.WriteString(" ")
	res.WriteString("i1")
	res.WriteString("i2")
	return res.String()
}

type avm_frame_dig_Ast struct {
	i1 string
}

func (a *avm_frame_dig_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("frame_dig")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_frame_bury_Ast struct {
	i1 string
}

func (a *avm_frame_bury_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("frame_bury")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_switch__Ast struct {
	s1 AstStatement
}

func (a *avm_switch__Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("switch")
	return res.String()
}

type avm_match_Ast struct {
}

func (a *avm_match_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("match")
	return res.String()
}

type avm_shl_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_shl_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("shl")
	return res.String()
}

type avm_shr_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_shr_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("shr")
	return res.String()
}

type avm_sqrt_Ast struct {
	s1 AstStatement
}

func (a *avm_sqrt_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("sqrt")
	return res.String()
}

type avm_bitlen_Ast struct {
	s1 AstStatement
}

func (a *avm_bitlen_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("bitlen")
	return res.String()
}

type avm_exp_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_exp_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("exp")
	return res.String()
}

type avm_expw_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_expw_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("expw")
	return res.String()
}

type avm_bsqrt_Ast struct {
	s1 AstStatement
}

func (a *avm_bsqrt_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("bsqrt")
	return res.String()
}

type avm_divw_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_divw_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("divw")
	return res.String()
}

type avm_sha3_256_Ast struct {
	s1 AstStatement
}

func (a *avm_sha3_256_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("sha3_256")
	return res.String()
}

type avm_bplus_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_bplus_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b+")
	return res.String()
}

type avm_bminus_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_bminus_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b-")
	return res.String()
}

type avm_bdiv_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_bdiv_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b/")
	return res.String()
}

type avm_bmul_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_bmul_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b*")
	return res.String()
}

type avm_blt_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_blt_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b<")
	return res.String()
}

type avm_bgt_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_bgt_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b>")
	return res.String()
}

type avm_blteq_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_blteq_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b<=")
	return res.String()
}

type avm_bgteq_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_bgteq_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b>=")
	return res.String()
}

type avm_beqeq_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_beqeq_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b==")
	return res.String()
}

type avm_bnoteq_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_bnoteq_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b!=")
	return res.String()
}

type avm_bmod_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_bmod_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b%")
	return res.String()
}

type avm_bor_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_bor_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b|")
	return res.String()
}

type avm_band_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_band_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b&")
	return res.String()
}

type avm_bxor_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_bxor_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("b^")
	return res.String()
}

type avm_binv_Ast struct {
	s1 AstStatement
}

func (a *avm_binv_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("b~")
	return res.String()
}

type avm_bzero_Ast struct {
	s1 AstStatement
}

func (a *avm_bzero_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("bzero")
	return res.String()
}

type avm_log_Ast struct {
	s1 AstStatement
}

func (a *avm_log_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("log")
	return res.String()
}

type avm_itxn_begin_Ast struct {
}

func (a *avm_itxn_begin_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("itxn_begin")
	return res.String()
}

type avm_itxn_field_Ast struct {
	i1 string
}

func (a *avm_itxn_field_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("itxn_field")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_itxn_submit_Ast struct {
}

func (a *avm_itxn_submit_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("itxn_submit")
	return res.String()
}

type avm_itxn_Ast struct {
	i1 string
}

func (a *avm_itxn_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("itxn")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_itxna_Ast struct {
	i1 string
	i2 string
}

func (a *avm_itxna_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("itxna")
	res.WriteString(" ")
	res.WriteString("i1")
	res.WriteString("i2")
	return res.String()
}

type avm_itxn_next_Ast struct {
}

func (a *avm_itxn_next_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("itxn_next")
	return res.String()
}

type avm_gitxn_Ast struct {
	i1 string
	i2 string
}

func (a *avm_gitxn_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gitxn")
	res.WriteString(" ")
	res.WriteString("i1")
	res.WriteString("i2")
	return res.String()
}

type avm_gitxna_Ast struct {
	i1 string
	i2 string
	i3 string
}

func (a *avm_gitxna_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gitxna")
	res.WriteString(" ")
	res.WriteString("i1")
	res.WriteString("i2")
	res.WriteString("i3")
	return res.String()
}

type avm_box_create_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_box_create_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("box_create")
	return res.String()
}

type avm_box_extract_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_box_extract_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("box_extract")
	return res.String()
}

type avm_box_replace_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
}

func (a *avm_box_replace_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("box_replace")
	return res.String()
}

type avm_box_del_Ast struct {
	s1 AstStatement
}

func (a *avm_box_del_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("box_del")
	return res.String()
}

type avm_box_len_Ast struct {
	s1 AstStatement
}

func (a *avm_box_len_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("box_len")
	return res.String()
}

type avm_box_get_Ast struct {
	s1 AstStatement
}

func (a *avm_box_get_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("box_get")
	return res.String()
}

type avm_box_put_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_box_put_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("box_put")
	return res.String()
}

type avm_txnas_Ast struct {
	i1 string
}

func (a *avm_txnas_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("txnas")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_gtxnas_Ast struct {
	i1 string
}

func (a *avm_gtxnas_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gtxnas")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_gtxnsas_Ast struct {
	s1 AstStatement
}

func (a *avm_gtxnsas_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("gtxnsas")
	return res.String()
}

type avm_args_Ast struct {
	s1 AstStatement
}

func (a *avm_args_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("args")
	return res.String()
}

type avm_gloadss_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_gloadss_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("gloadss")
	return res.String()
}

type avm_itxnas_Ast struct {
	i1 string
}

func (a *avm_itxnas_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("itxnas")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_gitxnas_Ast struct {
	i1 string
	s1 AstStatement
}

func (a *avm_gitxnas_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("gitxnas")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}

type avm_vrf_verify_Ast struct {
	s1 AstStatement
	s2 AstStatement
}

func (a *avm_vrf_verify_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("vrf_verify")
	return res.String()
}

type avm_block_Ast struct {
	i1 string
}

func (a *avm_block_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("block")
	res.WriteString(" ")
	res.WriteString("i1")
	return res.String()
}
