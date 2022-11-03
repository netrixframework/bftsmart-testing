package util

import (
	"github.com/netrixframework/netrix/sm"
	"github.com/netrixframework/netrix/types"
)

func IsPropose() sm.Condition {
	return func(e *types.Event, c *sm.Context) bool {
		bftMessage, ok := GetParsedMessage(e, c)
		if !ok {
			return false
		}
		return bftMessage.Type() == "Propose"
	}
}

func IsWrite() sm.Condition {
	return func(e *types.Event, c *sm.Context) bool {
		bftMessage, ok := GetParsedMessage(e, c)
		if !ok {
			return false
		}
		return bftMessage.Type() == "Write"
	}
}

func IsAccept() sm.Condition {
	return func(e *types.Event, c *sm.Context) bool {
		bftMessage, ok := GetParsedMessage(e, c)
		if !ok {
			return false
		}
		return bftMessage.Type() == "Accept"
	}
}

func IsEpoch(epoch int) sm.Condition {
	return func(e *types.Event, c *sm.Context) bool {
		bftMessage, ok := GetParsedMessage(e, c)
		if !ok {
			return false
		}
		return bftMessage.Epoch == epoch
	}
}

func IsView(view int) sm.Condition {
	return func(e *types.Event, c *sm.Context) bool {
		bftMessage, ok := GetParsedMessage(e, c)
		if !ok {
			return false
		}
		return bftMessage.Number == view
	}
}
