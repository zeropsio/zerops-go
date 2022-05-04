// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type UserId types.UuidShort

func NewUserIdFromString(value string) (out UserId, err error) {
	return UserId(value), nil
}

func (parameter UserId) Native() string {
	return string(parameter)
}

func (parameter UserId) TypedString() types.String {
	return types.String(parameter)
}

type UserIdNull struct {
	value  UserId
	filled bool
}

func (parameter UserId) UserIdNull() UserIdNull {
	return NewUserIdNull(parameter)
}

func NewUserIdNull(value UserId) UserIdNull {
	return UserIdNull{
		filled: true,
		value:  value,
	}
}

func NewUserIdNullFromString(value string) UserIdNull {
	return UserIdNull{
		filled: true,
		value:  UserId(value),
	}
}

func (parameter UserIdNull) Get() (UserId, bool) {
	return parameter.value, parameter.filled
}

func (parameter UserIdNull) Filled() bool {
	return parameter.filled
}

func (parameter UserIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *UserIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
