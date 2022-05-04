// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type MapNull struct {
	value  Map
	filled bool
}

func NewMapNull(value map[string]interface{}) MapNull {
	return MapNull{
		value:  NewMap(value),
		filled: true,
	}
}

func (parameter MapNull) Get() (Map, bool) {
	return parameter.value, parameter.filled
}

func (parameter MapNull) Filled() bool {
	return parameter.filled
}

func (parameter MapNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *MapNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
