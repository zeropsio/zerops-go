// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type MediumTextNull struct {
	value  MediumText
	filled bool
}

func NewMediumTextNull(value string) MediumTextNull {
	return MediumTextNull{
		value:  NewMediumText(value),
		filled: true,
	}
}

func (parameter MediumTextNull) Get() (MediumText, bool) {
	return parameter.value, parameter.filled
}

func (parameter MediumTextNull) Filled() bool {
	return parameter.filled
}

func (parameter MediumTextNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *MediumTextNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
