// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type FloatArrayNull struct {
	value  FloatArray
	filled bool
}

func NewFloatArrayNull(value []float64) FloatArrayNull {
	return FloatArrayNull{
		value:  NewFloatArray(value),
		filled: true,
	}
}

func (parameter FloatArrayNull) Get() (FloatArray, bool) {
	return parameter.value, parameter.filled
}

func (parameter FloatArrayNull) Filled() bool {
	return parameter.filled
}

func (parameter FloatArrayNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *FloatArrayNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
