#include "avm.hpp"

struct test_t
{
	uint64 f1;
	uint64 f2;
};

uint64 fun()
{
	test_t t;

	t.f1 = 5;
	t.f2 = 6;

	return t.f1 + t.f2;
}

uint64 avm_main()
{
	test_t t1;

	t1.f1 = 1;
	t1.f2 = 2;

	const uint64 sum = fun();

	test_t t2;

	t2.f1 = t1.f1;
	t2.f2 = t1.f2;

	return t2.f1 + t2.f2 + sum;
}