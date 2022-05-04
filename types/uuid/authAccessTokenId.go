// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type AuthAccessTokenId types.UuidShort

func NewAuthAccessTokenIdFromString(value string) (out AuthAccessTokenId, err error) {
	return AuthAccessTokenId(value), nil
}

func (parameter AuthAccessTokenId) Native() string {
	return string(parameter)
}

func (parameter AuthAccessTokenId) TypedString() types.String {
	return types.String(parameter)
}

type AuthAccessTokenIdNull struct {
	value  AuthAccessTokenId
	filled bool
}

func (parameter AuthAccessTokenId) AuthAccessTokenIdNull() AuthAccessTokenIdNull {
	return NewAuthAccessTokenIdNull(parameter)
}

func NewAuthAccessTokenIdNull(value AuthAccessTokenId) AuthAccessTokenIdNull {
	return AuthAccessTokenIdNull{
		filled: true,
		value:  value,
	}
}

func NewAuthAccessTokenIdNullFromString(value string) AuthAccessTokenIdNull {
	return AuthAccessTokenIdNull{
		filled: true,
		value:  AuthAccessTokenId(value),
	}
}

func (parameter AuthAccessTokenIdNull) Get() (AuthAccessTokenId, bool) {
	return parameter.value, parameter.filled
}

func (parameter AuthAccessTokenIdNull) Filled() bool {
	return parameter.filled
}

func (parameter AuthAccessTokenIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *AuthAccessTokenIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
