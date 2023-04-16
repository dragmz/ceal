#include "avm.hpp"

uint64 avm_main()
{
    asm("int 1"
        "return");
}