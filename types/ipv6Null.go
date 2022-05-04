// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type Ipv6Null struct {
	value  Ipv6
	filled bool
}

func NewIpv6Null(value string) Ipv6Null {
	return Ipv6Null{
		value:  NewIpv6(value),
		filled: true,
	}
}

func (parameter Ipv6Null) Get() (Ipv6, bool) {
	return parameter.value, parameter.filled
}

func (parameter Ipv6Null) Filled() bool {
	return parameter.filled
}

func (parameter Ipv6Null) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *Ipv6Null) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
