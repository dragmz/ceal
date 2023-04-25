#include "avm.hpp"

uint64 avm_main()
{
    switch (avm_txn.Fee)
    {
    case 0:
        avm_log("0");
        break;
    case 1:
        avm_log("1");
        break;
    default:
        avm_log("call");
        break;
    }

    switch (avm_txna.ApplicationArgs(0))
    {
    case const_string("test"):
        avm_log("test");
        break;
    }
}