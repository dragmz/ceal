#include "avm.hpp"

uint64 add(uint64 a, uint64 b)
{
    const uint64 v1 = a;
    const uint64 v2 = b;

    const uint64 sum = v1 + v2;

    return sum;
}

uint64 avm_main() {
    const uint64 v1 = 1;
    const uint64 v2 = 2;

    const uint64 sum = add(v1, v2);

    return sum;
}