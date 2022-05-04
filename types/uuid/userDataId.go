// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type UserDataId types.UuidShort

func NewUserDataIdFromString(value string) (out UserDataId, err error) {
	return UserDataId(value), nil
}

func (parameter UserDataId) Native() string {
	return string(parameter)
}

func (parameter UserDataId) TypedString() types.String {
	return types.String(parameter)
}

type UserDataIdNull struct {
	value  UserDataId
	filled bool
}

func (parameter UserDataId) UserDataIdNull() UserDataIdNull {
	return NewUserDataIdNull(parameter)
}

func NewUserDataIdNull(value UserDataId) UserDataIdNull {
	return UserDataIdNull{
		filled: true,
		value:  value,
	}
}

func NewUserDataIdNullFromString(value string) UserDataIdNull {
	return UserDataIdNull{
		filled: true,
		value:  UserDataId(value),
	}
}

func (parameter UserDataIdNull) Get() (UserDataId, bool) {
	return parameter.value, parameter.filled
}

func (parameter UserDataIdNull) Filled() bool {
	return parameter.filled
}

func (parameter UserDataIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *UserDataIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
