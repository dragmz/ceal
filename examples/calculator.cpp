#include "avm.hpp"

uint64 add(const uint64 a, const uint64 b)
{
    return a + b;
}

uint64 sub(const uint64 a, const uint64 b)
{
    return a - b;
}

uint64 avm_main()
{
    const uint64 a = avm_btoi(avm_txna.ApplicationArgs(1));
    const uint64 b = avm_btoi(avm_txna.ApplicationArgs(2));

    switch (avm_txna.ApplicationArgs(0))
    {
    case avm_method(add):
        add(a, b);
        break;

    case avm_method(sub):
        sub(a, b);
        break;
    }
}