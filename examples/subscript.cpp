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

    return sender[2] + test.data[3];
}