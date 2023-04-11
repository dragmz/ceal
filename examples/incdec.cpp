#include "avm.hpp"

void preinc()
{
    uint64 a = 1;
    uint64 b = ++a;
}

void predec()
{
    uint64 a = 1;
    uint64 b = --a;
}

void postinc()
{
    uint64 a = 1;
    uint64 b = a++;
}

void postdec()
{
    uint64 a = 1;
    uint64 b = --a;
}

uint64 avm_main()
{
}