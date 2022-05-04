// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type EmailNull struct {
	value  Email
	filled bool
}

func NewEmailNull(value string) EmailNull {
	return EmailNull{
		value:  NewEmail(value),
		filled: true,
	}
}

func (parameter EmailNull) Get() (Email, bool) {
	return parameter.value, parameter.filled
}

func (parameter EmailNull) Filled() bool {
	return parameter.filled
}

func (parameter EmailNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *EmailNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
