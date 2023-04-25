#include "avm.hpp"

struct test_t
{
    uint64 a;
    uint64 b;
};

uint64 add(test_t v)
{
    return v.a + v.b;
}

uint64 avm_main()
{
    test_t v;

    v.a = 1;
    v.b = 2;

    return add(v);
}