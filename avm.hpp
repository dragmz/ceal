#pragma once

#ifndef CEAL
#include <string>

constexpr uint64_t fnv1a_hash(const char* str, uint64_t hash = 14695981039346656037u) {
    return (*str == '\0') ? hash : fnv1a_hash(str + 1, (hash ^ static_cast<uint64_t>(*str)) * 1099511628211u);
}

#define const_string(x) fnv1a_hash(x)
#define avm_method(x) const_string(#x)

#define IMMEDIATE
#define STACK

using uint64 = unsigned long long;
using uint8 = unsigned char;
using int8 = signed char;
using none = void;

using uint64_t = uint64;
using none_t = none;
using name_t = bytes;
using method_t = bytes;
using key_t = bytes;
using bool_t = uint64;
using bigint_t = bytes;
using any_t = any;
using addr_t = bytes;
using r_byte_t = bytes;
using r32_byte_t = bytes;

struct bytes
{
	bytes() {}
	bytes(const char *) {}
	bytes(const unsigned char *) {}
	uint8& operator[](const uint64 index) const {}
};

using label = std::string;

template <typename T>
void abi_decode(bytes data, T &out){};

template <typename T>
bytes abi_encode(const T &in) { return {}; }

const uint64 NoOp = 0;
const uint64 OptIn = 1;
const uint64 CloseOut = 2;
const uint64 ClearState = 3;
const uint64 UpdateApplication = 4;
const uint64 DeleteApplication = 5;

struct any
{
	any() = default;

	template <typename T>
	any(T);

	template <typename T>
	any operator=(T);

	template <typename T>
	bool operator==(T);
};

void avm_err_op();
/*
err - Fail immediately.
*/
void avm_err();
void avm_sha256_op();
/*
sha256 - SHA256 hash of value A, yields [32]byte
*/
r32_byte_t avm_sha256(STACK r_byte_t STACK_1);
void avm_keccak256_op();
/*
keccak256 - Keccak256 hash of value A, yields [32]byte
*/
r32_byte_t avm_keccak256(STACK r_byte_t STACK_1);
void avm_sha512_256_op();
/*
sha512_256 - SHA512_256 hash of value A, yields [32]byte
*/
r32_byte_t avm_sha512_256(STACK r_byte_t STACK_1);
void avm_ed25519verify_op();
/*
ed25519verify - for (data A, signature B, pubkey C) verify the signature of ("ProgData" || program_hash || data) against the pubkey => {0 or 1}

The 32 byte public key is the last element on the stack, preceded by the 64 byte signature at the second-to-last element on the stack, preceded by the data which was signed at the third-to-last element on the stack.
*/
bool_t avm_ed25519verify(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2, STACK r_byte_t STACK_3);
void avm_ecdsa_verify_op(IMMEDIATE uint8 V1);
/*
ecdsa_verify - for (data A, signature B, C and pubkey D, E) verify the signature of the data against the pubkey => {0 or 1}

The 32 byte Y-component of a public key is the last element on the stack, preceded by X-component of a pubkey, preceded by S and R components of a signature, preceded by the data that is fifth element on the stack. All values are big-endian encoded. The signed data must be 32 bytes long, and signatures in lower-S form are only accepted.
*/
struct avm_ecdsa_verify_t
{
	void Secp256k1(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2, STACK r_byte_t STACK_3, STACK r_byte_t STACK_4, STACK r_byte_t STACK_5);
	void Secp256r1(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2, STACK r_byte_t STACK_3, STACK r_byte_t STACK_4, STACK r_byte_t STACK_5);
};
extern avm_ecdsa_verify_t avm_ecdsa_verify;
void avm_ecdsa_pk_decompress_op(IMMEDIATE uint8 V1);
struct avm_ecdsa_pk_decompress_result_t
{
	r_byte_t r1;
	r_byte_t r2;
};
/*
ecdsa_pk_decompress - decompress pubkey A into components X, Y

The 33 byte public key in a compressed form to be decompressed into X and Y (top) components. All values are big-endian encoded.
*/
struct avm_ecdsa_pk_decompress_t
{
	void Secp256k1(STACK r_byte_t STACK_1);
	void Secp256r1(STACK r_byte_t STACK_1);
};
extern avm_ecdsa_pk_decompress_t avm_ecdsa_pk_decompress;
void avm_ecdsa_pk_recover_op(IMMEDIATE uint8 V1);
struct avm_ecdsa_pk_recover_result_t
{
	r_byte_t r1;
	r_byte_t r2;
};
/*
ecdsa_pk_recover - for (data A, recovery id B, signature C, D) recover a public key

S (top) and R elements of a signature, recovery id and data (bottom) are expected on the stack and used to deriver a public key. All values are big-endian encoded. The signed data must be 32 bytes long.
*/
struct avm_ecdsa_pk_recover_t
{
	void Secp256k1(STACK r_byte_t STACK_1, STACK uint64_t STACK_2, STACK r_byte_t STACK_3, STACK r_byte_t STACK_4);
	void Secp256r1(STACK r_byte_t STACK_1, STACK uint64_t STACK_2, STACK r_byte_t STACK_3, STACK r_byte_t STACK_4);
};
extern avm_ecdsa_pk_recover_t avm_ecdsa_pk_recover;
void avm_plus_op();
/*
+ - A plus B. Fail on overflow.

Overflow is an error condition which halts execution and fails the transaction. Full precision is available from `addw`.
*/
uint64_t avm_plus(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_minus_op();
/*
- - A minus B. Fail if B > A.
*/
uint64_t avm_minus(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_div_op();
/*
/ - A divided by B (truncated division). Fail if B == 0.

`divmodw` is available to divide the two-element values produced by `mulw` and `addw`.
*/
uint64_t avm_div(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_mul_op();
/*
* - A times B. Fail on overflow.

Overflow is an error condition which halts execution and fails the transaction. Full precision is available from `mulw`.
*/
uint64_t avm_mul(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_lt_op();
/*
< - A less than B => {0 or 1}
*/
bool_t avm_lt(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_gt_op();
/*
> - A greater than B => {0 or 1}
*/
bool_t avm_gt(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_lteq_op();
/*
<= - A less than or equal to B => {0 or 1}
*/
bool_t avm_lteq(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_gteq_op();
/*
>= - A greater than or equal to B => {0 or 1}
*/
bool_t avm_gteq(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_andand_op();
/*
&& - A is not zero and B is not zero => {0 or 1}
*/
bool_t avm_andand(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_oror_op();
/*
|| - A is not zero or B is not zero => {0 or 1}
*/
bool_t avm_oror(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_eqeq_op();
/*
== - A is equal to B => {0 or 1}
*/
bool_t avm_eqeq(STACK any_t STACK_1, STACK any_t STACK_2);
void avm_noteq_op();
/*
!= - A is not equal to B => {0 or 1}
*/
bool_t avm_noteq(STACK any_t STACK_1, STACK any_t STACK_2);
void avm_not_op();
/*
! - A == 0 yields 1; else 0
*/
uint64_t avm_not(STACK uint64_t STACK_1);
void avm_len_op();
/*
len - yields length of byte value A
*/
uint64_t avm_len(STACK r_byte_t STACK_1);
void avm_itob_op();
/*
itob - converts uint64 A to big-endian byte array, always of length 8
*/
r_byte_t avm_itob(STACK uint64_t STACK_1);
void avm_btoi_op();
/*
btoi - converts big-endian byte array A to uint64. Fails if len(A) > 8. Padded by leading 0s if len(A) < 8.

`btoi` fails if the input is longer than 8 bytes.
*/
uint64_t avm_btoi(STACK r_byte_t STACK_1);
void avm_mod_op();
/*
% - A modulo B. Fail if B == 0.
*/
uint64_t avm_mod(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_or_op();
/*
| - A bitwise-or B
*/
uint64_t avm_or(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_and_op();
/*
& - A bitwise-and B
*/
uint64_t avm_and(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_xor_op();
/*
^ - A bitwise-xor B
*/
uint64_t avm_xor(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_inv_op();
/*
~ - bitwise invert value A
*/
uint64_t avm_inv(STACK uint64_t STACK_1);
void avm_mulw_op();
struct avm_mulw_result_t
{
	uint64_t r1;
	uint64_t r2;
};
/*
mulw - A times B as a 128-bit result in two uint64s. X is the high 64 bits, Y is the low
*/
avm_mulw_result_t avm_mulw(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_addw_op();
struct avm_addw_result_t
{
	uint64_t r1;
	uint64_t r2;
};
/*
addw - A plus B as a 128-bit result. X is the carry-bit, Y is the low-order 64 bits.
*/
avm_addw_result_t avm_addw(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_divmodw_op();
struct avm_divmodw_result_t
{
	uint64_t r1;
	uint64_t r2;
	uint64_t r3;
	uint64_t r4;
};
/*
divmodw - W,X = (A,B / C,D); Y,Z = (A,B modulo C,D)

The notation J,K indicates that two uint64 values J and K are interpreted as a uint128 value, with J as the high uint64 and K the low.
*/
avm_divmodw_result_t avm_divmodw(STACK uint64_t STACK_1, STACK uint64_t STACK_2, STACK uint64_t STACK_3, STACK uint64_t STACK_4);
void avm_intcblock_op(IMMEDIATE uint64 UINT1, ...);
/*
intcblock - prepare block of uint64 constants for use by intc

`intcblock` loads following program bytes into an array of integer constants in the evaluator. These integer constants can be referred to by `intc` and `intc_*` which will push the value onto the stack. Subsequent calls to `intcblock` reset and replace the integer constants available to the script.
*/
struct avm_intcblock_t
{
};
extern avm_intcblock_t avm_intcblock;
void avm_intc_op(IMMEDIATE uint8 I1);
/*
intc - Ith constant from intcblock
*/
struct avm_intc_t
{
};
extern avm_intc_t avm_intc;
void avm_intc_0_op();
/*
intc_0 - constant 0 from intcblock
*/
uint64_t avm_intc_0();
void avm_intc_1_op();
/*
intc_1 - constant 1 from intcblock
*/
uint64_t avm_intc_1();
void avm_intc_2_op();
/*
intc_2 - constant 2 from intcblock
*/
uint64_t avm_intc_2();
void avm_intc_3_op();
/*
intc_3 - constant 3 from intcblock
*/
uint64_t avm_intc_3();
void avm_bytecblock_op(IMMEDIATE bytes BYTES1, ...);
/*
bytecblock - prepare block of byte-array constants for use by bytec

`bytecblock` loads the following program bytes into an array of byte-array constants in the evaluator. These constants can be referred to by `bytec` and `bytec_*` which will push the value onto the stack. Subsequent calls to `bytecblock` reset and replace the bytes constants available to the script.
*/
struct avm_bytecblock_t
{
};
extern avm_bytecblock_t avm_bytecblock;
void avm_bytec_op(IMMEDIATE uint8 I1);
/*
bytec - Ith constant from bytecblock
*/
struct avm_bytec_t
{
};
extern avm_bytec_t avm_bytec;
void avm_bytec_0_op();
/*
bytec_0 - constant 0 from bytecblock
*/
r_byte_t avm_bytec_0();
void avm_bytec_1_op();
/*
bytec_1 - constant 1 from bytecblock
*/
r_byte_t avm_bytec_1();
void avm_bytec_2_op();
/*
bytec_2 - constant 2 from bytecblock
*/
r_byte_t avm_bytec_2();
void avm_bytec_3_op();
/*
bytec_3 - constant 3 from bytecblock
*/
r_byte_t avm_bytec_3();
void avm_arg_op(IMMEDIATE uint8 N1);
/*
arg - Nth LogicSig argument
*/
struct avm_arg_t
{
};
extern avm_arg_t avm_arg;
void avm_arg_0_op();
/*
arg_0 - LogicSig argument 0
*/
r_byte_t avm_arg_0();
void avm_arg_1_op();
/*
arg_1 - LogicSig argument 1
*/
r_byte_t avm_arg_1();
void avm_arg_2_op();
/*
arg_2 - LogicSig argument 2
*/
r_byte_t avm_arg_2();
void avm_arg_3_op();
/*
arg_3 - LogicSig argument 3
*/
r_byte_t avm_arg_3();
void avm_txn_op(IMMEDIATE uint8 F1);
/*
txn - field F of current transaction
*/
struct avm_txn_t
{
	const addr_t Sender;
	const uint64_t Fee;
	const uint64_t FirstValid;
	const uint64_t FirstValidTime;
	const uint64_t LastValid;
	const r_byte_t Note;
	const r32_byte_t Lease;
	const addr_t Receiver;
	const uint64_t Amount;
	const addr_t CloseRemainderTo;
	const r32_byte_t VotePK;
	const r32_byte_t SelectionPK;
	const uint64_t VoteFirst;
	const uint64_t VoteLast;
	const uint64_t VoteKeyDilution;
	const r_byte_t Type;
	const uint64_t TypeEnum;
	const uint64_t XferAsset;
	const uint64_t AssetAmount;
	const addr_t AssetSender;
	const addr_t AssetReceiver;
	const addr_t AssetCloseTo;
	const uint64_t GroupIndex;
	const r32_byte_t TxID;
	const uint64_t ApplicationID;
	const uint64_t OnCompletion;
	const r_byte_t ApplicationArgs;
	const uint64_t NumAppArgs;
	const addr_t Accounts;
	const uint64_t NumAccounts;
	const r_byte_t ApprovalProgram;
	const r_byte_t ClearStateProgram;
	const addr_t RekeyTo;
	const uint64_t ConfigAsset;
	const uint64_t ConfigAssetTotal;
	const uint64_t ConfigAssetDecimals;
	const bool_t ConfigAssetDefaultFrozen;
	const r_byte_t ConfigAssetUnitName;
	const r_byte_t ConfigAssetName;
	const r_byte_t ConfigAssetURL;
	const r32_byte_t ConfigAssetMetadataHash;
	const addr_t ConfigAssetManager;
	const addr_t ConfigAssetReserve;
	const addr_t ConfigAssetFreeze;
	const addr_t ConfigAssetClawback;
	const uint64_t FreezeAsset;
	const addr_t FreezeAssetAccount;
	const bool_t FreezeAssetFrozen;
	const uint64_t Assets;
	const uint64_t NumAssets;
	const uint64_t Applications;
	const uint64_t NumApplications;
	const uint64_t GlobalNumUint;
	const uint64_t GlobalNumByteSlice;
	const uint64_t LocalNumUint;
	const uint64_t LocalNumByteSlice;
	const uint64_t ExtraProgramPages;
	const bool_t Nonparticipation;
	const r_byte_t Logs;
	const uint64_t NumLogs;
	const uint64_t CreatedAssetID;
	const uint64_t CreatedApplicationID;
	const r_byte_t LastLog;
	const r_byte_t StateProofPK;
	const r_byte_t ApprovalProgramPages;
	const uint64_t NumApprovalProgramPages;
	const r_byte_t ClearStateProgramPages;
	const uint64_t NumClearStateProgramPages;
};
extern avm_txn_t avm_txn;
void avm_global_op(IMMEDIATE uint8 F1);
/*
global - global field F
*/
struct avm_global_t
{
	const uint64_t MinTxnFee;
	const uint64_t MinBalance;
	const uint64_t MaxTxnLife;
	const addr_t ZeroAddress;
	const uint64_t GroupSize;
	const uint64_t LogicSigVersion;
	const uint64_t Round;
	const uint64_t LatestTimestamp;
	const uint64_t CurrentApplicationID;
	const addr_t CreatorAddress;
	const addr_t CurrentApplicationAddress;
	const r32_byte_t GroupID;
	const uint64_t OpcodeBudget;
	const uint64_t CallerApplicationID;
	const addr_t CallerApplicationAddress;
};
extern avm_global_t avm_global;
void avm_gtxn_op(IMMEDIATE uint8 T1, IMMEDIATE uint8 F2);
/*
gtxn - field F of the Tth transaction in the current group

for notes on transaction fields available, see `txn`. If this transaction is _i_ in the group, `gtxn i field` is equivalent to `txn field`.
*/
struct avm_gtxn_t
{
	addr_t Sender(IMMEDIATE uint8 F2);
	uint64_t Fee(IMMEDIATE uint8 F2);
	uint64_t FirstValid(IMMEDIATE uint8 F2);
	uint64_t FirstValidTime(IMMEDIATE uint8 F2);
	uint64_t LastValid(IMMEDIATE uint8 F2);
	r_byte_t Note(IMMEDIATE uint8 F2);
	r32_byte_t Lease(IMMEDIATE uint8 F2);
	addr_t Receiver(IMMEDIATE uint8 F2);
	uint64_t Amount(IMMEDIATE uint8 F2);
	addr_t CloseRemainderTo(IMMEDIATE uint8 F2);
	r32_byte_t VotePK(IMMEDIATE uint8 F2);
	r32_byte_t SelectionPK(IMMEDIATE uint8 F2);
	uint64_t VoteFirst(IMMEDIATE uint8 F2);
	uint64_t VoteLast(IMMEDIATE uint8 F2);
	uint64_t VoteKeyDilution(IMMEDIATE uint8 F2);
	r_byte_t Type(IMMEDIATE uint8 F2);
	uint64_t TypeEnum(IMMEDIATE uint8 F2);
	uint64_t XferAsset(IMMEDIATE uint8 F2);
	uint64_t AssetAmount(IMMEDIATE uint8 F2);
	addr_t AssetSender(IMMEDIATE uint8 F2);
	addr_t AssetReceiver(IMMEDIATE uint8 F2);
	addr_t AssetCloseTo(IMMEDIATE uint8 F2);
	uint64_t GroupIndex(IMMEDIATE uint8 F2);
	r32_byte_t TxID(IMMEDIATE uint8 F2);
	uint64_t ApplicationID(IMMEDIATE uint8 F2);
	uint64_t OnCompletion(IMMEDIATE uint8 F2);
	r_byte_t ApplicationArgs(IMMEDIATE uint8 F2);
	uint64_t NumAppArgs(IMMEDIATE uint8 F2);
	addr_t Accounts(IMMEDIATE uint8 F2);
	uint64_t NumAccounts(IMMEDIATE uint8 F2);
	r_byte_t ApprovalProgram(IMMEDIATE uint8 F2);
	r_byte_t ClearStateProgram(IMMEDIATE uint8 F2);
	addr_t RekeyTo(IMMEDIATE uint8 F2);
	uint64_t ConfigAsset(IMMEDIATE uint8 F2);
	uint64_t ConfigAssetTotal(IMMEDIATE uint8 F2);
	uint64_t ConfigAssetDecimals(IMMEDIATE uint8 F2);
	bool_t ConfigAssetDefaultFrozen(IMMEDIATE uint8 F2);
	r_byte_t ConfigAssetUnitName(IMMEDIATE uint8 F2);
	r_byte_t ConfigAssetName(IMMEDIATE uint8 F2);
	r_byte_t ConfigAssetURL(IMMEDIATE uint8 F2);
	r32_byte_t ConfigAssetMetadataHash(IMMEDIATE uint8 F2);
	addr_t ConfigAssetManager(IMMEDIATE uint8 F2);
	addr_t ConfigAssetReserve(IMMEDIATE uint8 F2);
	addr_t ConfigAssetFreeze(IMMEDIATE uint8 F2);
	addr_t ConfigAssetClawback(IMMEDIATE uint8 F2);
	uint64_t FreezeAsset(IMMEDIATE uint8 F2);
	addr_t FreezeAssetAccount(IMMEDIATE uint8 F2);
	bool_t FreezeAssetFrozen(IMMEDIATE uint8 F2);
	uint64_t Assets(IMMEDIATE uint8 F2);
	uint64_t NumAssets(IMMEDIATE uint8 F2);
	uint64_t Applications(IMMEDIATE uint8 F2);
	uint64_t NumApplications(IMMEDIATE uint8 F2);
	uint64_t GlobalNumUint(IMMEDIATE uint8 F2);
	uint64_t GlobalNumByteSlice(IMMEDIATE uint8 F2);
	uint64_t LocalNumUint(IMMEDIATE uint8 F2);
	uint64_t LocalNumByteSlice(IMMEDIATE uint8 F2);
	uint64_t ExtraProgramPages(IMMEDIATE uint8 F2);
	bool_t Nonparticipation(IMMEDIATE uint8 F2);
	r_byte_t Logs(IMMEDIATE uint8 F2);
	uint64_t NumLogs(IMMEDIATE uint8 F2);
	uint64_t CreatedAssetID(IMMEDIATE uint8 F2);
	uint64_t CreatedApplicationID(IMMEDIATE uint8 F2);
	r_byte_t LastLog(IMMEDIATE uint8 F2);
	r_byte_t StateProofPK(IMMEDIATE uint8 F2);
	r_byte_t ApprovalProgramPages(IMMEDIATE uint8 F2);
	uint64_t NumApprovalProgramPages(IMMEDIATE uint8 F2);
	r_byte_t ClearStateProgramPages(IMMEDIATE uint8 F2);
	uint64_t NumClearStateProgramPages(IMMEDIATE uint8 F2);
};
extern avm_gtxn_t avm_gtxn;
void avm_load_op(IMMEDIATE uint8 I1);
/*
load - Ith scratch space value. All scratch spaces are 0 at program start.
*/
struct avm_load_t
{
};
extern avm_load_t avm_load;
void avm_store_op(IMMEDIATE uint8 I1);
/*
store - store A to the Ith scratch space
*/
struct avm_store_t
{
};
extern avm_store_t avm_store;
void avm_txna_op(IMMEDIATE uint8 F1, IMMEDIATE uint8 I2);
/*
txna - Ith value of the array field F of the current transaction
`txna` can be called using `txn` with 2 immediates.
*/
struct avm_txna_t
{
	r_byte_t ApplicationArgs(IMMEDIATE uint8 I2);
	addr_t Accounts(IMMEDIATE uint8 I2);
	uint64_t Assets(IMMEDIATE uint8 I2);
	uint64_t Applications(IMMEDIATE uint8 I2);
	r_byte_t Logs(IMMEDIATE uint8 I2);
	r_byte_t ApprovalProgramPages(IMMEDIATE uint8 I2);
	r_byte_t ClearStateProgramPages(IMMEDIATE uint8 I2);
};
extern avm_txna_t avm_txna;
void avm_gtxna_op(IMMEDIATE uint8 T1, IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
/*
gtxna - Ith value of the array field F from the Tth transaction in the current group
`gtxna` can be called using `gtxn` with 3 immediates.
*/
struct avm_gtxna_t
{
	r_byte_t ApplicationArgs(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
	addr_t Accounts(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
	uint64_t Assets(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
	uint64_t Applications(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
	r_byte_t Logs(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
	r_byte_t ApprovalProgramPages(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
	r_byte_t ClearStateProgramPages(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
};
extern avm_gtxna_t avm_gtxna;
void avm_gtxns_op(IMMEDIATE uint8 F1);
/*
gtxns - field F of the Ath transaction in the current group

for notes on transaction fields available, see `txn`. If top of stack is _i_, `gtxns field` is equivalent to `gtxn _i_ field`. gtxns exists so that _i_ can be calculated, often based on the index of the current transaction.
*/
struct avm_gtxns_t
{
	addr_t Sender(STACK uint64_t STACK_1);
	uint64_t Fee(STACK uint64_t STACK_1);
	uint64_t FirstValid(STACK uint64_t STACK_1);
	uint64_t FirstValidTime(STACK uint64_t STACK_1);
	uint64_t LastValid(STACK uint64_t STACK_1);
	r_byte_t Note(STACK uint64_t STACK_1);
	r32_byte_t Lease(STACK uint64_t STACK_1);
	addr_t Receiver(STACK uint64_t STACK_1);
	uint64_t Amount(STACK uint64_t STACK_1);
	addr_t CloseRemainderTo(STACK uint64_t STACK_1);
	r32_byte_t VotePK(STACK uint64_t STACK_1);
	r32_byte_t SelectionPK(STACK uint64_t STACK_1);
	uint64_t VoteFirst(STACK uint64_t STACK_1);
	uint64_t VoteLast(STACK uint64_t STACK_1);
	uint64_t VoteKeyDilution(STACK uint64_t STACK_1);
	r_byte_t Type(STACK uint64_t STACK_1);
	uint64_t TypeEnum(STACK uint64_t STACK_1);
	uint64_t XferAsset(STACK uint64_t STACK_1);
	uint64_t AssetAmount(STACK uint64_t STACK_1);
	addr_t AssetSender(STACK uint64_t STACK_1);
	addr_t AssetReceiver(STACK uint64_t STACK_1);
	addr_t AssetCloseTo(STACK uint64_t STACK_1);
	uint64_t GroupIndex(STACK uint64_t STACK_1);
	r32_byte_t TxID(STACK uint64_t STACK_1);
	uint64_t ApplicationID(STACK uint64_t STACK_1);
	uint64_t OnCompletion(STACK uint64_t STACK_1);
	r_byte_t ApplicationArgs(STACK uint64_t STACK_1);
	uint64_t NumAppArgs(STACK uint64_t STACK_1);
	addr_t Accounts(STACK uint64_t STACK_1);
	uint64_t NumAccounts(STACK uint64_t STACK_1);
	r_byte_t ApprovalProgram(STACK uint64_t STACK_1);
	r_byte_t ClearStateProgram(STACK uint64_t STACK_1);
	addr_t RekeyTo(STACK uint64_t STACK_1);
	uint64_t ConfigAsset(STACK uint64_t STACK_1);
	uint64_t ConfigAssetTotal(STACK uint64_t STACK_1);
	uint64_t ConfigAssetDecimals(STACK uint64_t STACK_1);
	bool_t ConfigAssetDefaultFrozen(STACK uint64_t STACK_1);
	r_byte_t ConfigAssetUnitName(STACK uint64_t STACK_1);
	r_byte_t ConfigAssetName(STACK uint64_t STACK_1);
	r_byte_t ConfigAssetURL(STACK uint64_t STACK_1);
	r32_byte_t ConfigAssetMetadataHash(STACK uint64_t STACK_1);
	addr_t ConfigAssetManager(STACK uint64_t STACK_1);
	addr_t ConfigAssetReserve(STACK uint64_t STACK_1);
	addr_t ConfigAssetFreeze(STACK uint64_t STACK_1);
	addr_t ConfigAssetClawback(STACK uint64_t STACK_1);
	uint64_t FreezeAsset(STACK uint64_t STACK_1);
	addr_t FreezeAssetAccount(STACK uint64_t STACK_1);
	bool_t FreezeAssetFrozen(STACK uint64_t STACK_1);
	uint64_t Assets(STACK uint64_t STACK_1);
	uint64_t NumAssets(STACK uint64_t STACK_1);
	uint64_t Applications(STACK uint64_t STACK_1);
	uint64_t NumApplications(STACK uint64_t STACK_1);
	uint64_t GlobalNumUint(STACK uint64_t STACK_1);
	uint64_t GlobalNumByteSlice(STACK uint64_t STACK_1);
	uint64_t LocalNumUint(STACK uint64_t STACK_1);
	uint64_t LocalNumByteSlice(STACK uint64_t STACK_1);
	uint64_t ExtraProgramPages(STACK uint64_t STACK_1);
	bool_t Nonparticipation(STACK uint64_t STACK_1);
	r_byte_t Logs(STACK uint64_t STACK_1);
	uint64_t NumLogs(STACK uint64_t STACK_1);
	uint64_t CreatedAssetID(STACK uint64_t STACK_1);
	uint64_t CreatedApplicationID(STACK uint64_t STACK_1);
	r_byte_t LastLog(STACK uint64_t STACK_1);
	r_byte_t StateProofPK(STACK uint64_t STACK_1);
	r_byte_t ApprovalProgramPages(STACK uint64_t STACK_1);
	uint64_t NumApprovalProgramPages(STACK uint64_t STACK_1);
	r_byte_t ClearStateProgramPages(STACK uint64_t STACK_1);
	uint64_t NumClearStateProgramPages(STACK uint64_t STACK_1);
};
extern avm_gtxns_t avm_gtxns;
void avm_gtxnsa_op(IMMEDIATE uint8 F1, IMMEDIATE uint8 I2);
/*
gtxnsa - Ith value of the array field F from the Ath transaction in the current group
`gtxnsa` can be called using `gtxns` with 2 immediates.
*/
struct avm_gtxnsa_t
{
	r_byte_t ApplicationArgs(STACK uint64_t STACK_1, IMMEDIATE uint8 I2);
	addr_t Accounts(STACK uint64_t STACK_1, IMMEDIATE uint8 I2);
	uint64_t Assets(STACK uint64_t STACK_1, IMMEDIATE uint8 I2);
	uint64_t Applications(STACK uint64_t STACK_1, IMMEDIATE uint8 I2);
	r_byte_t Logs(STACK uint64_t STACK_1, IMMEDIATE uint8 I2);
	r_byte_t ApprovalProgramPages(STACK uint64_t STACK_1, IMMEDIATE uint8 I2);
	r_byte_t ClearStateProgramPages(STACK uint64_t STACK_1, IMMEDIATE uint8 I2);
};
extern avm_gtxnsa_t avm_gtxnsa;
void avm_gload_op(IMMEDIATE uint8 T1, IMMEDIATE uint8 I2);
/*
gload - Ith scratch space value of the Tth transaction in the current group

`gload` fails unless the requested transaction is an ApplicationCall and T < GroupIndex.
*/
struct avm_gload_t
{
};
extern avm_gload_t avm_gload;
void avm_gloads_op(IMMEDIATE uint8 I1);
/*
gloads - Ith scratch space value of the Ath transaction in the current group

`gloads` fails unless the requested transaction is an ApplicationCall and A < GroupIndex.
*/
struct avm_gloads_t
{
};
extern avm_gloads_t avm_gloads;
void avm_gaid_op(IMMEDIATE uint8 T1);
/*
gaid - ID of the asset or application created in the Tth transaction of the current group

`gaid` fails unless the requested transaction created an asset or application and T < GroupIndex.
*/
struct avm_gaid_t
{
};
extern avm_gaid_t avm_gaid;
void avm_gaids_op();
/*
gaids - ID of the asset or application created in the Ath transaction of the current group

`gaids` fails unless the requested transaction created an asset or application and A < GroupIndex.
*/
uint64_t avm_gaids(STACK uint64_t STACK_1);
void avm_loads_op();
/*
loads - Ath scratch space value.  All scratch spaces are 0 at program start.
*/
any_t avm_loads(STACK uint64_t STACK_1);
void avm_stores_op();
/*
stores - store B to the Ath scratch space
*/
void avm_stores(STACK uint64_t STACK_1, STACK any_t STACK_2);
void avm_bnz_op(IMMEDIATE label TARGET1);
/*
bnz - branch to TARGET if value A is not zero

The `bnz` instruction opcode 0x40 is followed by two immediate data bytes which are a high byte first and low byte second which together form a 16 bit offset which the instruction may branch to. For a bnz instruction at `pc`, if the last element of the stack is not zero then branch to instruction at `pc + 3 + N`, else proceed to next instruction at `pc + 3`. Branch targets must be aligned instructions. (e.g. Branching to the second byte of a 2 byte op will be rejected.) Starting at v4, the offset is treated as a signed 16 bit integer allowing for backward branches and looping. In prior version (v1 to v3), branch offsets are limited to forward branches only, 0-0x7fff.

At v2 it became allowed to branch to the end of the program exactly after the last instruction: bnz to byte N (with 0-indexing) was illegal for a TEAL program with N bytes before v2, and is legal after it. This change eliminates the need for a last instruction of no-op as a branch target at the end. (Branching beyond the end--in other words, to a byte larger than N--is still illegal and will cause the program to fail.)
*/
struct avm_bnz_t
{
};
extern avm_bnz_t avm_bnz;
void avm_bz_op(IMMEDIATE label TARGET1);
/*
bz - branch to TARGET if value A is zero

See `bnz` for details on how branches work. `bz` inverts the behavior of `bnz`.
*/
struct avm_bz_t
{
};
extern avm_bz_t avm_bz;
void avm_b_op(IMMEDIATE label TARGET1);
/*
b - branch unconditionally to TARGET

See `bnz` for details on how branches work. `b` always jumps to the offset.
*/
struct avm_b_t
{
};
extern avm_b_t avm_b;
void avm_return_op();
/*
return - use A as success value; end
*/
void avm_return(STACK uint64_t STACK_1);
void avm_assert_op();
/*
assert - immediately fail unless A is a non-zero number
*/
void avm_assert(STACK uint64_t STACK_1);
void avm_bury_op(IMMEDIATE uint8 N1);
/*
bury - replace the Nth value from the top of the stack with A. bury 0 fails.
*/
struct avm_bury_t
{
};
extern avm_bury_t avm_bury;
void avm_popn_op(IMMEDIATE uint8 N1);
/*
popn - remove N values from the top of the stack
*/
struct avm_popn_t
{
};
extern avm_popn_t avm_popn;
void avm_dupn_op(IMMEDIATE uint8 N1);
/*
dupn - duplicate A, N times
*/
struct avm_dupn_t
{
};
extern avm_dupn_t avm_dupn;
void avm_pop_op();
/*
pop - discard A
*/
void avm_pop(STACK any_t STACK_1);
void avm_dup_op();
struct avm_dup_result_t
{
	any_t r1;
	any_t r2;
};
/*
dup - duplicate A
*/
avm_dup_result_t avm_dup(STACK any_t STACK_1);
void avm_dup2_op();
struct avm_dup2_result_t
{
	any_t r1;
	any_t r2;
	any_t r3;
	any_t r4;
};
/*
dup2 - duplicate A and B
*/
avm_dup2_result_t avm_dup2(STACK any_t STACK_1, STACK any_t STACK_2);
void avm_dig_op(IMMEDIATE uint8 N1);
struct avm_dig_result_t
{
	any_t r1;
	any_t r2;
};
/*
dig - Nth value from the top of the stack. dig 0 is equivalent to dup
*/
struct avm_dig_t
{
};
extern avm_dig_t avm_dig;
void avm_swap_op();
struct avm_swap_result_t
{
	any_t r1;
	any_t r2;
};
/*
swap - swaps A and B on stack
*/
avm_swap_result_t avm_swap(STACK any_t STACK_1, STACK any_t STACK_2);
void avm_select_op();
/*
select - selects one of two values based on top-of-stack: B if C != 0, else A
*/
any_t avm_select(STACK any_t STACK_1, STACK any_t STACK_2, STACK uint64_t STACK_3);
void avm_cover_op(IMMEDIATE uint8 N1);
/*
cover - remove top of stack, and place it deeper in the stack such that N elements are above it. Fails if stack depth <= N.
*/
struct avm_cover_t
{
};
extern avm_cover_t avm_cover;
void avm_uncover_op(IMMEDIATE uint8 N1);
/*
uncover - remove the value at depth N in the stack and shift above items down so the Nth deep value is on top of the stack. Fails if stack depth <= N.
*/
struct avm_uncover_t
{
};
extern avm_uncover_t avm_uncover;
void avm_concat_op();
/*
concat - join A and B

`concat` fails if the result would be greater than 4096 bytes.
*/
r_byte_t avm_concat(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2);
void avm_substring_op(IMMEDIATE uint8 S1, IMMEDIATE uint8 E2);
/*
substring - A range of bytes from A starting at S up to but not including E. If E < S, or either is larger than the array length, the program fails
*/
struct avm_substring_t
{
};
extern avm_substring_t avm_substring;
void avm_substring3_op();
/*
substring3 - A range of bytes from A starting at B up to but not including C. If C < B, or either is larger than the array length, the program fails
*/
r_byte_t avm_substring3(STACK r_byte_t STACK_1, STACK uint64_t STACK_2, STACK uint64_t STACK_3);
void avm_getbit_op();
/*
getbit - Bth bit of (byte-array or integer) A. If B is greater than or equal to the bit length of the value (8*byte length), the program fails

see explanation of bit ordering in setbit
*/
uint64_t avm_getbit(STACK any_t STACK_1, STACK uint64_t STACK_2);
void avm_setbit_op();
/*
setbit - Copy of (byte-array or integer) A, with the Bth bit set to (0 or 1) C. If B is greater than or equal to the bit length of the value (8*byte length), the program fails

When A is a uint64, index 0 is the least significant bit. Setting bit 3 to 1 on the integer 0 yields 8, or 2^3. When A is a byte array, index 0 is the leftmost bit of the leftmost byte. Setting bits 0 through 11 to 1 in a 4-byte-array of 0s yields the byte array 0xfff00000. Setting bit 3 to 1 on the 1-byte-array 0x00 yields the byte array 0x10.
*/
any_t avm_setbit(STACK any_t STACK_1, STACK uint64_t STACK_2, STACK uint64_t STACK_3);
void avm_getbyte_op();
/*
getbyte - Bth byte of A, as an integer. If B is greater than or equal to the array length, the program fails
*/
uint64_t avm_getbyte(STACK r_byte_t STACK_1, STACK uint64_t STACK_2);
void avm_setbyte_op();
/*
setbyte - Copy of A with the Bth byte set to small integer (between 0..255) C. If B is greater than or equal to the array length, the program fails
*/
r_byte_t avm_setbyte(STACK r_byte_t STACK_1, STACK uint64_t STACK_2, STACK uint64_t STACK_3);
void avm_extract_op(IMMEDIATE uint8 S1, IMMEDIATE uint8 L2);
/*
extract - A range of bytes from A starting at S up to but not including S+L. If L is 0, then extract to the end of the string. If S or S+L is larger than the array length, the program fails
*/
struct avm_extract_t
{
};
extern avm_extract_t avm_extract;
void avm_extract3_op();
/*
extract3 - A range of bytes from A starting at B up to but not including B+C. If B+C is larger than the array length, the program fails
`extract3` can be called using `extract` with no immediates.
*/
r_byte_t avm_extract3(STACK r_byte_t STACK_1, STACK uint64_t STACK_2, STACK uint64_t STACK_3);
void avm_extract_uint16_op();
/*
extract_uint16 - A uint16 formed from a range of big-endian bytes from A starting at B up to but not including B+2. If B+2 is larger than the array length, the program fails
*/
uint64_t avm_extract_uint16(STACK r_byte_t STACK_1, STACK uint64_t STACK_2);
void avm_extract_uint32_op();
/*
extract_uint32 - A uint32 formed from a range of big-endian bytes from A starting at B up to but not including B+4. If B+4 is larger than the array length, the program fails
*/
uint64_t avm_extract_uint32(STACK r_byte_t STACK_1, STACK uint64_t STACK_2);
void avm_extract_uint64_op();
/*
extract_uint64 - A uint64 formed from a range of big-endian bytes from A starting at B up to but not including B+8. If B+8 is larger than the array length, the program fails
*/
uint64_t avm_extract_uint64(STACK r_byte_t STACK_1, STACK uint64_t STACK_2);
void avm_replace2_op(IMMEDIATE uint8 S1);
/*
replace2 - Copy of A with the bytes starting at S replaced by the bytes of B. Fails if S+len(B) exceeds len(A)
`replace2` can be called using `replace` with 1 immediate.
*/
struct avm_replace2_t
{
};
extern avm_replace2_t avm_replace2;
void avm_replace3_op();
/*
replace3 - Copy of A with the bytes starting at B replaced by the bytes of C. Fails if B+len(C) exceeds len(A)
`replace3` can be called using `replace` with no immediates.
*/
r_byte_t avm_replace3(STACK r_byte_t STACK_1, STACK uint64_t STACK_2, STACK r_byte_t STACK_3);
void avm_base64_decode_op(IMMEDIATE uint8 E1);
/*
base64_decode - decode A which was base64-encoded using _encoding_ E. Fail if A is not base64 encoded with encoding E

*Warning*: Usage should be restricted to very rare use cases. In almost all cases, smart contracts should directly handle non-encoded byte-strings.	This opcode should only be used in cases where base64 is the only available option, e.g. interoperability with a third-party that only signs base64 strings.

 Decodes A using the base64 encoding E. Specify the encoding with an immediate arg either as URL and Filename Safe (`URLEncoding`) or Standard (`StdEncoding`). See [RFC 4648 sections 4 and 5](https://rfc-editor.org/rfc/rfc4648.html#section-4). It is assumed that the encoding ends with the exact number of `=` padding characters as required by the RFC. When padding occurs, any unused pad bits in the encoding must be set to zero or the decoding will fail. The special cases of `\n` and `\r` are allowed but completely ignored. An error will result when attempting to decode a string with a character that is not in the encoding alphabet or not one of `=`, `\r`, or `\n`.
*/
struct avm_base64_decode_t
{
	any_t URLEncoding(STACK r_byte_t STACK_1);
	any_t StdEncoding(STACK r_byte_t STACK_1);
};
extern avm_base64_decode_t avm_base64_decode;
void avm_json_ref_op(IMMEDIATE uint8 R1);
/*
json_ref - key B's value, of type R, from a [valid](jsonspec.md) utf-8 encoded json object A

*Warning*: Usage should be restricted to very rare use cases, as JSON decoding is expensive and quite limited. In addition, JSON objects are large and not optimized for size.

Almost all smart contracts should use simpler and smaller methods (such as the [ABI](https://arc.algorand.foundation/ARCs/arc-0004). This opcode should only be used in cases where JSON is only available option, e.g. when a third-party only signs JSON.
*/
struct avm_json_ref_t
{
	r_byte_t JSONString(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2);
	uint64_t JSONUint64(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2);
	r_byte_t JSONObject(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2);
};
extern avm_json_ref_t avm_json_ref;
void avm_balance_op();
/*
balance - balance for account A, in microalgos. The balance is observed after the effects of previous transactions in the group, and after the fee for the current transaction is deducted. Changes caused by inner transactions are observable immediately following `itxn_submit`

params: Txn.Accounts offset (or, since v4, an _available_ account address). Return: value.
*/
uint64_t avm_balance(STACK any_t STACK_1);
void avm_app_opted_in_op();
/*
app_opted_in - 1 if account A is opted in to application B, else 0

params: Txn.Accounts offset (or, since v4, an _available_ account address), _available_ application id (or, since v4, a Txn.ForeignApps offset). Return: 1 if opted in and 0 otherwise.
*/
bool_t avm_app_opted_in(STACK any_t STACK_1, STACK uint64_t STACK_2);
void avm_app_local_get_op();
/*
app_local_get - local state of the key B in the current application in account A

params: Txn.Accounts offset (or, since v4, an _available_ account address), state key. Return: value. The value is zero (of type uint64) if the key does not exist.
*/
any_t avm_app_local_get(STACK any_t STACK_1, STACK key_t STACK_2);
void avm_app_local_get_ex_op();
struct avm_app_local_get_ex_result_t
{
	any_t r1;
	bool_t r2;
};
/*
app_local_get_ex - X is the local state of application B, key C in account A. Y is 1 if key existed, else 0

params: Txn.Accounts offset (or, since v4, an _available_ account address), _available_ application id (or, since v4, a Txn.ForeignApps offset), state key. Return: did_exist flag (top of the stack, 1 if the application and key existed and 0 otherwise), value. The value is zero (of type uint64) if the key does not exist.
*/
avm_app_local_get_ex_result_t avm_app_local_get_ex(STACK any_t STACK_1, STACK uint64_t STACK_2, STACK key_t STACK_3);
void avm_app_global_get_op();
/*
app_global_get - global state of the key A in the current application

params: state key. Return: value. The value is zero (of type uint64) if the key does not exist.
*/
any_t avm_app_global_get(STACK key_t STACK_1);
void avm_app_global_get_ex_op();
struct avm_app_global_get_ex_result_t
{
	any_t r1;
	bool_t r2;
};
/*
app_global_get_ex - X is the global state of application A, key B. Y is 1 if key existed, else 0

params: Txn.ForeignApps offset (or, since v4, an _available_ application id), state key. Return: did_exist flag (top of the stack, 1 if the application and key existed and 0 otherwise), value. The value is zero (of type uint64) if the key does not exist.
*/
avm_app_global_get_ex_result_t avm_app_global_get_ex(STACK uint64_t STACK_1, STACK key_t STACK_2);
void avm_app_local_put_op();
/*
app_local_put - write C to key B in account A's local state of the current application

params: Txn.Accounts offset (or, since v4, an _available_ account address), state key, value.
*/
void avm_app_local_put(STACK any_t STACK_1, STACK key_t STACK_2, STACK any_t STACK_3);
void avm_app_global_put_op();
/*
app_global_put - write B to key A in the global state of the current application
*/
void avm_app_global_put(STACK key_t STACK_1, STACK any_t STACK_2);
void avm_app_local_del_op();
/*
app_local_del - delete key B from account A's local state of the current application

params: Txn.Accounts offset (or, since v4, an _available_ account address), state key.

Deleting a key which is already absent has no effect on the application local state. (In particular, it does _not_ cause the program to fail.)
*/
void avm_app_local_del(STACK any_t STACK_1, STACK key_t STACK_2);
void avm_app_global_del_op();
/*
app_global_del - delete key A from the global state of the current application

params: state key.

Deleting a key which is already absent has no effect on the application global state. (In particular, it does _not_ cause the program to fail.)
*/
void avm_app_global_del(STACK key_t STACK_1);
void avm_asset_holding_get_op(IMMEDIATE uint8 F1);
struct avm_asset_holding_get_result_t
{
	any_t r1;
	bool_t r2;
};
/*
asset_holding_get - X is field F from account A's holding of asset B. Y is 1 if A is opted into B, else 0

params: Txn.Accounts offset (or, since v4, an _available_ address), asset id (or, since v4, a Txn.ForeignAssets offset). Return: did_exist flag (1 if the asset existed and 0 otherwise), value.
*/
struct avm_asset_holding_get_t
{
	uint64_t AssetBalance(STACK any_t STACK_1, STACK uint64_t STACK_2);
	bool_t AssetFrozen(STACK any_t STACK_1, STACK uint64_t STACK_2);
};
extern avm_asset_holding_get_t avm_asset_holding_get;
void avm_asset_params_get_op(IMMEDIATE uint8 F1);
struct avm_asset_params_get_result_t
{
	any_t r1;
	bool_t r2;
};
/*
asset_params_get - X is field F from asset A. Y is 1 if A exists, else 0

params: Txn.ForeignAssets offset (or, since v4, an _available_ asset id. Return: did_exist flag (1 if the asset existed and 0 otherwise), value.
*/
struct avm_asset_params_get_t
{
	uint64_t AssetTotal(STACK uint64_t STACK_1);
	uint64_t AssetDecimals(STACK uint64_t STACK_1);
	bool_t AssetDefaultFrozen(STACK uint64_t STACK_1);
	r_byte_t AssetUnitName(STACK uint64_t STACK_1);
	r_byte_t AssetName(STACK uint64_t STACK_1);
	r_byte_t AssetURL(STACK uint64_t STACK_1);
	r32_byte_t AssetMetadataHash(STACK uint64_t STACK_1);
	addr_t AssetManager(STACK uint64_t STACK_1);
	addr_t AssetReserve(STACK uint64_t STACK_1);
	addr_t AssetFreeze(STACK uint64_t STACK_1);
	addr_t AssetClawback(STACK uint64_t STACK_1);
	addr_t AssetCreator(STACK uint64_t STACK_1);
};
extern avm_asset_params_get_t avm_asset_params_get;
void avm_app_params_get_op(IMMEDIATE uint8 F1);
struct avm_app_params_get_result_t
{
	any_t r1;
	bool_t r2;
};
/*
app_params_get - X is field F from app A. Y is 1 if A exists, else 0

params: Txn.ForeignApps offset or an _available_ app id. Return: did_exist flag (1 if the application existed and 0 otherwise), value.
*/
struct avm_app_params_get_t
{
	r_byte_t AppApprovalProgram(STACK uint64_t STACK_1);
	r_byte_t AppClearStateProgram(STACK uint64_t STACK_1);
	uint64_t AppGlobalNumUint(STACK uint64_t STACK_1);
	uint64_t AppGlobalNumByteSlice(STACK uint64_t STACK_1);
	uint64_t AppLocalNumUint(STACK uint64_t STACK_1);
	uint64_t AppLocalNumByteSlice(STACK uint64_t STACK_1);
	uint64_t AppExtraProgramPages(STACK uint64_t STACK_1);
	addr_t AppCreator(STACK uint64_t STACK_1);
	addr_t AppAddress(STACK uint64_t STACK_1);
};
extern avm_app_params_get_t avm_app_params_get;
void avm_acct_params_get_op(IMMEDIATE uint8 F1);
struct avm_acct_params_get_result_t
{
	any_t r1;
	bool_t r2;
};
/*
acct_params_get - X is field F from account A. Y is 1 if A owns positive algos, else 0
*/
struct avm_acct_params_get_t
{
	uint64_t AcctBalance(STACK any_t STACK_1);
	uint64_t AcctMinBalance(STACK any_t STACK_1);
	addr_t AcctAuthAddr(STACK any_t STACK_1);
	uint64_t AcctTotalNumUint(STACK any_t STACK_1);
	uint64_t AcctTotalNumByteSlice(STACK any_t STACK_1);
	uint64_t AcctTotalExtraAppPages(STACK any_t STACK_1);
	uint64_t AcctTotalAppsCreated(STACK any_t STACK_1);
	uint64_t AcctTotalAppsOptedIn(STACK any_t STACK_1);
	uint64_t AcctTotalAssetsCreated(STACK any_t STACK_1);
	uint64_t AcctTotalAssets(STACK any_t STACK_1);
	uint64_t AcctTotalBoxes(STACK any_t STACK_1);
	uint64_t AcctTotalBoxBytes(STACK any_t STACK_1);
};
extern avm_acct_params_get_t avm_acct_params_get;
void avm_min_balance_op();
/*
min_balance - minimum required balance for account A, in microalgos. Required balance is affected by ASA, App, and Box usage. When creating or opting into an app, the minimum balance grows before the app code runs, therefore the increase is visible there. When deleting or closing out, the minimum balance decreases after the app executes. Changes caused by inner transactions or box usage are observable immediately following the opcode effecting the change.

params: Txn.Accounts offset (or, since v4, an _available_ account address). Return: value.
*/
uint64_t avm_min_balance(STACK any_t STACK_1);
void avm_pushbytes_op(IMMEDIATE bytes BYTES1);
/*
pushbytes - immediate BYTES

pushbytes args are not added to the bytecblock during assembly processes
*/
struct avm_pushbytes_t
{
};
extern avm_pushbytes_t avm_pushbytes;
void avm_pushint_op(IMMEDIATE bytes UINT1);
/*
pushint - immediate UINT

pushint args are not added to the intcblock during assembly processes
*/
struct avm_pushint_t
{
};
extern avm_pushint_t avm_pushint;
void avm_pushbytess_op(IMMEDIATE bytes BYTES1, ...);
/*
pushbytess - push sequences of immediate byte arrays to stack (first byte array being deepest)

pushbytess args are not added to the bytecblock during assembly processes
*/
struct avm_pushbytess_t
{
};
extern avm_pushbytess_t avm_pushbytess;
void avm_pushints_op(IMMEDIATE uint64 UINT1, ...);
/*
pushints - push sequence of immediate uints to stack in the order they appear (first uint being deepest)

pushints args are not added to the intcblock during assembly processes
*/
struct avm_pushints_t
{
};
extern avm_pushints_t avm_pushints;
void avm_ed25519verify_bare_op();
/*
ed25519verify_bare - for (data A, signature B, pubkey C) verify the signature of the data against the pubkey => {0 or 1}
*/
bool_t avm_ed25519verify_bare(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2, STACK r_byte_t STACK_3);
void avm_callsub_op(IMMEDIATE label TARGET1);
/*
callsub - branch unconditionally to TARGET, saving the next instruction on the call stack

The call stack is separate from the data stack. Only `callsub`, `retsub`, and `proto` manipulate it.
*/
struct avm_callsub_t
{
};
extern avm_callsub_t avm_callsub;
void avm_retsub_op();
/*
retsub - pop the top instruction from the call stack and branch to it

If the current frame was prepared by `proto A R`, `retsub` will remove the 'A' arguments from the stack, move the `R` return values down, and pop any stack locations above the relocated return values.
*/
void avm_retsub();
void avm_proto_op(IMMEDIATE uint8 A1, IMMEDIATE uint8 R2);
/*
proto - Prepare top call frame for a retsub that will assume A args and R return values.

Fails unless the last instruction executed was a `callsub`.
*/
struct avm_proto_t
{
};
extern avm_proto_t avm_proto;
void avm_frame_dig_op(IMMEDIATE int8 I1);
/*
frame_dig - Nth (signed) value from the frame pointer.
*/
struct avm_frame_dig_t
{
};
extern avm_frame_dig_t avm_frame_dig;
void avm_frame_bury_op(IMMEDIATE int8 I1);
/*
frame_bury - replace the Nth (signed) value from the frame pointer in the stack with A
*/
struct avm_frame_bury_t
{
};
extern avm_frame_bury_t avm_frame_bury;
void avm_switch_op(IMMEDIATE label TARGET1, ...);
/*
switch - branch to the Ath label. Continue at following instruction if index A exceeds the number of labels.
*/
struct avm_switch_t
{
};
extern avm_switch_t avm_switch;
void avm_match_op(IMMEDIATE label TARGET1, ...);
/*
match - given match cases from A[1] to A[N], branch to the Ith label where A[I] = B. Continue to the following instruction if no matches are found.

`match` consumes N+1 values from the stack. Let the top stack value be B. The following N values represent an ordered list of match cases/constants (A), where the first value (A[0]) is the deepest in the stack. The immediate arguments are an ordered list of N labels (T). `match` will branch to target T[I], where A[I] = B. If there are no matches then execution continues on to the next instruction.
*/
struct avm_match_t
{
};
extern avm_match_t avm_match;
void avm_shl_op();
/*
shl - A times 2^B, modulo 2^64
*/
uint64_t avm_shl(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_shr_op();
/*
shr - A divided by 2^B
*/
uint64_t avm_shr(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_sqrt_op();
/*
sqrt - The largest integer I such that I^2 <= A
*/
uint64_t avm_sqrt(STACK uint64_t STACK_1);
void avm_bitlen_op();
/*
bitlen - The highest set bit in A. If A is a byte-array, it is interpreted as a big-endian unsigned integer. bitlen of 0 is 0, bitlen of 8 is 4

bitlen interprets arrays as big-endian integers, unlike setbit/getbit
*/
uint64_t avm_bitlen(STACK any_t STACK_1);
void avm_exp_op();
/*
exp - A raised to the Bth power. Fail if A == B == 0 and on overflow
*/
uint64_t avm_exp(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_expw_op();
struct avm_expw_result_t
{
	uint64_t r1;
	uint64_t r2;
};
/*
expw - A raised to the Bth power as a 128-bit result in two uint64s. X is the high 64 bits, Y is the low. Fail if A == B == 0 or if the results exceeds 2^128-1
*/
avm_expw_result_t avm_expw(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_bsqrt_op();
/*
bsqrt - The largest integer I such that I^2 <= A. A and I are interpreted as big-endian unsigned integers
*/
r_byte_t avm_bsqrt(STACK r_byte_t STACK_1);
void avm_divw_op();
/*
divw - A,B / C. Fail if C == 0 or if result overflows.

The notation A,B indicates that A and B are interpreted as a uint128 value, with A as the high uint64 and B the low.
*/
uint64_t avm_divw(STACK uint64_t STACK_1, STACK uint64_t STACK_2, STACK uint64_t STACK_3);
void avm_sha3_256_op();
/*
sha3_256 - SHA3_256 hash of value A, yields [32]byte
*/
r_byte_t avm_sha3_256(STACK r_byte_t STACK_1);
void avm_bplus_op();
/*
b+ - A plus B. A and B are interpreted as big-endian unsigned integers
*/
r_byte_t avm_bplus(STACK bigint_t STACK_1, STACK bigint_t STACK_2);
void avm_bminus_op();
/*
b- - A minus B. A and B are interpreted as big-endian unsigned integers. Fail on underflow.
*/
bigint_t avm_bminus(STACK bigint_t STACK_1, STACK bigint_t STACK_2);
void avm_bdiv_op();
/*
b/ - A divided by B (truncated division). A and B are interpreted as big-endian unsigned integers. Fail if B is zero.
*/
bigint_t avm_bdiv(STACK bigint_t STACK_1, STACK bigint_t STACK_2);
void avm_bmul_op();
/*
b* - A times B. A and B are interpreted as big-endian unsigned integers.
*/
r_byte_t avm_bmul(STACK bigint_t STACK_1, STACK bigint_t STACK_2);
void avm_blt_op();
/*
b< - 1 if A is less than B, else 0. A and B are interpreted as big-endian unsigned integers
*/
bool_t avm_blt(STACK bigint_t STACK_1, STACK bigint_t STACK_2);
void avm_bgt_op();
/*
b> - 1 if A is greater than B, else 0. A and B are interpreted as big-endian unsigned integers
*/
bool_t avm_bgt(STACK bigint_t STACK_1, STACK bigint_t STACK_2);
void avm_blteq_op();
/*
b<= - 1 if A is less than or equal to B, else 0. A and B are interpreted as big-endian unsigned integers
*/
bool_t avm_blteq(STACK bigint_t STACK_1, STACK bigint_t STACK_2);
void avm_bgteq_op();
/*
b>= - 1 if A is greater than or equal to B, else 0. A and B are interpreted as big-endian unsigned integers
*/
bool_t avm_bgteq(STACK bigint_t STACK_1, STACK bigint_t STACK_2);
void avm_beqeq_op();
/*
b== - 1 if A is equal to B, else 0. A and B are interpreted as big-endian unsigned integers
*/
bool_t avm_beqeq(STACK bigint_t STACK_1, STACK bigint_t STACK_2);
void avm_bnoteq_op();
/*
b!= - 0 if A is equal to B, else 1. A and B are interpreted as big-endian unsigned integers
*/
bool_t avm_bnoteq(STACK bigint_t STACK_1, STACK bigint_t STACK_2);
void avm_bmod_op();
/*
b% - A modulo B. A and B are interpreted as big-endian unsigned integers. Fail if B is zero.
*/
r_byte_t avm_bmod(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2);
void avm_bor_op();
/*
b| - A bitwise-or B. A and B are zero-left extended to the greater of their lengths
*/
r_byte_t avm_bor(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2);
void avm_band_op();
/*
b& - A bitwise-and B. A and B are zero-left extended to the greater of their lengths
*/
r_byte_t avm_band(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2);
void avm_bxor_op();
/*
b^ - A bitwise-xor B. A and B are zero-left extended to the greater of their lengths
*/
r_byte_t avm_bxor(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2);
void avm_binv_op();
/*
b~ - A with all bits inverted
*/
r_byte_t avm_binv(STACK r_byte_t STACK_1);
void avm_bzero_op();
/*
bzero - zero filled byte-array of length A
*/
r_byte_t avm_bzero(STACK uint64_t STACK_1);
void avm_log_op();
/*
log - write A to log state of the current application

`log` fails if called more than MaxLogCalls times in a program, or if the sum of logged bytes exceeds 1024 bytes.
*/
void avm_log(STACK r_byte_t STACK_1);
void avm_itxn_begin_op();
/*
itxn_begin - begin preparation of a new inner transaction in a new transaction group

`itxn_begin` initializes Sender to the application address; Fee to the minimum allowable, taking into account MinTxnFee and credit from overpaying in earlier transactions; FirstValid/LastValid to the values in the invoking transaction, and all other fields to zero or empty values.
*/
void avm_itxn_begin();
void avm_itxn_field_op(IMMEDIATE uint8 F1);
/*
itxn_field - set field F of the current inner transaction to A

`itxn_field` fails if A is of the wrong type for F, including a byte array of the wrong size for use as an address when F is an address field. `itxn_field` also fails if A is an account, asset, or app that is not _available_, or an attempt is made extend an array field beyond the limit imposed by consensus parameters. (Addresses set into asset params of acfg transactions need not be _available_.)
*/
struct avm_itxn_field_t
{
	void Sender(STACK any_t STACK_1);
	void Fee(STACK any_t STACK_1);
	void Note(STACK any_t STACK_1);
	void Receiver(STACK any_t STACK_1);
	void Amount(STACK any_t STACK_1);
	void CloseRemainderTo(STACK any_t STACK_1);
	void VotePK(STACK any_t STACK_1);
	void SelectionPK(STACK any_t STACK_1);
	void VoteFirst(STACK any_t STACK_1);
	void VoteLast(STACK any_t STACK_1);
	void VoteKeyDilution(STACK any_t STACK_1);
	void Type(STACK any_t STACK_1);
	void TypeEnum(STACK any_t STACK_1);
	void XferAsset(STACK any_t STACK_1);
	void AssetAmount(STACK any_t STACK_1);
	void AssetSender(STACK any_t STACK_1);
	void AssetReceiver(STACK any_t STACK_1);
	void AssetCloseTo(STACK any_t STACK_1);
	void ApplicationID(STACK any_t STACK_1);
	void OnCompletion(STACK any_t STACK_1);
	void ApplicationArgs(STACK any_t STACK_1);
	void Accounts(STACK any_t STACK_1);
	void ApprovalProgram(STACK any_t STACK_1);
	void ClearStateProgram(STACK any_t STACK_1);
	void RekeyTo(STACK any_t STACK_1);
	void ConfigAsset(STACK any_t STACK_1);
	void ConfigAssetTotal(STACK any_t STACK_1);
	void ConfigAssetDecimals(STACK any_t STACK_1);
	void ConfigAssetDefaultFrozen(STACK any_t STACK_1);
	void ConfigAssetUnitName(STACK any_t STACK_1);
	void ConfigAssetName(STACK any_t STACK_1);
	void ConfigAssetURL(STACK any_t STACK_1);
	void ConfigAssetMetadataHash(STACK any_t STACK_1);
	void ConfigAssetManager(STACK any_t STACK_1);
	void ConfigAssetReserve(STACK any_t STACK_1);
	void ConfigAssetFreeze(STACK any_t STACK_1);
	void ConfigAssetClawback(STACK any_t STACK_1);
	void FreezeAsset(STACK any_t STACK_1);
	void FreezeAssetAccount(STACK any_t STACK_1);
	void FreezeAssetFrozen(STACK any_t STACK_1);
	void Assets(STACK any_t STACK_1);
	void Applications(STACK any_t STACK_1);
	void GlobalNumUint(STACK any_t STACK_1);
	void GlobalNumByteSlice(STACK any_t STACK_1);
	void LocalNumUint(STACK any_t STACK_1);
	void LocalNumByteSlice(STACK any_t STACK_1);
	void ExtraProgramPages(STACK any_t STACK_1);
	void Nonparticipation(STACK any_t STACK_1);
	void StateProofPK(STACK any_t STACK_1);
	void ApprovalProgramPages(STACK any_t STACK_1);
	void ClearStateProgramPages(STACK any_t STACK_1);
};
extern avm_itxn_field_t avm_itxn_field;
void avm_itxn_submit_op();
/*
itxn_submit - execute the current inner transaction group. Fail if executing this group would exceed the inner transaction limit, or if any transaction in the group fails.

`itxn_submit` resets the current transaction so that it can not be resubmitted. A new `itxn_begin` is required to prepare another inner transaction.
*/
void avm_itxn_submit();
void avm_itxn_op(IMMEDIATE uint8 F1);
/*
itxn - field F of the last inner transaction
*/
struct avm_itxn_t
{
	const addr_t Sender;
	const uint64_t Fee;
	const uint64_t FirstValid;
	const uint64_t FirstValidTime;
	const uint64_t LastValid;
	const r_byte_t Note;
	const r32_byte_t Lease;
	const addr_t Receiver;
	const uint64_t Amount;
	const addr_t CloseRemainderTo;
	const r32_byte_t VotePK;
	const r32_byte_t SelectionPK;
	const uint64_t VoteFirst;
	const uint64_t VoteLast;
	const uint64_t VoteKeyDilution;
	const r_byte_t Type;
	const uint64_t TypeEnum;
	const uint64_t XferAsset;
	const uint64_t AssetAmount;
	const addr_t AssetSender;
	const addr_t AssetReceiver;
	const addr_t AssetCloseTo;
	const uint64_t GroupIndex;
	const r32_byte_t TxID;
	const uint64_t ApplicationID;
	const uint64_t OnCompletion;
	const r_byte_t ApplicationArgs;
	const uint64_t NumAppArgs;
	const addr_t Accounts;
	const uint64_t NumAccounts;
	const r_byte_t ApprovalProgram;
	const r_byte_t ClearStateProgram;
	const addr_t RekeyTo;
	const uint64_t ConfigAsset;
	const uint64_t ConfigAssetTotal;
	const uint64_t ConfigAssetDecimals;
	const bool_t ConfigAssetDefaultFrozen;
	const r_byte_t ConfigAssetUnitName;
	const r_byte_t ConfigAssetName;
	const r_byte_t ConfigAssetURL;
	const r32_byte_t ConfigAssetMetadataHash;
	const addr_t ConfigAssetManager;
	const addr_t ConfigAssetReserve;
	const addr_t ConfigAssetFreeze;
	const addr_t ConfigAssetClawback;
	const uint64_t FreezeAsset;
	const addr_t FreezeAssetAccount;
	const bool_t FreezeAssetFrozen;
	const uint64_t Assets;
	const uint64_t NumAssets;
	const uint64_t Applications;
	const uint64_t NumApplications;
	const uint64_t GlobalNumUint;
	const uint64_t GlobalNumByteSlice;
	const uint64_t LocalNumUint;
	const uint64_t LocalNumByteSlice;
	const uint64_t ExtraProgramPages;
	const bool_t Nonparticipation;
	const r_byte_t Logs;
	const uint64_t NumLogs;
	const uint64_t CreatedAssetID;
	const uint64_t CreatedApplicationID;
	const r_byte_t LastLog;
	const r_byte_t StateProofPK;
	const r_byte_t ApprovalProgramPages;
	const uint64_t NumApprovalProgramPages;
	const r_byte_t ClearStateProgramPages;
	const uint64_t NumClearStateProgramPages;
};
extern avm_itxn_t avm_itxn;
void avm_itxna_op(IMMEDIATE uint8 F1, IMMEDIATE uint8 I2);
/*
itxna - Ith value of the array field F of the last inner transaction
*/
struct avm_itxna_t
{
	r_byte_t ApplicationArgs(IMMEDIATE uint8 I2);
	addr_t Accounts(IMMEDIATE uint8 I2);
	uint64_t Assets(IMMEDIATE uint8 I2);
	uint64_t Applications(IMMEDIATE uint8 I2);
	r_byte_t Logs(IMMEDIATE uint8 I2);
	r_byte_t ApprovalProgramPages(IMMEDIATE uint8 I2);
	r_byte_t ClearStateProgramPages(IMMEDIATE uint8 I2);
};
extern avm_itxna_t avm_itxna;
void avm_itxn_next_op();
/*
itxn_next - begin preparation of a new inner transaction in the same transaction group

`itxn_next` initializes the transaction exactly as `itxn_begin` does
*/
void avm_itxn_next();
void avm_gitxn_op(IMMEDIATE uint8 T1, IMMEDIATE uint8 F2);
/*
gitxn - field F of the Tth transaction in the last inner group submitted
*/
struct avm_gitxn_t
{
	addr_t Sender(IMMEDIATE uint8 F2);
	uint64_t Fee(IMMEDIATE uint8 F2);
	uint64_t FirstValid(IMMEDIATE uint8 F2);
	uint64_t FirstValidTime(IMMEDIATE uint8 F2);
	uint64_t LastValid(IMMEDIATE uint8 F2);
	r_byte_t Note(IMMEDIATE uint8 F2);
	r32_byte_t Lease(IMMEDIATE uint8 F2);
	addr_t Receiver(IMMEDIATE uint8 F2);
	uint64_t Amount(IMMEDIATE uint8 F2);
	addr_t CloseRemainderTo(IMMEDIATE uint8 F2);
	r32_byte_t VotePK(IMMEDIATE uint8 F2);
	r32_byte_t SelectionPK(IMMEDIATE uint8 F2);
	uint64_t VoteFirst(IMMEDIATE uint8 F2);
	uint64_t VoteLast(IMMEDIATE uint8 F2);
	uint64_t VoteKeyDilution(IMMEDIATE uint8 F2);
	r_byte_t Type(IMMEDIATE uint8 F2);
	uint64_t TypeEnum(IMMEDIATE uint8 F2);
	uint64_t XferAsset(IMMEDIATE uint8 F2);
	uint64_t AssetAmount(IMMEDIATE uint8 F2);
	addr_t AssetSender(IMMEDIATE uint8 F2);
	addr_t AssetReceiver(IMMEDIATE uint8 F2);
	addr_t AssetCloseTo(IMMEDIATE uint8 F2);
	uint64_t GroupIndex(IMMEDIATE uint8 F2);
	r32_byte_t TxID(IMMEDIATE uint8 F2);
	uint64_t ApplicationID(IMMEDIATE uint8 F2);
	uint64_t OnCompletion(IMMEDIATE uint8 F2);
	r_byte_t ApplicationArgs(IMMEDIATE uint8 F2);
	uint64_t NumAppArgs(IMMEDIATE uint8 F2);
	addr_t Accounts(IMMEDIATE uint8 F2);
	uint64_t NumAccounts(IMMEDIATE uint8 F2);
	r_byte_t ApprovalProgram(IMMEDIATE uint8 F2);
	r_byte_t ClearStateProgram(IMMEDIATE uint8 F2);
	addr_t RekeyTo(IMMEDIATE uint8 F2);
	uint64_t ConfigAsset(IMMEDIATE uint8 F2);
	uint64_t ConfigAssetTotal(IMMEDIATE uint8 F2);
	uint64_t ConfigAssetDecimals(IMMEDIATE uint8 F2);
	bool_t ConfigAssetDefaultFrozen(IMMEDIATE uint8 F2);
	r_byte_t ConfigAssetUnitName(IMMEDIATE uint8 F2);
	r_byte_t ConfigAssetName(IMMEDIATE uint8 F2);
	r_byte_t ConfigAssetURL(IMMEDIATE uint8 F2);
	r32_byte_t ConfigAssetMetadataHash(IMMEDIATE uint8 F2);
	addr_t ConfigAssetManager(IMMEDIATE uint8 F2);
	addr_t ConfigAssetReserve(IMMEDIATE uint8 F2);
	addr_t ConfigAssetFreeze(IMMEDIATE uint8 F2);
	addr_t ConfigAssetClawback(IMMEDIATE uint8 F2);
	uint64_t FreezeAsset(IMMEDIATE uint8 F2);
	addr_t FreezeAssetAccount(IMMEDIATE uint8 F2);
	bool_t FreezeAssetFrozen(IMMEDIATE uint8 F2);
	uint64_t Assets(IMMEDIATE uint8 F2);
	uint64_t NumAssets(IMMEDIATE uint8 F2);
	uint64_t Applications(IMMEDIATE uint8 F2);
	uint64_t NumApplications(IMMEDIATE uint8 F2);
	uint64_t GlobalNumUint(IMMEDIATE uint8 F2);
	uint64_t GlobalNumByteSlice(IMMEDIATE uint8 F2);
	uint64_t LocalNumUint(IMMEDIATE uint8 F2);
	uint64_t LocalNumByteSlice(IMMEDIATE uint8 F2);
	uint64_t ExtraProgramPages(IMMEDIATE uint8 F2);
	bool_t Nonparticipation(IMMEDIATE uint8 F2);
	r_byte_t Logs(IMMEDIATE uint8 F2);
	uint64_t NumLogs(IMMEDIATE uint8 F2);
	uint64_t CreatedAssetID(IMMEDIATE uint8 F2);
	uint64_t CreatedApplicationID(IMMEDIATE uint8 F2);
	r_byte_t LastLog(IMMEDIATE uint8 F2);
	r_byte_t StateProofPK(IMMEDIATE uint8 F2);
	r_byte_t ApprovalProgramPages(IMMEDIATE uint8 F2);
	uint64_t NumApprovalProgramPages(IMMEDIATE uint8 F2);
	r_byte_t ClearStateProgramPages(IMMEDIATE uint8 F2);
	uint64_t NumClearStateProgramPages(IMMEDIATE uint8 F2);
};
extern avm_gitxn_t avm_gitxn;
void avm_gitxna_op(IMMEDIATE uint8 T1, IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
/*
gitxna - Ith value of the array field F from the Tth transaction in the last inner group submitted
*/
struct avm_gitxna_t
{
	r_byte_t ApplicationArgs(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
	addr_t Accounts(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
	uint64_t Assets(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
	uint64_t Applications(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
	r_byte_t Logs(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
	r_byte_t ApprovalProgramPages(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
	r_byte_t ClearStateProgramPages(IMMEDIATE uint8 F2, IMMEDIATE uint8 I3);
};
extern avm_gitxna_t avm_gitxna;
void avm_box_create_op();
/*
box_create - create a box named A, of length B. Fail if A is empty or B exceeds 32,768. Returns 0 if A already existed, else 1

Newly created boxes are filled with 0 bytes. `box_create` will fail if the referenced box already exists with a different size. Otherwise, existing boxes are unchanged by `box_create`.
*/
bool_t avm_box_create(STACK name_t STACK_1, STACK uint64_t STACK_2);
void avm_box_extract_op();
/*
box_extract - read C bytes from box A, starting at offset B. Fail if A does not exist, or the byte range is outside A's size.
*/
r_byte_t avm_box_extract(STACK name_t STACK_1, STACK uint64_t STACK_2, STACK uint64_t STACK_3);
void avm_box_replace_op();
/*
box_replace - write byte-array C into box A, starting at offset B. Fail if A does not exist, or the byte range is outside A's size.
*/
void avm_box_replace(STACK name_t STACK_1, STACK uint64_t STACK_2, STACK r_byte_t STACK_3);
void avm_box_del_op();
/*
box_del - delete box named A if it exists. Return 1 if A existed, 0 otherwise
*/
bool_t avm_box_del(STACK name_t STACK_1);
void avm_box_len_op();
struct avm_box_len_result_t
{
	uint64_t r1;
	bool_t r2;
};
/*
box_len - X is the length of box A if A exists, else 0. Y is 1 if A exists, else 0.
*/
avm_box_len_result_t avm_box_len(STACK name_t STACK_1);
void avm_box_get_op();
struct avm_box_get_result_t
{
	r_byte_t r1;
	bool_t r2;
};
/*
box_get - X is the contents of box A if A exists, else ''. Y is 1 if A exists, else 0.

For boxes that exceed 4,096 bytes, consider `box_create`, `box_extract`, and `box_replace`
*/
avm_box_get_result_t avm_box_get(STACK name_t STACK_1);
void avm_box_put_op();
/*
box_put - replaces the contents of box A with byte-array B. Fails if A exists and len(B) != len(box A). Creates A if it does not exist

For boxes that exceed 4,096 bytes, consider `box_create`, `box_extract`, and `box_replace`
*/
void avm_box_put(STACK name_t STACK_1, STACK r_byte_t STACK_2);
void avm_txnas_op(IMMEDIATE uint8 F1);
/*
txnas - Ath value of the array field F of the current transaction
*/
struct avm_txnas_t
{
	r_byte_t ApplicationArgs(STACK uint64_t STACK_1);
	addr_t Accounts(STACK uint64_t STACK_1);
	uint64_t Assets(STACK uint64_t STACK_1);
	uint64_t Applications(STACK uint64_t STACK_1);
	r_byte_t Logs(STACK uint64_t STACK_1);
	r_byte_t ApprovalProgramPages(STACK uint64_t STACK_1);
	r_byte_t ClearStateProgramPages(STACK uint64_t STACK_1);
};
extern avm_txnas_t avm_txnas;
void avm_gtxnas_op(IMMEDIATE uint8 T1, IMMEDIATE uint8 F2);
/*
gtxnas - Ath value of the array field F from the Tth transaction in the current group
*/
struct avm_gtxnas_t
{
	r_byte_t ApplicationArgs(STACK uint64_t STACK_1, IMMEDIATE uint8 F2);
	addr_t Accounts(STACK uint64_t STACK_1, IMMEDIATE uint8 F2);
	uint64_t Assets(STACK uint64_t STACK_1, IMMEDIATE uint8 F2);
	uint64_t Applications(STACK uint64_t STACK_1, IMMEDIATE uint8 F2);
	r_byte_t Logs(STACK uint64_t STACK_1, IMMEDIATE uint8 F2);
	r_byte_t ApprovalProgramPages(STACK uint64_t STACK_1, IMMEDIATE uint8 F2);
	r_byte_t ClearStateProgramPages(STACK uint64_t STACK_1, IMMEDIATE uint8 F2);
};
extern avm_gtxnas_t avm_gtxnas;
void avm_gtxnsas_op(IMMEDIATE uint8 F1);
/*
gtxnsas - Bth value of the array field F from the Ath transaction in the current group
*/
struct avm_gtxnsas_t
{
	r_byte_t ApplicationArgs(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
	addr_t Accounts(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
	uint64_t Assets(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
	uint64_t Applications(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
	r_byte_t Logs(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
	r_byte_t ApprovalProgramPages(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
	r_byte_t ClearStateProgramPages(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
};
extern avm_gtxnsas_t avm_gtxnsas;
void avm_args_op();
/*
args - Ath LogicSig argument
*/
r_byte_t avm_args(STACK uint64_t STACK_1);
void avm_gloadss_op();
/*
gloadss - Bth scratch space value of the Ath transaction in the current group
*/
any_t avm_gloadss(STACK uint64_t STACK_1, STACK uint64_t STACK_2);
void avm_itxnas_op(IMMEDIATE uint8 F1);
/*
itxnas - Ath value of the array field F of the last inner transaction
*/
struct avm_itxnas_t
{
};
extern avm_itxnas_t avm_itxnas;
void avm_gitxnas_op(IMMEDIATE uint8 T1, IMMEDIATE uint8 F2);
/*
gitxnas - Ath value of the array field F from the Tth transaction in the last inner group submitted
*/
struct avm_gitxnas_t
{
};
extern avm_gitxnas_t avm_gitxnas;
void avm_vrf_verify_op(IMMEDIATE uint8 S1);
struct avm_vrf_verify_result_t
{
	r_byte_t r1;
	bool_t r2;
};
/*
vrf_verify - Verify the proof B of message A against pubkey C. Returns vrf output and verification flag.

`VrfAlgorand` is the VRF used in Algorand. It is ECVRF-ED25519-SHA512-Elligator2, specified in the IETF internet draft [draft-irtf-cfrg-vrf-03](https://datatracker.ietf.org/doc/draft-irtf-cfrg-vrf/03/).
*/
struct avm_vrf_verify_t
{
	void VrfAlgorand(STACK r_byte_t STACK_1, STACK r_byte_t STACK_2, STACK r_byte_t STACK_3);
};
extern avm_vrf_verify_t avm_vrf_verify;
void avm_block_op(IMMEDIATE uint8 F1);
/*
block - field F of block A. Fail unless A falls between txn.LastValid-1002 and txn.FirstValid (exclusive)
*/
struct avm_block_t
{
	r_byte_t BlkSeed(STACK uint64_t STACK_1);
	uint64_t BlkTimestamp(STACK uint64_t STACK_1);
};
extern avm_block_t avm_block;
#endif
