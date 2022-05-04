// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type AccessTokenKey types.UuidShort

func NewAccessTokenKeyFromString(value string) (out AccessTokenKey, err error) {
	return AccessTokenKey(value), nil
}

func (parameter AccessTokenKey) Native() string {
	return string(parameter)
}

func (parameter AccessTokenKey) TypedString() types.String {
	return types.String(parameter)
}

type AccessTokenKeyNull struct {
	value  AccessTokenKey
	filled bool
}

func (parameter AccessTokenKey) AccessTokenKeyNull() AccessTokenKeyNull {
	return NewAccessTokenKeyNull(parameter)
}

func NewAccessTokenKeyNull(value AccessTokenKey) AccessTokenKeyNull {
	return AccessTokenKeyNull{
		filled: true,
		value:  value,
	}
}

func NewAccessTokenKeyNullFromString(value string) AccessTokenKeyNull {
	return AccessTokenKeyNull{
		filled: true,
		value:  AccessTokenKey(value),
	}
}

func (parameter AccessTokenKeyNull) Get() (AccessTokenKey, bool) {
	return parameter.value, parameter.filled
}

func (parameter AccessTokenKeyNull) Filled() bool {
	return parameter.filled
}

func (parameter AccessTokenKeyNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *AccessTokenKeyNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
