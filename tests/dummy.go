package tests

import (
	"time"

	"github.com/netrixframework/bftsmart-testing/util"
	"github.com/netrixframework/netrix/sm"
	"github.com/netrixframework/netrix/testlib"
)

func DummyTest() *testlib.TestCase {
	stateMachine := sm.NewStateMachine()
	filters := testlib.NewFilterSet()

	// filters.AddFilter(
	// 	testlib.If(util.IsWrite().And(util.IsEpoch(1))).Then(util.PrintMessage(), testlib.DeliverMessage()),
	// )

	filters.AddFilter(
		testlib.If(util.IsPropose().And(util.IsEpoch(0))).Then(testlib.DropMessage()),
	)

	testcase := testlib.NewTestCase("Dummy", 2*time.Minute, stateMachine, filters)
	return testcase
}
