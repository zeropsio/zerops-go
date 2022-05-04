// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type GenericId types.UuidShort

func NewGenericIdFromString(value string) (out GenericId, err error) {
	return GenericId(value), nil
}

func (parameter GenericId) Native() string {
	return string(parameter)
}

func (parameter GenericId) TypedString() types.String {
	return types.String(parameter)
}

type GenericIdNull struct {
	value  GenericId
	filled bool
}

func (parameter GenericId) GenericIdNull() GenericIdNull {
	return NewGenericIdNull(parameter)
}

func NewGenericIdNull(value GenericId) GenericIdNull {
	return GenericIdNull{
		filled: true,
		value:  value,
	}
}

func NewGenericIdNullFromString(value string) GenericIdNull {
	return GenericIdNull{
		filled: true,
		value:  GenericId(value),
	}
}

func (parameter GenericIdNull) Get() (GenericId, bool) {
	return parameter.value, parameter.filled
}

func (parameter GenericIdNull) Filled() bool {
	return parameter.filled
}

func (parameter GenericIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *GenericIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
