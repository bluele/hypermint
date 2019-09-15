package event

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/bluele/hypermint/pkg/abci/types"
	"github.com/ethereum/go-ethereum/common"
	tmcmn "github.com/tendermint/tendermint/libs/common"
)

type Event struct {
	address common.Address
	entries []*Entry
}

func NewEvent(address common.Address, entries []*Entry) *Event {
	return &Event{
		address: address,
		entries: entries,
	}
}

func (es Event) Address() common.Address {
	return es.address
}

func (es Event) Entries() []*Entry {
	return es.entries
}

type Entry struct {
	Name  []byte
	Value []byte
}

func (e Entry) String() string {
	return fmt.Sprintf("%v{0x%X}", string(e.Name), e.Value)
}

func (e Entry) Validate() error {
	if len(e.Name) == 0 {
		return fmt.Errorf("length of Name must be greater than 0")
	}
	if len(e.Name) > 32 {
		return fmt.Errorf("length of Name must be not greater than 32")
	}
	if len(e.Value) == 0 {
		return fmt.Errorf("length of Value must be greater than 0")
	}
	if len(e.Value) > 1024 {
		return fmt.Errorf("length of Value must be not greater than 1024")
	}
	return nil
}

func (e Entry) Bytes() []byte {
	var buf bytes.Buffer
	buf.WriteByte(byte(len(e.Name)))
	buf.Write(e.Name)
	buf.Write(e.Value)
	return []byte(hex.EncodeToString(buf.Bytes()))
}

func MakeTMEvent(contractAddr common.Address, es []*Entry) (*types.Event, error) {
	pairs, err := eventsToPairs(es)
	if err != nil {
		return nil, err
	}
	pairs = append(pairs, tmcmn.KVPair{
		Key:   []byte("address"),
		Value: []byte(contractAddr.Hex()),
	})
	e := types.Event{Type: "contract"}
	e.Attributes = pairs
	return &e, nil
}

func eventsToPairs(es []*Entry) (tmcmn.KVPairs, error) {
	var pairs tmcmn.KVPairs
	for _, e := range es {
		if err := e.Validate(); err != nil {
			return nil, err
		}
		key := []byte("event.name")
		pairs = append(pairs, tmcmn.KVPair{Key: key, Value: e.Name})
		dataKey := []byte("event.data")
		pairs = append(pairs, tmcmn.KVPair{Key: dataKey, Value: e.Bytes()})
	}
	return pairs, nil
}

func ParseEntry(hexStrBytes []byte) (*Entry, error) {
	b, err := hex.DecodeString(string(hexStrBytes))
	if err != nil {
		return nil, err
	}

	size := b[0]
	name := b[1 : 1+size]
	value := b[1+size : len(b)]

	return &Entry{Name: name, Value: value}, nil
}

func MakeEntry(name, value string) (*Entry, error) {
	var v []byte
	if strings.Contains(value, "0x") {
		h, err := hex.DecodeString(value[2:])
		if err != nil {
			return nil, err
		}
		v = h
	} else {
		v = []byte(value)
	}
	return &Entry{Name: []byte(name), Value: v}, nil
}

func MakeEntryBytes(name, value string) ([]byte, error) {
	e, err := MakeEntry(name, value)
	if err != nil {
		return nil, err
	}
	return e.Bytes(), nil
}

// MakeEventSearchQuery returns a query for searching events on transaction
// contractAddr: target contract address
// eventName: event name
// eventValue: value corresponding to event name. NOTE: if value has a prefix "0x", value will be decoded into []byte
func MakeEventSearchQuery(contractAddr string, eventName, eventValue string) (string, error) {
	var parts []string

	parts = append(parts, fmt.Sprintf("contract.address='%v'", contractAddr))
	parts = append(parts, fmt.Sprintf("contract.event.name='%v'", eventName))
	if len(eventValue) > 0 {
		ev, err := MakeEntryBytes(eventName, eventValue)
		if err != nil {
			return "", err
		}
		parts = append(parts, fmt.Sprintf("contract.event.data='%v'", string(ev)))
	}
	return strings.Join(parts, " AND "), nil
}