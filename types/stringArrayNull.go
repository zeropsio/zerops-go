// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type StringArrayNull struct {
	value  StringArray
	filled bool
}

func NewStringArrayNull(value []string) StringArrayNull {
	return StringArrayNull{
		value:  NewStringArray(value),
		filled: true,
	}
}

func (parameter StringArrayNull) Get() (StringArray, bool) {
	return parameter.value, parameter.filled
}

func (parameter StringArrayNull) Filled() bool {
	return parameter.filled
}

func (parameter StringArrayNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *StringArrayNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
