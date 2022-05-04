// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type AuthUserId types.UuidShort

func NewAuthUserIdFromString(value string) (out AuthUserId, err error) {
	return AuthUserId(value), nil
}

func (parameter AuthUserId) Native() string {
	return string(parameter)
}

func (parameter AuthUserId) TypedString() types.String {
	return types.String(parameter)
}

type AuthUserIdNull struct {
	value  AuthUserId
	filled bool
}

func (parameter AuthUserId) AuthUserIdNull() AuthUserIdNull {
	return NewAuthUserIdNull(parameter)
}

func NewAuthUserIdNull(value AuthUserId) AuthUserIdNull {
	return AuthUserIdNull{
		filled: true,
		value:  value,
	}
}

func NewAuthUserIdNullFromString(value string) AuthUserIdNull {
	return AuthUserIdNull{
		filled: true,
		value:  AuthUserId(value),
	}
}

func (parameter AuthUserIdNull) Get() (AuthUserId, bool) {
	return parameter.value, parameter.filled
}

func (parameter AuthUserIdNull) Filled() bool {
	return parameter.filled
}

func (parameter AuthUserIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *AuthUserIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
