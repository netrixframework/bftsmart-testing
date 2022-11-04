package tests

import (
	"time"

	"github.com/netrixframework/bftsmart-testing/util"
	"github.com/netrixframework/netrix/sm"
	"github.com/netrixframework/netrix/testlib"
	"github.com/netrixframework/netrix/types"
)

func RecordProposal(as string) testlib.Action {
	return func(e *types.Event, ctx *testlib.Context) (message []*types.Message) {
		bftMessage, ok := util.GetParsedMessage(e, ctx.Context)
		if !ok {
			return
		}
		if bftMessage.Type != util.ProposeMessageType {
			return
		}
		ctx.Vars.Set(as, string(bftMessage.Value))
		return
	}
}

func DelayedPropose() *testlib.TestCase {
	filters := testlib.NewFilterSet()
	filters.AddFilter(
		testlib.If(util.IsNewEpoch()).Then(testlib.DeliverAllFromSet(sm.Set("delayedProposals"))),
	)

	testCase := testlib.NewTestCase("DelayedPropose", 2*time.Minute, DelayProposeProperty(), filters)
	return testCase
}

func DelayedProposeProperty() *sm.StateMachine {
	property := sm.NewStateMachine()

	property.Builder().On(
		util.IsNewEpochOf(1),
		"NewEpoch",
	).On(
		util.IsDecided(),
		sm.SuccessStateLabel,
	)

	return property
}
