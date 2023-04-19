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
}

type BuiltinFunctionData struct {
	t    string
	name string
	op   string

	stack   []BuiltinFunctionParamData
	imm     []BuiltinFunctionParamData
	returns int
}

var builtin_functions = []BuiltinFunctionData{
	{
		t: "void", name: "avm_err", op: "err",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_err_op", op: "err",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_sha256", op: "sha256",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_sha256_op", op: "sha256",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_keccak256", op: "keccak256",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_keccak256_op", op: "keccak256",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_sha512_256", op: "sha512_256",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_sha512_256_op", op: "sha512_256",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_ed25519verify", op: "ed25519verify",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
			{t: "bytes", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_ed25519verify_op", op: "ed25519verify",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_ecdsa_verify", op: "ecdsa_verify",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
			{t: "bytes", name: "STACK_3", array: false},
			{t: "bytes", name: "STACK_4", array: false},
			{t: "bytes", name: "STACK_5", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1", array: false},
		},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_ecdsa_verify_op", op: "ecdsa_verify",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1", array: false},
		},
	},
	{
		t: "avm_ecdsa_pk_decompress_result_t", name: "avm_ecdsa_pk_decompress", op: "ecdsa_pk_decompress",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1", array: false},
		},
		returns: 2,
	},
	{
		t: "avm_ecdsa_pk_decompress_result_t", name: "avm_ecdsa_pk_decompress_op", op: "ecdsa_pk_decompress",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1", array: false},
		},
	},
	{
		t: "avm_ecdsa_pk_recover_result_t", name: "avm_ecdsa_pk_recover", op: "ecdsa_pk_recover",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
			{t: "bytes", name: "STACK_3", array: false},
			{t: "bytes", name: "STACK_4", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1", array: false},
		},
		returns: 2,
	},
	{
		t: "avm_ecdsa_pk_recover_result_t", name: "avm_ecdsa_pk_recover_op", op: "ecdsa_pk_recover",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "V1", array: false},
		},
	},
	{
		t: "uint64", name: "avm_plus", op: "+",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_plus_op", op: "+",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_minus", op: "-",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_minus_op", op: "-",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_div", op: "/",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_div_op", op: "/",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_mul", op: "*",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_mul_op", op: "*",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_lt", op: "<",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_lt_op", op: "<",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_gt", op: ">",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_gt_op", op: ">",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_lteq", op: "<=",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_lteq_op", op: "<=",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_gteq", op: ">=",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_gteq_op", op: ">=",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_andand", op: "&&",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_andand_op", op: "&&",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_oror", op: "||",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_oror_op", op: "||",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_eqeq", op: "==",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "any", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_eqeq_op", op: "==",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_noteq", op: "!=",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "any", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_noteq_op", op: "!=",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_not", op: "!",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_not_op", op: "!",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_len", op: "len",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_len_op", op: "len",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_itob", op: "itob",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_itob_op", op: "itob",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_btoi", op: "btoi",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_btoi_op", op: "btoi",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_mod", op: "%",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_mod_op", op: "%",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_or", op: "|",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_or_op", op: "|",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_and", op: "&",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_and_op", op: "&",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_xor", op: "^",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_xor_op", op: "^",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_inv", op: "~",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_inv_op", op: "~",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_mulw_result_t", name: "avm_mulw", op: "mulw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 2,
	},
	{
		t: "avm_mulw_result_t", name: "avm_mulw_op", op: "mulw",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_addw_result_t", name: "avm_addw", op: "addw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 2,
	},
	{
		t: "avm_addw_result_t", name: "avm_addw_op", op: "addw",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_divmodw_result_t", name: "avm_divmodw", op: "divmodw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
			{t: "uint64", name: "STACK_3", array: false},
			{t: "uint64", name: "STACK_4", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 4,
	},
	{
		t: "avm_divmodw_result_t", name: "avm_divmodw_op", op: "divmodw",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_intcblock", op: "intcblock",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint64", name: "UINT1", array: true},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_intcblock_op", op: "intcblock",
		imm: []BuiltinFunctionParamData{
			{t: "uint64", name: "UINT1", array: true},
		},
	},
	{
		t: "uint64", name: "avm_intc", op: "intc",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false},
		},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_intc_op", op: "intc",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false},
		},
	},
	{
		t: "uint64", name: "avm_intc_0", op: "intc_0",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_intc_0_op", op: "intc_0",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_intc_1", op: "intc_1",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_intc_1_op", op: "intc_1",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_intc_2", op: "intc_2",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_intc_2_op", op: "intc_2",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_intc_3", op: "intc_3",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_intc_3_op", op: "intc_3",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_bytecblock", op: "bytecblock",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1", array: true},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_bytecblock_op", op: "bytecblock",
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1", array: true},
		},
	},
	{
		t: "bytes", name: "avm_bytec", op: "bytec",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false},
		},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bytec_op", op: "bytec",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false},
		},
	},
	{
		t: "bytes", name: "avm_bytec_0", op: "bytec_0",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bytec_0_op", op: "bytec_0",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bytec_1", op: "bytec_1",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bytec_1_op", op: "bytec_1",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bytec_2", op: "bytec_2",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bytec_2_op", op: "bytec_2",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bytec_3", op: "bytec_3",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bytec_3_op", op: "bytec_3",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_arg", op: "arg",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_arg_op", op: "arg",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
	},
	{
		t: "bytes", name: "avm_arg_0", op: "arg_0",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_arg_0_op", op: "arg_0",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_arg_1", op: "arg_1",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_arg_1_op", op: "arg_1",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_arg_2", op: "arg_2",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_arg_2_op", op: "arg_2",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_arg_3", op: "arg_3",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_arg_3_op", op: "arg_3",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_txn", op: "txn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_txn_op", op: "txn",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
	},
	{
		t: "any", name: "avm_global", op: "global",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_global_op", op: "global",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
	},
	{
		t: "any", name: "avm_gtxn", op: "gtxn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "F2", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_gtxn_op", op: "gtxn",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "F2", array: false},
		},
	},
	{
		t: "any", name: "avm_load", op: "load",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_load_op", op: "load",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false},
		},
	},
	{
		t: "void", name: "avm_store", op: "store",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_store_op", op: "store",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false},
		},
	},
	{
		t: "any", name: "avm_txna", op: "txna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
			{t: "uint8", name: "I2", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_txna_op", op: "txna",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
			{t: "uint8", name: "I2", array: false},
		},
	},
	{
		t: "any", name: "avm_gtxna", op: "gtxna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "F2", array: false},
			{t: "uint8", name: "I3", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_gtxna_op", op: "gtxna",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "F2", array: false},
			{t: "uint8", name: "I3", array: false},
		},
	},
	{
		t: "any", name: "avm_gtxns", op: "gtxns",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_gtxns_op", op: "gtxns",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
	},
	{
		t: "any", name: "avm_gtxnsa", op: "gtxnsa",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
			{t: "uint8", name: "I2", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_gtxnsa_op", op: "gtxnsa",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
			{t: "uint8", name: "I2", array: false},
		},
	},
	{
		t: "any", name: "avm_gload", op: "gload",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "I2", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_gload_op", op: "gload",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "I2", array: false},
		},
	},
	{
		t: "any", name: "avm_gloads", op: "gloads",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_gloads_op", op: "gloads",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "I1", array: false},
		},
	},
	{
		t: "uint64", name: "avm_gaid", op: "gaid",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
		},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_gaid_op", op: "gaid",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
		},
	},
	{
		t: "uint64", name: "avm_gaids", op: "gaids",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_gaids_op", op: "gaids",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_loads", op: "loads",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "any", name: "avm_loads_op", op: "loads",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_stores", op: "stores",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "any", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_stores_op", op: "stores",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_bnz", op: "bnz",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_bnz_op", op: "bnz",
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false},
		},
	},
	{
		t: "void", name: "avm_bz", op: "bz",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_bz_op", op: "bz",
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false},
		},
	},
	{
		t: "void", name: "avm_b", op: "b",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_b_op", op: "b",
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false},
		},
	},
	{
		t: "void", name: "avm_return", op: "return",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_return_op", op: "return",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_assert", op: "assert",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_assert_op", op: "assert",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_bury", op: "bury",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_bury_op", op: "bury",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
	},
	{
		t: "void", name: "avm_popn", op: "popn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_popn_op", op: "popn",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
	},
	{
		t: "void", name: "avm_dupn", op: "dupn",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_dupn_op", op: "dupn",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
	},
	{
		t: "void", name: "avm_pop", op: "pop",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_pop_op", op: "pop",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_dup_result_t", name: "avm_dup", op: "dup",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 2,
	},
	{
		t: "avm_dup_result_t", name: "avm_dup_op", op: "dup",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_dup2_result_t", name: "avm_dup2", op: "dup2",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "any", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 4,
	},
	{
		t: "avm_dup2_result_t", name: "avm_dup2_op", op: "dup2",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_dig_result_t", name: "avm_dig", op: "dig",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
		returns: 2,
	},
	{
		t: "avm_dig_result_t", name: "avm_dig_op", op: "dig",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
	},
	{
		t: "avm_swap_result_t", name: "avm_swap", op: "swap",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "any", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 2,
	},
	{
		t: "avm_swap_result_t", name: "avm_swap_op", op: "swap",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_select", op: "select",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "any", name: "STACK_2", array: false},
			{t: "uint64", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "any", name: "avm_select_op", op: "select",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_cover", op: "cover",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_cover_op", op: "cover",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
	},
	{
		t: "any", name: "avm_uncover", op: "uncover",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_uncover_op", op: "uncover",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "N1", array: false},
		},
	},
	{
		t: "bytes", name: "avm_concat", op: "concat",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_concat_op", op: "concat",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_substring", op: "substring",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false},
			{t: "uint8", name: "E2", array: false},
		},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_substring_op", op: "substring",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false},
			{t: "uint8", name: "E2", array: false},
		},
	},
	{
		t: "bytes", name: "avm_substring3", op: "substring3",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
			{t: "uint64", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_substring3_op", op: "substring3",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_getbit", op: "getbit",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_getbit_op", op: "getbit",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_setbit", op: "setbit",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
			{t: "uint64", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "any", name: "avm_setbit_op", op: "setbit",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_getbyte", op: "getbyte",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_getbyte_op", op: "getbyte",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_setbyte", op: "setbyte",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
			{t: "uint64", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_setbyte_op", op: "setbyte",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_extract", op: "extract",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false},
			{t: "uint8", name: "L2", array: false},
		},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_extract_op", op: "extract",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false},
			{t: "uint8", name: "L2", array: false},
		},
	},
	{
		t: "bytes", name: "avm_extract3", op: "extract3",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
			{t: "uint64", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_extract3_op", op: "extract3",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_extract_uint16", op: "extract_uint16",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_extract_uint16_op", op: "extract_uint16",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_extract_uint32", op: "extract_uint32",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_extract_uint32_op", op: "extract_uint32",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_extract_uint64", op: "extract_uint64",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_extract_uint64_op", op: "extract_uint64",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_replace2", op: "replace2",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false},
		},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_replace2_op", op: "replace2",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false},
		},
	},
	{
		t: "bytes", name: "avm_replace3", op: "replace3",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
			{t: "bytes", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_replace3_op", op: "replace3",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_base64_decode", op: "base64_decode",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "E1", array: false},
		},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_base64_decode_op", op: "base64_decode",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "E1", array: false},
		},
	},
	{
		t: "any", name: "avm_json_ref", op: "json_ref",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "R1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_json_ref_op", op: "json_ref",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "R1", array: false},
		},
	},
	{
		t: "uint64", name: "avm_balance", op: "balance",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_balance_op", op: "balance",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_app_opted_in", op: "app_opted_in",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_app_opted_in_op", op: "app_opted_in",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_app_local_get", op: "app_local_get",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "any", name: "avm_app_local_get_op", op: "app_local_get",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_app_local_get_ex_result_t", name: "avm_app_local_get_ex", op: "app_local_get_ex",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
			{t: "bytes", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 2,
	},
	{
		t: "avm_app_local_get_ex_result_t", name: "avm_app_local_get_ex_op", op: "app_local_get_ex",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_app_global_get", op: "app_global_get",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "any", name: "avm_app_global_get_op", op: "app_global_get",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_app_global_get_ex_result_t", name: "avm_app_global_get_ex", op: "app_global_get_ex",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 2,
	},
	{
		t: "avm_app_global_get_ex_result_t", name: "avm_app_global_get_ex_op", op: "app_global_get_ex",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_app_local_put", op: "app_local_put",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
			{t: "any", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_app_local_put_op", op: "app_local_put",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_app_global_put", op: "app_global_put",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "any", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_app_global_put_op", op: "app_global_put",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_app_local_del", op: "app_local_del",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_app_local_del_op", op: "app_local_del",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_app_global_del", op: "app_global_del",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_app_global_del_op", op: "app_global_del",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_asset_holding_get_result_t", name: "avm_asset_holding_get", op: "asset_holding_get",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 2,
	},
	{
		t: "avm_asset_holding_get_result_t", name: "avm_asset_holding_get_op", op: "asset_holding_get",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
	},
	{
		t: "avm_asset_params_get_result_t", name: "avm_asset_params_get", op: "asset_params_get",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 2,
	},
	{
		t: "avm_asset_params_get_result_t", name: "avm_asset_params_get_op", op: "asset_params_get",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
	},
	{
		t: "avm_app_params_get_result_t", name: "avm_app_params_get", op: "app_params_get",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 2,
	},
	{
		t: "avm_app_params_get_result_t", name: "avm_app_params_get_op", op: "app_params_get",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
	},
	{
		t: "avm_acct_params_get_result_t", name: "avm_acct_params_get", op: "acct_params_get",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 2,
	},
	{
		t: "avm_acct_params_get_result_t", name: "avm_acct_params_get_op", op: "acct_params_get",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
	},
	{
		t: "uint64", name: "avm_min_balance", op: "min_balance",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_min_balance_op", op: "min_balance",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_pushbytes", op: "pushbytes",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1", array: false},
		},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_pushbytes_op", op: "pushbytes",
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1", array: false},
		},
	},
	{
		t: "uint64", name: "avm_pushint", op: "pushint",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "UINT1", array: false},
		},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_pushint_op", op: "pushint",
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "UINT1", array: false},
		},
	},
	{
		t: "void", name: "avm_pushbytess", op: "pushbytess",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1", array: true},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_pushbytess_op", op: "pushbytess",
		imm: []BuiltinFunctionParamData{
			{t: "bytes", name: "BYTES1", array: true},
		},
	},
	{
		t: "void", name: "avm_pushints", op: "pushints",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint64", name: "UINT1", array: true},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_pushints_op", op: "pushints",
		imm: []BuiltinFunctionParamData{
			{t: "uint64", name: "UINT1", array: true},
		},
	},
	{
		t: "uint64", name: "avm_ed25519verify_bare", op: "ed25519verify_bare",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
			{t: "bytes", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_ed25519verify_bare_op", op: "ed25519verify_bare",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_callsub", op: "callsub",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_callsub_op", op: "callsub",
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: false},
		},
	},
	{
		t: "void", name: "avm_retsub", op: "retsub",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_retsub_op", op: "retsub",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_proto", op: "proto",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "A1", array: false},
			{t: "uint8", name: "R2", array: false},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_proto_op", op: "proto",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "A1", array: false},
			{t: "uint8", name: "R2", array: false},
		},
	},
	{
		t: "any", name: "avm_frame_dig", op: "frame_dig",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "int8", name: "I1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_frame_dig_op", op: "frame_dig",
		imm: []BuiltinFunctionParamData{
			{t: "int8", name: "I1", array: false},
		},
	},
	{
		t: "void", name: "avm_frame_bury", op: "frame_bury",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "int8", name: "I1", array: false},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_frame_bury_op", op: "frame_bury",
		imm: []BuiltinFunctionParamData{
			{t: "int8", name: "I1", array: false},
		},
	},
	{
		t: "void", name: "avm_switch", op: "switch",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: true},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_switch_op", op: "switch",
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: true},
		},
	},
	{
		t: "void", name: "avm_match", op: "match",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: true},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_match_op", op: "match",
		imm: []BuiltinFunctionParamData{
			{t: "label", name: "TARGET1", array: true},
		},
	},
	{
		t: "uint64", name: "avm_shl", op: "shl",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_shl_op", op: "shl",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_shr", op: "shr",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_shr_op", op: "shr",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_sqrt", op: "sqrt",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_sqrt_op", op: "sqrt",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_bitlen", op: "bitlen",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_bitlen_op", op: "bitlen",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_exp", op: "exp",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_exp_op", op: "exp",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_expw_result_t", name: "avm_expw", op: "expw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 2,
	},
	{
		t: "avm_expw_result_t", name: "avm_expw_op", op: "expw",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bsqrt", op: "bsqrt",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bsqrt_op", op: "bsqrt",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_divw", op: "divw",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
			{t: "uint64", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_divw_op", op: "divw",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_sha3_256", op: "sha3_256",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_sha3_256_op", op: "sha3_256",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bplus", op: "b+",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bplus_op", op: "b+",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bminus", op: "b-",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bminus_op", op: "b-",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bdiv", op: "b/",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bdiv_op", op: "b/",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bmul", op: "b*",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bmul_op", op: "b*",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_blt", op: "b<",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_blt_op", op: "b<",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_bgt", op: "b>",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_bgt_op", op: "b>",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_blteq", op: "b<=",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_blteq_op", op: "b<=",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_bgteq", op: "b>=",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_bgteq_op", op: "b>=",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_beqeq", op: "b==",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_beqeq_op", op: "b==",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_bnoteq", op: "b!=",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_bnoteq_op", op: "b!=",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bmod", op: "b%",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bmod_op", op: "b%",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bor", op: "b|",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bor_op", op: "b|",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_band", op: "b&",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_band_op", op: "b&",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bxor", op: "b^",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bxor_op", op: "b^",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_binv", op: "b~",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_binv_op", op: "b~",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_bzero", op: "bzero",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_bzero_op", op: "bzero",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_log", op: "log",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_log_op", op: "log",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_itxn_begin", op: "itxn_begin",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_itxn_begin_op", op: "itxn_begin",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_itxn_field", op: "itxn_field",
		stack: []BuiltinFunctionParamData{
			{t: "any", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 0,
	},
	{
		t: "void", name: "avm_itxn_field_op", op: "itxn_field",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
	},
	{
		t: "void", name: "avm_itxn_submit", op: "itxn_submit",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_itxn_submit_op", op: "itxn_submit",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_itxn", op: "itxn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_itxn_op", op: "itxn",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
	},
	{
		t: "any", name: "avm_itxna", op: "itxna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
			{t: "uint8", name: "I2", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_itxna_op", op: "itxna",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
			{t: "uint8", name: "I2", array: false},
		},
	},
	{
		t: "void", name: "avm_itxn_next", op: "itxn_next",
		stack:   []BuiltinFunctionParamData{},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_itxn_next_op", op: "itxn_next",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_gitxn", op: "gitxn",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "F2", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_gitxn_op", op: "gitxn",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "F2", array: false},
		},
	},
	{
		t: "any", name: "avm_gitxna", op: "gitxna",
		stack: []BuiltinFunctionParamData{},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "F2", array: false},
			{t: "uint8", name: "I3", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_gitxna_op", op: "gitxna",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "F2", array: false},
			{t: "uint8", name: "I3", array: false},
		},
	},
	{
		t: "uint64", name: "avm_box_create", op: "box_create",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_box_create_op", op: "box_create",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "bytes", name: "avm_box_extract", op: "box_extract",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
			{t: "uint64", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_box_extract_op", op: "box_extract",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_box_replace", op: "box_replace",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
			{t: "bytes", name: "STACK_3", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_box_replace_op", op: "box_replace",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "uint64", name: "avm_box_del", op: "box_del",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "uint64", name: "avm_box_del_op", op: "box_del",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_box_len_result_t", name: "avm_box_len", op: "box_len",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 2,
	},
	{
		t: "avm_box_len_result_t", name: "avm_box_len_op", op: "box_len",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "avm_box_get_result_t", name: "avm_box_get", op: "box_get",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 2,
	},
	{
		t: "avm_box_get_result_t", name: "avm_box_get_op", op: "box_get",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "void", name: "avm_box_put", op: "box_put",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 0,
	},
	{
		t: "void", name: "avm_box_put_op", op: "box_put",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_txnas", op: "txnas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_txnas_op", op: "txnas",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
	},
	{
		t: "any", name: "avm_gtxnas", op: "gtxnas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "F2", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_gtxnas_op", op: "gtxnas",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "F2", array: false},
		},
	},
	{
		t: "any", name: "avm_gtxnsas", op: "gtxnsas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_gtxnsas_op", op: "gtxnsas",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
	},
	{
		t: "bytes", name: "avm_args", op: "args",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "bytes", name: "avm_args_op", op: "args",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_gloadss", op: "gloadss",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
			{t: "uint64", name: "STACK_2", array: false},
		},
		imm:     []BuiltinFunctionParamData{},
		returns: 1,
	},
	{
		t: "any", name: "avm_gloadss_op", op: "gloadss",
		imm: []BuiltinFunctionParamData{},
	},
	{
		t: "any", name: "avm_itxnas", op: "itxnas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_itxnas_op", op: "itxnas",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
	},
	{
		t: "any", name: "avm_gitxnas", op: "gitxnas",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "F2", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_gitxnas_op", op: "gitxnas",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "T1", array: false},
			{t: "uint8", name: "F2", array: false},
		},
	},
	{
		t: "avm_vrf_verify_result_t", name: "avm_vrf_verify", op: "vrf_verify",
		stack: []BuiltinFunctionParamData{
			{t: "bytes", name: "STACK_1", array: false},
			{t: "bytes", name: "STACK_2", array: false},
			{t: "bytes", name: "STACK_3", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false},
		},
		returns: 2,
	},
	{
		t: "avm_vrf_verify_result_t", name: "avm_vrf_verify_op", op: "vrf_verify",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "S1", array: false},
		},
	},
	{
		t: "any", name: "avm_block", op: "block",
		stack: []BuiltinFunctionParamData{
			{t: "uint64", name: "STACK_1", array: false},
		},
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
		},
		returns: 1,
	},
	{
		t: "any", name: "avm_block_op", op: "block",
		imm: []BuiltinFunctionParamData{
			{t: "uint8", name: "F1", array: false},
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
