// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type AuthActualUserId types.UuidShort

func NewAuthActualUserIdFromString(value string) (out AuthActualUserId, err error) {
	return AuthActualUserId(value), nil
}

func (parameter AuthActualUserId) Native() string {
	return string(parameter)
}

func (parameter AuthActualUserId) TypedString() types.String {
	return types.String(parameter)
}

type AuthActualUserIdNull struct {
	value  AuthActualUserId
	filled bool
}

func (parameter AuthActualUserId) AuthActualUserIdNull() AuthActualUserIdNull {
	return NewAuthActualUserIdNull(parameter)
}

func NewAuthActualUserIdNull(value AuthActualUserId) AuthActualUserIdNull {
	return AuthActualUserIdNull{
		filled: true,
		value:  value,
	}
}

func NewAuthActualUserIdNullFromString(value string) AuthActualUserIdNull {
	return AuthActualUserIdNull{
		filled: true,
		value:  AuthActualUserId(value),
	}
}

func (parameter AuthActualUserIdNull) Get() (AuthActualUserId, bool) {
	return parameter.value, parameter.filled
}

func (parameter AuthActualUserIdNull) Filled() bool {
	return parameter.filled
}

func (parameter AuthActualUserIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *AuthActualUserIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
