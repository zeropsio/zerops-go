// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type UserTokenId types.UuidShort

func NewUserTokenIdFromString(value string) (out UserTokenId, err error) {
	return UserTokenId(value), nil
}

func (parameter UserTokenId) Native() string {
	return string(parameter)
}

func (parameter UserTokenId) TypedString() types.String {
	return types.String(parameter)
}

type UserTokenIdNull struct {
	value  UserTokenId
	filled bool
}

func (parameter UserTokenId) UserTokenIdNull() UserTokenIdNull {
	return NewUserTokenIdNull(parameter)
}

func NewUserTokenIdNull(value UserTokenId) UserTokenIdNull {
	return UserTokenIdNull{
		filled: true,
		value:  value,
	}
}

func NewUserTokenIdNullFromString(value string) UserTokenIdNull {
	return UserTokenIdNull{
		filled: true,
		value:  UserTokenId(value),
	}
}

func (parameter UserTokenIdNull) Get() (UserTokenId, bool) {
	return parameter.value, parameter.filled
}

func (parameter UserTokenIdNull) Filled() bool {
	return parameter.filled
}

func (parameter UserTokenIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *UserTokenIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
