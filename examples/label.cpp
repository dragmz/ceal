#include "avm.hpp"

uint64 avm_main()
{
    asm("int 1");
    avm_bnz_op("test");
    asm("test:");

    asm("int 1");
    asm("int 2");
    asm("int 1");
    avm_match_op("a", "b");
    asm("a:");
    asm("b:");
}