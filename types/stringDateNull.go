// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type StringDateNull struct {
	value  StringDate
	filled bool
}

func NewStringDateNull(value string) StringDateNull {
	return StringDateNull{
		value:  NewStringDate(value),
		filled: true,
	}
}

func NewStringDateNullFromString(value string) (StringDateNull, error) {
	val, err := NewStringDateFromString(value)
	if err != nil {
		return StringDateNull{}, err
	}
	return StringDateNull{
		value:  NewStringDate(val.Native()),
		filled: true,
	}, nil
}

func (parameter StringDateNull) Get() (StringDate, bool) {
	return parameter.value, parameter.filled
}

func (parameter StringDateNull) Filled() bool {
	return parameter.filled
}

func (parameter StringDateNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *StringDateNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
