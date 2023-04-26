#include "avm.hpp"

struct test_t
{
    uint64 f1;
    uint64 f2;
};

void test(uint64 a, test_t b)
{
    a = a;
    b = b;
}

void avm_main()
{
    test_t v;
    test(1, v);
}