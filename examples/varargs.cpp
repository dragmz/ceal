#include "avm.hpp"

uint64 avm_main()
{
    avm_pushints_op(1, 2, 3);
    avm_pushbytess_op("test1", "test2", "test3");
}