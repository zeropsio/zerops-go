// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ClientUserValidateClientId types.UuidShort

func NewClientUserValidateClientIdFromString(value string) (out ClientUserValidateClientId, err error) {
	return ClientUserValidateClientId(value), nil
}

func (parameter ClientUserValidateClientId) Native() string {
	return string(parameter)
}

func (parameter ClientUserValidateClientId) TypedString() types.String {
	return types.String(parameter)
}

type ClientUserValidateClientIdNull struct {
	value  ClientUserValidateClientId
	filled bool
}

func (parameter ClientUserValidateClientId) ClientUserValidateClientIdNull() ClientUserValidateClientIdNull {
	return NewClientUserValidateClientIdNull(parameter)
}

func NewClientUserValidateClientIdNull(value ClientUserValidateClientId) ClientUserValidateClientIdNull {
	return ClientUserValidateClientIdNull{
		filled: true,
		value:  value,
	}
}

func NewClientUserValidateClientIdNullFromString(value string) ClientUserValidateClientIdNull {
	return ClientUserValidateClientIdNull{
		filled: true,
		value:  ClientUserValidateClientId(value),
	}
}

func (parameter ClientUserValidateClientIdNull) Get() (ClientUserValidateClientId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ClientUserValidateClientIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ClientUserValidateClientIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ClientUserValidateClientIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
