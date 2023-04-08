#include "avm.hpp"

uint64 avm_main()
{
	if (avm_txn.ApplicationID == 0)
	{
		avm_box_create(avm_txn.Sender, 4);
	}
	else
	{
		if (avm_txn.OnCompletion != NoOp)
			avm_err();

		avm_box_get(avm_txn.Sender);
	}
}