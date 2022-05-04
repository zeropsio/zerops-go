// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type StringNull struct {
	value  String
	filled bool
}

func NewStringNull(value string) StringNull {
	return StringNull{
		value:  NewString(value),
		filled: true,
	}
}

func NewStringNullFromString(value string) (StringNull, error) {
	val, err := NewStringFromString(value)
	if err != nil {
		return StringNull{}, err
	}
	return StringNull{
		value:  NewString(val.Native()),
		filled: true,
	}, nil
}

func (parameter StringNull) Get() (String, bool) {
	return parameter.value, parameter.filled
}

func (parameter StringNull) Filled() bool {
	return parameter.filled
}

func (parameter StringNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *StringNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
