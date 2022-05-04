// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ServiceId types.UuidShort

func NewServiceIdFromString(value string) (out ServiceId, err error) {
	return ServiceId(value), nil
}

func (parameter ServiceId) Native() string {
	return string(parameter)
}

func (parameter ServiceId) TypedString() types.String {
	return types.String(parameter)
}

type ServiceIdNull struct {
	value  ServiceId
	filled bool
}

func (parameter ServiceId) ServiceIdNull() ServiceIdNull {
	return NewServiceIdNull(parameter)
}

func NewServiceIdNull(value ServiceId) ServiceIdNull {
	return ServiceIdNull{
		filled: true,
		value:  value,
	}
}

func NewServiceIdNullFromString(value string) ServiceIdNull {
	return ServiceIdNull{
		filled: true,
		value:  ServiceId(value),
	}
}

func (parameter ServiceIdNull) Get() (ServiceId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ServiceIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ServiceIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ServiceIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
