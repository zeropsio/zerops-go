// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type Int64Null struct {
	value  Int64
	filled bool
}

func NewInt64Null(value int64) Int64Null {
	return Int64Null{
		value:  NewInt64(value),
		filled: true,
	}
}

func NewInt64NullFromString(value string) (Int64Null, error) {
	val, err := NewInt64FromString(value)
	if err != nil {
		return Int64Null{}, err
	}
	return Int64Null{
		value:  NewInt64(val.Native()),
		filled: true,
	}, nil
}

func (parameter Int64Null) Get() (Int64, bool) {
	return parameter.value, parameter.filled
}

func (parameter Int64Null) Filled() bool {
	return parameter.filled
}

func (parameter Int64Null) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *Int64Null) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
