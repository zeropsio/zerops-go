// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ServiceStackId types.UuidShort

func NewServiceStackIdFromString(value string) (out ServiceStackId, err error) {
	return ServiceStackId(value), nil
}

func (parameter ServiceStackId) Native() string {
	return string(parameter)
}

func (parameter ServiceStackId) TypedString() types.String {
	return types.String(parameter)
}

type ServiceStackIdNull struct {
	value  ServiceStackId
	filled bool
}

func (parameter ServiceStackId) ServiceStackIdNull() ServiceStackIdNull {
	return NewServiceStackIdNull(parameter)
}

func NewServiceStackIdNull(value ServiceStackId) ServiceStackIdNull {
	return ServiceStackIdNull{
		filled: true,
		value:  value,
	}
}

func NewServiceStackIdNullFromString(value string) ServiceStackIdNull {
	return ServiceStackIdNull{
		filled: true,
		value:  ServiceStackId(value),
	}
}

func (parameter ServiceStackIdNull) Get() (ServiceStackId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ServiceStackIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ServiceStackIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ServiceStackIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
