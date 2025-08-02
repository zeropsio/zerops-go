// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type InstanceId types.UuidShort

func NewInstanceIdFromString(value string) (out InstanceId, err error) {
	return InstanceId(value), nil
}

func (parameter InstanceId) Native() string {
	return string(parameter)
}

func (parameter InstanceId) TypedString() types.String {
	return types.String(parameter)
}

type InstanceIdNull struct {
	value  InstanceId
	filled bool
}

func (parameter InstanceId) InstanceIdNull() InstanceIdNull {
	return NewInstanceIdNull(parameter)
}

func NewInstanceIdNull(value InstanceId) InstanceIdNull {
	return InstanceIdNull{
		filled: true,
		value:  value,
	}
}

func NewInstanceIdNullFromString(value string) InstanceIdNull {
	return InstanceIdNull{
		filled: true,
		value:  InstanceId(value),
	}
}

func (parameter InstanceIdNull) Get() (InstanceId, bool) {
	return parameter.value, parameter.filled
}

func (parameter InstanceIdNull) Filled() bool {
	return parameter.filled
}

func (parameter InstanceIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *InstanceIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
