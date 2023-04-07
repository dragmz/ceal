
#pragma once

#include <variant>

#define IMMEDIATE
#define STACK

using uint64 = unsigned long long;
using bytes = std::variant<const char*, const unsigned char*>;

const uint64 NoOp = 0;
const uint64 OptIn = 1;
const uint64 CloseOut = 2;
const uint64 ClearState = 3;
const uint64 UpdateApplication = 4;
const uint64 DeleteApplication = 5;

struct any
{
	bool operator==(const any);
};

/*
err - Fail immediately.
*/
void avm_err();
/*
sha256 - SHA256 hash of value A, yields [32]byte
*/
bytes avm_sha256(STACK bytes s1);
/*
keccak256 - Keccak256 hash of value A, yields [32]byte
*/
bytes avm_keccak256(STACK bytes s1);
/*
sha512_256 - SHA512_256 hash of value A, yields [32]byte
*/
bytes avm_sha512_256(STACK bytes s1);
/*
ed25519verify - for (data A, signature B, pubkey C) verify the signature of ("ProgData" || program_hash || data) against the pubkey => {0 or 1}
*/
uint64 avm_ed25519verify(STACK bytes s1, STACK bytes s2, STACK bytes s3);
/*
ecdsa_verify - for (data A, signature B, C and pubkey D, E) verify the signature of the data against the pubkey => {0 or 1}
*/
uint64 avm_ecdsa_verify(STACK bytes s1, STACK bytes s2, STACK bytes s3, STACK bytes s4);
struct avm_ecdsa_pk_decompress_result_t
{
	bytes r1;
	bytes r2;
};
/*
ecdsa_pk_decompress - decompress pubkey A into components X, Y
*/
struct avm_ecdsa_pk_decompress_t
{
	const void Secp256k1;
	const void Secp256r1;
};
const extern avm_ecdsa_pk_decompress_t avm_ecdsa_pk_decompress;
struct avm_ecdsa_pk_recover_result_t
{
	bytes r1;
	bytes r2;
};
/*
ecdsa_pk_recover - for (data A, recovery id B, signature C, D) recover a public key
*/
avm_ecdsa_pk_recover_result_t avm_ecdsa_pk_recover(STACK bytes s1, STACK uint64 s2, STACK bytes s3);
/*
+ - A plus B. Fail on overflow.
*/
uint64 avm_plus(STACK uint64 s1, STACK uint64 s2);
/*
- - A minus B. Fail if B > A.
*/
uint64 avm_minus(STACK uint64 s1, STACK uint64 s2);
/*
/ - A divided by B (truncated division). Fail if B == 0.
*/
uint64 avm_div(STACK uint64 s1, STACK uint64 s2);
/*
* - A times B. Fail on overflow.
*/
uint64 avm_mul(STACK uint64 s1, STACK uint64 s2);
/*
< - A less than B => {0 or 1}
*/
uint64 avm_lt(STACK uint64 s1, STACK uint64 s2);
/*
> - A greater than B => {0 or 1}
*/
uint64 avm_gt(STACK uint64 s1, STACK uint64 s2);
/*
<= - A less than or equal to B => {0 or 1}
*/
uint64 avm_lteq(STACK uint64 s1, STACK uint64 s2);
/*
>= - A greater than or equal to B => {0 or 1}
*/
uint64 avm_gteq(STACK uint64 s1, STACK uint64 s2);
/*
&& - A is not zero and B is not zero => {0 or 1}
*/
uint64 avm_andand(STACK uint64 s1, STACK uint64 s2);
/*
|| - A is not zero or B is not zero => {0 or 1}
*/
uint64 avm_oror(STACK uint64 s1, STACK uint64 s2);
/*
== - A is equal to B => {0 or 1}
*/
uint64 avm_eqeq(STACK any s1, STACK any s2);
/*
!= - A is not equal to B => {0 or 1}
*/
uint64 avm_noteq(STACK any s1, STACK any s2);
/*
! - A == 0 yields 1; else 0
*/
uint64 avm_not(STACK uint64 s1);
/*
len - yields length of byte value A
*/
uint64 avm_len(STACK bytes s1);
/*
itob - converts uint64 A to big-endian byte array, always of length 8
*/
bytes avm_itob(STACK uint64 s1);
/*
btoi - converts big-endian byte array A to uint64. Fails if len(A) > 8. Padded by leading 0s if len(A) < 8.
*/
uint64 avm_btoi(STACK bytes s1);
/*
% - A modulo B. Fail if B == 0.
*/
uint64 avm_mod(STACK uint64 s1, STACK uint64 s2);
/*
| - A bitwise-or B
*/
uint64 avm_or(STACK uint64 s1, STACK uint64 s2);
/*
& - A bitwise-and B
*/
uint64 avm_and(STACK uint64 s1, STACK uint64 s2);
/*
^ - A bitwise-xor B
*/
uint64 avm_xor(STACK uint64 s1, STACK uint64 s2);
/*
~ - bitwise invert value A
*/
uint64 avm_inv(STACK uint64 s1);
struct avm_mulw_result_t
{
	uint64 r1;
	uint64 r2;
};
/*
mulw - A times B as a 128-bit result in two uint64s. X is the high 64 bits, Y is the low
*/
avm_mulw_result_t avm_mulw(STACK uint64 s1, STACK uint64 s2);
struct avm_addw_result_t
{
	uint64 r1;
	uint64 r2;
};
/*
addw - A plus B as a 128-bit result. X is the carry-bit, Y is the low-order 64 bits.
*/
avm_addw_result_t avm_addw(STACK uint64 s1, STACK uint64 s2);
struct avm_divmodw_result_t
{
	uint64 r1;
	uint64 r2;
	uint64 r3;
	uint64 r4;
};
/*
divmodw - W,X = (A,B / C,D); Y,Z = (A,B modulo C,D)
*/
avm_divmodw_result_t avm_divmodw(STACK uint64 s1, STACK uint64 s2, STACK uint64 s3, STACK uint64 s4);
/*
intcblock - prepare block of uint64 constants for use by intc
*/
void avm_intcblock();
/*
intc - Ith constant from intcblock
*/
struct avm_intc_t
{
};
const extern avm_intc_t avm_intc;
/*
intc_0 - constant 0 from intcblock
*/
uint64 avm_intc_0();
/*
intc_1 - constant 1 from intcblock
*/
uint64 avm_intc_1();
/*
intc_2 - constant 2 from intcblock
*/
uint64 avm_intc_2();
/*
intc_3 - constant 3 from intcblock
*/
uint64 avm_intc_3();
/*
bytecblock - prepare block of byte-array constants for use by bytec
*/
void avm_bytecblock();
/*
bytec - Ith constant from bytecblock
*/
struct avm_bytec_t
{
};
const extern avm_bytec_t avm_bytec;
/*
bytec_0 - constant 0 from bytecblock
*/
bytes avm_bytec_0();
/*
bytec_1 - constant 1 from bytecblock
*/
bytes avm_bytec_1();
/*
bytec_2 - constant 2 from bytecblock
*/
bytes avm_bytec_2();
/*
bytec_3 - constant 3 from bytecblock
*/
bytes avm_bytec_3();
/*
arg - Nth LogicSig argument
*/
struct avm_arg_t
{
};
const extern avm_arg_t avm_arg;
/*
arg_0 - LogicSig argument 0
*/
bytes avm_arg_0();
/*
arg_1 - LogicSig argument 1
*/
bytes avm_arg_1();
/*
arg_2 - LogicSig argument 2
*/
bytes avm_arg_2();
/*
arg_3 - LogicSig argument 3
*/
bytes avm_arg_3();
/*
txn - field F of current transaction
*/
struct avm_txn_t
{
	const bytes Sender;
	const uint64 Fee;
	const uint64 FirstValid;
	const uint64 FirstValidTime;
	const uint64 LastValid;
	const bytes Note;
	const bytes Lease;
	const bytes Receiver;
	const uint64 Amount;
	const bytes CloseRemainderTo;
	const bytes VotePK;
	const bytes SelectionPK;
	const uint64 VoteFirst;
	const uint64 VoteLast;
	const uint64 VoteKeyDilution;
	const bytes Type;
	const uint64 TypeEnum;
	const uint64 XferAsset;
	const uint64 AssetAmount;
	const bytes AssetSender;
	const bytes AssetReceiver;
	const bytes AssetCloseTo;
	const uint64 GroupIndex;
	const bytes TxID;
	const uint64 ApplicationID;
	const uint64 OnCompletion;
	const bytes ApplicationArgs;
	const uint64 NumAppArgs;
	const bytes Accounts;
	const uint64 NumAccounts;
	const bytes ApprovalProgram;
	const bytes ClearStateProgram;
	const bytes RekeyTo;
	const uint64 ConfigAsset;
	const uint64 ConfigAssetTotal;
	const uint64 ConfigAssetDecimals;
	const uint64 ConfigAssetDefaultFrozen;
	const bytes ConfigAssetUnitName;
	const bytes ConfigAssetName;
	const bytes ConfigAssetURL;
	const bytes ConfigAssetMetadataHash;
	const bytes ConfigAssetManager;
	const bytes ConfigAssetReserve;
	const bytes ConfigAssetFreeze;
	const bytes ConfigAssetClawback;
	const uint64 FreezeAsset;
	const bytes FreezeAssetAccount;
	const uint64 FreezeAssetFrozen;
	const uint64 Assets;
	const uint64 NumAssets;
	const uint64 Applications;
	const uint64 NumApplications;
	const uint64 GlobalNumUint;
	const uint64 GlobalNumByteSlice;
	const uint64 LocalNumUint;
	const uint64 LocalNumByteSlice;
	const uint64 ExtraProgramPages;
	const uint64 Nonparticipation;
	const bytes Logs;
	const uint64 NumLogs;
	const uint64 CreatedAssetID;
	const uint64 CreatedApplicationID;
	const bytes LastLog;
	const bytes StateProofPK;
	const bytes ApprovalProgramPages;
	const uint64 NumApprovalProgramPages;
	const bytes ClearStateProgramPages;
	const uint64 NumClearStateProgramPages;
};
const extern avm_txn_t avm_txn;
/*
global - global field F
*/
struct avm_global_t
{
	const uint64 MinTxnFee;
	const uint64 MinBalance;
	const uint64 MaxTxnLife;
	const bytes ZeroAddress;
	const uint64 GroupSize;
	const uint64 LogicSigVersion;
	const uint64 Round;
	const uint64 LatestTimestamp;
	const uint64 CurrentApplicationID;
	const bytes CreatorAddress;
	const bytes CurrentApplicationAddress;
	const bytes GroupID;
	const uint64 OpcodeBudget;
	const uint64 CallerApplicationID;
	const bytes CallerApplicationAddress;
};
const extern avm_global_t avm_global;
/*
gtxn - field F of the Tth transaction in the current group
*/
struct avm_gtxn_t
{
	bytes Sender(uint64 i2);
	uint64 Fee(uint64 i2);
	uint64 FirstValid(uint64 i2);
	uint64 FirstValidTime(uint64 i2);
	uint64 LastValid(uint64 i2);
	bytes Note(uint64 i2);
	bytes Lease(uint64 i2);
	bytes Receiver(uint64 i2);
	uint64 Amount(uint64 i2);
	bytes CloseRemainderTo(uint64 i2);
	bytes VotePK(uint64 i2);
	bytes SelectionPK(uint64 i2);
	uint64 VoteFirst(uint64 i2);
	uint64 VoteLast(uint64 i2);
	uint64 VoteKeyDilution(uint64 i2);
	bytes Type(uint64 i2);
	uint64 TypeEnum(uint64 i2);
	uint64 XferAsset(uint64 i2);
	uint64 AssetAmount(uint64 i2);
	bytes AssetSender(uint64 i2);
	bytes AssetReceiver(uint64 i2);
	bytes AssetCloseTo(uint64 i2);
	uint64 GroupIndex(uint64 i2);
	bytes TxID(uint64 i2);
	uint64 ApplicationID(uint64 i2);
	uint64 OnCompletion(uint64 i2);
	bytes ApplicationArgs(uint64 i2);
	uint64 NumAppArgs(uint64 i2);
	bytes Accounts(uint64 i2);
	uint64 NumAccounts(uint64 i2);
	bytes ApprovalProgram(uint64 i2);
	bytes ClearStateProgram(uint64 i2);
	bytes RekeyTo(uint64 i2);
	uint64 ConfigAsset(uint64 i2);
	uint64 ConfigAssetTotal(uint64 i2);
	uint64 ConfigAssetDecimals(uint64 i2);
	uint64 ConfigAssetDefaultFrozen(uint64 i2);
	bytes ConfigAssetUnitName(uint64 i2);
	bytes ConfigAssetName(uint64 i2);
	bytes ConfigAssetURL(uint64 i2);
	bytes ConfigAssetMetadataHash(uint64 i2);
	bytes ConfigAssetManager(uint64 i2);
	bytes ConfigAssetReserve(uint64 i2);
	bytes ConfigAssetFreeze(uint64 i2);
	bytes ConfigAssetClawback(uint64 i2);
	uint64 FreezeAsset(uint64 i2);
	bytes FreezeAssetAccount(uint64 i2);
	uint64 FreezeAssetFrozen(uint64 i2);
	uint64 Assets(uint64 i2);
	uint64 NumAssets(uint64 i2);
	uint64 Applications(uint64 i2);
	uint64 NumApplications(uint64 i2);
	uint64 GlobalNumUint(uint64 i2);
	uint64 GlobalNumByteSlice(uint64 i2);
	uint64 LocalNumUint(uint64 i2);
	uint64 LocalNumByteSlice(uint64 i2);
	uint64 ExtraProgramPages(uint64 i2);
	uint64 Nonparticipation(uint64 i2);
	bytes Logs(uint64 i2);
	uint64 NumLogs(uint64 i2);
	uint64 CreatedAssetID(uint64 i2);
	uint64 CreatedApplicationID(uint64 i2);
	bytes LastLog(uint64 i2);
	bytes StateProofPK(uint64 i2);
	bytes ApprovalProgramPages(uint64 i2);
	uint64 NumApprovalProgramPages(uint64 i2);
	bytes ClearStateProgramPages(uint64 i2);
	uint64 NumClearStateProgramPages(uint64 i2);
};
const extern avm_gtxn_t avm_gtxn;
/*
load - Ith scratch space value. All scratch spaces are 0 at program start.
*/
struct avm_load_t
{
};
const extern avm_load_t avm_load;
/*
store - store A to the Ith scratch space
*/
struct avm_store_t
{
};
const extern avm_store_t avm_store;
/*
txna - Ith value of the array field F of the current transaction
`txna` can be called using `txn` with 2 immediates.
*/
struct avm_txna_t
{
	bytes ApplicationArgs(bytes i2);
	bytes Accounts(bytes i2);
	uint64 Assets(bytes i2);
	uint64 Applications(bytes i2);
	bytes Logs(bytes i2);
	bytes ApprovalProgramPages(bytes i2);
	bytes ClearStateProgramPages(bytes i2);
};
const extern avm_txna_t avm_txna;
/*
gtxna - Ith value of the array field F from the Tth transaction in the current group
`gtxna` can be called using `gtxn` with 3 immediates.
*/
struct avm_gtxna_t
{
	bytes ApplicationArgs(bytes i2);
	bytes Accounts(bytes i2);
	uint64 Assets(bytes i2);
	uint64 Applications(bytes i2);
	bytes Logs(bytes i2);
	bytes ApprovalProgramPages(bytes i2);
	bytes ClearStateProgramPages(bytes i2);
};
const extern avm_gtxna_t avm_gtxna;
/*
gtxns - field F of the Ath transaction in the current group
*/
struct avm_gtxns_t
{
	const bytes Sender;
	const uint64 Fee;
	const uint64 FirstValid;
	const uint64 FirstValidTime;
	const uint64 LastValid;
	const bytes Note;
	const bytes Lease;
	const bytes Receiver;
	const uint64 Amount;
	const bytes CloseRemainderTo;
	const bytes VotePK;
	const bytes SelectionPK;
	const uint64 VoteFirst;
	const uint64 VoteLast;
	const uint64 VoteKeyDilution;
	const bytes Type;
	const uint64 TypeEnum;
	const uint64 XferAsset;
	const uint64 AssetAmount;
	const bytes AssetSender;
	const bytes AssetReceiver;
	const bytes AssetCloseTo;
	const uint64 GroupIndex;
	const bytes TxID;
	const uint64 ApplicationID;
	const uint64 OnCompletion;
	const bytes ApplicationArgs;
	const uint64 NumAppArgs;
	const bytes Accounts;
	const uint64 NumAccounts;
	const bytes ApprovalProgram;
	const bytes ClearStateProgram;
	const bytes RekeyTo;
	const uint64 ConfigAsset;
	const uint64 ConfigAssetTotal;
	const uint64 ConfigAssetDecimals;
	const uint64 ConfigAssetDefaultFrozen;
	const bytes ConfigAssetUnitName;
	const bytes ConfigAssetName;
	const bytes ConfigAssetURL;
	const bytes ConfigAssetMetadataHash;
	const bytes ConfigAssetManager;
	const bytes ConfigAssetReserve;
	const bytes ConfigAssetFreeze;
	const bytes ConfigAssetClawback;
	const uint64 FreezeAsset;
	const bytes FreezeAssetAccount;
	const uint64 FreezeAssetFrozen;
	const uint64 Assets;
	const uint64 NumAssets;
	const uint64 Applications;
	const uint64 NumApplications;
	const uint64 GlobalNumUint;
	const uint64 GlobalNumByteSlice;
	const uint64 LocalNumUint;
	const uint64 LocalNumByteSlice;
	const uint64 ExtraProgramPages;
	const uint64 Nonparticipation;
	const bytes Logs;
	const uint64 NumLogs;
	const uint64 CreatedAssetID;
	const uint64 CreatedApplicationID;
	const bytes LastLog;
	const bytes StateProofPK;
	const bytes ApprovalProgramPages;
	const uint64 NumApprovalProgramPages;
	const bytes ClearStateProgramPages;
	const uint64 NumClearStateProgramPages;
};
const extern avm_gtxns_t avm_gtxns;
/*
gtxnsa - Ith value of the array field F from the Ath transaction in the current group
`gtxnsa` can be called using `gtxns` with 2 immediates.
*/
struct avm_gtxnsa_t
{
	const bytes ApplicationArgs;
	const bytes Accounts;
	const uint64 Assets;
	const uint64 Applications;
	const bytes Logs;
	const bytes ApprovalProgramPages;
	const bytes ClearStateProgramPages;
};
const extern avm_gtxnsa_t avm_gtxnsa;
/*
gload - Ith scratch space value of the Tth transaction in the current group
*/
struct avm_gload_t
{
};
const extern avm_gload_t avm_gload;
/*
gloads - Ith scratch space value of the Ath transaction in the current group
*/
struct avm_gloads_t
{
};
const extern avm_gloads_t avm_gloads;
/*
gaid - ID of the asset or application created in the Tth transaction of the current group
*/
struct avm_gaid_t
{
};
const extern avm_gaid_t avm_gaid;
/*
gaids - ID of the asset or application created in the Ath transaction of the current group
*/
uint64 avm_gaids(STACK uint64 s1);
/*
loads - Ath scratch space value.  All scratch spaces are 0 at program start.
*/
any avm_loads(STACK uint64 s1);
/*
stores - store B to the Ath scratch space
*/
void avm_stores(STACK uint64 s1, STACK any s2);
/*
bnz - branch to TARGET if value A is not zero
*/
struct avm_bnz_t
{
};
const extern avm_bnz_t avm_bnz;
/*
bz - branch to TARGET if value A is zero
*/
struct avm_bz_t
{
};
const extern avm_bz_t avm_bz;
/*
b - branch unconditionally to TARGET
*/
struct avm_b_t
{
};
const extern avm_b_t avm_b;
/*
return - use A as success value; end
*/
void avm_return_(STACK uint64 s1);
/*
assert - immediately fail unless A is a non-zero number
*/
void avm_assert(STACK uint64 s1);
/*
bury - replace the Nth value from the top of the stack with A. bury 0 fails.
*/
struct avm_bury_t
{
};
const extern avm_bury_t avm_bury;
/*
popn - remove N values from the top of the stack
*/
struct avm_popn_t
{
};
const extern avm_popn_t avm_popn;
/*
dupn - duplicate A, N times
*/
struct avm_dupn_t
{
};
const extern avm_dupn_t avm_dupn;
/*
pop - discard A
*/
void avm_pop(STACK any s1);
struct avm_dup_result_t
{
	any r1;
	any r2;
};
/*
dup - duplicate A
*/
avm_dup_result_t avm_dup(STACK any s1);
struct avm_dup2_result_t
{
	any r1;
	any r2;
	any r3;
	any r4;
};
/*
dup2 - duplicate A and B
*/
avm_dup2_result_t avm_dup2(STACK any s1, STACK any s2);
struct avm_dig_result_t
{
	any r1;
	any r2;
};
/*
dig - Nth value from the top of the stack. dig 0 is equivalent to dup
*/
struct avm_dig_t
{
};
const extern avm_dig_t avm_dig;
struct avm_swap_result_t
{
	any r1;
	any r2;
};
/*
swap - swaps A and B on stack
*/
avm_swap_result_t avm_swap(STACK any s1, STACK any s2);
/*
select - selects one of two values based on top-of-stack: B if C != 0, else A
*/
any avm_select(STACK any s1, STACK any s2, STACK uint64 s3);
/*
cover - remove top of stack, and place it deeper in the stack such that N elements are above it. Fails if stack depth <= N.
*/
struct avm_cover_t
{
};
const extern avm_cover_t avm_cover;
/*
uncover - remove the value at depth N in the stack and shift above items down so the Nth deep value is on top of the stack. Fails if stack depth <= N.
*/
struct avm_uncover_t
{
};
const extern avm_uncover_t avm_uncover;
/*
concat - join A and B
*/
bytes avm_concat(STACK bytes s1, STACK bytes s2);
/*
substring - A range of bytes from A starting at S up to but not including E. If E < S, or either is larger than the array length, the program fails
*/
struct avm_substring_t
{
};
const extern avm_substring_t avm_substring;
/*
substring3 - A range of bytes from A starting at B up to but not including C. If C < B, or either is larger than the array length, the program fails
*/
bytes avm_substring3(STACK bytes s1, STACK uint64 s2, STACK uint64 s3);
/*
getbit - Bth bit of (byte-array or integer) A. If B is greater than or equal to the bit length of the value (8*byte length), the program fails
*/
uint64 avm_getbit(STACK any s1, STACK uint64 s2);
/*
setbit - Copy of (byte-array or integer) A, with the Bth bit set to (0 or 1) C. If B is greater than or equal to the bit length of the value (8*byte length), the program fails
*/
any avm_setbit(STACK any s1, STACK uint64 s2, STACK uint64 s3);
/*
getbyte - Bth byte of A, as an integer. If B is greater than or equal to the array length, the program fails
*/
uint64 avm_getbyte(STACK bytes s1, STACK uint64 s2);
/*
setbyte - Copy of A with the Bth byte set to small integer (between 0..255) C. If B is greater than or equal to the array length, the program fails
*/
bytes avm_setbyte(STACK bytes s1, STACK uint64 s2, STACK uint64 s3);
/*
extract - A range of bytes from A starting at S up to but not including S+L. If L is 0, then extract to the end of the string. If S or S+L is larger than the array length, the program fails
*/
struct avm_extract_t
{
};
const extern avm_extract_t avm_extract;
/*
extract3 - A range of bytes from A starting at B up to but not including B+C. If B+C is larger than the array length, the program fails
`extract3` can be called using `extract` with no immediates.
*/
bytes avm_extract3(STACK bytes s1, STACK uint64 s2, STACK uint64 s3);
/*
extract_uint16 - A uint16 formed from a range of big-endian bytes from A starting at B up to but not including B+2. If B+2 is larger than the array length, the program fails
*/
uint64 avm_extract_uint16(STACK bytes s1, STACK uint64 s2);
/*
extract_uint32 - A uint32 formed from a range of big-endian bytes from A starting at B up to but not including B+4. If B+4 is larger than the array length, the program fails
*/
uint64 avm_extract_uint32(STACK bytes s1, STACK uint64 s2);
/*
extract_uint64 - A uint64 formed from a range of big-endian bytes from A starting at B up to but not including B+8. If B+8 is larger than the array length, the program fails
*/
uint64 avm_extract_uint64(STACK bytes s1, STACK uint64 s2);
/*
replace2 - Copy of A with the bytes starting at S replaced by the bytes of B. Fails if S+len(B) exceeds len(A)
`replace2` can be called using `replace` with 1 immediate.
*/
bytes avm_replace2(STACK bytes s1, STACK bytes s2);
/*
replace3 - Copy of A with the bytes starting at B replaced by the bytes of C. Fails if B+len(C) exceeds len(A)
`replace3` can be called using `replace` with no immediates.
*/
bytes avm_replace3(STACK bytes s1, STACK uint64 s2, STACK bytes s3);
/*
base64_decode - decode A which was base64-encoded using _encoding_ E. Fail if A is not base64 encoded with encoding E
*/
struct avm_base64_decode_t
{
	const any URLEncoding;
	const any StdEncoding;
};
const extern avm_base64_decode_t avm_base64_decode;
/*
json_ref - key B's value, of type R, from a [valid](jsonspec.md) utf-8 encoded json object A
*/
any avm_json_ref(STACK bytes s1);
/*
balance - balance for account A, in microalgos. The balance is observed after the effects of previous transactions in the group, and after the fee for the current transaction is deducted. Changes caused by inner transactions are observable immediately following `itxn_submit`
*/
uint64 avm_balance(STACK any s1);
/*
app_opted_in - 1 if account A is opted in to application B, else 0
*/
uint64 avm_app_opted_in(STACK any s1, STACK uint64 s2);
/*
app_local_get - local state of the key B in the current application in account A
*/
any avm_app_local_get(STACK any s1, STACK bytes s2);
struct avm_app_local_get_ex_result_t
{
	any r1;
	uint64 r2;
};
/*
app_local_get_ex - X is the local state of application B, key C in account A. Y is 1 if key existed, else 0
*/
avm_app_local_get_ex_result_t avm_app_local_get_ex(STACK any s1, STACK uint64 s2, STACK bytes s3);
/*
app_global_get - global state of the key A in the current application
*/
any avm_app_global_get(STACK bytes s1);
struct avm_app_global_get_ex_result_t
{
	any r1;
	uint64 r2;
};
/*
app_global_get_ex - X is the global state of application A, key B. Y is 1 if key existed, else 0
*/
avm_app_global_get_ex_result_t avm_app_global_get_ex(STACK uint64 s1, STACK bytes s2);
/*
app_local_put - write C to key B in account A's local state of the current application
*/
void avm_app_local_put(STACK any s1, STACK bytes s2, STACK any s3);
/*
app_global_put - write B to key A in the global state of the current application
*/
void avm_app_global_put(STACK bytes s1, STACK any s2);
/*
app_local_del - delete key B from account A's local state of the current application
*/
void avm_app_local_del(STACK any s1, STACK bytes s2);
/*
app_global_del - delete key A from the global state of the current application
*/
void avm_app_global_del(STACK bytes s1);
struct avm_asset_holding_get_result_t
{
	any r1;
	uint64 r2;
};
/*
asset_holding_get - X is field F from account A's holding of asset B. Y is 1 if A is opted into B, else 0
*/
avm_asset_holding_get_result_t avm_asset_holding_get(STACK any s1);
struct avm_asset_params_get_result_t
{
	any r1;
	uint64 r2;
};
/*
asset_params_get - X is field F from asset A. Y is 1 if A exists, else 0
*/
struct avm_asset_params_get_t
{
	const uint64 AssetTotal;
	const uint64 AssetDecimals;
	const uint64 AssetDefaultFrozen;
	const bytes AssetUnitName;
	const bytes AssetName;
	const bytes AssetURL;
	const bytes AssetMetadataHash;
	const bytes AssetManager;
	const bytes AssetReserve;
	const bytes AssetFreeze;
	const bytes AssetClawback;
	const bytes AssetCreator;
};
const extern avm_asset_params_get_t avm_asset_params_get;
struct avm_app_params_get_result_t
{
	any r1;
	uint64 r2;
};
/*
app_params_get - X is field F from app A. Y is 1 if A exists, else 0
*/
struct avm_app_params_get_t
{
	const bytes AppApprovalProgram;
	const bytes AppClearStateProgram;
	const uint64 AppGlobalNumUint;
	const uint64 AppGlobalNumByteSlice;
	const uint64 AppLocalNumUint;
	const uint64 AppLocalNumByteSlice;
	const uint64 AppExtraProgramPages;
	const bytes AppCreator;
	const bytes AppAddress;
};
const extern avm_app_params_get_t avm_app_params_get;
struct avm_acct_params_get_result_t
{
	any r1;
	uint64 r2;
};
/*
acct_params_get - X is field F from account A. Y is 1 if A owns positive algos, else 0
*/
struct avm_acct_params_get_t
{
	const uint64 AcctBalance;
	const uint64 AcctMinBalance;
	const bytes AcctAuthAddr;
	const uint64 AcctTotalNumUint;
	const uint64 AcctTotalNumByteSlice;
	const uint64 AcctTotalExtraAppPages;
	const uint64 AcctTotalAppsCreated;
	const uint64 AcctTotalAppsOptedIn;
	const uint64 AcctTotalAssetsCreated;
	const uint64 AcctTotalAssets;
	const uint64 AcctTotalBoxes;
	const uint64 AcctTotalBoxBytes;
};
const extern avm_acct_params_get_t avm_acct_params_get;
/*
min_balance - minimum required balance for account A, in microalgos. Required balance is affected by ASA, App, and Box usage. When creating or opting into an app, the minimum balance grows before the app code runs, therefore the increase is visible there. When deleting or closing out, the minimum balance decreases after the app executes. Changes caused by inner transactions or box usage are observable immediately following the opcode effecting the change.
*/
uint64 avm_min_balance(STACK any s1);
/*
pushbytes - immediate BYTES
*/
bytes avm_pushbytes();
/*
pushint - immediate UINT
*/
uint64 avm_pushint();
/*
pushbytess - push sequences of immediate byte arrays to stack (first byte array being deepest)
*/
void avm_pushbytess();
/*
pushints - push sequence of immediate uints to stack in the order they appear (first uint being deepest)
*/
void avm_pushints();
/*
ed25519verify_bare - for (data A, signature B, pubkey C) verify the signature of the data against the pubkey => {0 or 1}
*/
uint64 avm_ed25519verify_bare(STACK bytes s1, STACK bytes s2, STACK bytes s3);
/*
callsub - branch unconditionally to TARGET, saving the next instruction on the call stack
*/
struct avm_callsub_t
{
};
const extern avm_callsub_t avm_callsub;
/*
retsub - pop the top instruction from the call stack and branch to it
*/
void avm_retsub();
/*
proto - Prepare top call frame for a retsub that will assume A args and R return values.
*/
struct avm_proto_t
{
};
const extern avm_proto_t avm_proto;
/*
frame_dig - Nth (signed) value from the frame pointer.
*/
struct avm_frame_dig_t
{
};
const extern avm_frame_dig_t avm_frame_dig;
/*
frame_bury - replace the Nth (signed) value from the frame pointer in the stack with A
*/
struct avm_frame_bury_t
{
};
const extern avm_frame_bury_t avm_frame_bury;
/*
switch - branch to the Ath label. Continue at following instruction if index A exceeds the number of labels.
*/
void avm_switch_(STACK uint64 s1);
/*
match - given match cases from A[1] to A[N], branch to the Ith label where A[I] = B. Continue to the following instruction if no matches are found.
*/
void avm_match();
/*
shl - A times 2^B, modulo 2^64
*/
uint64 avm_shl(STACK uint64 s1, STACK uint64 s2);
/*
shr - A divided by 2^B
*/
uint64 avm_shr(STACK uint64 s1, STACK uint64 s2);
/*
sqrt - The largest integer I such that I^2 <= A
*/
uint64 avm_sqrt(STACK uint64 s1);
/*
bitlen - The highest set bit in A. If A is a byte-array, it is interpreted as a big-endian unsigned integer. bitlen of 0 is 0, bitlen of 8 is 4
*/
uint64 avm_bitlen(STACK any s1);
/*
exp - A raised to the Bth power. Fail if A == B == 0 and on overflow
*/
uint64 avm_exp(STACK uint64 s1, STACK uint64 s2);
struct avm_expw_result_t
{
	uint64 r1;
	uint64 r2;
};
/*
expw - A raised to the Bth power as a 128-bit result in two uint64s. X is the high 64 bits, Y is the low. Fail if A == B == 0 or if the results exceeds 2^128-1
*/
avm_expw_result_t avm_expw(STACK uint64 s1, STACK uint64 s2);
/*
bsqrt - The largest integer I such that I^2 <= A. A and I are interpreted as big-endian unsigned integers
*/
bytes avm_bsqrt(STACK bytes s1);
/*
divw - A,B / C. Fail if C == 0 or if result overflows.
*/
uint64 avm_divw(STACK uint64 s1, STACK uint64 s2, STACK uint64 s3);
/*
sha3_256 - SHA3_256 hash of value A, yields [32]byte
*/
bytes avm_sha3_256(STACK bytes s1);
/*
b+ - A plus B. A and B are interpreted as big-endian unsigned integers
*/
bytes avm_bplus(STACK bytes s1, STACK bytes s2);
/*
b- - A minus B. A and B are interpreted as big-endian unsigned integers. Fail on underflow.
*/
bytes avm_bminus(STACK bytes s1, STACK bytes s2);
/*
b/ - A divided by B (truncated division). A and B are interpreted as big-endian unsigned integers. Fail if B is zero.
*/
bytes avm_bdiv(STACK bytes s1, STACK bytes s2);
/*
b* - A times B. A and B are interpreted as big-endian unsigned integers.
*/
bytes avm_bmul(STACK bytes s1, STACK bytes s2);
/*
b< - 1 if A is less than B, else 0. A and B are interpreted as big-endian unsigned integers
*/
uint64 avm_blt(STACK bytes s1, STACK bytes s2);
/*
b> - 1 if A is greater than B, else 0. A and B are interpreted as big-endian unsigned integers
*/
uint64 avm_bgt(STACK bytes s1, STACK bytes s2);
/*
b<= - 1 if A is less than or equal to B, else 0. A and B are interpreted as big-endian unsigned integers
*/
uint64 avm_blteq(STACK bytes s1, STACK bytes s2);
/*
b>= - 1 if A is greater than or equal to B, else 0. A and B are interpreted as big-endian unsigned integers
*/
uint64 avm_bgteq(STACK bytes s1, STACK bytes s2);
/*
b== - 1 if A is equal to B, else 0. A and B are interpreted as big-endian unsigned integers
*/
uint64 avm_beqeq(STACK bytes s1, STACK bytes s2);
/*
b!= - 0 if A is equal to B, else 1. A and B are interpreted as big-endian unsigned integers
*/
uint64 avm_bnoteq(STACK bytes s1, STACK bytes s2);
/*
b% - A modulo B. A and B are interpreted as big-endian unsigned integers. Fail if B is zero.
*/
bytes avm_bmod(STACK bytes s1, STACK bytes s2);
/*
b| - A bitwise-or B. A and B are zero-left extended to the greater of their lengths
*/
bytes avm_bor(STACK bytes s1, STACK bytes s2);
/*
b& - A bitwise-and B. A and B are zero-left extended to the greater of their lengths
*/
bytes avm_band(STACK bytes s1, STACK bytes s2);
/*
b^ - A bitwise-xor B. A and B are zero-left extended to the greater of their lengths
*/
bytes avm_bxor(STACK bytes s1, STACK bytes s2);
/*
b~ - A with all bits inverted
*/
bytes avm_binv(STACK bytes s1);
/*
bzero - zero filled byte-array of length A
*/
bytes avm_bzero(STACK uint64 s1);
/*
log - write A to log state of the current application
*/
void avm_log(STACK bytes s1);
/*
itxn_begin - begin preparation of a new inner transaction in a new transaction group
*/
void avm_itxn_begin();
/*
itxn_field - set field F of the current inner transaction to A
*/
struct avm_itxn_field_t
{
	const bytes Sender;
	const uint64 Fee;
	const bytes Note;
	const bytes Receiver;
	const uint64 Amount;
	const bytes CloseRemainderTo;
	const bytes VotePK;
	const bytes SelectionPK;
	const uint64 VoteFirst;
	const uint64 VoteLast;
	const uint64 VoteKeyDilution;
	const bytes Type;
	const uint64 TypeEnum;
	const uint64 XferAsset;
	const uint64 AssetAmount;
	const bytes AssetSender;
	const bytes AssetReceiver;
	const bytes AssetCloseTo;
	const uint64 ApplicationID;
	const uint64 OnCompletion;
	const bytes ApplicationArgs;
	const bytes Accounts;
	const bytes ApprovalProgram;
	const bytes ClearStateProgram;
	const bytes RekeyTo;
	const uint64 ConfigAsset;
	const uint64 ConfigAssetTotal;
	const uint64 ConfigAssetDecimals;
	const uint64 ConfigAssetDefaultFrozen;
	const bytes ConfigAssetUnitName;
	const bytes ConfigAssetName;
	const bytes ConfigAssetURL;
	const bytes ConfigAssetMetadataHash;
	const bytes ConfigAssetManager;
	const bytes ConfigAssetReserve;
	const bytes ConfigAssetFreeze;
	const bytes ConfigAssetClawback;
	const uint64 FreezeAsset;
	const bytes FreezeAssetAccount;
	const uint64 FreezeAssetFrozen;
	const uint64 Assets;
	const uint64 Applications;
	const uint64 GlobalNumUint;
	const uint64 GlobalNumByteSlice;
	const uint64 LocalNumUint;
	const uint64 LocalNumByteSlice;
	const uint64 ExtraProgramPages;
	const uint64 Nonparticipation;
	const bytes StateProofPK;
	const bytes ApprovalProgramPages;
	const bytes ClearStateProgramPages;
};
const extern avm_itxn_field_t avm_itxn_field;
/*
itxn_submit - execute the current inner transaction group. Fail if executing this group would exceed the inner transaction limit, or if any transaction in the group fails.
*/
void avm_itxn_submit();
/*
itxn - field F of the last inner transaction
*/
struct avm_itxn_t
{
	const bytes Sender;
	const uint64 Fee;
	const uint64 FirstValid;
	const uint64 FirstValidTime;
	const uint64 LastValid;
	const bytes Note;
	const bytes Lease;
	const bytes Receiver;
	const uint64 Amount;
	const bytes CloseRemainderTo;
	const bytes VotePK;
	const bytes SelectionPK;
	const uint64 VoteFirst;
	const uint64 VoteLast;
	const uint64 VoteKeyDilution;
	const bytes Type;
	const uint64 TypeEnum;
	const uint64 XferAsset;
	const uint64 AssetAmount;
	const bytes AssetSender;
	const bytes AssetReceiver;
	const bytes AssetCloseTo;
	const uint64 GroupIndex;
	const bytes TxID;
	const uint64 ApplicationID;
	const uint64 OnCompletion;
	const bytes ApplicationArgs;
	const uint64 NumAppArgs;
	const bytes Accounts;
	const uint64 NumAccounts;
	const bytes ApprovalProgram;
	const bytes ClearStateProgram;
	const bytes RekeyTo;
	const uint64 ConfigAsset;
	const uint64 ConfigAssetTotal;
	const uint64 ConfigAssetDecimals;
	const uint64 ConfigAssetDefaultFrozen;
	const bytes ConfigAssetUnitName;
	const bytes ConfigAssetName;
	const bytes ConfigAssetURL;
	const bytes ConfigAssetMetadataHash;
	const bytes ConfigAssetManager;
	const bytes ConfigAssetReserve;
	const bytes ConfigAssetFreeze;
	const bytes ConfigAssetClawback;
	const uint64 FreezeAsset;
	const bytes FreezeAssetAccount;
	const uint64 FreezeAssetFrozen;
	const uint64 Assets;
	const uint64 NumAssets;
	const uint64 Applications;
	const uint64 NumApplications;
	const uint64 GlobalNumUint;
	const uint64 GlobalNumByteSlice;
	const uint64 LocalNumUint;
	const uint64 LocalNumByteSlice;
	const uint64 ExtraProgramPages;
	const uint64 Nonparticipation;
	const bytes Logs;
	const uint64 NumLogs;
	const uint64 CreatedAssetID;
	const uint64 CreatedApplicationID;
	const bytes LastLog;
	const bytes StateProofPK;
	const bytes ApprovalProgramPages;
	const uint64 NumApprovalProgramPages;
	const bytes ClearStateProgramPages;
	const uint64 NumClearStateProgramPages;
};
const extern avm_itxn_t avm_itxn;
/*
itxna - Ith value of the array field F of the last inner transaction
*/
struct avm_itxna_t
{
	bytes ApplicationArgs(bytes i2);
	bytes Accounts(bytes i2);
	uint64 Assets(bytes i2);
	uint64 Applications(bytes i2);
	bytes Logs(bytes i2);
	bytes ApprovalProgramPages(bytes i2);
	bytes ClearStateProgramPages(bytes i2);
};
const extern avm_itxna_t avm_itxna;
/*
itxn_next - begin preparation of a new inner transaction in the same transaction group
*/
void avm_itxn_next();
/*
gitxn - field F of the Tth transaction in the last inner group submitted
*/
struct avm_gitxn_t
{
	bytes Sender(uint64 i2);
	uint64 Fee(uint64 i2);
	uint64 FirstValid(uint64 i2);
	uint64 FirstValidTime(uint64 i2);
	uint64 LastValid(uint64 i2);
	bytes Note(uint64 i2);
	bytes Lease(uint64 i2);
	bytes Receiver(uint64 i2);
	uint64 Amount(uint64 i2);
	bytes CloseRemainderTo(uint64 i2);
	bytes VotePK(uint64 i2);
	bytes SelectionPK(uint64 i2);
	uint64 VoteFirst(uint64 i2);
	uint64 VoteLast(uint64 i2);
	uint64 VoteKeyDilution(uint64 i2);
	bytes Type(uint64 i2);
	uint64 TypeEnum(uint64 i2);
	uint64 XferAsset(uint64 i2);
	uint64 AssetAmount(uint64 i2);
	bytes AssetSender(uint64 i2);
	bytes AssetReceiver(uint64 i2);
	bytes AssetCloseTo(uint64 i2);
	uint64 GroupIndex(uint64 i2);
	bytes TxID(uint64 i2);
	uint64 ApplicationID(uint64 i2);
	uint64 OnCompletion(uint64 i2);
	bytes ApplicationArgs(uint64 i2);
	uint64 NumAppArgs(uint64 i2);
	bytes Accounts(uint64 i2);
	uint64 NumAccounts(uint64 i2);
	bytes ApprovalProgram(uint64 i2);
	bytes ClearStateProgram(uint64 i2);
	bytes RekeyTo(uint64 i2);
	uint64 ConfigAsset(uint64 i2);
	uint64 ConfigAssetTotal(uint64 i2);
	uint64 ConfigAssetDecimals(uint64 i2);
	uint64 ConfigAssetDefaultFrozen(uint64 i2);
	bytes ConfigAssetUnitName(uint64 i2);
	bytes ConfigAssetName(uint64 i2);
	bytes ConfigAssetURL(uint64 i2);
	bytes ConfigAssetMetadataHash(uint64 i2);
	bytes ConfigAssetManager(uint64 i2);
	bytes ConfigAssetReserve(uint64 i2);
	bytes ConfigAssetFreeze(uint64 i2);
	bytes ConfigAssetClawback(uint64 i2);
	uint64 FreezeAsset(uint64 i2);
	bytes FreezeAssetAccount(uint64 i2);
	uint64 FreezeAssetFrozen(uint64 i2);
	uint64 Assets(uint64 i2);
	uint64 NumAssets(uint64 i2);
	uint64 Applications(uint64 i2);
	uint64 NumApplications(uint64 i2);
	uint64 GlobalNumUint(uint64 i2);
	uint64 GlobalNumByteSlice(uint64 i2);
	uint64 LocalNumUint(uint64 i2);
	uint64 LocalNumByteSlice(uint64 i2);
	uint64 ExtraProgramPages(uint64 i2);
	uint64 Nonparticipation(uint64 i2);
	bytes Logs(uint64 i2);
	uint64 NumLogs(uint64 i2);
	uint64 CreatedAssetID(uint64 i2);
	uint64 CreatedApplicationID(uint64 i2);
	bytes LastLog(uint64 i2);
	bytes StateProofPK(uint64 i2);
	bytes ApprovalProgramPages(uint64 i2);
	uint64 NumApprovalProgramPages(uint64 i2);
	bytes ClearStateProgramPages(uint64 i2);
	uint64 NumClearStateProgramPages(uint64 i2);
};
const extern avm_gitxn_t avm_gitxn;
/*
gitxna - Ith value of the array field F from the Tth transaction in the last inner group submitted
*/
struct avm_gitxna_t
{
	bytes ApplicationArgs(bytes i2);
	bytes Accounts(bytes i2);
	uint64 Assets(bytes i2);
	uint64 Applications(bytes i2);
	bytes Logs(bytes i2);
	bytes ApprovalProgramPages(bytes i2);
	bytes ClearStateProgramPages(bytes i2);
};
const extern avm_gitxna_t avm_gitxna;
/*
box_create - create a box named A, of length B. Fail if A is empty or B exceeds 32,768. Returns 0 if A already existed, else 1
*/
uint64 avm_box_create(STACK bytes s1, STACK uint64 s2);
/*
box_extract - read C bytes from box A, starting at offset B. Fail if A does not exist, or the byte range is outside A's size.
*/
bytes avm_box_extract(STACK bytes s1, STACK uint64 s2, STACK uint64 s3);
/*
box_replace - write byte-array C into box A, starting at offset B. Fail if A does not exist, or the byte range is outside A's size.
*/
void avm_box_replace(STACK bytes s1, STACK uint64 s2, STACK bytes s3);
/*
box_del - delete box named A if it exists. Return 1 if A existed, 0 otherwise
*/
uint64 avm_box_del(STACK bytes s1);
struct avm_box_len_result_t
{
	uint64 r1;
	uint64 r2;
};
/*
box_len - X is the length of box A if A exists, else 0. Y is 1 if A exists, else 0.
*/
avm_box_len_result_t avm_box_len(STACK bytes s1);
struct avm_box_get_result_t
{
	bytes r1;
	uint64 r2;
};
/*
box_get - X is the contents of box A if A exists, else ''. Y is 1 if A exists, else 0.
*/
avm_box_get_result_t avm_box_get(STACK bytes s1);
/*
box_put - replaces the contents of box A with byte-array B. Fails if A exists and len(B) != len(box A). Creates A if it does not exist
*/
void avm_box_put(STACK bytes s1, STACK bytes s2);
/*
txnas - Ath value of the array field F of the current transaction
*/
struct avm_txnas_t
{
	const bytes ApplicationArgs;
	const bytes Accounts;
	const uint64 Assets;
	const uint64 Applications;
	const bytes Logs;
	const bytes ApprovalProgramPages;
	const bytes ClearStateProgramPages;
};
const extern avm_txnas_t avm_txnas;
/*
gtxnas - Ath value of the array field F from the Tth transaction in the current group
*/
struct avm_gtxnas_t
{
	const bytes ApplicationArgs;
	const bytes Accounts;
	const uint64 Assets;
	const uint64 Applications;
	const bytes Logs;
	const bytes ApprovalProgramPages;
	const bytes ClearStateProgramPages;
};
const extern avm_gtxnas_t avm_gtxnas;
/*
gtxnsas - Bth value of the array field F from the Ath transaction in the current group
*/
any avm_gtxnsas(STACK uint64 s1);
/*
args - Ath LogicSig argument
*/
bytes avm_args(STACK uint64 s1);
/*
gloadss - Bth scratch space value of the Ath transaction in the current group
*/
any avm_gloadss(STACK uint64 s1, STACK uint64 s2);
/*
itxnas - Ath value of the array field F of the last inner transaction
*/
struct avm_itxnas_t
{
};
const extern avm_itxnas_t avm_itxnas;
/*
gitxnas - Ath value of the array field F from the Tth transaction in the last inner group submitted
*/
struct avm_gitxnas_t
{
};
const extern avm_gitxnas_t avm_gitxnas;
struct avm_vrf_verify_result_t
{
	bytes r1;
	uint64 r2;
};
/*
vrf_verify - Verify the proof B of message A against pubkey C. Returns vrf output and verification flag.
*/
avm_vrf_verify_result_t avm_vrf_verify(STACK bytes s1, STACK bytes s2);
/*
block - field F of block A. Fail unless A falls between txn.LastValid-1002 and txn.FirstValid (exclusive)
*/
struct avm_block_t
{
	const bytes BlkSeed;
	const uint64 BlkTimestamp;
};
const extern avm_block_t avm_block;
