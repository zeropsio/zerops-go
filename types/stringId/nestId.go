// Generated ZEROPS sdk

package stringId

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type NestId types.StringId

func (parameter NestId) NestIdNull() NestIdNull {
	return NewNestIdNull(parameter)
}

func NewNestIdNullFromString(value string) NestIdNull {
	return NestIdNull{
		filled: true,
		value:  NestId(value),
	}
}

func NewNestIdNull(value NestId) NestIdNull {
	return NestIdNull{
		filled: true,
		value:  value,
	}
}

func NewNestIdFromString(value string) (out NestId, err error) {
	return NestId(value), nil
}

func (parameter NestId) Native() string {
	return string(parameter)
}

type NestIdNull struct {
	value  NestId
	filled bool
}

func (parameter NestIdNull) Get() (NestId, bool) {
	return parameter.value, parameter.filled
}

func (parameter NestIdNull) Filled() bool {
	return parameter.filled
}

func (parameter NestIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *NestIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
