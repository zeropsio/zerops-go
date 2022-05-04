// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type IntNull struct {
	value  Int
	filled bool
}

func NewIntNull(value int) IntNull {
	return IntNull{
		value:  NewInt(value),
		filled: true,
	}
}

func NewIntNullFromString(value string) (IntNull, error) {
	val, err := NewIntFromString(value)
	if err != nil {
		return IntNull{}, err
	}
	return IntNull{
		value:  NewInt(val.Native()),
		filled: true,
	}, nil
}

func (parameter IntNull) Get() (Int, bool) {
	return parameter.value, parameter.filled
}

func (parameter IntNull) Filled() bool {
	return parameter.filled
}

func (parameter IntNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *IntNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
