// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type EmptyStringNull struct {
	value  EmptyString
	filled bool
}

func NewEmptyStringNull(value string) EmptyStringNull {
	return EmptyStringNull{
		value:  NewEmptyString(value),
		filled: true,
	}
}

func NewEmptyStringNullFromString(value string) (EmptyStringNull, error) {
	val, err := NewEmptyStringFromString(value)
	if err != nil {
		return EmptyStringNull{}, err
	}
	return EmptyStringNull{
		value:  NewEmptyString(val.Native()),
		filled: true,
	}, nil
}

func (parameter EmptyStringNull) Get() (EmptyString, bool) {
	return parameter.value, parameter.filled
}

func (parameter EmptyStringNull) Filled() bool {
	return parameter.filled
}

func (parameter EmptyStringNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *EmptyStringNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
