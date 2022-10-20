package util

import (
	"encoding/json"
	"fmt"

	"github.com/netrixframework/netrix/types"
)

type BFTSmartParser struct {
}

func (*BFTSmartParser) Parse(data []byte) (types.ParsedMessage, error) {
	var m BFTSmartMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("error parsing: %s", err)
	}
	return &m, nil
}

var _ types.MessageParser = &BFTSmartParser{}

type BFTSmartMessage struct {
	Number int
	Epoch  int
	Type   string
	Value  []byte
	Proof  []byte
}

var _ types.ParsedMessage = &BFTSmartMessage{}

func (m *BFTSmartMessage) Clone() types.ParsedMessage {
	return &BFTSmartMessage{
		Number: m.Number,
		Epoch:  m.Epoch,
		Type:   m.Type,
		Value:  m.Value,
		Proof:  m.Proof,
	}
}

func (m *BFTSmartMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *BFTSmartMessage) String() string {
	return fmt.Sprintf("{Number: %d, Epoch: %d, Type: %s, Value: %v}", m.Number, m.Epoch, m.Type, m.Value)
}
