// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type UserPutPasswordViaTokenToken types.UuidShort

func NewUserPutPasswordViaTokenTokenFromString(value string) (out UserPutPasswordViaTokenToken, err error) {
	return UserPutPasswordViaTokenToken(value), nil
}

func (parameter UserPutPasswordViaTokenToken) Native() string {
	return string(parameter)
}

func (parameter UserPutPasswordViaTokenToken) TypedString() types.String {
	return types.String(parameter)
}

type UserPutPasswordViaTokenTokenNull struct {
	value  UserPutPasswordViaTokenToken
	filled bool
}

func (parameter UserPutPasswordViaTokenToken) UserPutPasswordViaTokenTokenNull() UserPutPasswordViaTokenTokenNull {
	return NewUserPutPasswordViaTokenTokenNull(parameter)
}

func NewUserPutPasswordViaTokenTokenNull(value UserPutPasswordViaTokenToken) UserPutPasswordViaTokenTokenNull {
	return UserPutPasswordViaTokenTokenNull{
		filled: true,
		value:  value,
	}
}

func NewUserPutPasswordViaTokenTokenNullFromString(value string) UserPutPasswordViaTokenTokenNull {
	return UserPutPasswordViaTokenTokenNull{
		filled: true,
		value:  UserPutPasswordViaTokenToken(value),
	}
}

func (parameter UserPutPasswordViaTokenTokenNull) Get() (UserPutPasswordViaTokenToken, bool) {
	return parameter.value, parameter.filled
}

func (parameter UserPutPasswordViaTokenTokenNull) Filled() bool {
	return parameter.filled
}

func (parameter UserPutPasswordViaTokenTokenNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *UserPutPasswordViaTokenTokenNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
