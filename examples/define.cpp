#include "avm.hpp"

#define VALUE 123
#define TEST VALUE
#define CONST

#ifdef CONST
const uint64 ConstValue = TEST;
#endif

uint64 avm_main()
{
    #ifdef ZERO
    return 0;
    #endif
    
    #ifdef CONST
    return ConstValue;
    #endif
}