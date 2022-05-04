// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ClientId types.UuidShort

func NewClientIdFromString(value string) (out ClientId, err error) {
	return ClientId(value), nil
}

func (parameter ClientId) Native() string {
	return string(parameter)
}

func (parameter ClientId) TypedString() types.String {
	return types.String(parameter)
}

type ClientIdNull struct {
	value  ClientId
	filled bool
}

func (parameter ClientId) ClientIdNull() ClientIdNull {
	return NewClientIdNull(parameter)
}

func NewClientIdNull(value ClientId) ClientIdNull {
	return ClientIdNull{
		filled: true,
		value:  value,
	}
}

func NewClientIdNullFromString(value string) ClientIdNull {
	return ClientIdNull{
		filled: true,
		value:  ClientId(value),
	}
}

func (parameter ClientIdNull) Get() (ClientId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ClientIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ClientIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ClientIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
