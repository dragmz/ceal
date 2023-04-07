#include "avm.hpp"

void init(const bytes sender, const uint64 ts)
{
	avm_box_create(sender, 4);
	avm_box_put(sender, avm_itob(ts));
}

void noop()
{
	bytes ts; // timestamp bytes
	
	ts = avm_box_extract(avm_txn.Sender, 0, 4);
	avm_log(ts);
}

uint64 avm_main()
{
	/* 
		this is AVM entry point

		it checks if the application already exists
		- initializes box if it doesn't
		- logs box contents if it does
	*/
	if (avm_txn.ApplicationID == 0)
	{
		init(avm_txn.Sender, avm_block.BlkTimestamp);
		return 1;
	}

	if (avm_txn.OnCompletion != NoOp)
		avm_err();

	noop();

	return 1;
}
