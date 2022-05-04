// Generated ZEROPS sdk

package stringId

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type DriverId types.StringId

func (parameter DriverId) DriverIdNull() DriverIdNull {
	return NewDriverIdNull(parameter)
}

func NewDriverIdNullFromString(value string) DriverIdNull {
	return DriverIdNull{
		filled: true,
		value:  DriverId(value),
	}
}

func NewDriverIdNull(value DriverId) DriverIdNull {
	return DriverIdNull{
		filled: true,
		value:  value,
	}
}

func NewDriverIdFromString(value string) (out DriverId, err error) {
	return DriverId(value), nil
}

func (parameter DriverId) Native() string {
	return string(parameter)
}

type DriverIdNull struct {
	value  DriverId
	filled bool
}

func (parameter DriverIdNull) Get() (DriverId, bool) {
	return parameter.value, parameter.filled
}

func (parameter DriverIdNull) Filled() bool {
	return parameter.filled
}

func (parameter DriverIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *DriverIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
