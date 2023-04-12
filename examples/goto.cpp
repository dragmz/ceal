#include "avm.hpp"

uint64 avm_main()
{
    if(avm_txn.ApplicationID == 0)
        goto Quit;
        
    return 0;

Quit:

    return 1;
}