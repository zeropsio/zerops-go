// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type Ipv4Null struct {
	value  Ipv4
	filled bool
}

func NewIpv4Null(value string) Ipv4Null {
	return Ipv4Null{
		value:  NewIpv4(value),
		filled: true,
	}
}

func (parameter Ipv4Null) Get() (Ipv4, bool) {
	return parameter.value, parameter.filled
}

func (parameter Ipv4Null) Filled() bool {
	return parameter.filled
}

func (parameter Ipv4Null) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *Ipv4Null) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
