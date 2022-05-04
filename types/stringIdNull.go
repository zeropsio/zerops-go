// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type StringIdNull struct {
	value  StringId
	filled bool
}

func NewStringIdNull(value string) StringIdNull {
	return StringIdNull{
		value:  NewStringId(value),
		filled: true,
	}
}

func NewStringIdNullFromString(value string) (StringIdNull, error) {
	val, err := NewStringIdFromString(value)
	if err != nil {
		return StringIdNull{}, err
	}
	return StringIdNull{
		value:  NewStringId(val.Native()),
		filled: true,
	}, nil
}

func (parameter StringIdNull) Get() (StringId, bool) {
	return parameter.value, parameter.filled
}

func (parameter StringIdNull) Filled() bool {
	return parameter.filled
}

func (parameter StringIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *StringIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
