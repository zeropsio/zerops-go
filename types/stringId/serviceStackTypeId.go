// Generated ZEROPS sdk

package stringId

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ServiceStackTypeId types.StringId

func (parameter ServiceStackTypeId) ServiceStackTypeIdNull() ServiceStackTypeIdNull {
	return NewServiceStackTypeIdNull(parameter)
}

func NewServiceStackTypeIdNullFromString(value string) ServiceStackTypeIdNull {
	return ServiceStackTypeIdNull{
		filled: true,
		value:  ServiceStackTypeId(value),
	}
}

func NewServiceStackTypeIdNull(value ServiceStackTypeId) ServiceStackTypeIdNull {
	return ServiceStackTypeIdNull{
		filled: true,
		value:  value,
	}
}

func NewServiceStackTypeIdFromString(value string) (out ServiceStackTypeId, err error) {
	return ServiceStackTypeId(value), nil
}

func (parameter ServiceStackTypeId) Native() string {
	return string(parameter)
}

type ServiceStackTypeIdNull struct {
	value  ServiceStackTypeId
	filled bool
}

func (parameter ServiceStackTypeIdNull) Get() (ServiceStackTypeId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ServiceStackTypeIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ServiceStackTypeIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ServiceStackTypeIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
