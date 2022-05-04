// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type UserNotificationId types.UuidShort

func NewUserNotificationIdFromString(value string) (out UserNotificationId, err error) {
	return UserNotificationId(value), nil
}

func (parameter UserNotificationId) Native() string {
	return string(parameter)
}

func (parameter UserNotificationId) TypedString() types.String {
	return types.String(parameter)
}

type UserNotificationIdNull struct {
	value  UserNotificationId
	filled bool
}

func (parameter UserNotificationId) UserNotificationIdNull() UserNotificationIdNull {
	return NewUserNotificationIdNull(parameter)
}

func NewUserNotificationIdNull(value UserNotificationId) UserNotificationIdNull {
	return UserNotificationIdNull{
		filled: true,
		value:  value,
	}
}

func NewUserNotificationIdNullFromString(value string) UserNotificationIdNull {
	return UserNotificationIdNull{
		filled: true,
		value:  UserNotificationId(value),
	}
}

func (parameter UserNotificationIdNull) Get() (UserNotificationId, bool) {
	return parameter.value, parameter.filled
}

func (parameter UserNotificationIdNull) Filled() bool {
	return parameter.filled
}

func (parameter UserNotificationIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *UserNotificationIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
