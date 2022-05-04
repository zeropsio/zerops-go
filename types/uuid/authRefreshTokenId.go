// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type AuthRefreshTokenId types.UuidShort

func NewAuthRefreshTokenIdFromString(value string) (out AuthRefreshTokenId, err error) {
	return AuthRefreshTokenId(value), nil
}

func (parameter AuthRefreshTokenId) Native() string {
	return string(parameter)
}

func (parameter AuthRefreshTokenId) TypedString() types.String {
	return types.String(parameter)
}

type AuthRefreshTokenIdNull struct {
	value  AuthRefreshTokenId
	filled bool
}

func (parameter AuthRefreshTokenId) AuthRefreshTokenIdNull() AuthRefreshTokenIdNull {
	return NewAuthRefreshTokenIdNull(parameter)
}

func NewAuthRefreshTokenIdNull(value AuthRefreshTokenId) AuthRefreshTokenIdNull {
	return AuthRefreshTokenIdNull{
		filled: true,
		value:  value,
	}
}

func NewAuthRefreshTokenIdNullFromString(value string) AuthRefreshTokenIdNull {
	return AuthRefreshTokenIdNull{
		filled: true,
		value:  AuthRefreshTokenId(value),
	}
}

func (parameter AuthRefreshTokenIdNull) Get() (AuthRefreshTokenId, bool) {
	return parameter.value, parameter.filled
}

func (parameter AuthRefreshTokenIdNull) Filled() bool {
	return parameter.filled
}

func (parameter AuthRefreshTokenIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *AuthRefreshTokenIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
