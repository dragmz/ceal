#include "avm.hpp"

struct test_t
{
    bytes data;
};

uint64 avm_main()
{
    const bytes sender = avm_txn.Sender;

    test_t test;
    test.data = sender;

    // test.data[3] = test.data[4]
    test.data[3] = test.data[4];
    
    // test.data[4] += 2
    test.data[4] += 2;
    
    // test.data[5]++
    test.data[5]++;

    // ++test.data[6];
    ++test.data[6];

    // return sender[2] + test.data[3]
    return sender[2] + test.data[3];
}