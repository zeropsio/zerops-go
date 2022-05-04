// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type FloatNull struct {
	value  Float
	filled bool
}

func NewFloatNull(value float64) FloatNull {
	return FloatNull{
		value:  NewFloat(value),
		filled: true,
	}
}

func (parameter FloatNull) Get() (Float, bool) {
	return parameter.value, parameter.filled
}

func (parameter FloatNull) Filled() bool {
	return parameter.filled
}

func (parameter FloatNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *FloatNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
