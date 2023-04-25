#include "avm.hpp"

struct inner_t
{
    uint64 a;
    uint64 b;
};

struct test_t
{
    uint64 a;
    inner_t b;
};

uint64 add(test_t v)
{
    return v.a + v.b.a + v.b.b;
}

uint64 avm_main()
{
    test_t v;

    v.a = 1;
    v.b.a = 2;
    v.b.b = 2;

    return add(v);
}