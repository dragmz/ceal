package compiler

const AvmMainName = "avm_main"
const AvmVersion = 8

type AstStatement interface {
	String() string
}

type BuiltinFunctionParamData struct {
	t     string
	name  string
	array bool
	field bool
}

type BuiltinFunctionReturnData struct {
	t    string
	name string
}

type BuiltinFunctionData struct {
	t    string
	name string
	op   string

	stack   []BuiltinFunctionParamData
	imm     []BuiltinFunctionParamData
	returns []BuiltinFunctionReturnData
}

var builtin_functions = []BuiltinFunctionData{
	{
		t: "void", name: "avm_err", op: "err",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_err_op", op: "err",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r32_byte_t", name: "avm_sha256", op: "sha256",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r32_byte_t", name: "r1"},
		},
	},
	{
		t: "r32_byte_t", name: "avm_sha256_op", op: "sha256",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r32_byte_t", name: "avm_keccak256", op: "keccak256",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r32_byte_t", name: "r1"},
		},
	},
	{
		t: "r32_byte_t", name: "avm_keccak256_op", op: "keccak256",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r32_byte_t", name: "avm_sha512_256", op: "sha512_256",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r32_byte_t", name: "r1"},
		},
	},
	{
		t: "r32_byte_t", name: "avm_sha512_256_op", op: "sha512_256",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_ed25519verify", op: "ed25519verify",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "r_byte_t", name: "STACK_2", array: false, field: false},
			{t: "r_byte_t", name: "STACK_3", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_ed25519verify_op", op: "ed25519verify",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_ecdsa_verify", op: "ecdsa_verify",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "r_byte_t", name: "STACK_2", array: false, field: false},
			{t: "r_byte_t", name: "STACK_3", array: false, field: false},
			{t: "r_byte_t", name: "STACK_4", array: false, field: false},
			{t: "r_byte_t", name: "STACK_5", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_ecdsa_verify_op", op: "ecdsa_verify",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1", array: false, field: true},
		},
	},
	{
		t: "avm_ecdsa_pk_decompress_result_t", name: "avm_ecdsa_pk_decompress", op: "ecdsa_pk_decompress",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
			{t: "r_byte_t", name: "r2"},
		},
	},
	{
		t: "avm_ecdsa_pk_decompress_result_t", name: "avm_ecdsa_pk_decompress_op", op: "ecdsa_pk_decompress",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1", array: false, field: true},
		},
	},
	{
		t: "avm_ecdsa_pk_recover_result_t", name: "avm_ecdsa_pk_recover", op: "ecdsa_pk_recover",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
			{t: "r_byte_t", name: "STACK_3", array: false, field: false},
			{t: "r_byte_t", name: "STACK_4", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
			{t: "r_byte_t", name: "r2"},
		},
	},
	{
		t: "avm_ecdsa_pk_recover_result_t", name: "avm_ecdsa_pk_recover_op", op: "ecdsa_pk_recover",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1", array: false, field: true},
		},
	},
	{
		t: "uint64_t", name: "avm_plus", op: "+",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_plus_op", op: "+",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_minus", op: "-",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_minus_op", op: "-",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_div", op: "/",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_div_op", op: "/",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_mul", op: "*",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_mul_op", op: "*",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_lt", op: "<",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_lt_op", op: "<",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_gt", op: ">",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_gt_op", op: ">",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_lteq", op: "<=",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_lteq_op", op: "<=",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_gteq", op: ">=",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_gteq_op", op: ">=",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_andand", op: "&&",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_andand_op", op: "&&",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_oror", op: "||",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_oror_op", op: "||",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_eqeq", op: "==",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "any_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_eqeq_op", op: "==",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_noteq", op: "!=",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "any_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_noteq_op", op: "!=",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_not", op: "!",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_not_op", op: "!",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_len", op: "len",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_len_op", op: "len",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_itob", op: "itob",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_itob_op", op: "itob",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_btoi", op: "btoi",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_btoi_op", op: "btoi",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_mod", op: "%",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_mod_op", op: "%",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_or", op: "|",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_or_op", op: "|",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_and", op: "&",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_and_op", op: "&",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_xor", op: "^",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_xor_op", op: "^",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_inv", op: "~",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_inv_op", op: "~",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_mulw_result_t", name: "avm_mulw", op: "mulw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
			{t: "uint64_t", name: "r2"},
		},
	},
	{
		t: "avm_mulw_result_t", name: "avm_mulw_op", op: "mulw",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_addw_result_t", name: "avm_addw", op: "addw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
			{t: "uint64_t", name: "r2"},
		},
	},
	{
		t: "avm_addw_result_t", name: "avm_addw_op", op: "addw",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_divmodw_result_t", name: "avm_divmodw", op: "divmodw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
			{t: "uint64_t", name: "STACK_3", array: false, field: false},
			{t: "uint64_t", name: "STACK_4", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
			{t: "uint64_t", name: "r2"},
			{t: "uint64_t", name: "r3"},
			{t: "uint64_t", name: "r4"},
		},
	},
	{
		t: "avm_divmodw_result_t", name: "avm_divmodw_op", op: "divmodw",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_intcblock", op: "intcblock",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint64", name: "UINT1", array: true, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_intcblock_op", op: "intcblock",
		imm: []BuiltinFunctionParamData{
			{t: "uint64", name: "UINT1", array: true, field: false},
		},
	},
	{
		t: "uint64_t", name: "avm_intc", op: "intc",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_intc_op", op: "intc",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false, field: false},
		},
	},
	{
		t: "uint64_t", name: "avm_intc_0", op: "intc_0",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_intc_0_op", op: "intc_0",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_intc_1", op: "intc_1",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_intc_1_op", op: "intc_1",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_intc_2", op: "intc_2",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_intc_2_op", op: "intc_2",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_intc_3", op: "intc_3",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_intc_3_op", op: "intc_3",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_bytecblock", op: "bytecblock",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1", array: true, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_bytecblock_op", op: "bytecblock",
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1", array: true, field: false},
		},
	},
	{
		t: "r_byte_t", name: "avm_bytec", op: "bytec",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_bytec_op", op: "bytec",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false, field: false},
		},
	},
	{
		t: "r_byte_t", name: "avm_bytec_0", op: "bytec_0",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_bytec_0_op", op: "bytec_0",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_bytec_1", op: "bytec_1",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_bytec_1_op", op: "bytec_1",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_bytec_2", op: "bytec_2",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_bytec_2_op", op: "bytec_2",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_bytec_3", op: "bytec_3",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_bytec_3_op", op: "bytec_3",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_arg", op: "arg",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_arg_op", op: "arg",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
	},
	{
		t: "r_byte_t", name: "avm_arg_0", op: "arg_0",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_arg_0_op", op: "arg_0",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_arg_1", op: "arg_1",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_arg_1_op", op: "arg_1",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_arg_2", op: "arg_2",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_arg_2_op", op: "arg_2",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_arg_3", op: "arg_3",
		stack: []BuiltinFunctionParamData{},
		imm:   []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_arg_3_op", op: "arg_3",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any_t", name: "avm_txn", op: "txn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_txn_op", op: "txn",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
	{
		t: "any_t", name: "avm_global", op: "global",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_global_op", op: "global",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
	{
		t: "any_t", name: "avm_gtxn", op: "gtxn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "F2", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_gtxn_op", op: "gtxn",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "F2", array: false, field: true},
		},
	},
	{
		t: "any_t", name: "avm_load", op: "load",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_load_op", op: "load",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false, field: false},
		},
	},
	{
		t: "void", name: "avm_store", op: "store",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_store_op", op: "store",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false, field: false},
		},
	},
	{
		t: "any_t", name: "avm_txna", op: "txna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
			{t: "uint8", name: "I2", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_txna_op", op: "txna",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
			{t: "uint8", name: "I2", array: false, field: false},
		},
	},
	{
		t: "any_t", name: "avm_gtxna", op: "gtxna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "F2", array: false, field: true},
			{t: "uint8", name: "I3", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_gtxna_op", op: "gtxna",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "F2", array: false, field: true},
			{t: "uint8", name: "I3", array: false, field: false},
		},
	},
	{
		t: "any_t", name: "avm_gtxns", op: "gtxns",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_gtxns_op", op: "gtxns",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
	{
		t: "any_t", name: "avm_gtxnsa", op: "gtxnsa",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
			{t: "uint8", name: "I2", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_gtxnsa_op", op: "gtxnsa",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
			{t: "uint8", name: "I2", array: false, field: false},
		},
	},
	{
		t: "any_t", name: "avm_gload", op: "gload",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "I2", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_gload_op", op: "gload",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "I2", array: false, field: false},
		},
	},
	{
		t: "any_t", name: "avm_gloads", op: "gloads",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_gloads_op", op: "gloads",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false, field: false},
		},
	},
	{
		t: "uint64_t", name: "avm_gaid", op: "gaid",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_gaid_op", op: "gaid",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
		},
	},
	{
		t: "uint64_t", name: "avm_gaids", op: "gaids",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_gaids_op", op: "gaids",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any_t", name: "avm_loads", op: "loads",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_loads_op", op: "loads",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_stores", op: "stores",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "any_t", name: "STACK_2", array: false, field: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_stores_op", op: "stores",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_bnz", op: "bnz",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_bnz_op", op: "bnz",
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false, field: false},
		},
	},
	{
		t: "void", name: "avm_bz", op: "bz",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_bz_op", op: "bz",
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false, field: false},
		},
	},
	{
		t: "void", name: "avm_b", op: "b",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_b_op", op: "b",
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false, field: false},
		},
	},
	{
		t: "void", name: "avm_return", op: "return",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_return_op", op: "return",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_assert", op: "assert",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_assert_op", op: "assert",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_bury", op: "bury",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_bury_op", op: "bury",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
	},
	{
		t: "void", name: "avm_popn", op: "popn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_popn_op", op: "popn",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
	},
	{
		t: "void", name: "avm_dupn", op: "dupn",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_dupn_op", op: "dupn",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
	},
	{
		t: "void", name: "avm_pop", op: "pop",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_pop_op", op: "pop",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_dup_result_t", name: "avm_dup", op: "dup",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
			{t: "any_t", name: "r2"},
		},
	},
	{
		t: "avm_dup_result_t", name: "avm_dup_op", op: "dup",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_dup2_result_t", name: "avm_dup2", op: "dup2",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "any_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
			{t: "any_t", name: "r2"},
			{t: "any_t", name: "r3"},
			{t: "any_t", name: "r4"},
		},
	},
	{
		t: "avm_dup2_result_t", name: "avm_dup2_op", op: "dup2",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_dig_result_t", name: "avm_dig", op: "dig",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
			{t: "any_t", name: "r2"},
		},
	},
	{
		t: "avm_dig_result_t", name: "avm_dig_op", op: "dig",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
	},
	{
		t: "avm_swap_result_t", name: "avm_swap", op: "swap",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "any_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
			{t: "any_t", name: "r2"},
		},
	},
	{
		t: "avm_swap_result_t", name: "avm_swap_op", op: "swap",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any_t", name: "avm_select", op: "select",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "any_t", name: "STACK_2", array: false, field: false},
			{t: "uint64_t", name: "STACK_3", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_select_op", op: "select",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any_t", name: "avm_cover", op: "cover",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_cover_op", op: "cover",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
	},
	{
		t: "any_t", name: "avm_uncover", op: "uncover",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_uncover_op", op: "uncover",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false, field: false},
		},
	},
	{
		t: "r_byte_t", name: "avm_concat", op: "concat",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "r_byte_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_concat_op", op: "concat",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_substring", op: "substring",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false, field: false},
			{t: "uint8", name: "E2", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_substring_op", op: "substring",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false, field: false},
			{t: "uint8", name: "E2", array: false, field: false},
		},
	},
	{
		t: "r_byte_t", name: "avm_substring3", op: "substring3",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
			{t: "uint64_t", name: "STACK_3", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_substring3_op", op: "substring3",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_getbit", op: "getbit",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_getbit_op", op: "getbit",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any_t", name: "avm_setbit", op: "setbit",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
			{t: "uint64_t", name: "STACK_3", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_setbit_op", op: "setbit",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_getbyte", op: "getbyte",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_getbyte_op", op: "getbyte",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_setbyte", op: "setbyte",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
			{t: "uint64_t", name: "STACK_3", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_setbyte_op", op: "setbyte",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_extract", op: "extract",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false, field: false},
			{t: "uint8", name: "L2", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_extract_op", op: "extract",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false, field: false},
			{t: "uint8", name: "L2", array: false, field: false},
		},
	},
	{
		t: "r_byte_t", name: "avm_extract3", op: "extract3",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
			{t: "uint64_t", name: "STACK_3", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_extract3_op", op: "extract3",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_extract_uint16", op: "extract_uint16",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_extract_uint16_op", op: "extract_uint16",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_extract_uint32", op: "extract_uint32",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_extract_uint32_op", op: "extract_uint32",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_extract_uint64", op: "extract_uint64",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_extract_uint64_op", op: "extract_uint64",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_replace2", op: "replace2",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "r_byte_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_replace2_op", op: "replace2",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false, field: false},
		},
	},
	{
		t: "r_byte_t", name: "avm_replace3", op: "replace3",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
			{t: "r_byte_t", name: "STACK_3", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_replace3_op", op: "replace3",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_base64_decode", op: "base64_decode",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "E1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_base64_decode_op", op: "base64_decode",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "E1", array: false, field: true},
		},
	},
	{
		t: "any_t", name: "avm_json_ref", op: "json_ref",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "r_byte_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "R1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_json_ref_op", op: "json_ref",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "R1", array: false, field: true},
		},
	},
	{
		t: "uint64_t", name: "avm_balance", op: "balance",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_balance_op", op: "balance",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_app_opted_in", op: "app_opted_in",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_app_opted_in_op", op: "app_opted_in",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any_t", name: "avm_app_local_get", op: "app_local_get",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "key_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_app_local_get_op", op: "app_local_get",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_app_local_get_ex_result_t", name: "avm_app_local_get_ex", op: "app_local_get_ex",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
			{t: "key_t", name: "STACK_3", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
			{t: "bool_t", name: "r2"},
		},
	},
	{
		t: "avm_app_local_get_ex_result_t", name: "avm_app_local_get_ex_op", op: "app_local_get_ex",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any_t", name: "avm_app_global_get", op: "app_global_get",
		stack: []BuiltinFunctionParamData{
			{t: "key_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_app_global_get_op", op: "app_global_get",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_app_global_get_ex_result_t", name: "avm_app_global_get_ex", op: "app_global_get_ex",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "key_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
			{t: "bool_t", name: "r2"},
		},
	},
	{
		t: "avm_app_global_get_ex_result_t", name: "avm_app_global_get_ex_op", op: "app_global_get_ex",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_app_local_put", op: "app_local_put",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "key_t", name: "STACK_2", array: false, field: false},
			{t: "any_t", name: "STACK_3", array: false, field: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_app_local_put_op", op: "app_local_put",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_app_global_put", op: "app_global_put",
		stack: []BuiltinFunctionParamData{
			{t: "key_t", name: "STACK_1", array: false, field: false},
			{t: "any_t", name: "STACK_2", array: false, field: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_app_global_put_op", op: "app_global_put",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_app_local_del", op: "app_local_del",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "key_t", name: "STACK_2", array: false, field: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_app_local_del_op", op: "app_local_del",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_app_global_del", op: "app_global_del",
		stack: []BuiltinFunctionParamData{
			{t: "key_t", name: "STACK_1", array: false, field: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_app_global_del_op", op: "app_global_del",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_asset_holding_get_result_t", name: "avm_asset_holding_get", op: "asset_holding_get",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
			{t: "bool_t", name: "r2"},
		},
	},
	{
		t: "avm_asset_holding_get_result_t", name: "avm_asset_holding_get_op", op: "asset_holding_get",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
	{
		t: "avm_asset_params_get_result_t", name: "avm_asset_params_get", op: "asset_params_get",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
			{t: "bool_t", name: "r2"},
		},
	},
	{
		t: "avm_asset_params_get_result_t", name: "avm_asset_params_get_op", op: "asset_params_get",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
	{
		t: "avm_app_params_get_result_t", name: "avm_app_params_get", op: "app_params_get",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
			{t: "bool_t", name: "r2"},
		},
	},
	{
		t: "avm_app_params_get_result_t", name: "avm_app_params_get_op", op: "app_params_get",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
	{
		t: "avm_acct_params_get_result_t", name: "avm_acct_params_get", op: "acct_params_get",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
			{t: "bool_t", name: "r2"},
		},
	},
	{
		t: "avm_acct_params_get_result_t", name: "avm_acct_params_get_op", op: "acct_params_get",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
	{
		t: "uint64_t", name: "avm_min_balance", op: "min_balance",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_min_balance_op", op: "min_balance",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_pushbytes", op: "pushbytes",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_pushbytes_op", op: "pushbytes",
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1", array: false, field: false},
		},
	},
	{
		t: "uint64_t", name: "avm_pushint", op: "pushint",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "UINT1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_pushint_op", op: "pushint",
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "UINT1", array: false, field: false},
		},
	},
	{
		t: "void", name: "avm_pushbytess", op: "pushbytess",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1", array: true, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_pushbytess_op", op: "pushbytess",
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1", array: true, field: false},
		},
	},
	{
		t: "void", name: "avm_pushints", op: "pushints",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint64", name: "UINT1", array: true, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_pushints_op", op: "pushints",
		imm: []BuiltinFunctionParamData{
			{t: "uint64", name: "UINT1", array: true, field: false},
		},
	},
	{
		t: "bool_t", name: "avm_ed25519verify_bare", op: "ed25519verify_bare",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "r_byte_t", name: "STACK_2", array: false, field: false},
			{t: "r_byte_t", name: "STACK_3", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_ed25519verify_bare_op", op: "ed25519verify_bare",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_callsub", op: "callsub",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_callsub_op", op: "callsub",
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false, field: false},
		},
	},
	{
		t: "void", name: "avm_retsub", op: "retsub",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_retsub_op", op: "retsub",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_proto", op: "proto",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "A1", array: false, field: false},
			{t: "uint8", name: "R2", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_proto_op", op: "proto",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "A1", array: false, field: false},
			{t: "uint8", name: "R2", array: false, field: false},
		},
	},
	{
		t: "any_t", name: "avm_frame_dig", op: "frame_dig",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "int8", name: "I1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_frame_dig_op", op: "frame_dig",
		imm: []BuiltinFunctionParamData{
			{t: "int8", name: "I1", array: false, field: false},
		},
	},
	{
		t: "void", name: "avm_frame_bury", op: "frame_bury",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "int8", name: "I1", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_frame_bury_op", op: "frame_bury",
		imm: []BuiltinFunctionParamData{
			{t: "int8", name: "I1", array: false, field: false},
		},
	},
	{
		t: "void", name: "avm_switch", op: "switch",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: true, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_switch_op", op: "switch",
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: true, field: false},
		},
	},
	{
		t: "void", name: "avm_match", op: "match",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: true, field: false},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_match_op", op: "match",
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: true, field: false},
		},
	},
	{
		t: "uint64_t", name: "avm_shl", op: "shl",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_shl_op", op: "shl",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_shr", op: "shr",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_shr_op", op: "shr",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_sqrt", op: "sqrt",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_sqrt_op", op: "sqrt",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_bitlen", op: "bitlen",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_bitlen_op", op: "bitlen",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_exp", op: "exp",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_exp_op", op: "exp",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_expw_result_t", name: "avm_expw", op: "expw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
			{t: "uint64_t", name: "r2"},
		},
	},
	{
		t: "avm_expw_result_t", name: "avm_expw_op", op: "expw",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_bsqrt", op: "bsqrt",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_bsqrt_op", op: "bsqrt",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64_t", name: "avm_divw", op: "divw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
			{t: "uint64_t", name: "STACK_3", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
		},
	},
	{
		t: "uint64_t", name: "avm_divw_op", op: "divw",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_sha3_256", op: "sha3_256",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_sha3_256_op", op: "sha3_256",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_bplus", op: "b+",
		stack: []BuiltinFunctionParamData{
			{t: "bigint_t", name: "STACK_1", array: false, field: false},
			{t: "bigint_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_bplus_op", op: "b+",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bigint_t", name: "avm_bminus", op: "b-",
		stack: []BuiltinFunctionParamData{
			{t: "bigint_t", name: "STACK_1", array: false, field: false},
			{t: "bigint_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bigint_t", name: "r1"},
		},
	},
	{
		t: "bigint_t", name: "avm_bminus_op", op: "b-",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bigint_t", name: "avm_bdiv", op: "b/",
		stack: []BuiltinFunctionParamData{
			{t: "bigint_t", name: "STACK_1", array: false, field: false},
			{t: "bigint_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bigint_t", name: "r1"},
		},
	},
	{
		t: "bigint_t", name: "avm_bdiv_op", op: "b/",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_bmul", op: "b*",
		stack: []BuiltinFunctionParamData{
			{t: "bigint_t", name: "STACK_1", array: false, field: false},
			{t: "bigint_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_bmul_op", op: "b*",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_blt", op: "b<",
		stack: []BuiltinFunctionParamData{
			{t: "bigint_t", name: "STACK_1", array: false, field: false},
			{t: "bigint_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_blt_op", op: "b<",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_bgt", op: "b>",
		stack: []BuiltinFunctionParamData{
			{t: "bigint_t", name: "STACK_1", array: false, field: false},
			{t: "bigint_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_bgt_op", op: "b>",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_blteq", op: "b<=",
		stack: []BuiltinFunctionParamData{
			{t: "bigint_t", name: "STACK_1", array: false, field: false},
			{t: "bigint_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_blteq_op", op: "b<=",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_bgteq", op: "b>=",
		stack: []BuiltinFunctionParamData{
			{t: "bigint_t", name: "STACK_1", array: false, field: false},
			{t: "bigint_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_bgteq_op", op: "b>=",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_beqeq", op: "b==",
		stack: []BuiltinFunctionParamData{
			{t: "bigint_t", name: "STACK_1", array: false, field: false},
			{t: "bigint_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_beqeq_op", op: "b==",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_bnoteq", op: "b!=",
		stack: []BuiltinFunctionParamData{
			{t: "bigint_t", name: "STACK_1", array: false, field: false},
			{t: "bigint_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_bnoteq_op", op: "b!=",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_bmod", op: "b%",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "r_byte_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_bmod_op", op: "b%",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_bor", op: "b|",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "r_byte_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_bor_op", op: "b|",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_band", op: "b&",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "r_byte_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_band_op", op: "b&",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_bxor", op: "b^",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "r_byte_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_bxor_op", op: "b^",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_binv", op: "b~",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_binv_op", op: "b~",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_bzero", op: "bzero",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_bzero_op", op: "bzero",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_log", op: "log",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_log_op", op: "log",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_itxn_begin", op: "itxn_begin",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_itxn_begin_op", op: "itxn_begin",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_itxn_field", op: "itxn_field",
		stack: []BuiltinFunctionParamData{
			{t: "any_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_itxn_field_op", op: "itxn_field",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
	{
		t: "void", name: "avm_itxn_submit", op: "itxn_submit",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_itxn_submit_op", op: "itxn_submit",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any_t", name: "avm_itxn", op: "itxn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_itxn_op", op: "itxn",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
	{
		t: "any_t", name: "avm_itxna", op: "itxna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
			{t: "uint8", name: "I2", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_itxna_op", op: "itxna",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
			{t: "uint8", name: "I2", array: false, field: false},
		},
	},
	{
		t: "void", name: "avm_itxn_next", op: "itxn_next",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_itxn_next_op", op: "itxn_next",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any_t", name: "avm_gitxn", op: "gitxn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "F2", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_gitxn_op", op: "gitxn",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "F2", array: false, field: true},
		},
	},
	{
		t: "any_t", name: "avm_gitxna", op: "gitxna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "F2", array: false, field: true},
			{t: "uint8", name: "I3", array: false, field: false},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_gitxna_op", op: "gitxna",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "F2", array: false, field: true},
			{t: "uint8", name: "I3", array: false, field: false},
		},
	},
	{
		t: "bool_t", name: "avm_box_create", op: "box_create",
		stack: []BuiltinFunctionParamData{
			{t: "name_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_box_create_op", op: "box_create",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "r_byte_t", name: "avm_box_extract", op: "box_extract",
		stack: []BuiltinFunctionParamData{
			{t: "name_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
			{t: "uint64_t", name: "STACK_3", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_box_extract_op", op: "box_extract",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_box_replace", op: "box_replace",
		stack: []BuiltinFunctionParamData{
			{t: "name_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
			{t: "r_byte_t", name: "STACK_3", array: false, field: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_box_replace_op", op: "box_replace",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bool_t", name: "avm_box_del", op: "box_del",
		stack: []BuiltinFunctionParamData{
			{t: "name_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "bool_t", name: "r1"},
		},
	},
	{
		t: "bool_t", name: "avm_box_del_op", op: "box_del",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_box_len_result_t", name: "avm_box_len", op: "box_len",
		stack: []BuiltinFunctionParamData{
			{t: "name_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "uint64_t", name: "r1"},
			{t: "bool_t", name: "r2"},
		},
	},
	{
		t: "avm_box_len_result_t", name: "avm_box_len_op", op: "box_len",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_box_get_result_t", name: "avm_box_get", op: "box_get",
		stack: []BuiltinFunctionParamData{
			{t: "name_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
			{t: "bool_t", name: "r2"},
		},
	},
	{
		t: "avm_box_get_result_t", name: "avm_box_get_op", op: "box_get",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_box_put", op: "box_put",
		stack: []BuiltinFunctionParamData{
			{t: "name_t", name: "STACK_1", array: false, field: false},
			{t: "r_byte_t", name: "STACK_2", array: false, field: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{},
	},
	{
		t: "void", name: "avm_box_put_op", op: "box_put",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any_t", name: "avm_txnas", op: "txnas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_txnas_op", op: "txnas",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
	{
		t: "any_t", name: "avm_gtxnas", op: "gtxnas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "F2", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_gtxnas_op", op: "gtxnas",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "F2", array: false, field: true},
		},
	},
	{
		t: "any_t", name: "avm_gtxnsas", op: "gtxnsas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_gtxnsas_op", op: "gtxnsas",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
	{
		t: "r_byte_t", name: "avm_args", op: "args",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
		},
	},
	{
		t: "r_byte_t", name: "avm_args_op", op: "args",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any_t", name: "avm_gloadss", op: "gloadss",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
			{t: "uint64_t", name: "STACK_2", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_gloadss_op", op: "gloadss",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any_t", name: "avm_itxnas", op: "itxnas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_itxnas_op", op: "itxnas",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
	{
		t: "any_t", name: "avm_gitxnas", op: "gitxnas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "F2", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_gitxnas_op", op: "gitxnas",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false, field: false},
			{t: "uint8", name: "F2", array: false, field: true},
		},
	},
	{
		t: "avm_vrf_verify_result_t", name: "avm_vrf_verify", op: "vrf_verify",
		stack: []BuiltinFunctionParamData{
			{t: "r_byte_t", name: "STACK_1", array: false, field: false},
			{t: "r_byte_t", name: "STACK_2", array: false, field: false},
			{t: "r_byte_t", name: "STACK_3", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "r_byte_t", name: "r1"},
			{t: "bool_t", name: "r2"},
		},
	},
	{
		t: "avm_vrf_verify_result_t", name: "avm_vrf_verify_op", op: "vrf_verify",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false, field: true},
		},
	},
	{
		t: "any_t", name: "avm_block", op: "block",
		stack: []BuiltinFunctionParamData{
			{t: "uint64_t", name: "STACK_1", array: false, field: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
		returns: []BuiltinFunctionReturnData{
			{t: "any_t", name: "r1"},
		},
	},
	{
		t: "any_t", name: "avm_block_op", op: "block",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false, field: true},
		},
	},
}

type BuiltinStructFieldData struct {
	t    string
	name string
	fun  string
}

type BuiltinStructFunctionParamData struct {
	t     string
	name  string
	array bool
	field bool
}

type BuiltinStructFunctionData struct {
	t    string
	name string
	fun  string
}

type BuiltinStructData struct {
	name      string
	fields    []BuiltinStructFieldData
	functions []BuiltinStructFunctionData
}

var builtin_structs = []BuiltinStructData{
	{
		name:   "avm_ecdsa_verify_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "void", name: "Secp256k1", fun: "avm_ecdsa_verify",
			},
			{
				t: "void", name: "Secp256r1", fun: "avm_ecdsa_verify",
			},
		},
	},
	{
		name:   "avm_ecdsa_pk_decompress_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "void", name: "Secp256k1", fun: "avm_ecdsa_pk_decompress",
			},
			{
				t: "void", name: "Secp256r1", fun: "avm_ecdsa_pk_decompress",
			},
		},
	},
	{
		name: "avm_ecdsa_pk_decompress_result_t",
		fields: []BuiltinStructFieldData{
			{t: "r_byte_t", name: "r1", fun: "avm_ecdsa_pk_decompress"},
			{t: "r_byte_t", name: "r2", fun: "avm_ecdsa_pk_decompress"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_ecdsa_pk_recover_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "void", name: "Secp256k1", fun: "avm_ecdsa_pk_recover",
			},
			{
				t: "void", name: "Secp256r1", fun: "avm_ecdsa_pk_recover",
			},
		},
	},
	{
		name: "avm_ecdsa_pk_recover_result_t",
		fields: []BuiltinStructFieldData{
			{t: "r_byte_t", name: "r1", fun: "avm_ecdsa_pk_recover"},
			{t: "r_byte_t", name: "r2", fun: "avm_ecdsa_pk_recover"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_mulw_result_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64_t", name: "r1", fun: "avm_mulw"},
			{t: "uint64_t", name: "r2", fun: "avm_mulw"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_addw_result_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64_t", name: "r1", fun: "avm_addw"},
			{t: "uint64_t", name: "r2", fun: "avm_addw"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_divmodw_result_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64_t", name: "r1", fun: "avm_divmodw"},
			{t: "uint64_t", name: "r2", fun: "avm_divmodw"},
			{t: "uint64_t", name: "r3", fun: "avm_divmodw"},
			{t: "uint64_t", name: "r4", fun: "avm_divmodw"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_txn_t",
		fields: []BuiltinStructFieldData{
			{t: "addr_t", name: "Sender", fun: "avm_txn"},
			{t: "uint64_t", name: "Fee", fun: "avm_txn"},
			{t: "uint64_t", name: "FirstValid", fun: "avm_txn"},
			{t: "uint64_t", name: "FirstValidTime", fun: "avm_txn"},
			{t: "uint64_t", name: "LastValid", fun: "avm_txn"},
			{t: "r_byte_t", name: "Note", fun: "avm_txn"},
			{t: "r32_byte_t", name: "Lease", fun: "avm_txn"},
			{t: "addr_t", name: "Receiver", fun: "avm_txn"},
			{t: "uint64_t", name: "Amount", fun: "avm_txn"},
			{t: "addr_t", name: "CloseRemainderTo", fun: "avm_txn"},
			{t: "r32_byte_t", name: "VotePK", fun: "avm_txn"},
			{t: "r32_byte_t", name: "SelectionPK", fun: "avm_txn"},
			{t: "uint64_t", name: "VoteFirst", fun: "avm_txn"},
			{t: "uint64_t", name: "VoteLast", fun: "avm_txn"},
			{t: "uint64_t", name: "VoteKeyDilution", fun: "avm_txn"},
			{t: "r_byte_t", name: "Type", fun: "avm_txn"},
			{t: "uint64_t", name: "TypeEnum", fun: "avm_txn"},
			{t: "uint64_t", name: "XferAsset", fun: "avm_txn"},
			{t: "uint64_t", name: "AssetAmount", fun: "avm_txn"},
			{t: "addr_t", name: "AssetSender", fun: "avm_txn"},
			{t: "addr_t", name: "AssetReceiver", fun: "avm_txn"},
			{t: "addr_t", name: "AssetCloseTo", fun: "avm_txn"},
			{t: "uint64_t", name: "GroupIndex", fun: "avm_txn"},
			{t: "r32_byte_t", name: "TxID", fun: "avm_txn"},
			{t: "uint64_t", name: "ApplicationID", fun: "avm_txn"},
			{t: "uint64_t", name: "OnCompletion", fun: "avm_txn"},
			{t: "r_byte_t", name: "ApplicationArgs", fun: "avm_txn"},
			{t: "uint64_t", name: "NumAppArgs", fun: "avm_txn"},
			{t: "addr_t", name: "Accounts", fun: "avm_txn"},
			{t: "uint64_t", name: "NumAccounts", fun: "avm_txn"},
			{t: "r_byte_t", name: "ApprovalProgram", fun: "avm_txn"},
			{t: "r_byte_t", name: "ClearStateProgram", fun: "avm_txn"},
			{t: "addr_t", name: "RekeyTo", fun: "avm_txn"},
			{t: "uint64_t", name: "ConfigAsset", fun: "avm_txn"},
			{t: "uint64_t", name: "ConfigAssetTotal", fun: "avm_txn"},
			{t: "uint64_t", name: "ConfigAssetDecimals", fun: "avm_txn"},
			{t: "bool_t", name: "ConfigAssetDefaultFrozen", fun: "avm_txn"},
			{t: "r_byte_t", name: "ConfigAssetUnitName", fun: "avm_txn"},
			{t: "r_byte_t", name: "ConfigAssetName", fun: "avm_txn"},
			{t: "r_byte_t", name: "ConfigAssetURL", fun: "avm_txn"},
			{t: "r32_byte_t", name: "ConfigAssetMetadataHash", fun: "avm_txn"},
			{t: "addr_t", name: "ConfigAssetManager", fun: "avm_txn"},
			{t: "addr_t", name: "ConfigAssetReserve", fun: "avm_txn"},
			{t: "addr_t", name: "ConfigAssetFreeze", fun: "avm_txn"},
			{t: "addr_t", name: "ConfigAssetClawback", fun: "avm_txn"},
			{t: "uint64_t", name: "FreezeAsset", fun: "avm_txn"},
			{t: "addr_t", name: "FreezeAssetAccount", fun: "avm_txn"},
			{t: "bool_t", name: "FreezeAssetFrozen", fun: "avm_txn"},
			{t: "uint64_t", name: "Assets", fun: "avm_txn"},
			{t: "uint64_t", name: "NumAssets", fun: "avm_txn"},
			{t: "uint64_t", name: "Applications", fun: "avm_txn"},
			{t: "uint64_t", name: "NumApplications", fun: "avm_txn"},
			{t: "uint64_t", name: "GlobalNumUint", fun: "avm_txn"},
			{t: "uint64_t", name: "GlobalNumByteSlice", fun: "avm_txn"},
			{t: "uint64_t", name: "LocalNumUint", fun: "avm_txn"},
			{t: "uint64_t", name: "LocalNumByteSlice", fun: "avm_txn"},
			{t: "uint64_t", name: "ExtraProgramPages", fun: "avm_txn"},
			{t: "bool_t", name: "Nonparticipation", fun: "avm_txn"},
			{t: "r_byte_t", name: "Logs", fun: "avm_txn"},
			{t: "uint64_t", name: "NumLogs", fun: "avm_txn"},
			{t: "uint64_t", name: "CreatedAssetID", fun: "avm_txn"},
			{t: "uint64_t", name: "CreatedApplicationID", fun: "avm_txn"},
			{t: "r_byte_t", name: "LastLog", fun: "avm_txn"},
			{t: "r_byte_t", name: "StateProofPK", fun: "avm_txn"},
			{t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_txn"},
			{t: "uint64_t", name: "NumApprovalProgramPages", fun: "avm_txn"},
			{t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_txn"},
			{t: "uint64_t", name: "NumClearStateProgramPages", fun: "avm_txn"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_global_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64_t", name: "MinTxnFee", fun: "avm_global"},
			{t: "uint64_t", name: "MinBalance", fun: "avm_global"},
			{t: "uint64_t", name: "MaxTxnLife", fun: "avm_global"},
			{t: "addr_t", name: "ZeroAddress", fun: "avm_global"},
			{t: "uint64_t", name: "GroupSize", fun: "avm_global"},
			{t: "uint64_t", name: "LogicSigVersion", fun: "avm_global"},
			{t: "uint64_t", name: "Round", fun: "avm_global"},
			{t: "uint64_t", name: "LatestTimestamp", fun: "avm_global"},
			{t: "uint64_t", name: "CurrentApplicationID", fun: "avm_global"},
			{t: "addr_t", name: "CreatorAddress", fun: "avm_global"},
			{t: "addr_t", name: "CurrentApplicationAddress", fun: "avm_global"},
			{t: "r32_byte_t", name: "GroupID", fun: "avm_global"},
			{t: "uint64_t", name: "OpcodeBudget", fun: "avm_global"},
			{t: "uint64_t", name: "CallerApplicationID", fun: "avm_global"},
			{t: "addr_t", name: "CallerApplicationAddress", fun: "avm_global"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_gtxn_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "addr_t", name: "Sender", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "Fee", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "FirstValid", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "FirstValidTime", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "LastValid", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "Note", fun: "avm_gtxn",
			},
			{
				t: "r32_byte_t", name: "Lease", fun: "avm_gtxn",
			},
			{
				t: "addr_t", name: "Receiver", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "Amount", fun: "avm_gtxn",
			},
			{
				t: "addr_t", name: "CloseRemainderTo", fun: "avm_gtxn",
			},
			{
				t: "r32_byte_t", name: "VotePK", fun: "avm_gtxn",
			},
			{
				t: "r32_byte_t", name: "SelectionPK", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "VoteFirst", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "VoteLast", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "VoteKeyDilution", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "Type", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "TypeEnum", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "XferAsset", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "AssetAmount", fun: "avm_gtxn",
			},
			{
				t: "addr_t", name: "AssetSender", fun: "avm_gtxn",
			},
			{
				t: "addr_t", name: "AssetReceiver", fun: "avm_gtxn",
			},
			{
				t: "addr_t", name: "AssetCloseTo", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "GroupIndex", fun: "avm_gtxn",
			},
			{
				t: "r32_byte_t", name: "TxID", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "ApplicationID", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "OnCompletion", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "ApplicationArgs", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "NumAppArgs", fun: "avm_gtxn",
			},
			{
				t: "addr_t", name: "Accounts", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "NumAccounts", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "ApprovalProgram", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "ClearStateProgram", fun: "avm_gtxn",
			},
			{
				t: "addr_t", name: "RekeyTo", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "ConfigAsset", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "ConfigAssetTotal", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "ConfigAssetDecimals", fun: "avm_gtxn",
			},
			{
				t: "bool_t", name: "ConfigAssetDefaultFrozen", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "ConfigAssetUnitName", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "ConfigAssetName", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "ConfigAssetURL", fun: "avm_gtxn",
			},
			{
				t: "r32_byte_t", name: "ConfigAssetMetadataHash", fun: "avm_gtxn",
			},
			{
				t: "addr_t", name: "ConfigAssetManager", fun: "avm_gtxn",
			},
			{
				t: "addr_t", name: "ConfigAssetReserve", fun: "avm_gtxn",
			},
			{
				t: "addr_t", name: "ConfigAssetFreeze", fun: "avm_gtxn",
			},
			{
				t: "addr_t", name: "ConfigAssetClawback", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "FreezeAsset", fun: "avm_gtxn",
			},
			{
				t: "addr_t", name: "FreezeAssetAccount", fun: "avm_gtxn",
			},
			{
				t: "bool_t", name: "FreezeAssetFrozen", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "Assets", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "NumAssets", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "Applications", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "NumApplications", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "GlobalNumUint", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "GlobalNumByteSlice", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "LocalNumUint", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "LocalNumByteSlice", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "ExtraProgramPages", fun: "avm_gtxn",
			},
			{
				t: "bool_t", name: "Nonparticipation", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "Logs", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "NumLogs", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "CreatedAssetID", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "CreatedApplicationID", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "LastLog", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "StateProofPK", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "NumApprovalProgramPages", fun: "avm_gtxn",
			},
			{
				t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_gtxn",
			},
			{
				t: "uint64_t", name: "NumClearStateProgramPages", fun: "avm_gtxn",
			},
		},
	},
	{
		name:   "avm_txna_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "r_byte_t", name: "ApplicationArgs", fun: "avm_txna",
			},
			{
				t: "addr_t", name: "Accounts", fun: "avm_txna",
			},
			{
				t: "uint64_t", name: "Assets", fun: "avm_txna",
			},
			{
				t: "uint64_t", name: "Applications", fun: "avm_txna",
			},
			{
				t: "r_byte_t", name: "Logs", fun: "avm_txna",
			},
			{
				t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_txna",
			},
			{
				t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_txna",
			},
		},
	},
	{
		name:   "avm_gtxna_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "r_byte_t", name: "ApplicationArgs", fun: "avm_gtxna",
			},
			{
				t: "addr_t", name: "Accounts", fun: "avm_gtxna",
			},
			{
				t: "uint64_t", name: "Assets", fun: "avm_gtxna",
			},
			{
				t: "uint64_t", name: "Applications", fun: "avm_gtxna",
			},
			{
				t: "r_byte_t", name: "Logs", fun: "avm_gtxna",
			},
			{
				t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_gtxna",
			},
			{
				t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_gtxna",
			},
		},
	},
	{
		name:   "avm_gtxns_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "addr_t", name: "Sender", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "Fee", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "FirstValid", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "FirstValidTime", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "LastValid", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "Note", fun: "avm_gtxns",
			},
			{
				t: "r32_byte_t", name: "Lease", fun: "avm_gtxns",
			},
			{
				t: "addr_t", name: "Receiver", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "Amount", fun: "avm_gtxns",
			},
			{
				t: "addr_t", name: "CloseRemainderTo", fun: "avm_gtxns",
			},
			{
				t: "r32_byte_t", name: "VotePK", fun: "avm_gtxns",
			},
			{
				t: "r32_byte_t", name: "SelectionPK", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "VoteFirst", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "VoteLast", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "VoteKeyDilution", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "Type", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "TypeEnum", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "XferAsset", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "AssetAmount", fun: "avm_gtxns",
			},
			{
				t: "addr_t", name: "AssetSender", fun: "avm_gtxns",
			},
			{
				t: "addr_t", name: "AssetReceiver", fun: "avm_gtxns",
			},
			{
				t: "addr_t", name: "AssetCloseTo", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "GroupIndex", fun: "avm_gtxns",
			},
			{
				t: "r32_byte_t", name: "TxID", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "ApplicationID", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "OnCompletion", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "ApplicationArgs", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "NumAppArgs", fun: "avm_gtxns",
			},
			{
				t: "addr_t", name: "Accounts", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "NumAccounts", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "ApprovalProgram", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "ClearStateProgram", fun: "avm_gtxns",
			},
			{
				t: "addr_t", name: "RekeyTo", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "ConfigAsset", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "ConfigAssetTotal", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "ConfigAssetDecimals", fun: "avm_gtxns",
			},
			{
				t: "bool_t", name: "ConfigAssetDefaultFrozen", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "ConfigAssetUnitName", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "ConfigAssetName", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "ConfigAssetURL", fun: "avm_gtxns",
			},
			{
				t: "r32_byte_t", name: "ConfigAssetMetadataHash", fun: "avm_gtxns",
			},
			{
				t: "addr_t", name: "ConfigAssetManager", fun: "avm_gtxns",
			},
			{
				t: "addr_t", name: "ConfigAssetReserve", fun: "avm_gtxns",
			},
			{
				t: "addr_t", name: "ConfigAssetFreeze", fun: "avm_gtxns",
			},
			{
				t: "addr_t", name: "ConfigAssetClawback", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "FreezeAsset", fun: "avm_gtxns",
			},
			{
				t: "addr_t", name: "FreezeAssetAccount", fun: "avm_gtxns",
			},
			{
				t: "bool_t", name: "FreezeAssetFrozen", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "Assets", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "NumAssets", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "Applications", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "NumApplications", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "GlobalNumUint", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "GlobalNumByteSlice", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "LocalNumUint", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "LocalNumByteSlice", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "ExtraProgramPages", fun: "avm_gtxns",
			},
			{
				t: "bool_t", name: "Nonparticipation", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "Logs", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "NumLogs", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "CreatedAssetID", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "CreatedApplicationID", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "LastLog", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "StateProofPK", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "NumApprovalProgramPages", fun: "avm_gtxns",
			},
			{
				t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_gtxns",
			},
			{
				t: "uint64_t", name: "NumClearStateProgramPages", fun: "avm_gtxns",
			},
		},
	},
	{
		name:   "avm_gtxnsa_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "r_byte_t", name: "ApplicationArgs", fun: "avm_gtxnsa",
			},
			{
				t: "addr_t", name: "Accounts", fun: "avm_gtxnsa",
			},
			{
				t: "uint64_t", name: "Assets", fun: "avm_gtxnsa",
			},
			{
				t: "uint64_t", name: "Applications", fun: "avm_gtxnsa",
			},
			{
				t: "r_byte_t", name: "Logs", fun: "avm_gtxnsa",
			},
			{
				t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_gtxnsa",
			},
			{
				t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_gtxnsa",
			},
		},
	},
	{
		name: "avm_dup_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any_t", name: "r1", fun: "avm_dup"},
			{t: "any_t", name: "r2", fun: "avm_dup"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_dup2_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any_t", name: "r1", fun: "avm_dup2"},
			{t: "any_t", name: "r2", fun: "avm_dup2"},
			{t: "any_t", name: "r3", fun: "avm_dup2"},
			{t: "any_t", name: "r4", fun: "avm_dup2"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_dig_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any_t", name: "r1", fun: "avm_dig"},
			{t: "any_t", name: "r2", fun: "avm_dig"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_swap_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any_t", name: "r1", fun: "avm_swap"},
			{t: "any_t", name: "r2", fun: "avm_swap"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_base64_decode_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "any_t", name: "URLEncoding", fun: "avm_base64_decode",
			},
			{
				t: "any_t", name: "StdEncoding", fun: "avm_base64_decode",
			},
		},
	},
	{
		name:   "avm_json_ref_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "r_byte_t", name: "JSONString", fun: "avm_json_ref",
			},
			{
				t: "uint64_t", name: "JSONUint64", fun: "avm_json_ref",
			},
			{
				t: "r_byte_t", name: "JSONObject", fun: "avm_json_ref",
			},
		},
	},
	{
		name: "avm_app_local_get_ex_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any_t", name: "r1", fun: "avm_app_local_get_ex"},
			{t: "bool_t", name: "r2", fun: "avm_app_local_get_ex"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_app_global_get_ex_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any_t", name: "r1", fun: "avm_app_global_get_ex"},
			{t: "bool_t", name: "r2", fun: "avm_app_global_get_ex"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_asset_holding_get_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "uint64_t", name: "AssetBalance", fun: "avm_asset_holding_get",
			},
			{
				t: "bool_t", name: "AssetFrozen", fun: "avm_asset_holding_get",
			},
		},
	},
	{
		name: "avm_asset_holding_get_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any_t", name: "r1", fun: "avm_asset_holding_get"},
			{t: "bool_t", name: "r2", fun: "avm_asset_holding_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_asset_params_get_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "uint64_t", name: "AssetTotal", fun: "avm_asset_params_get",
			},
			{
				t: "uint64_t", name: "AssetDecimals", fun: "avm_asset_params_get",
			},
			{
				t: "bool_t", name: "AssetDefaultFrozen", fun: "avm_asset_params_get",
			},
			{
				t: "r_byte_t", name: "AssetUnitName", fun: "avm_asset_params_get",
			},
			{
				t: "r_byte_t", name: "AssetName", fun: "avm_asset_params_get",
			},
			{
				t: "r_byte_t", name: "AssetURL", fun: "avm_asset_params_get",
			},
			{
				t: "r32_byte_t", name: "AssetMetadataHash", fun: "avm_asset_params_get",
			},
			{
				t: "addr_t", name: "AssetManager", fun: "avm_asset_params_get",
			},
			{
				t: "addr_t", name: "AssetReserve", fun: "avm_asset_params_get",
			},
			{
				t: "addr_t", name: "AssetFreeze", fun: "avm_asset_params_get",
			},
			{
				t: "addr_t", name: "AssetClawback", fun: "avm_asset_params_get",
			},
			{
				t: "addr_t", name: "AssetCreator", fun: "avm_asset_params_get",
			},
		},
	},
	{
		name: "avm_asset_params_get_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any_t", name: "r1", fun: "avm_asset_params_get"},
			{t: "bool_t", name: "r2", fun: "avm_asset_params_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_app_params_get_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "r_byte_t", name: "AppApprovalProgram", fun: "avm_app_params_get",
			},
			{
				t: "r_byte_t", name: "AppClearStateProgram", fun: "avm_app_params_get",
			},
			{
				t: "uint64_t", name: "AppGlobalNumUint", fun: "avm_app_params_get",
			},
			{
				t: "uint64_t", name: "AppGlobalNumByteSlice", fun: "avm_app_params_get",
			},
			{
				t: "uint64_t", name: "AppLocalNumUint", fun: "avm_app_params_get",
			},
			{
				t: "uint64_t", name: "AppLocalNumByteSlice", fun: "avm_app_params_get",
			},
			{
				t: "uint64_t", name: "AppExtraProgramPages", fun: "avm_app_params_get",
			},
			{
				t: "addr_t", name: "AppCreator", fun: "avm_app_params_get",
			},
			{
				t: "addr_t", name: "AppAddress", fun: "avm_app_params_get",
			},
		},
	},
	{
		name: "avm_app_params_get_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any_t", name: "r1", fun: "avm_app_params_get"},
			{t: "bool_t", name: "r2", fun: "avm_app_params_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_acct_params_get_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "uint64_t", name: "AcctBalance", fun: "avm_acct_params_get",
			},
			{
				t: "uint64_t", name: "AcctMinBalance", fun: "avm_acct_params_get",
			},
			{
				t: "addr_t", name: "AcctAuthAddr", fun: "avm_acct_params_get",
			},
			{
				t: "uint64_t", name: "AcctTotalNumUint", fun: "avm_acct_params_get",
			},
			{
				t: "uint64_t", name: "AcctTotalNumByteSlice", fun: "avm_acct_params_get",
			},
			{
				t: "uint64_t", name: "AcctTotalExtraAppPages", fun: "avm_acct_params_get",
			},
			{
				t: "uint64_t", name: "AcctTotalAppsCreated", fun: "avm_acct_params_get",
			},
			{
				t: "uint64_t", name: "AcctTotalAppsOptedIn", fun: "avm_acct_params_get",
			},
			{
				t: "uint64_t", name: "AcctTotalAssetsCreated", fun: "avm_acct_params_get",
			},
			{
				t: "uint64_t", name: "AcctTotalAssets", fun: "avm_acct_params_get",
			},
			{
				t: "uint64_t", name: "AcctTotalBoxes", fun: "avm_acct_params_get",
			},
			{
				t: "uint64_t", name: "AcctTotalBoxBytes", fun: "avm_acct_params_get",
			},
		},
	},
	{
		name: "avm_acct_params_get_result_t",
		fields: []BuiltinStructFieldData{
			{t: "any_t", name: "r1", fun: "avm_acct_params_get"},
			{t: "bool_t", name: "r2", fun: "avm_acct_params_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_expw_result_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64_t", name: "r1", fun: "avm_expw"},
			{t: "uint64_t", name: "r2", fun: "avm_expw"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_itxn_field_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "addr_t", name: "Sender", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "Fee", fun: "avm_itxn_field",
			},
			{
				t: "r_byte_t", name: "Note", fun: "avm_itxn_field",
			},
			{
				t: "addr_t", name: "Receiver", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "Amount", fun: "avm_itxn_field",
			},
			{
				t: "addr_t", name: "CloseRemainderTo", fun: "avm_itxn_field",
			},
			{
				t: "r32_byte_t", name: "VotePK", fun: "avm_itxn_field",
			},
			{
				t: "r32_byte_t", name: "SelectionPK", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "VoteFirst", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "VoteLast", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "VoteKeyDilution", fun: "avm_itxn_field",
			},
			{
				t: "r_byte_t", name: "Type", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "TypeEnum", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "XferAsset", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "AssetAmount", fun: "avm_itxn_field",
			},
			{
				t: "addr_t", name: "AssetSender", fun: "avm_itxn_field",
			},
			{
				t: "addr_t", name: "AssetReceiver", fun: "avm_itxn_field",
			},
			{
				t: "addr_t", name: "AssetCloseTo", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "ApplicationID", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "OnCompletion", fun: "avm_itxn_field",
			},
			{
				t: "r_byte_t", name: "ApplicationArgs", fun: "avm_itxn_field",
			},
			{
				t: "addr_t", name: "Accounts", fun: "avm_itxn_field",
			},
			{
				t: "r_byte_t", name: "ApprovalProgram", fun: "avm_itxn_field",
			},
			{
				t: "r_byte_t", name: "ClearStateProgram", fun: "avm_itxn_field",
			},
			{
				t: "addr_t", name: "RekeyTo", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "ConfigAsset", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "ConfigAssetTotal", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "ConfigAssetDecimals", fun: "avm_itxn_field",
			},
			{
				t: "bool_t", name: "ConfigAssetDefaultFrozen", fun: "avm_itxn_field",
			},
			{
				t: "r_byte_t", name: "ConfigAssetUnitName", fun: "avm_itxn_field",
			},
			{
				t: "r_byte_t", name: "ConfigAssetName", fun: "avm_itxn_field",
			},
			{
				t: "r_byte_t", name: "ConfigAssetURL", fun: "avm_itxn_field",
			},
			{
				t: "r32_byte_t", name: "ConfigAssetMetadataHash", fun: "avm_itxn_field",
			},
			{
				t: "addr_t", name: "ConfigAssetManager", fun: "avm_itxn_field",
			},
			{
				t: "addr_t", name: "ConfigAssetReserve", fun: "avm_itxn_field",
			},
			{
				t: "addr_t", name: "ConfigAssetFreeze", fun: "avm_itxn_field",
			},
			{
				t: "addr_t", name: "ConfigAssetClawback", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "FreezeAsset", fun: "avm_itxn_field",
			},
			{
				t: "addr_t", name: "FreezeAssetAccount", fun: "avm_itxn_field",
			},
			{
				t: "bool_t", name: "FreezeAssetFrozen", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "Assets", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "Applications", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "GlobalNumUint", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "GlobalNumByteSlice", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "LocalNumUint", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "LocalNumByteSlice", fun: "avm_itxn_field",
			},
			{
				t: "uint64_t", name: "ExtraProgramPages", fun: "avm_itxn_field",
			},
			{
				t: "bool_t", name: "Nonparticipation", fun: "avm_itxn_field",
			},
			{
				t: "r_byte_t", name: "StateProofPK", fun: "avm_itxn_field",
			},
			{
				t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_itxn_field",
			},
			{
				t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_itxn_field",
			},
		},
	},
	{
		name: "avm_itxn_t",
		fields: []BuiltinStructFieldData{
			{t: "addr_t", name: "Sender", fun: "avm_itxn"},
			{t: "uint64_t", name: "Fee", fun: "avm_itxn"},
			{t: "uint64_t", name: "FirstValid", fun: "avm_itxn"},
			{t: "uint64_t", name: "FirstValidTime", fun: "avm_itxn"},
			{t: "uint64_t", name: "LastValid", fun: "avm_itxn"},
			{t: "r_byte_t", name: "Note", fun: "avm_itxn"},
			{t: "r32_byte_t", name: "Lease", fun: "avm_itxn"},
			{t: "addr_t", name: "Receiver", fun: "avm_itxn"},
			{t: "uint64_t", name: "Amount", fun: "avm_itxn"},
			{t: "addr_t", name: "CloseRemainderTo", fun: "avm_itxn"},
			{t: "r32_byte_t", name: "VotePK", fun: "avm_itxn"},
			{t: "r32_byte_t", name: "SelectionPK", fun: "avm_itxn"},
			{t: "uint64_t", name: "VoteFirst", fun: "avm_itxn"},
			{t: "uint64_t", name: "VoteLast", fun: "avm_itxn"},
			{t: "uint64_t", name: "VoteKeyDilution", fun: "avm_itxn"},
			{t: "r_byte_t", name: "Type", fun: "avm_itxn"},
			{t: "uint64_t", name: "TypeEnum", fun: "avm_itxn"},
			{t: "uint64_t", name: "XferAsset", fun: "avm_itxn"},
			{t: "uint64_t", name: "AssetAmount", fun: "avm_itxn"},
			{t: "addr_t", name: "AssetSender", fun: "avm_itxn"},
			{t: "addr_t", name: "AssetReceiver", fun: "avm_itxn"},
			{t: "addr_t", name: "AssetCloseTo", fun: "avm_itxn"},
			{t: "uint64_t", name: "GroupIndex", fun: "avm_itxn"},
			{t: "r32_byte_t", name: "TxID", fun: "avm_itxn"},
			{t: "uint64_t", name: "ApplicationID", fun: "avm_itxn"},
			{t: "uint64_t", name: "OnCompletion", fun: "avm_itxn"},
			{t: "r_byte_t", name: "ApplicationArgs", fun: "avm_itxn"},
			{t: "uint64_t", name: "NumAppArgs", fun: "avm_itxn"},
			{t: "addr_t", name: "Accounts", fun: "avm_itxn"},
			{t: "uint64_t", name: "NumAccounts", fun: "avm_itxn"},
			{t: "r_byte_t", name: "ApprovalProgram", fun: "avm_itxn"},
			{t: "r_byte_t", name: "ClearStateProgram", fun: "avm_itxn"},
			{t: "addr_t", name: "RekeyTo", fun: "avm_itxn"},
			{t: "uint64_t", name: "ConfigAsset", fun: "avm_itxn"},
			{t: "uint64_t", name: "ConfigAssetTotal", fun: "avm_itxn"},
			{t: "uint64_t", name: "ConfigAssetDecimals", fun: "avm_itxn"},
			{t: "bool_t", name: "ConfigAssetDefaultFrozen", fun: "avm_itxn"},
			{t: "r_byte_t", name: "ConfigAssetUnitName", fun: "avm_itxn"},
			{t: "r_byte_t", name: "ConfigAssetName", fun: "avm_itxn"},
			{t: "r_byte_t", name: "ConfigAssetURL", fun: "avm_itxn"},
			{t: "r32_byte_t", name: "ConfigAssetMetadataHash", fun: "avm_itxn"},
			{t: "addr_t", name: "ConfigAssetManager", fun: "avm_itxn"},
			{t: "addr_t", name: "ConfigAssetReserve", fun: "avm_itxn"},
			{t: "addr_t", name: "ConfigAssetFreeze", fun: "avm_itxn"},
			{t: "addr_t", name: "ConfigAssetClawback", fun: "avm_itxn"},
			{t: "uint64_t", name: "FreezeAsset", fun: "avm_itxn"},
			{t: "addr_t", name: "FreezeAssetAccount", fun: "avm_itxn"},
			{t: "bool_t", name: "FreezeAssetFrozen", fun: "avm_itxn"},
			{t: "uint64_t", name: "Assets", fun: "avm_itxn"},
			{t: "uint64_t", name: "NumAssets", fun: "avm_itxn"},
			{t: "uint64_t", name: "Applications", fun: "avm_itxn"},
			{t: "uint64_t", name: "NumApplications", fun: "avm_itxn"},
			{t: "uint64_t", name: "GlobalNumUint", fun: "avm_itxn"},
			{t: "uint64_t", name: "GlobalNumByteSlice", fun: "avm_itxn"},
			{t: "uint64_t", name: "LocalNumUint", fun: "avm_itxn"},
			{t: "uint64_t", name: "LocalNumByteSlice", fun: "avm_itxn"},
			{t: "uint64_t", name: "ExtraProgramPages", fun: "avm_itxn"},
			{t: "bool_t", name: "Nonparticipation", fun: "avm_itxn"},
			{t: "r_byte_t", name: "Logs", fun: "avm_itxn"},
			{t: "uint64_t", name: "NumLogs", fun: "avm_itxn"},
			{t: "uint64_t", name: "CreatedAssetID", fun: "avm_itxn"},
			{t: "uint64_t", name: "CreatedApplicationID", fun: "avm_itxn"},
			{t: "r_byte_t", name: "LastLog", fun: "avm_itxn"},
			{t: "r_byte_t", name: "StateProofPK", fun: "avm_itxn"},
			{t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_itxn"},
			{t: "uint64_t", name: "NumApprovalProgramPages", fun: "avm_itxn"},
			{t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_itxn"},
			{t: "uint64_t", name: "NumClearStateProgramPages", fun: "avm_itxn"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_itxna_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "r_byte_t", name: "ApplicationArgs", fun: "avm_itxna",
			},
			{
				t: "addr_t", name: "Accounts", fun: "avm_itxna",
			},
			{
				t: "uint64_t", name: "Assets", fun: "avm_itxna",
			},
			{
				t: "uint64_t", name: "Applications", fun: "avm_itxna",
			},
			{
				t: "r_byte_t", name: "Logs", fun: "avm_itxna",
			},
			{
				t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_itxna",
			},
			{
				t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_itxna",
			},
		},
	},
	{
		name:   "avm_gitxn_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "addr_t", name: "Sender", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "Fee", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "FirstValid", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "FirstValidTime", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "LastValid", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "Note", fun: "avm_gitxn",
			},
			{
				t: "r32_byte_t", name: "Lease", fun: "avm_gitxn",
			},
			{
				t: "addr_t", name: "Receiver", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "Amount", fun: "avm_gitxn",
			},
			{
				t: "addr_t", name: "CloseRemainderTo", fun: "avm_gitxn",
			},
			{
				t: "r32_byte_t", name: "VotePK", fun: "avm_gitxn",
			},
			{
				t: "r32_byte_t", name: "SelectionPK", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "VoteFirst", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "VoteLast", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "VoteKeyDilution", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "Type", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "TypeEnum", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "XferAsset", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "AssetAmount", fun: "avm_gitxn",
			},
			{
				t: "addr_t", name: "AssetSender", fun: "avm_gitxn",
			},
			{
				t: "addr_t", name: "AssetReceiver", fun: "avm_gitxn",
			},
			{
				t: "addr_t", name: "AssetCloseTo", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "GroupIndex", fun: "avm_gitxn",
			},
			{
				t: "r32_byte_t", name: "TxID", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "ApplicationID", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "OnCompletion", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "ApplicationArgs", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "NumAppArgs", fun: "avm_gitxn",
			},
			{
				t: "addr_t", name: "Accounts", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "NumAccounts", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "ApprovalProgram", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "ClearStateProgram", fun: "avm_gitxn",
			},
			{
				t: "addr_t", name: "RekeyTo", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "ConfigAsset", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "ConfigAssetTotal", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "ConfigAssetDecimals", fun: "avm_gitxn",
			},
			{
				t: "bool_t", name: "ConfigAssetDefaultFrozen", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "ConfigAssetUnitName", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "ConfigAssetName", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "ConfigAssetURL", fun: "avm_gitxn",
			},
			{
				t: "r32_byte_t", name: "ConfigAssetMetadataHash", fun: "avm_gitxn",
			},
			{
				t: "addr_t", name: "ConfigAssetManager", fun: "avm_gitxn",
			},
			{
				t: "addr_t", name: "ConfigAssetReserve", fun: "avm_gitxn",
			},
			{
				t: "addr_t", name: "ConfigAssetFreeze", fun: "avm_gitxn",
			},
			{
				t: "addr_t", name: "ConfigAssetClawback", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "FreezeAsset", fun: "avm_gitxn",
			},
			{
				t: "addr_t", name: "FreezeAssetAccount", fun: "avm_gitxn",
			},
			{
				t: "bool_t", name: "FreezeAssetFrozen", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "Assets", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "NumAssets", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "Applications", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "NumApplications", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "GlobalNumUint", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "GlobalNumByteSlice", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "LocalNumUint", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "LocalNumByteSlice", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "ExtraProgramPages", fun: "avm_gitxn",
			},
			{
				t: "bool_t", name: "Nonparticipation", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "Logs", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "NumLogs", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "CreatedAssetID", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "CreatedApplicationID", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "LastLog", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "StateProofPK", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "NumApprovalProgramPages", fun: "avm_gitxn",
			},
			{
				t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_gitxn",
			},
			{
				t: "uint64_t", name: "NumClearStateProgramPages", fun: "avm_gitxn",
			},
		},
	},
	{
		name:   "avm_gitxna_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "r_byte_t", name: "ApplicationArgs", fun: "avm_gitxna",
			},
			{
				t: "addr_t", name: "Accounts", fun: "avm_gitxna",
			},
			{
				t: "uint64_t", name: "Assets", fun: "avm_gitxna",
			},
			{
				t: "uint64_t", name: "Applications", fun: "avm_gitxna",
			},
			{
				t: "r_byte_t", name: "Logs", fun: "avm_gitxna",
			},
			{
				t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_gitxna",
			},
			{
				t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_gitxna",
			},
		},
	},
	{
		name: "avm_box_len_result_t",
		fields: []BuiltinStructFieldData{
			{t: "uint64_t", name: "r1", fun: "avm_box_len"},
			{t: "bool_t", name: "r2", fun: "avm_box_len"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name: "avm_box_get_result_t",
		fields: []BuiltinStructFieldData{
			{t: "r_byte_t", name: "r1", fun: "avm_box_get"},
			{t: "bool_t", name: "r2", fun: "avm_box_get"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_txnas_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "r_byte_t", name: "ApplicationArgs", fun: "avm_txnas",
			},
			{
				t: "addr_t", name: "Accounts", fun: "avm_txnas",
			},
			{
				t: "uint64_t", name: "Assets", fun: "avm_txnas",
			},
			{
				t: "uint64_t", name: "Applications", fun: "avm_txnas",
			},
			{
				t: "r_byte_t", name: "Logs", fun: "avm_txnas",
			},
			{
				t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_txnas",
			},
			{
				t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_txnas",
			},
		},
	},
	{
		name:   "avm_gtxnas_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "r_byte_t", name: "ApplicationArgs", fun: "avm_gtxnas",
			},
			{
				t: "addr_t", name: "Accounts", fun: "avm_gtxnas",
			},
			{
				t: "uint64_t", name: "Assets", fun: "avm_gtxnas",
			},
			{
				t: "uint64_t", name: "Applications", fun: "avm_gtxnas",
			},
			{
				t: "r_byte_t", name: "Logs", fun: "avm_gtxnas",
			},
			{
				t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_gtxnas",
			},
			{
				t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_gtxnas",
			},
		},
	},
	{
		name:   "avm_gtxnsas_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "r_byte_t", name: "ApplicationArgs", fun: "avm_gtxnsas",
			},
			{
				t: "addr_t", name: "Accounts", fun: "avm_gtxnsas",
			},
			{
				t: "uint64_t", name: "Assets", fun: "avm_gtxnsas",
			},
			{
				t: "uint64_t", name: "Applications", fun: "avm_gtxnsas",
			},
			{
				t: "r_byte_t", name: "Logs", fun: "avm_gtxnsas",
			},
			{
				t: "r_byte_t", name: "ApprovalProgramPages", fun: "avm_gtxnsas",
			},
			{
				t: "r_byte_t", name: "ClearStateProgramPages", fun: "avm_gtxnsas",
			},
		},
	},
	{
		name:   "avm_vrf_verify_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "void", name: "VrfAlgorand", fun: "avm_vrf_verify",
			},
		},
	},
	{
		name: "avm_vrf_verify_result_t",
		fields: []BuiltinStructFieldData{
			{t: "r_byte_t", name: "r1", fun: "avm_vrf_verify"},
			{t: "bool_t", name: "r2", fun: "avm_vrf_verify"},
		},
		functions: []BuiltinStructFunctionData{},
	},
	{
		name:   "avm_block_t",
		fields: []BuiltinStructFieldData{},
		functions: []BuiltinStructFunctionData{
			{
				t: "r_byte_t", name: "BlkSeed", fun: "avm_block",
			},
			{
				t: "uint64_t", name: "BlkTimestamp", fun: "avm_block",
			},
		},
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
