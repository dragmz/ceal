#include "avm.hpp"

struct point_t
{
    uint64 x;
    uint64 y;
};

uint64 avm_main()
{
    const bytes data = avm_box_extract("point", 0, 16);

    point_t pt;
    abi_decode(data, pt);
}