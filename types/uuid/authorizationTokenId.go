// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type AuthorizationTokenId types.UuidShort

func NewAuthorizationTokenIdFromString(value string) (out AuthorizationTokenId, err error) {
	return AuthorizationTokenId(value), nil
}

func (parameter AuthorizationTokenId) Native() string {
	return string(parameter)
}

func (parameter AuthorizationTokenId) TypedString() types.String {
	return types.String(parameter)
}

type AuthorizationTokenIdNull struct {
	value  AuthorizationTokenId
	filled bool
}

func (parameter AuthorizationTokenId) AuthorizationTokenIdNull() AuthorizationTokenIdNull {
	return NewAuthorizationTokenIdNull(parameter)
}

func NewAuthorizationTokenIdNull(value AuthorizationTokenId) AuthorizationTokenIdNull {
	return AuthorizationTokenIdNull{
		filled: true,
		value:  value,
	}
}

func NewAuthorizationTokenIdNullFromString(value string) AuthorizationTokenIdNull {
	return AuthorizationTokenIdNull{
		filled: true,
		value:  AuthorizationTokenId(value),
	}
}

func (parameter AuthorizationTokenIdNull) Get() (AuthorizationTokenId, bool) {
	return parameter.value, parameter.filled
}

func (parameter AuthorizationTokenIdNull) Filled() bool {
	return parameter.filled
}

func (parameter AuthorizationTokenIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *AuthorizationTokenIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
