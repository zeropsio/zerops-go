// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type ObjectNull struct {
	value  Object
	filled bool
}

func NewObjectNull(value string) ObjectNull {
	return ObjectNull{
		value:  NewObject(value),
		filled: true,
	}
}

func (parameter ObjectNull) Get() (Object, bool) {
	return parameter.value, parameter.filled
}

func (parameter ObjectNull) Filled() bool {
	return parameter.filled
}

func (parameter ObjectNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *ObjectNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
