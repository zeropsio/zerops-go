// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type BoolNull struct {
	value  Bool
	filled bool
}

func NewBoolNull(value bool) BoolNull {
	return BoolNull{
		value:  NewBool(value),
		filled: true,
	}
}

func NewBoolNullFromString(value string) (BoolNull, error) {
	val, err := NewBoolFromString(value)
	if err != nil {
		return BoolNull{}, err
	}
	return BoolNull{
		value:  NewBool(val.Native()),
		filled: true,
	}, nil
}

func (parameter BoolNull) Get() (Bool, bool) {
	return parameter.value, parameter.filled
}

func (parameter BoolNull) Filled() bool {
	return parameter.filled
}

func (parameter BoolNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *BoolNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
