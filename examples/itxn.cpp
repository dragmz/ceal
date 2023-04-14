#include "avm.hpp"

uint64 avm_main()
{
    avm_itxn_begin();
    {
        avm_itxn_field.XferAsset(1);
        avm_itxn_field.AssetAmount(2);
    }
    avm_itxn_submit();
}