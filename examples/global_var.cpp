#include "avm.hpp"

const uint64 A = 123;
const uint64 B = 234;

const bytes C = "CValue";
const bytes D = "DValue";

uint64 avm_main()
{
    avm_log(C);
    avm_log(D);
    
    return A + B;
}