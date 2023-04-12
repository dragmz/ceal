#include "avm.hpp"

uint64 avm_main()
{
    const uint64 a = 1 * 2 + 3;
    const uint64 b = 1 * (2 + 3);
    const uint64 c = (1 * 2 + 3);
    const uint64 d = 1 && 2 || 3;
    const uint64 e = 1 && (2 || 3);
    const uint64 f = 1 + 2 * 2 + 3;
    const uint64 g = (1 + 2) * (2 + 3);
}