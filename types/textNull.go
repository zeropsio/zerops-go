// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type TextNull struct {
	value  Text
	filled bool
}

func NewTextNull(value string) TextNull {
	return TextNull{
		value:  NewText(value),
		filled: true,
	}
}

func (parameter TextNull) Get() (Text, bool) {
	return parameter.value, parameter.filled
}

func (parameter TextNull) Filled() bool {
	return parameter.filled
}

func (parameter TextNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *TextNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
