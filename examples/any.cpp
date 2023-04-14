#include "avm.hpp"

uint64 avm_main()
{
    const uint64 v1 = avm_bitlen(1);
    const uint64 v2 = avm_bitlen("test");
    const uint64 v3 = avm_bitlen(avm_box_extract("test", 0, 4));
    const uint64 v4 = avm_loads(0) == avm_loads(1);
}