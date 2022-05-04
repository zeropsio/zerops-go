// Generated ZEROPS sdk

package stringId

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ServiceStackTypeVersionId types.StringId

func (parameter ServiceStackTypeVersionId) ServiceStackTypeVersionIdNull() ServiceStackTypeVersionIdNull {
	return NewServiceStackTypeVersionIdNull(parameter)
}

func NewServiceStackTypeVersionIdNullFromString(value string) ServiceStackTypeVersionIdNull {
	return ServiceStackTypeVersionIdNull{
		filled: true,
		value:  ServiceStackTypeVersionId(value),
	}
}

func NewServiceStackTypeVersionIdNull(value ServiceStackTypeVersionId) ServiceStackTypeVersionIdNull {
	return ServiceStackTypeVersionIdNull{
		filled: true,
		value:  value,
	}
}

func NewServiceStackTypeVersionIdFromString(value string) (out ServiceStackTypeVersionId, err error) {
	return ServiceStackTypeVersionId(value), nil
}

func (parameter ServiceStackTypeVersionId) Native() string {
	return string(parameter)
}

type ServiceStackTypeVersionIdNull struct {
	value  ServiceStackTypeVersionId
	filled bool
}

func (parameter ServiceStackTypeVersionIdNull) Get() (ServiceStackTypeVersionId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ServiceStackTypeVersionIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ServiceStackTypeVersionIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ServiceStackTypeVersionIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
