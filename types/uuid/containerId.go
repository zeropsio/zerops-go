// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ContainerId types.UuidShort

func NewContainerIdFromString(value string) (out ContainerId, err error) {
	return ContainerId(value), nil
}

func (parameter ContainerId) Native() string {
	return string(parameter)
}

func (parameter ContainerId) TypedString() types.String {
	return types.String(parameter)
}

type ContainerIdNull struct {
	value  ContainerId
	filled bool
}

func (parameter ContainerId) ContainerIdNull() ContainerIdNull {
	return NewContainerIdNull(parameter)
}

func NewContainerIdNull(value ContainerId) ContainerIdNull {
	return ContainerIdNull{
		filled: true,
		value:  value,
	}
}

func NewContainerIdNullFromString(value string) ContainerIdNull {
	return ContainerIdNull{
		filled: true,
		value:  ContainerId(value),
	}
}

func (parameter ContainerIdNull) Get() (ContainerId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ContainerIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ContainerIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ContainerIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
