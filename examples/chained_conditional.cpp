#include "avm.hpp"

uint64 avm_main()
{

    uint64 x = 5;

    // this should short-circuit
    // for any falsy conditions
    if (x > 0 && x < 10 && x != 5)
    {
        return 1;
    }

    return 0;
}