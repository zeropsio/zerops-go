// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type UrlEncodedBase64EmailNull struct {
	value  UrlEncodedBase64Email
	filled bool
}

func NewUrlEncodedBase64EmailNull(value string) UrlEncodedBase64EmailNull {
	return UrlEncodedBase64EmailNull{
		value:  NewUrlEncodedBase64Email(value),
		filled: true,
	}
}

func (parameter UrlEncodedBase64EmailNull) Get() (UrlEncodedBase64Email, bool) {
	return parameter.value, parameter.filled
}

func (parameter UrlEncodedBase64EmailNull) Filled() bool {
	return parameter.filled
}

func (parameter UrlEncodedBase64EmailNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *UrlEncodedBase64EmailNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
