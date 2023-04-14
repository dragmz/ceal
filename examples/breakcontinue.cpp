#include "avm.hpp"

uint64 avm_main()
{
    for(uint64 i = 0; i < 10; ++i)
    {
        if(i < 5)
            continue;
        else
            break;
    }

    uint64 i = 0;
    while(i < 10)
    {
        if(i++ < 5)
            continue;
        else
            break;
    }

    i = 0;
    do
    {
        if(i < 5)
            continue;
        else
            break;
    } while (i < 10);
}