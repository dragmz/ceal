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
			{t: "bytes", name: "s5"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1"},
		},
	},
	{
		t: "avm_ecdsa_pk_decompress_result_t", name: "avm_ecdsa_pk_decompress", op: "ecdsa_pk_decompress",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1"},
		},
	},
	{
		t: "avm_ecdsa_pk_recover_result_t", name: "avm_ecdsa_pk_recover", op: "ecdsa_pk_recover",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "uint64", name: "s2"},
			{t: "bytes", name: "s3"},
			{t: "bytes", name: "s4"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1"},
		},
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
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "UINT1"},
		},
	},
	{
		t: "uint64", name: "avm_intc", op: "intc",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1"},
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
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1"},
		},
	},
	{
		t: "bytes", name: "avm_bytec", op: "bytec",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1"},
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
			{t: "uint8", name: "N1"},
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
			{t: "uint8", name: "F1"},
		},
	},
	{
		t: "any", name: "avm_global", op: "global",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
		},
	},
	{
		t: "any", name: "avm_gtxn", op: "gtxn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1"},
			{t: "uint8", name: "F2"},
		},
	},
	{
		t: "any", name: "avm_load", op: "load",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1"},
		},
	},
	{
		t: "void", name: "avm_store", op: "store",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1"},
		},
	},
	{
		t: "any", name: "avm_txna", op: "txna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
			{t: "uint8", name: "I2"},
		},
	},
	{
		t: "any", name: "avm_gtxna", op: "gtxna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1"},
			{t: "uint8", name: "F2"},
			{t: "uint8", name: "I3"},
		},
	},
	{
		t: "any", name: "avm_gtxns", op: "gtxns",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
		},
	},
	{
		t: "any", name: "avm_gtxnsa", op: "gtxnsa",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
			{t: "uint8", name: "I2"},
		},
	},
	{
		t: "any", name: "avm_gload", op: "gload",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1"},
			{t: "uint8", name: "I2"},
		},
	},
	{
		t: "any", name: "avm_gloads", op: "gloads",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1"},
		},
	},
	{
		t: "uint64", name: "avm_gaid", op: "gaid",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1"},
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
			{t: "int16", name: "TARGET1"},
		},
	},
	{
		t: "void", name: "avm_bz", op: "bz",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "int16", name: "TARGET1"},
		},
	},
	{
		t: "void", name: "avm_b", op: "b",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "int16", name: "TARGET1"},
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
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1"},
		},
	},
	{
		t: "void", name: "avm_popn", op: "popn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1"},
		},
	},
	{
		t: "void", name: "avm_dupn", op: "dupn",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1"},
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
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1"},
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
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1"},
		},
	},
	{
		t: "any", name: "avm_uncover", op: "uncover",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1"},
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
			{t: "uint8", name: "S1"},
			{t: "uint8", name: "E2"},
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
			{t: "uint8", name: "S1"},
			{t: "uint8", name: "L2"},
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
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1"},
		},
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
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "E1"},
		},
	},
	{
		t: "any", name: "avm_json_ref", op: "json_ref",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "R1"},
		},
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
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
		},
	},
	{
		t: "avm_asset_params_get_result_t", name: "avm_asset_params_get", op: "asset_params_get",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
		},
	},
	{
		t: "avm_app_params_get_result_t", name: "avm_app_params_get", op: "app_params_get",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
		},
	},
	{
		t: "avm_acct_params_get_result_t", name: "avm_acct_params_get", op: "acct_params_get",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
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
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1"},
		},
	},
	{
		t: "uint64", name: "avm_pushint", op: "pushint",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "UINT1"},
		},
	},
	{
		t: "void", name: "avm_pushbytess", op: "pushbytess",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1"},
		},
	},
	{
		t: "void", name: "avm_pushints", op: "pushints",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "UINT1"},
		},
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
			{t: "int16", name: "TARGET1"},
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
			{t: "uint8", name: "A1"},
			{t: "uint8", name: "R2"},
		},
	},
	{
		t: "any", name: "avm_frame_dig", op: "frame_dig",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "int8", name: "I1"},
		},
	},
	{
		t: "void", name: "avm_frame_bury", op: "frame_bury",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "int8", name: "I1"},
		},
	},
	{
		t: "void", name: "avm_switch_", op: "switch",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "TARGET1"},
		},
	},
	{
		t: "void", name: "avm_match", op: "match",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "TARGET1"},
		},
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
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
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
			{t: "uint8", name: "F1"},
		},
	},
	{
		t: "any", name: "avm_itxna", op: "itxna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
			{t: "uint8", name: "I2"},
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
			{t: "uint8", name: "T1"},
			{t: "uint8", name: "F2"},
		},
	},
	{
		t: "any", name: "avm_gitxna", op: "gitxna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1"},
			{t: "uint8", name: "F2"},
			{t: "uint8", name: "I3"},
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
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
		},
	},
	{
		t: "any", name: "avm_gtxnas", op: "gtxnas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1"},
			{t: "uint8", name: "F2"},
		},
	},
	{
		t: "any", name: "avm_gtxnsas", op: "gtxnsas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
			{t: "uint64", name: "s2"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
		},
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
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
		},
	},
	{
		t: "any", name: "avm_gitxnas", op: "gitxnas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1"},
			{t: "uint8", name: "F2"},
		},
	},
	{
		t: "avm_vrf_verify_result_t", name: "avm_vrf_verify", op: "vrf_verify",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "s1"},
			{t: "bytes", name: "s2"},
			{t: "bytes", name: "s3"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1"},
		},
	},
	{
		t: "any", name: "avm_block", op: "block",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "s1"},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1"},
		},
	},
}

type BuiltinStructFieldData struct {
	t    string
	name string
	fun  string
}

type BuiltinStructFunctionParamData struct {
	t    string
	name string
}

type BuiltinStructFunctionData struct {
	t      string
	name   string
	fun    string
	params []BuiltinStructFunctionParamData
}

type BuiltinStructData struct {
	name      string
	fields    []BuiltinStructFieldData
	functions []BuiltinStructFunctionData
}

var builtin_structs = []BuiltinStructData{
	{
		name: "avm_ecdsa_verify_t",
		fields: []BuiltinStructFieldData{
			{t: "void", name: "Secp256k1", fun: "avm_ecdsa_verify"},
			{t: "void", name: "Secp256r1", fun: "avm_ecdsa_verify"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_ecdsa_pk_decompress_t",
		fields: []BuiltinStructFieldData{
			{t: "void", name: "Secp256k1", fun: "avm_ecdsa_pk_decompress"},
			{t: "void", name: "Secp256r1", fun: "avm_ecdsa_pk_decompress"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_ecdsa_pk_decompress_result_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "r1", fun: "avm_ecdsa_pk_decompress"},
			{t: "bytes", name: "r2", fun: "avm_ecdsa_pk_decompress"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_ecdsa_pk_recover_t",
		fields: []BuiltinStructFieldData{
			{t: "void", name: "Secp256k1", fun: "avm_ecdsa_pk_recover"},
			{t: "void", name: "Secp256r1", fun: "avm_ecdsa_pk_recover"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_ecdsa_pk_recover_result_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "r1", fun: "avm_ecdsa_pk_recover"},
			{t: "bytes", name: "r2", fun: "avm_ecdsa_pk_recover"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_mulw_result_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64", name: "r1", fun: "avm_mulw"},
			{t: "uint64", name: "r2", fun: "avm_mulw"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_addw_result_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64", name: "r1", fun: "avm_addw"},
			{t: "uint64", name: "r2", fun: "avm_addw"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_divmodw_result_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64", name: "r1", fun: "avm_divmodw"},
			{t: "uint64", name: "r2", fun: "avm_divmodw"},
			{t: "uint64", name: "r3", fun: "avm_divmodw"},
			{t: "uint64", name: "r4", fun: "avm_divmodw"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_txn_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "Sender", fun: "avm_txn"},
			{t: "uint64", name: "Fee", fun: "avm_txn"},
			{t: "uint64", name: "FirstValid", fun: "avm_txn"},
			{t: "uint64", name: "FirstValidTime", fun: "avm_txn"},
			{t: "uint64", name: "LastValid", fun: "avm_txn"},
			{t: "bytes", name: "Note", fun: "avm_txn"},
			{t: "bytes", name: "Lease", fun: "avm_txn"},
			{t: "bytes", name: "Receiver", fun: "avm_txn"},
			{t: "uint64", name: "Amount", fun: "avm_txn"},
			{t: "bytes", name: "CloseRemainderTo", fun: "avm_txn"},
			{t: "bytes", name: "VotePK", fun: "avm_txn"},
			{t: "bytes", name: "SelectionPK", fun: "avm_txn"},
			{t: "uint64", name: "VoteFirst", fun: "avm_txn"},
			{t: "uint64", name: "VoteLast", fun: "avm_txn"},
			{t: "uint64", name: "VoteKeyDilution", fun: "avm_txn"},
			{t: "bytes", name: "Type", fun: "avm_txn"},
			{t: "uint64", name: "TypeEnum", fun: "avm_txn"},
			{t: "uint64", name: "XferAsset", fun: "avm_txn"},
			{t: "uint64", name: "AssetAmount", fun: "avm_txn"},
			{t: "bytes", name: "AssetSender", fun: "avm_txn"},
			{t: "bytes", name: "AssetReceiver", fun: "avm_txn"},
			{t: "bytes", name: "AssetCloseTo", fun: "avm_txn"},
			{t: "uint64", name: "GroupIndex", fun: "avm_txn"},
			{t: "bytes", name: "TxID", fun: "avm_txn"},
			{t: "uint64", name: "ApplicationID", fun: "avm_txn"},
			{t: "uint64", name: "OnCompletion", fun: "avm_txn"},
			{t: "bytes", name: "ApplicationArgs", fun: "avm_txn"},
			{t: "uint64", name: "NumAppArgs", fun: "avm_txn"},
			{t: "bytes", name: "Accounts", fun: "avm_txn"},
			{t: "uint64", name: "NumAccounts", fun: "avm_txn"},
			{t: "bytes", name: "ApprovalProgram", fun: "avm_txn"},
			{t: "bytes", name: "ClearStateProgram", fun: "avm_txn"},
			{t: "bytes", name: "RekeyTo", fun: "avm_txn"},
			{t: "uint64", name: "ConfigAsset", fun: "avm_txn"},
			{t: "uint64", name: "ConfigAssetTotal", fun: "avm_txn"},
			{t: "uint64", name: "ConfigAssetDecimals", fun: "avm_txn"},
			{t: "uint64", name: "ConfigAssetDefaultFrozen", fun: "avm_txn"},
			{t: "bytes", name: "ConfigAssetUnitName", fun: "avm_txn"},
			{t: "bytes", name: "ConfigAssetName", fun: "avm_txn"},
			{t: "bytes", name: "ConfigAssetURL", fun: "avm_txn"},
			{t: "bytes", name: "ConfigAssetMetadataHash", fun: "avm_txn"},
			{t: "bytes", name: "ConfigAssetManager", fun: "avm_txn"},
			{t: "bytes", name: "ConfigAssetReserve", fun: "avm_txn"},
			{t: "bytes", name: "ConfigAssetFreeze", fun: "avm_txn"},
			{t: "bytes", name: "ConfigAssetClawback", fun: "avm_txn"},
			{t: "uint64", name: "FreezeAsset", fun: "avm_txn"},
			{t: "bytes", name: "FreezeAssetAccount", fun: "avm_txn"},
			{t: "uint64", name: "FreezeAssetFrozen", fun: "avm_txn"},
			{t: "uint64", name: "Assets", fun: "avm_txn"},
			{t: "uint64", name: "NumAssets", fun: "avm_txn"},
			{t: "uint64", name: "Applications", fun: "avm_txn"},
			{t: "uint64", name: "NumApplications", fun: "avm_txn"},
			{t: "uint64", name: "GlobalNumUint", fun: "avm_txn"},
			{t: "uint64", name: "GlobalNumByteSlice", fun: "avm_txn"},
			{t: "uint64", name: "LocalNumUint", fun: "avm_txn"},
			{t: "uint64", name: "LocalNumByteSlice", fun: "avm_txn"},
			{t: "uint64", name: "ExtraProgramPages", fun: "avm_txn"},
			{t: "uint64", name: "Nonparticipation", fun: "avm_txn"},
			{t: "bytes", name: "Logs", fun: "avm_txn"},
			{t: "uint64", name: "NumLogs", fun: "avm_txn"},
			{t: "uint64", name: "CreatedAssetID", fun: "avm_txn"},
			{t: "uint64", name: "CreatedApplicationID", fun: "avm_txn"},
			{t: "bytes", name: "LastLog", fun: "avm_txn"},
			{t: "bytes", name: "StateProofPK", fun: "avm_txn"},
			{t: "bytes", name: "ApprovalProgramPages", fun: "avm_txn"},
			{t: "uint64", name: "NumApprovalProgramPages", fun: "avm_txn"},
			{t: "bytes", name: "ClearStateProgramPages", fun: "avm_txn"},
			{t: "uint64", name: "NumClearStateProgramPages", fun: "avm_txn"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_global_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64", name: "MinTxnFee", fun: "avm_global"},
			{t: "uint64", name: "MinBalance", fun: "avm_global"},
			{t: "uint64", name: "MaxTxnLife", fun: "avm_global"},
			{t: "bytes", name: "ZeroAddress", fun: "avm_global"},
			{t: "uint64", name: "GroupSize", fun: "avm_global"},
			{t: "uint64", name: "LogicSigVersion", fun: "avm_global"},
			{t: "uint64", name: "Round", fun: "avm_global"},
			{t: "uint64", name: "LatestTimestamp", fun: "avm_global"},
			{t: "uint64", name: "CurrentApplicationID", fun: "avm_global"},
			{t: "bytes", name: "CreatorAddress", fun: "avm_global"},
			{t: "bytes", name: "CurrentApplicationAddress", fun: "avm_global"},
			{t: "bytes", name: "GroupID", fun: "avm_global"},
			{t: "uint64", name: "OpcodeBudget", fun: "avm_global"},
			{t: "uint64", name: "CallerApplicationID", fun: "avm_global"},
			{t: "bytes", name: "CallerApplicationAddress", fun: "avm_global"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_gtxn_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "Sender", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "Fee", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "FirstValid", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "FirstValidTime", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "LastValid", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Note", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Lease", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Receiver", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "Amount", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "CloseRemainderTo", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "VotePK", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "SelectionPK", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "VoteFirst", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "VoteLast", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "VoteKeyDilution", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Type", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "TypeEnum", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "XferAsset", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "AssetAmount", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "AssetSender", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "AssetReceiver", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "AssetCloseTo", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "GroupIndex", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "TxID", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "ApplicationID", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "OnCompletion", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ApplicationArgs", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumAppArgs", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Accounts", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumAccounts", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgram", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgram", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "RekeyTo", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "ConfigAsset", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "ConfigAssetTotal", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "ConfigAssetDecimals", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "ConfigAssetDefaultFrozen", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetUnitName", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetName", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetURL", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetMetadataHash", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetManager", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetReserve", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetFreeze", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetClawback", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "FreezeAsset", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "FreezeAssetAccount", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "FreezeAssetFrozen", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "Assets", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumAssets", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "Applications", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumApplications", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "GlobalNumUint", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "GlobalNumByteSlice", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "LocalNumUint", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "LocalNumByteSlice", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "ExtraProgramPages", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "Nonparticipation", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Logs", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumLogs", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "CreatedAssetID", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "CreatedApplicationID", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "LastLog", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "StateProofPK", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumApprovalProgramPages", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumClearStateProgramPages", fun: "avm_gtxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
		},
	},
	{
		name:   "avm_txna_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "ApplicationArgs", fun: "avm_txna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "bytes", name: "Accounts", fun: "avm_txna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "uint64", name: "Assets", fun: "avm_txna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "uint64", name: "Applications", fun: "avm_txna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "bytes", name: "Logs", fun: "avm_txna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages", fun: "avm_txna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages", fun: "avm_txna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
		},
	},
	{
		name:   "avm_gtxna_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "ApplicationArgs", fun: "avm_gtxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
			{
				t: "bytes", name: "Accounts", fun: "avm_gtxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
			{
				t: "uint64", name: "Assets", fun: "avm_gtxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
			{
				t: "uint64", name: "Applications", fun: "avm_gtxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
			{
				t: "bytes", name: "Logs", fun: "avm_gtxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages", fun: "avm_gtxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages", fun: "avm_gtxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
		},
	},
	{
		name: "avm_gtxns_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "Sender", fun: "avm_gtxns"},
			{t: "uint64", name: "Fee", fun: "avm_gtxns"},
			{t: "uint64", name: "FirstValid", fun: "avm_gtxns"},
			{t: "uint64", name: "FirstValidTime", fun: "avm_gtxns"},
			{t: "uint64", name: "LastValid", fun: "avm_gtxns"},
			{t: "bytes", name: "Note", fun: "avm_gtxns"},
			{t: "bytes", name: "Lease", fun: "avm_gtxns"},
			{t: "bytes", name: "Receiver", fun: "avm_gtxns"},
			{t: "uint64", name: "Amount", fun: "avm_gtxns"},
			{t: "bytes", name: "CloseRemainderTo", fun: "avm_gtxns"},
			{t: "bytes", name: "VotePK", fun: "avm_gtxns"},
			{t: "bytes", name: "SelectionPK", fun: "avm_gtxns"},
			{t: "uint64", name: "VoteFirst", fun: "avm_gtxns"},
			{t: "uint64", name: "VoteLast", fun: "avm_gtxns"},
			{t: "uint64", name: "VoteKeyDilution", fun: "avm_gtxns"},
			{t: "bytes", name: "Type", fun: "avm_gtxns"},
			{t: "uint64", name: "TypeEnum", fun: "avm_gtxns"},
			{t: "uint64", name: "XferAsset", fun: "avm_gtxns"},
			{t: "uint64", name: "AssetAmount", fun: "avm_gtxns"},
			{t: "bytes", name: "AssetSender", fun: "avm_gtxns"},
			{t: "bytes", name: "AssetReceiver", fun: "avm_gtxns"},
			{t: "bytes", name: "AssetCloseTo", fun: "avm_gtxns"},
			{t: "uint64", name: "GroupIndex", fun: "avm_gtxns"},
			{t: "bytes", name: "TxID", fun: "avm_gtxns"},
			{t: "uint64", name: "ApplicationID", fun: "avm_gtxns"},
			{t: "uint64", name: "OnCompletion", fun: "avm_gtxns"},
			{t: "bytes", name: "ApplicationArgs", fun: "avm_gtxns"},
			{t: "uint64", name: "NumAppArgs", fun: "avm_gtxns"},
			{t: "bytes", name: "Accounts", fun: "avm_gtxns"},
			{t: "uint64", name: "NumAccounts", fun: "avm_gtxns"},
			{t: "bytes", name: "ApprovalProgram", fun: "avm_gtxns"},
			{t: "bytes", name: "ClearStateProgram", fun: "avm_gtxns"},
			{t: "bytes", name: "RekeyTo", fun: "avm_gtxns"},
			{t: "uint64", name: "ConfigAsset", fun: "avm_gtxns"},
			{t: "uint64", name: "ConfigAssetTotal", fun: "avm_gtxns"},
			{t: "uint64", name: "ConfigAssetDecimals", fun: "avm_gtxns"},
			{t: "uint64", name: "ConfigAssetDefaultFrozen", fun: "avm_gtxns"},
			{t: "bytes", name: "ConfigAssetUnitName", fun: "avm_gtxns"},
			{t: "bytes", name: "ConfigAssetName", fun: "avm_gtxns"},
			{t: "bytes", name: "ConfigAssetURL", fun: "avm_gtxns"},
			{t: "bytes", name: "ConfigAssetMetadataHash", fun: "avm_gtxns"},
			{t: "bytes", name: "ConfigAssetManager", fun: "avm_gtxns"},
			{t: "bytes", name: "ConfigAssetReserve", fun: "avm_gtxns"},
			{t: "bytes", name: "ConfigAssetFreeze", fun: "avm_gtxns"},
			{t: "bytes", name: "ConfigAssetClawback", fun: "avm_gtxns"},
			{t: "uint64", name: "FreezeAsset", fun: "avm_gtxns"},
			{t: "bytes", name: "FreezeAssetAccount", fun: "avm_gtxns"},
			{t: "uint64", name: "FreezeAssetFrozen", fun: "avm_gtxns"},
			{t: "uint64", name: "Assets", fun: "avm_gtxns"},
			{t: "uint64", name: "NumAssets", fun: "avm_gtxns"},
			{t: "uint64", name: "Applications", fun: "avm_gtxns"},
			{t: "uint64", name: "NumApplications", fun: "avm_gtxns"},
			{t: "uint64", name: "GlobalNumUint", fun: "avm_gtxns"},
			{t: "uint64", name: "GlobalNumByteSlice", fun: "avm_gtxns"},
			{t: "uint64", name: "LocalNumUint", fun: "avm_gtxns"},
			{t: "uint64", name: "LocalNumByteSlice", fun: "avm_gtxns"},
			{t: "uint64", name: "ExtraProgramPages", fun: "avm_gtxns"},
			{t: "uint64", name: "Nonparticipation", fun: "avm_gtxns"},
			{t: "bytes", name: "Logs", fun: "avm_gtxns"},
			{t: "uint64", name: "NumLogs", fun: "avm_gtxns"},
			{t: "uint64", name: "CreatedAssetID", fun: "avm_gtxns"},
			{t: "uint64", name: "CreatedApplicationID", fun: "avm_gtxns"},
			{t: "bytes", name: "LastLog", fun: "avm_gtxns"},
			{t: "bytes", name: "StateProofPK", fun: "avm_gtxns"},
			{t: "bytes", name: "ApprovalProgramPages", fun: "avm_gtxns"},
			{t: "uint64", name: "NumApprovalProgramPages", fun: "avm_gtxns"},
			{t: "bytes", name: "ClearStateProgramPages", fun: "avm_gtxns"},
			{t: "uint64", name: "NumClearStateProgramPages", fun: "avm_gtxns"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_gtxnsa_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "ApplicationArgs", fun: "avm_gtxnsa",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "bytes", name: "Accounts", fun: "avm_gtxnsa",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "uint64", name: "Assets", fun: "avm_gtxnsa",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "uint64", name: "Applications", fun: "avm_gtxnsa",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "bytes", name: "Logs", fun: "avm_gtxnsa",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages", fun: "avm_gtxnsa",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages", fun: "avm_gtxnsa",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
		},
	},
	{
		name: "avm_dup_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any", name: "r1", fun: "avm_dup"},
			{t: "any", name: "r2", fun: "avm_dup"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_dup2_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any", name: "r1", fun: "avm_dup2"},
			{t: "any", name: "r2", fun: "avm_dup2"},
			{t: "any", name: "r3", fun: "avm_dup2"},
			{t: "any", name: "r4", fun: "avm_dup2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_dig_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any", name: "r1", fun: "avm_dig"},
			{t: "any", name: "r2", fun: "avm_dig"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_swap_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any", name: "r1", fun: "avm_swap"},
			{t: "any", name: "r2", fun: "avm_swap"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_base64_decode_t",
		fields: []BuiltinStructFieldData{
			{t: "any", name: "URLEncoding", fun: "avm_base64_decode"},
			{t: "any", name: "StdEncoding", fun: "avm_base64_decode"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_json_ref_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "JSONString", fun: "avm_json_ref"},
			{t: "uint64", name: "JSONUint64", fun: "avm_json_ref"},
			{t: "bytes", name: "JSONObject", fun: "avm_json_ref"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_app_local_get_ex_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any", name: "r1", fun: "avm_app_local_get_ex"},
			{t: "uint64", name: "r2", fun: "avm_app_local_get_ex"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_app_global_get_ex_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any", name: "r1", fun: "avm_app_global_get_ex"},
			{t: "uint64", name: "r2", fun: "avm_app_global_get_ex"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_asset_holding_get_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64", name: "AssetBalance", fun: "avm_asset_holding_get"},
			{t: "uint64", name: "AssetFrozen", fun: "avm_asset_holding_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_asset_holding_get_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any", name: "r1", fun: "avm_asset_holding_get"},
			{t: "uint64", name: "r2", fun: "avm_asset_holding_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_asset_params_get_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64", name: "AssetTotal", fun: "avm_asset_params_get"},
			{t: "uint64", name: "AssetDecimals", fun: "avm_asset_params_get"},
			{t: "uint64", name: "AssetDefaultFrozen", fun: "avm_asset_params_get"},
			{t: "bytes", name: "AssetUnitName", fun: "avm_asset_params_get"},
			{t: "bytes", name: "AssetName", fun: "avm_asset_params_get"},
			{t: "bytes", name: "AssetURL", fun: "avm_asset_params_get"},
			{t: "bytes", name: "AssetMetadataHash", fun: "avm_asset_params_get"},
			{t: "bytes", name: "AssetManager", fun: "avm_asset_params_get"},
			{t: "bytes", name: "AssetReserve", fun: "avm_asset_params_get"},
			{t: "bytes", name: "AssetFreeze", fun: "avm_asset_params_get"},
			{t: "bytes", name: "AssetClawback", fun: "avm_asset_params_get"},
			{t: "bytes", name: "AssetCreator", fun: "avm_asset_params_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_asset_params_get_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any", name: "r1", fun: "avm_asset_params_get"},
			{t: "uint64", name: "r2", fun: "avm_asset_params_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_app_params_get_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "AppApprovalProgram", fun: "avm_app_params_get"},
			{t: "bytes", name: "AppClearStateProgram", fun: "avm_app_params_get"},
			{t: "uint64", name: "AppGlobalNumUint", fun: "avm_app_params_get"},
			{t: "uint64", name: "AppGlobalNumByteSlice", fun: "avm_app_params_get"},
			{t: "uint64", name: "AppLocalNumUint", fun: "avm_app_params_get"},
			{t: "uint64", name: "AppLocalNumByteSlice", fun: "avm_app_params_get"},
			{t: "uint64", name: "AppExtraProgramPages", fun: "avm_app_params_get"},
			{t: "bytes", name: "AppCreator", fun: "avm_app_params_get"},
			{t: "bytes", name: "AppAddress", fun: "avm_app_params_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_app_params_get_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any", name: "r1", fun: "avm_app_params_get"},
			{t: "uint64", name: "r2", fun: "avm_app_params_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_acct_params_get_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64", name: "AcctBalance", fun: "avm_acct_params_get"},
			{t: "uint64", name: "AcctMinBalance", fun: "avm_acct_params_get"},
			{t: "bytes", name: "AcctAuthAddr", fun: "avm_acct_params_get"},
			{t: "uint64", name: "AcctTotalNumUint", fun: "avm_acct_params_get"},
			{t: "uint64", name: "AcctTotalNumByteSlice", fun: "avm_acct_params_get"},
			{t: "uint64", name: "AcctTotalExtraAppPages", fun: "avm_acct_params_get"},
			{t: "uint64", name: "AcctTotalAppsCreated", fun: "avm_acct_params_get"},
			{t: "uint64", name: "AcctTotalAppsOptedIn", fun: "avm_acct_params_get"},
			{t: "uint64", name: "AcctTotalAssetsCreated", fun: "avm_acct_params_get"},
			{t: "uint64", name: "AcctTotalAssets", fun: "avm_acct_params_get"},
			{t: "uint64", name: "AcctTotalBoxes", fun: "avm_acct_params_get"},
			{t: "uint64", name: "AcctTotalBoxBytes", fun: "avm_acct_params_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_acct_params_get_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any", name: "r1", fun: "avm_acct_params_get"},
			{t: "uint64", name: "r2", fun: "avm_acct_params_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_expw_result_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64", name: "r1", fun: "avm_expw"},
			{t: "uint64", name: "r2", fun: "avm_expw"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_itxn_field_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "Sender", fun: "avm_itxn_field"},
			{t: "uint64", name: "Fee", fun: "avm_itxn_field"},
			{t: "bytes", name: "Note", fun: "avm_itxn_field"},
			{t: "bytes", name: "Receiver", fun: "avm_itxn_field"},
			{t: "uint64", name: "Amount", fun: "avm_itxn_field"},
			{t: "bytes", name: "CloseRemainderTo", fun: "avm_itxn_field"},
			{t: "bytes", name: "VotePK", fun: "avm_itxn_field"},
			{t: "bytes", name: "SelectionPK", fun: "avm_itxn_field"},
			{t: "uint64", name: "VoteFirst", fun: "avm_itxn_field"},
			{t: "uint64", name: "VoteLast", fun: "avm_itxn_field"},
			{t: "uint64", name: "VoteKeyDilution", fun: "avm_itxn_field"},
			{t: "bytes", name: "Type", fun: "avm_itxn_field"},
			{t: "uint64", name: "TypeEnum", fun: "avm_itxn_field"},
			{t: "uint64", name: "XferAsset", fun: "avm_itxn_field"},
			{t: "uint64", name: "AssetAmount", fun: "avm_itxn_field"},
			{t: "bytes", name: "AssetSender", fun: "avm_itxn_field"},
			{t: "bytes", name: "AssetReceiver", fun: "avm_itxn_field"},
			{t: "bytes", name: "AssetCloseTo", fun: "avm_itxn_field"},
			{t: "uint64", name: "ApplicationID", fun: "avm_itxn_field"},
			{t: "uint64", name: "OnCompletion", fun: "avm_itxn_field"},
			{t: "bytes", name: "ApplicationArgs", fun: "avm_itxn_field"},
			{t: "bytes", name: "Accounts", fun: "avm_itxn_field"},
			{t: "bytes", name: "ApprovalProgram", fun: "avm_itxn_field"},
			{t: "bytes", name: "ClearStateProgram", fun: "avm_itxn_field"},
			{t: "bytes", name: "RekeyTo", fun: "avm_itxn_field"},
			{t: "uint64", name: "ConfigAsset", fun: "avm_itxn_field"},
			{t: "uint64", name: "ConfigAssetTotal", fun: "avm_itxn_field"},
			{t: "uint64", name: "ConfigAssetDecimals", fun: "avm_itxn_field"},
			{t: "uint64", name: "ConfigAssetDefaultFrozen", fun: "avm_itxn_field"},
			{t: "bytes", name: "ConfigAssetUnitName", fun: "avm_itxn_field"},
			{t: "bytes", name: "ConfigAssetName", fun: "avm_itxn_field"},
			{t: "bytes", name: "ConfigAssetURL", fun: "avm_itxn_field"},
			{t: "bytes", name: "ConfigAssetMetadataHash", fun: "avm_itxn_field"},
			{t: "bytes", name: "ConfigAssetManager", fun: "avm_itxn_field"},
			{t: "bytes", name: "ConfigAssetReserve", fun: "avm_itxn_field"},
			{t: "bytes", name: "ConfigAssetFreeze", fun: "avm_itxn_field"},
			{t: "bytes", name: "ConfigAssetClawback", fun: "avm_itxn_field"},
			{t: "uint64", name: "FreezeAsset", fun: "avm_itxn_field"},
			{t: "bytes", name: "FreezeAssetAccount", fun: "avm_itxn_field"},
			{t: "uint64", name: "FreezeAssetFrozen", fun: "avm_itxn_field"},
			{t: "uint64", name: "Assets", fun: "avm_itxn_field"},
			{t: "uint64", name: "Applications", fun: "avm_itxn_field"},
			{t: "uint64", name: "GlobalNumUint", fun: "avm_itxn_field"},
			{t: "uint64", name: "GlobalNumByteSlice", fun: "avm_itxn_field"},
			{t: "uint64", name: "LocalNumUint", fun: "avm_itxn_field"},
			{t: "uint64", name: "LocalNumByteSlice", fun: "avm_itxn_field"},
			{t: "uint64", name: "ExtraProgramPages", fun: "avm_itxn_field"},
			{t: "uint64", name: "Nonparticipation", fun: "avm_itxn_field"},
			{t: "bytes", name: "StateProofPK", fun: "avm_itxn_field"},
			{t: "bytes", name: "ApprovalProgramPages", fun: "avm_itxn_field"},
			{t: "bytes", name: "ClearStateProgramPages", fun: "avm_itxn_field"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_itxn_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "Sender", fun: "avm_itxn"},
			{t: "uint64", name: "Fee", fun: "avm_itxn"},
			{t: "uint64", name: "FirstValid", fun: "avm_itxn"},
			{t: "uint64", name: "FirstValidTime", fun: "avm_itxn"},
			{t: "uint64", name: "LastValid", fun: "avm_itxn"},
			{t: "bytes", name: "Note", fun: "avm_itxn"},
			{t: "bytes", name: "Lease", fun: "avm_itxn"},
			{t: "bytes", name: "Receiver", fun: "avm_itxn"},
			{t: "uint64", name: "Amount", fun: "avm_itxn"},
			{t: "bytes", name: "CloseRemainderTo", fun: "avm_itxn"},
			{t: "bytes", name: "VotePK", fun: "avm_itxn"},
			{t: "bytes", name: "SelectionPK", fun: "avm_itxn"},
			{t: "uint64", name: "VoteFirst", fun: "avm_itxn"},
			{t: "uint64", name: "VoteLast", fun: "avm_itxn"},
			{t: "uint64", name: "VoteKeyDilution", fun: "avm_itxn"},
			{t: "bytes", name: "Type", fun: "avm_itxn"},
			{t: "uint64", name: "TypeEnum", fun: "avm_itxn"},
			{t: "uint64", name: "XferAsset", fun: "avm_itxn"},
			{t: "uint64", name: "AssetAmount", fun: "avm_itxn"},
			{t: "bytes", name: "AssetSender", fun: "avm_itxn"},
			{t: "bytes", name: "AssetReceiver", fun: "avm_itxn"},
			{t: "bytes", name: "AssetCloseTo", fun: "avm_itxn"},
			{t: "uint64", name: "GroupIndex", fun: "avm_itxn"},
			{t: "bytes", name: "TxID", fun: "avm_itxn"},
			{t: "uint64", name: "ApplicationID", fun: "avm_itxn"},
			{t: "uint64", name: "OnCompletion", fun: "avm_itxn"},
			{t: "bytes", name: "ApplicationArgs", fun: "avm_itxn"},
			{t: "uint64", name: "NumAppArgs", fun: "avm_itxn"},
			{t: "bytes", name: "Accounts", fun: "avm_itxn"},
			{t: "uint64", name: "NumAccounts", fun: "avm_itxn"},
			{t: "bytes", name: "ApprovalProgram", fun: "avm_itxn"},
			{t: "bytes", name: "ClearStateProgram", fun: "avm_itxn"},
			{t: "bytes", name: "RekeyTo", fun: "avm_itxn"},
			{t: "uint64", name: "ConfigAsset", fun: "avm_itxn"},
			{t: "uint64", name: "ConfigAssetTotal", fun: "avm_itxn"},
			{t: "uint64", name: "ConfigAssetDecimals", fun: "avm_itxn"},
			{t: "uint64", name: "ConfigAssetDefaultFrozen", fun: "avm_itxn"},
			{t: "bytes", name: "ConfigAssetUnitName", fun: "avm_itxn"},
			{t: "bytes", name: "ConfigAssetName", fun: "avm_itxn"},
			{t: "bytes", name: "ConfigAssetURL", fun: "avm_itxn"},
			{t: "bytes", name: "ConfigAssetMetadataHash", fun: "avm_itxn"},
			{t: "bytes", name: "ConfigAssetManager", fun: "avm_itxn"},
			{t: "bytes", name: "ConfigAssetReserve", fun: "avm_itxn"},
			{t: "bytes", name: "ConfigAssetFreeze", fun: "avm_itxn"},
			{t: "bytes", name: "ConfigAssetClawback", fun: "avm_itxn"},
			{t: "uint64", name: "FreezeAsset", fun: "avm_itxn"},
			{t: "bytes", name: "FreezeAssetAccount", fun: "avm_itxn"},
			{t: "uint64", name: "FreezeAssetFrozen", fun: "avm_itxn"},
			{t: "uint64", name: "Assets", fun: "avm_itxn"},
			{t: "uint64", name: "NumAssets", fun: "avm_itxn"},
			{t: "uint64", name: "Applications", fun: "avm_itxn"},
			{t: "uint64", name: "NumApplications", fun: "avm_itxn"},
			{t: "uint64", name: "GlobalNumUint", fun: "avm_itxn"},
			{t: "uint64", name: "GlobalNumByteSlice", fun: "avm_itxn"},
			{t: "uint64", name: "LocalNumUint", fun: "avm_itxn"},
			{t: "uint64", name: "LocalNumByteSlice", fun: "avm_itxn"},
			{t: "uint64", name: "ExtraProgramPages", fun: "avm_itxn"},
			{t: "uint64", name: "Nonparticipation", fun: "avm_itxn"},
			{t: "bytes", name: "Logs", fun: "avm_itxn"},
			{t: "uint64", name: "NumLogs", fun: "avm_itxn"},
			{t: "uint64", name: "CreatedAssetID", fun: "avm_itxn"},
			{t: "uint64", name: "CreatedApplicationID", fun: "avm_itxn"},
			{t: "bytes", name: "LastLog", fun: "avm_itxn"},
			{t: "bytes", name: "StateProofPK", fun: "avm_itxn"},
			{t: "bytes", name: "ApprovalProgramPages", fun: "avm_itxn"},
			{t: "uint64", name: "NumApprovalProgramPages", fun: "avm_itxn"},
			{t: "bytes", name: "ClearStateProgramPages", fun: "avm_itxn"},
			{t: "uint64", name: "NumClearStateProgramPages", fun: "avm_itxn"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_itxna_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "ApplicationArgs", fun: "avm_itxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "bytes", name: "Accounts", fun: "avm_itxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "uint64", name: "Assets", fun: "avm_itxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "uint64", name: "Applications", fun: "avm_itxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "bytes", name: "Logs", fun: "avm_itxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages", fun: "avm_itxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages", fun: "avm_itxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "I2"},
				},
			},
		},
	},
	{
		name:   "avm_gitxn_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "Sender", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "Fee", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "FirstValid", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "FirstValidTime", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "LastValid", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Note", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Lease", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Receiver", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "Amount", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "CloseRemainderTo", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "VotePK", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "SelectionPK", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "VoteFirst", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "VoteLast", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "VoteKeyDilution", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Type", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "TypeEnum", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "XferAsset", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "AssetAmount", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "AssetSender", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "AssetReceiver", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "AssetCloseTo", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "GroupIndex", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "TxID", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "ApplicationID", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "OnCompletion", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ApplicationArgs", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumAppArgs", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Accounts", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumAccounts", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgram", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgram", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "RekeyTo", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "ConfigAsset", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "ConfigAssetTotal", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "ConfigAssetDecimals", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "ConfigAssetDefaultFrozen", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetUnitName", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetName", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetURL", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetMetadataHash", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetManager", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetReserve", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetFreeze", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ConfigAssetClawback", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "FreezeAsset", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "FreezeAssetAccount", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "FreezeAssetFrozen", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "Assets", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumAssets", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "Applications", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumApplications", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "GlobalNumUint", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "GlobalNumByteSlice", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "LocalNumUint", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "LocalNumByteSlice", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "ExtraProgramPages", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "Nonparticipation", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Logs", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumLogs", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "CreatedAssetID", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "CreatedApplicationID", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "LastLog", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "StateProofPK", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumApprovalProgramPages", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "NumClearStateProgramPages", fun: "avm_gitxn",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
		},
	},
	{
		name:   "avm_gitxna_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "ApplicationArgs", fun: "avm_gitxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
			{
				t: "bytes", name: "Accounts", fun: "avm_gitxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
			{
				t: "uint64", name: "Assets", fun: "avm_gitxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
			{
				t: "uint64", name: "Applications", fun: "avm_gitxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
			{
				t: "bytes", name: "Logs", fun: "avm_gitxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages", fun: "avm_gitxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages", fun: "avm_gitxna",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
					{t: "uint8", name: "I3"},
				},
			},
		},
	},
	{
		name: "avm_box_len_result_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64", name: "r1", fun: "avm_box_len"},
			{t: "uint64", name: "r2", fun: "avm_box_len"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_box_get_result_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "r1", fun: "avm_box_get"},
			{t: "uint64", name: "r2", fun: "avm_box_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_txnas_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "ApplicationArgs", fun: "avm_txnas"},
			{t: "bytes", name: "Accounts", fun: "avm_txnas"},
			{t: "uint64", name: "Assets", fun: "avm_txnas"},
			{t: "uint64", name: "Applications", fun: "avm_txnas"},
			{t: "bytes", name: "Logs", fun: "avm_txnas"},
			{t: "bytes", name: "ApprovalProgramPages", fun: "avm_txnas"},
			{t: "bytes", name: "ClearStateProgramPages", fun: "avm_txnas"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_gtxnas_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "bytes", name: "ApplicationArgs", fun: "avm_gtxnas",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Accounts", fun: "avm_gtxnas",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "Assets", fun: "avm_gtxnas",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "uint64", name: "Applications", fun: "avm_gtxnas",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "Logs", fun: "avm_gtxnas",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ApprovalProgramPages", fun: "avm_gtxnas",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
			{
				t: "bytes", name: "ClearStateProgramPages", fun: "avm_gtxnas",
				params: []BuiltinStructFunctionParamData{
					{t: "uint8", name: "F2"},
				},
			},
		},
	},
	{
		name: "avm_gtxnsas_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "ApplicationArgs", fun: "avm_gtxnsas"},
			{t: "bytes", name: "Accounts", fun: "avm_gtxnsas"},
			{t: "uint64", name: "Assets", fun: "avm_gtxnsas"},
			{t: "uint64", name: "Applications", fun: "avm_gtxnsas"},
			{t: "bytes", name: "Logs", fun: "avm_gtxnsas"},
			{t: "bytes", name: "ApprovalProgramPages", fun: "avm_gtxnsas"},
			{t: "bytes", name: "ClearStateProgramPages", fun: "avm_gtxnsas"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_vrf_verify_t",
		fields: []BuiltinStructFieldData{
			{t: "void", name: "VrfAlgorand", fun: "avm_vrf_verify"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_vrf_verify_result_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "r1", fun: "avm_vrf_verify"},
			{t: "uint64", name: "r2", fun: "avm_vrf_verify"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_block_t",
		fields: []BuiltinStructFieldData{
			{t: "bytes", name: "BlkSeed", fun: "avm_block"},
			{t: "uint64", name: "BlkTimestamp", fun: "avm_block"},
		},
		functions: []BuiltinStructFunctionData{},
	},
}

type BuiltinVariableData struct {
	t    string
	name string
}

var builtin_variables = []BuiltinVariableData{
	{"avm_ecdsa_verify_t", "avm_ecdsa_verify"},
	{"avm_ecdsa_pk_decompress_t", "avm_ecdsa_pk_decompress"},
	{"avm_ecdsa_pk_recover_t", "avm_ecdsa_pk_recover"},
	{"avm_txn_t", "avm_txn"},
	{"avm_global_t", "avm_global"},
	{"avm_gtxn_t", "avm_gtxn"},
	{"avm_txna_t", "avm_txna"},
	{"avm_gtxna_t", "avm_gtxna"},
	{"avm_gtxns_t", "avm_gtxns"},
	{"avm_gtxnsa_t", "avm_gtxnsa"},
	{"avm_base64_decode_t", "avm_base64_decode"},
	{"avm_json_ref_t", "avm_json_ref"},
	{"avm_asset_holding_get_t", "avm_asset_holding_get"},
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
	{"avm_gtxnsas_t", "avm_gtxnsas"},
	{"avm_vrf_verify_t", "avm_vrf_verify"},
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
	s5 AstStatement
	V1 string
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
	res.WriteString(a.s5.String())
	res.WriteString("\n")
	res.WriteString("ecdsa_verify")
	res.WriteString(" ")
	res.WriteString(a.V1)
	return res.String()
}

type avm_ecdsa_pk_decompress_Ast struct {
	s1 AstStatement
	V1 string
}

func (a *avm_ecdsa_pk_decompress_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("ecdsa_pk_decompress")
	res.WriteString(" ")
	res.WriteString(a.V1)
	return res.String()
}

type avm_ecdsa_pk_recover_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
	s4 AstStatement
	V1 string
}

func (a *avm_ecdsa_pk_recover_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString(a.s4.String())
	res.WriteString("\n")
	res.WriteString("ecdsa_pk_recover")
	res.WriteString(" ")
	res.WriteString(a.V1)
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
	UINT1 string
}

func (a *avm_intcblock_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("intcblock")
	res.WriteString(" ")
	res.WriteString(a.UINT1)
	return res.String()
}

type avm_intc_Ast struct {
	I1 string
}

func (a *avm_intc_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("intc")
	res.WriteString(" ")
	res.WriteString(a.I1)
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
	BYTES1 string
}

func (a *avm_bytecblock_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("bytecblock")
	res.WriteString(" ")
	res.WriteString(a.BYTES1)
	return res.String()
}

type avm_bytec_Ast struct {
	I1 string
}

func (a *avm_bytec_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("bytec")
	res.WriteString(" ")
	res.WriteString(a.I1)
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
	N1 string
}

func (a *avm_arg_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("arg")
	res.WriteString(" ")
	res.WriteString(a.N1)
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
	F1 string
}

func (a *avm_txn_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("txn")
	res.WriteString(" ")
	res.WriteString(a.F1)
	return res.String()
}

type avm_global_Ast struct {
	F1 string
}

func (a *avm_global_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("global")
	res.WriteString(" ")
	res.WriteString(a.F1)
	return res.String()
}

type avm_gtxn_Ast struct {
	T1 string
	F2 string
}

func (a *avm_gtxn_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gtxn")
	res.WriteString(" ")
	res.WriteString(a.T1)
	res.WriteString(" ")
	res.WriteString(a.F2)
	return res.String()
}

type avm_load_Ast struct {
	I1 string
}

func (a *avm_load_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("load")
	res.WriteString(" ")
	res.WriteString(a.I1)
	return res.String()
}

type avm_store_Ast struct {
	s1 AstStatement
	I1 string
}

func (a *avm_store_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("store")
	res.WriteString(" ")
	res.WriteString(a.I1)
	return res.String()
}

type avm_txna_Ast struct {
	F1 string
	I2 string
}

func (a *avm_txna_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("txna")
	res.WriteString(" ")
	res.WriteString(a.F1)
	res.WriteString(" ")
	res.WriteString(a.I2)
	return res.String()
}

type avm_gtxna_Ast struct {
	T1 string
	F2 string
	I3 string
}

func (a *avm_gtxna_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gtxna")
	res.WriteString(" ")
	res.WriteString(a.T1)
	res.WriteString(" ")
	res.WriteString(a.F2)
	res.WriteString(" ")
	res.WriteString(a.I3)
	return res.String()
}

type avm_gtxns_Ast struct {
	s1 AstStatement
	F1 string
}

func (a *avm_gtxns_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("gtxns")
	res.WriteString(" ")
	res.WriteString(a.F1)
	return res.String()
}

type avm_gtxnsa_Ast struct {
	s1 AstStatement
	F1 string
	I2 string
}

func (a *avm_gtxnsa_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("gtxnsa")
	res.WriteString(" ")
	res.WriteString(a.F1)
	res.WriteString(" ")
	res.WriteString(a.I2)
	return res.String()
}

type avm_gload_Ast struct {
	T1 string
	I2 string
}

func (a *avm_gload_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gload")
	res.WriteString(" ")
	res.WriteString(a.T1)
	res.WriteString(" ")
	res.WriteString(a.I2)
	return res.String()
}

type avm_gloads_Ast struct {
	s1 AstStatement
	I1 string
}

func (a *avm_gloads_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("gloads")
	res.WriteString(" ")
	res.WriteString(a.I1)
	return res.String()
}

type avm_gaid_Ast struct {
	T1 string
}

func (a *avm_gaid_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gaid")
	res.WriteString(" ")
	res.WriteString(a.T1)
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
	s1      AstStatement
	TARGET1 string
}

func (a *avm_bnz_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("bnz")
	res.WriteString(" ")
	res.WriteString(a.TARGET1)
	return res.String()
}

type avm_bz_Ast struct {
	s1      AstStatement
	TARGET1 string
}

func (a *avm_bz_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("bz")
	res.WriteString(" ")
	res.WriteString(a.TARGET1)
	return res.String()
}

type avm_b_Ast struct {
	TARGET1 string
}

func (a *avm_b_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("b")
	res.WriteString(" ")
	res.WriteString(a.TARGET1)
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
	s1 AstStatement
	N1 string
}

func (a *avm_bury_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("bury")
	res.WriteString(" ")
	res.WriteString(a.N1)
	return res.String()
}

type avm_popn_Ast struct {
	N1 string
}

func (a *avm_popn_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("popn")
	res.WriteString(" ")
	res.WriteString(a.N1)
	return res.String()
}

type avm_dupn_Ast struct {
	s1 AstStatement
	N1 string
}

func (a *avm_dupn_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("dupn")
	res.WriteString(" ")
	res.WriteString(a.N1)
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
	s1 AstStatement
	N1 string
}

func (a *avm_dig_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("dig")
	res.WriteString(" ")
	res.WriteString(a.N1)
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
	s1 AstStatement
	N1 string
}

func (a *avm_cover_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("cover")
	res.WriteString(" ")
	res.WriteString(a.N1)
	return res.String()
}

type avm_uncover_Ast struct {
	s1 AstStatement
	N1 string
}

func (a *avm_uncover_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("uncover")
	res.WriteString(" ")
	res.WriteString(a.N1)
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
	s1 AstStatement
	S1 string
	E2 string
}

func (a *avm_substring_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("substring")
	res.WriteString(" ")
	res.WriteString(a.S1)
	res.WriteString(" ")
	res.WriteString(a.E2)
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
	s1 AstStatement
	S1 string
	L2 string
}

func (a *avm_extract_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("extract")
	res.WriteString(" ")
	res.WriteString(a.S1)
	res.WriteString(" ")
	res.WriteString(a.L2)
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
	S1 string
}

func (a *avm_replace2_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("replace2")
	res.WriteString(" ")
	res.WriteString(a.S1)
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
	s1 AstStatement
	E1 string
}

func (a *avm_base64_decode_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("base64_decode")
	res.WriteString(" ")
	res.WriteString(a.E1)
	return res.String()
}

type avm_json_ref_Ast struct {
	s1 AstStatement
	s2 AstStatement
	R1 string
}

func (a *avm_json_ref_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("json_ref")
	res.WriteString(" ")
	res.WriteString(a.R1)
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
	s2 AstStatement
	F1 string
}

func (a *avm_asset_holding_get_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("asset_holding_get")
	res.WriteString(" ")
	res.WriteString(a.F1)
	return res.String()
}

type avm_asset_params_get_Ast struct {
	s1 AstStatement
	F1 string
}

func (a *avm_asset_params_get_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("asset_params_get")
	res.WriteString(" ")
	res.WriteString(a.F1)
	return res.String()
}

type avm_app_params_get_Ast struct {
	s1 AstStatement
	F1 string
}

func (a *avm_app_params_get_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("app_params_get")
	res.WriteString(" ")
	res.WriteString(a.F1)
	return res.String()
}

type avm_acct_params_get_Ast struct {
	s1 AstStatement
	F1 string
}

func (a *avm_acct_params_get_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("acct_params_get")
	res.WriteString(" ")
	res.WriteString(a.F1)
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
	BYTES1 string
}

func (a *avm_pushbytes_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("pushbytes")
	res.WriteString(" ")
	res.WriteString(a.BYTES1)
	return res.String()
}

type avm_pushint_Ast struct {
	UINT1 string
}

func (a *avm_pushint_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("pushint")
	res.WriteString(" ")
	res.WriteString(a.UINT1)
	return res.String()
}

type avm_pushbytess_Ast struct {
	BYTES1 string
}

func (a *avm_pushbytess_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("pushbytess")
	res.WriteString(" ")
	res.WriteString(a.BYTES1)
	return res.String()
}

type avm_pushints_Ast struct {
	UINT1 string
}

func (a *avm_pushints_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("pushints")
	res.WriteString(" ")
	res.WriteString(a.UINT1)
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
	TARGET1 string
}

func (a *avm_callsub_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("callsub")
	res.WriteString(" ")
	res.WriteString(a.TARGET1)
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
	A1 string
	R2 string
}

func (a *avm_proto_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("proto")
	res.WriteString(" ")
	res.WriteString(a.A1)
	res.WriteString(" ")
	res.WriteString(a.R2)
	return res.String()
}

type avm_frame_dig_Ast struct {
	I1 string
}

func (a *avm_frame_dig_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("frame_dig")
	res.WriteString(" ")
	res.WriteString(a.I1)
	return res.String()
}

type avm_frame_bury_Ast struct {
	s1 AstStatement
	I1 string
}

func (a *avm_frame_bury_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("frame_bury")
	res.WriteString(" ")
	res.WriteString(a.I1)
	return res.String()
}

type avm_switch__Ast struct {
	s1      AstStatement
	TARGET1 string
}

func (a *avm_switch__Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("switch")
	res.WriteString(" ")
	res.WriteString(a.TARGET1)
	return res.String()
}

type avm_match_Ast struct {
	TARGET1 string
}

func (a *avm_match_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("match")
	res.WriteString(" ")
	res.WriteString(a.TARGET1)
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
	s1 AstStatement
	F1 string
}

func (a *avm_itxn_field_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("itxn_field")
	res.WriteString(" ")
	res.WriteString(a.F1)
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
	F1 string
}

func (a *avm_itxn_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("itxn")
	res.WriteString(" ")
	res.WriteString(a.F1)
	return res.String()
}

type avm_itxna_Ast struct {
	F1 string
	I2 string
}

func (a *avm_itxna_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("itxna")
	res.WriteString(" ")
	res.WriteString(a.F1)
	res.WriteString(" ")
	res.WriteString(a.I2)
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
	T1 string
	F2 string
}

func (a *avm_gitxn_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gitxn")
	res.WriteString(" ")
	res.WriteString(a.T1)
	res.WriteString(" ")
	res.WriteString(a.F2)
	return res.String()
}

type avm_gitxna_Ast struct {
	T1 string
	F2 string
	I3 string
}

func (a *avm_gitxna_Ast) String() string {
	res := strings.Builder{}
	res.WriteString("gitxna")
	res.WriteString(" ")
	res.WriteString(a.T1)
	res.WriteString(" ")
	res.WriteString(a.F2)
	res.WriteString(" ")
	res.WriteString(a.I3)
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
	s1 AstStatement
	F1 string
}

func (a *avm_txnas_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("txnas")
	res.WriteString(" ")
	res.WriteString(a.F1)
	return res.String()
}

type avm_gtxnas_Ast struct {
	s1 AstStatement
	T1 string
	F2 string
}

func (a *avm_gtxnas_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("gtxnas")
	res.WriteString(" ")
	res.WriteString(a.T1)
	res.WriteString(" ")
	res.WriteString(a.F2)
	return res.String()
}

type avm_gtxnsas_Ast struct {
	s1 AstStatement
	s2 AstStatement
	F1 string
}

func (a *avm_gtxnsas_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString("gtxnsas")
	res.WriteString(" ")
	res.WriteString(a.F1)
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
	s1 AstStatement
	F1 string
}

func (a *avm_itxnas_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("itxnas")
	res.WriteString(" ")
	res.WriteString(a.F1)
	return res.String()
}

type avm_gitxnas_Ast struct {
	s1 AstStatement
	T1 string
	F2 string
}

func (a *avm_gitxnas_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("gitxnas")
	res.WriteString(" ")
	res.WriteString(a.T1)
	res.WriteString(" ")
	res.WriteString(a.F2)
	return res.String()
}

type avm_vrf_verify_Ast struct {
	s1 AstStatement
	s2 AstStatement
	s3 AstStatement
	S1 string
}

func (a *avm_vrf_verify_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString(a.s2.String())
	res.WriteString("\n")
	res.WriteString(a.s3.String())
	res.WriteString("\n")
	res.WriteString("vrf_verify")
	res.WriteString(" ")
	res.WriteString(a.S1)
	return res.String()
}

type avm_block_Ast struct {
	s1 AstStatement
	F1 string
}

func (a *avm_block_Ast) String() string {
	res := strings.Builder{}
	res.WriteString(a.s1.String())
	res.WriteString("\n")
	res.WriteString("block")
	res.WriteString(" ")
	res.WriteString(a.F1)
	return res.String()
}
