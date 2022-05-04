// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ClientUserId types.UuidShort

func NewClientUserIdFromString(value string) (out ClientUserId, err error) {
	return ClientUserId(value), nil
}

func (parameter ClientUserId) Native() string {
	return string(parameter)
}

func (parameter ClientUserId) TypedString() types.String {
	return types.String(parameter)
}

type ClientUserIdNull struct {
	value  ClientUserId
	filled bool
}

func (parameter ClientUserId) ClientUserIdNull() ClientUserIdNull {
	return NewClientUserIdNull(parameter)
}

func NewClientUserIdNull(value ClientUserId) ClientUserIdNull {
	return ClientUserIdNull{
		filled: true,
		value:  value,
	}
}

func NewClientUserIdNullFromString(value string) ClientUserIdNull {
	return ClientUserIdNull{
		filled: true,
		value:  ClientUserId(value),
	}
}

func (parameter ClientUserIdNull) Get() (ClientUserId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ClientUserIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ClientUserIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ClientUserIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
