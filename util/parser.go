package util

import (
	"encoding/json"
	"fmt"

	"github.com/netrixframework/netrix/sm"
	"github.com/netrixframework/netrix/types"
)

var WriteMessageType string = "WRITE"
var ProposeMessageType string = "PROPOSE"
var AcceptMessageType string = "ACCEPT"
var StopMessageType string = "STOP"
var StopDataMessageType string = "STOPDATA"
var SyncMessageType string = "SYNC"

type BFTSmartParser struct {
}

func (*BFTSmartParser) Parse(data []byte) (types.ParsedMessage, error) {
	var m BFTSmartMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("error parsing: %s", err)
	}
	if m.Type == "" && m.PaxosType != 0 {
		switch m.PaxosType {
		case 44781:
			m.Type = ProposeMessageType
		case 44782:
			m.Type = WriteMessageType
		case 44783:
			m.Type = AcceptMessageType
		}
	}
	return &m, nil
}

var _ types.MessageParser = &BFTSmartParser{}

type BFTSmartMessage struct {
	Type      string
	Ts        int
	Payload   []byte
	Number    int
	Epoch     int
	PaxosType int
	Value     []byte
	Proof     []byte
}

var _ types.ParsedMessage = &BFTSmartMessage{}

func (m *BFTSmartMessage) Clone() types.ParsedMessage {
	return &BFTSmartMessage{
		Number:    m.Number,
		Epoch:     m.Epoch,
		PaxosType: m.PaxosType,
		Value:     m.Value,
		Proof:     m.Proof,
	}
}

func (m *BFTSmartMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *BFTSmartMessage) String() string {
	return fmt.Sprintf("{Number: %d, Epoch: %d, Type: %s, Value: %v}", m.Number, m.Epoch, m.Type, m.Value)
}

func GetParsedMessage(e *types.Event, c *sm.Context) (*BFTSmartMessage, bool) {
	messageID, ok := e.MessageID()
	if !ok {
		return nil, false
	}
	message, ok := c.MessagePool.Get(messageID)
	if !ok {
		return nil, false
	}
	bftMessage, ok := message.ParsedMessage.(*BFTSmartMessage)
	return bftMessage, ok
}
