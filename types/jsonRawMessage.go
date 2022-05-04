// Generated ZEROPS sdk

package types

import (
	"encoding/json"
	"errors"
)

type JsonRawMessage []byte

func NewJsonRawMessage(value string) JsonRawMessage {
	return JsonRawMessage(value)
}

func (parameter JsonRawMessage) Native() []byte {
	return parameter
}

func (parameter JsonRawMessage) MarshalJSON() ([]byte, error) {
	if parameter == nil {
		return []byte("null"), nil
	}
	return parameter, nil
}

func (parameter *JsonRawMessage) UnmarshalJSON(data []byte) error {
	if parameter == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*parameter = append((*parameter)[0:0], data...)
	return nil
}

var _ json.Marshaler = (*JsonRawMessage)(nil)
var _ json.Unmarshaler = (*JsonRawMessage)(nil)
