// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type UuidShortNull struct {
	value  UuidShort
	filled bool
}

func NewUuidShortNull(value string) UuidShortNull {
	return UuidShortNull{
		value:  NewUuidShort(value),
		filled: true,
	}
}

func NewUuidShortNullFromString(value string) (UuidShortNull, error) {
	val, err := NewUuidShortFromString(value)
	if err != nil {
		return UuidShortNull{}, err
	}
	return UuidShortNull{
		value:  NewUuidShort(val.Native()),
		filled: true,
	}, nil
}

func (parameter UuidShortNull) Get() (UuidShort, bool) {
	return parameter.value, parameter.filled
}

func (parameter UuidShortNull) Filled() bool {
	return parameter.filled
}

func (parameter UuidShortNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *UuidShortNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
