// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type GenericUuidId types.UuidShort

func NewGenericUuidIdFromString(value string) (out GenericUuidId, err error) {
	return GenericUuidId(value), nil
}

func (parameter GenericUuidId) Native() string {
	return string(parameter)
}

func (parameter GenericUuidId) TypedString() types.String {
	return types.String(parameter)
}

type GenericUuidIdNull struct {
	value  GenericUuidId
	filled bool
}

func (parameter GenericUuidId) GenericUuidIdNull() GenericUuidIdNull {
	return NewGenericUuidIdNull(parameter)
}

func NewGenericUuidIdNull(value GenericUuidId) GenericUuidIdNull {
	return GenericUuidIdNull{
		filled: true,
		value:  value,
	}
}

func NewGenericUuidIdNullFromString(value string) GenericUuidIdNull {
	return GenericUuidIdNull{
		filled: true,
		value:  GenericUuidId(value),
	}
}

func (parameter GenericUuidIdNull) Get() (GenericUuidId, bool) {
	return parameter.value, parameter.filled
}

func (parameter GenericUuidIdNull) Filled() bool {
	return parameter.filled
}

func (parameter GenericUuidIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *GenericUuidIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
