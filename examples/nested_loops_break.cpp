#include "avm.hpp"

uint64 avm_main()
{
    for(uint64 a = 0; a < 2; ++a)
    {
        for(uint64 b = 0; b < 3; ++b)
        {
            if(b > 0)
                break;
        }
    }
}