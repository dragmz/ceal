#include "avm.hpp"

struct test_t
{
    uint64 f1;
    uint64 f2;
};

void avm_main()
{
    // test_t a
    test_t a;

    // set a fields
    a.f1 = 1;
    a.f2 = 2;
    
    // test_t b = a
    test_t b = a;

    // a = b
    a = b;
}