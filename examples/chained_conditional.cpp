#include "avm.hpp"

uint64 avm_main()
{

    uint64 x = 0;
    x += 1;

    // ok
    if(x>0 && x<10 && x!=5){
        return 1;
    }else{
        return 0;
    }
}