#include "avm.hpp"

uint64 a_eq_b(const uint64 a, const uint64 b)
{
    return a == b;
}

uint64 a_neq_b(const uint64 a, const uint64 b)
{
    return a != b;
}

uint64 a_lt_b(const uint64 a, const uint64 b)
{
    return a < b;
}

uint64 a_gt_b(const uint64 a, const uint64 b)
{
    return a > b;
}

uint64 a_lteq_b(const uint64 a, const uint64 b)
{
    return a <= b;
}

uint64 a_gteq_b(const uint64 a, const uint64 b)
{
    return a >= b;
}

uint64 not_a(const uint64 a)
{
    return !a;
}

uint64 avm_main()
{
}