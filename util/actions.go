package util

import (
	"github.com/netrixframework/netrix/log"
	"github.com/netrixframework/netrix/testlib"
	"github.com/netrixframework/netrix/types"
)

func PrintMessage() testlib.Action {
	return func(e *types.Event, ctx *testlib.Context) (out []*types.Message) {
		bftMessage, ok := GetParsedMessage(e, ctx.Context)
		if !ok {
			return
		}
		ctx.Logger.With(log.LogParams{"message": bftMessage.String()}).Info("Observed Message")
		return
	}
}
