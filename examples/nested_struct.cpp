#include "avm.hpp"

struct value_t
{
    uint64 v1;
    uint64 v2;
};

struct nested_t
{
    uint64 v1;
    uint64 v2;

    value_t n;
};

uint64 avm_main()
{
    nested_t n;

    n.v1 = 1;
    n.v2 = 2;

    n.n.v1 = 1;
    n.n.v2 = n.v1 + n.v2;

    return n.v1 + n.v2 + n.n.v1 + n.n.v2;
}