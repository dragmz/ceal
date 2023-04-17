#include "avm.hpp"

struct point_t
{
    uint64 x;
    uint64 y;
};

uint64 avm_main()
{
    point_t pt1;

    pt1.x = 1;
    pt1.y = 2;

    const bytes data = abi_encode(pt1);

    point_t pt2;
    abi_decode(data, pt2);
}