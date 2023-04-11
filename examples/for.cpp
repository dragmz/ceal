#include "avm.hpp"

uint64 avm_main()
{
    for (uint64 i = 0; i < 10; ++i)
    {
        avm_log(avm_itob(i));
        
        if (i == 5)
            break;
    }
}