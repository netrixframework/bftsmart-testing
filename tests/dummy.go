package tests

import (
	"time"

	"github.com/netrixframework/netrix/sm"
	"github.com/netrixframework/netrix/testlib"
)

func DummyTest() *testlib.TestCase {
	stateMachine := sm.NewStateMachine()
	filters := testlib.NewFilterSet()
	testcase := testlib.NewTestCase("Dummy", 2*time.Minute, stateMachine, filters)
	return testcase
}
