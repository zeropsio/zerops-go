// Generated ZEROPS sdk

package types

import (
	"encoding/json"
	"errors"
)

type JsonRawMessageArray []byte

func NewJsonRawMessageArray(value string) JsonRawMessageArray {
	return JsonRawMessageArray(value)
}

func (parameter JsonRawMessageArray) Native() []byte {
	return parameter
}

func (parameter JsonRawMessageArray) MarshalJSON() ([]byte, error) {
	if parameter == nil {
		return []byte("null"), nil
	}
	return parameter, nil
}

func (parameter *JsonRawMessageArray) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return errors.Wrap(err)
	}
	if arr == nil {
		arr = make([]json.RawMessage, 0)
	}
	*parameter = JsonRawMessageArray(arr)
	return nil
}

var _ json.Marshaler = (*JsonRawMessageArray)(nil)
var _ json.Unmarshaler = (*JsonRawMessageArray)(nil)
