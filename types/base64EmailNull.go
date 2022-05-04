// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type Base64EmailNull struct {
	value  Base64Email
	filled bool
}

func NewBase64EmailNull(value string) Base64EmailNull {
	return Base64EmailNull{
		value:  NewBase64Email(value),
		filled: true,
	}
}

func (parameter Base64EmailNull) Get() (Base64Email, bool) {
	return parameter.value, parameter.filled
}

func (parameter Base64EmailNull) Filled() bool {
	return parameter.filled
}

func (parameter Base64EmailNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *Base64EmailNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
