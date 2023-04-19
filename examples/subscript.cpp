#include "avm.hpp"

uint64 avm_main()
{
    const bytes sender = avm_txn.Sender;
    return sender[2];
}