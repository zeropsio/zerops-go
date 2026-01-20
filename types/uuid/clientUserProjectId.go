// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ClientUserProjectId types.UuidShort

func NewClientUserProjectIdFromString(value string) (out ClientUserProjectId, err error) {
	return ClientUserProjectId(value), nil
}

func (parameter ClientUserProjectId) Native() string {
	return string(parameter)
}

func (parameter ClientUserProjectId) TypedString() types.String {
	return types.String(parameter)
}

type ClientUserProjectIdNull struct {
	value  ClientUserProjectId
	filled bool
}

func (parameter ClientUserProjectId) ClientUserProjectIdNull() ClientUserProjectIdNull {
	return NewClientUserProjectIdNull(parameter)
}

func NewClientUserProjectIdNull(value ClientUserProjectId) ClientUserProjectIdNull {
	return ClientUserProjectIdNull{
		filled: true,
		value:  value,
	}
}

func NewClientUserProjectIdNullFromString(value string) ClientUserProjectIdNull {
	return ClientUserProjectIdNull{
		filled: true,
		value:  ClientUserProjectId(value),
	}
}

func (parameter ClientUserProjectIdNull) Get() (ClientUserProjectId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ClientUserProjectIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ClientUserProjectIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ClientUserProjectIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
