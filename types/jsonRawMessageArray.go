// Generated ZEROPS sdk

package types

import (
	"encoding/json"
	"errors"
)

type JsonRawMessageArray []json.RawMessage

func NewJsonRawMessageArray(value []json.RawMessage) JsonRawMessageArray {
	return JsonRawMessageArray(value)
}

func (parameter JsonRawMessageArray) Native() []json.RawMessage {
	return []json.RawMessage(parameter)
}

func (parameter JsonRawMessageArray) MarshalJSON() ([]byte, error) {
	if parameter == nil {
		return nil, errors.New("value can't be nil")
	}
	return json.Marshal([]json.RawMessage(parameter))
}

func (parameter *JsonRawMessageArray) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	if arr == nil {
		arr = make([]json.RawMessage, 0)
	}
	*parameter = JsonRawMessageArray(arr)
	return nil
}

var _ json.Marshaler = (*JsonRawMessageArray)(nil)
var _ json.Unmarshaler = (*JsonRawMessageArray)(nil)
