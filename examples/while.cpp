#include "avm.hpp"

uint64 avm_main()
{
    uint64 a = 3;

    while(a > 0)
    {
        a = a - 1;
    }

    while(a < 3)
    {
        a = a + 1;
    }

    return a;
}