// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ServiceStackConnectionId types.UuidShort

func NewServiceStackConnectionIdFromString(value string) (out ServiceStackConnectionId, err error) {
	return ServiceStackConnectionId(value), nil
}

func (parameter ServiceStackConnectionId) Native() string {
	return string(parameter)
}

func (parameter ServiceStackConnectionId) TypedString() types.String {
	return types.String(parameter)
}

type ServiceStackConnectionIdNull struct {
	value  ServiceStackConnectionId
	filled bool
}

func (parameter ServiceStackConnectionId) ServiceStackConnectionIdNull() ServiceStackConnectionIdNull {
	return NewServiceStackConnectionIdNull(parameter)
}

func NewServiceStackConnectionIdNull(value ServiceStackConnectionId) ServiceStackConnectionIdNull {
	return ServiceStackConnectionIdNull{
		filled: true,
		value:  value,
	}
}

func NewServiceStackConnectionIdNullFromString(value string) ServiceStackConnectionIdNull {
	return ServiceStackConnectionIdNull{
		filled: true,
		value:  ServiceStackConnectionId(value),
	}
}

func (parameter ServiceStackConnectionIdNull) Get() (ServiceStackConnectionId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ServiceStackConnectionIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ServiceStackConnectionIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ServiceStackConnectionIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
