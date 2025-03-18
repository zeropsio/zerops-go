// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type EnvId types.UuidShort

func NewEnvIdFromString(value string) (out EnvId, err error) {
	return EnvId(value), nil
}

func (parameter EnvId) Native() string {
	return string(parameter)
}

func (parameter EnvId) TypedString() types.String {
	return types.String(parameter)
}

type EnvIdNull struct {
	value  EnvId
	filled bool
}

func (parameter EnvId) EnvIdNull() EnvIdNull {
	return NewEnvIdNull(parameter)
}

func NewEnvIdNull(value EnvId) EnvIdNull {
	return EnvIdNull{
		filled: true,
		value:  value,
	}
}

func NewEnvIdNullFromString(value string) EnvIdNull {
	return EnvIdNull{
		filled: true,
		value:  EnvId(value),
	}
}

func (parameter EnvIdNull) Get() (EnvId, bool) {
	return parameter.value, parameter.filled
}

func (parameter EnvIdNull) Filled() bool {
	return parameter.filled
}

func (parameter EnvIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *EnvIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
