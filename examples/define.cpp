#include "avm.hpp"

#if 1

#define VALUE 123
#define TEST VALUE
#define CONST

#ifdef CONST
#if TEST
const uint64 ConstValue = TEST;
#endif
#endif

#endif

uint64 avm_main()
{
#if 1
#ifdef ZERO
    return 0;
#endif
#endif

#if 0
#if TEST
return 1;
#endif
#endif

#if TEST
#ifdef CONST
    return ConstValue;
#endif
#endif
}