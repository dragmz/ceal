#include "avm.hpp"

uint64 avm_main()
{
    asm("int 1"
        "callsub fun");
    avm_return__op();

    asm("fun:");

    avm_proto_op(1, 1);
    asm("frame_dig -1");
    asm("int 1");
    avm_plus_op();
    avm_retsub_op();
}