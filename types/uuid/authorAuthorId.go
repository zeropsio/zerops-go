// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type AuthorAuthorId types.UuidShort

func NewAuthorAuthorIdFromString(value string) (out AuthorAuthorId, err error) {
	return AuthorAuthorId(value), nil
}

func (parameter AuthorAuthorId) Native() string {
	return string(parameter)
}

func (parameter AuthorAuthorId) TypedString() types.String {
	return types.String(parameter)
}

type AuthorAuthorIdNull struct {
	value  AuthorAuthorId
	filled bool
}

func (parameter AuthorAuthorId) AuthorAuthorIdNull() AuthorAuthorIdNull {
	return NewAuthorAuthorIdNull(parameter)
}

func NewAuthorAuthorIdNull(value AuthorAuthorId) AuthorAuthorIdNull {
	return AuthorAuthorIdNull{
		filled: true,
		value:  value,
	}
}

func NewAuthorAuthorIdNullFromString(value string) AuthorAuthorIdNull {
	return AuthorAuthorIdNull{
		filled: true,
		value:  AuthorAuthorId(value),
	}
}

func (parameter AuthorAuthorIdNull) Get() (AuthorAuthorId, bool) {
	return parameter.value, parameter.filled
}

func (parameter AuthorAuthorIdNull) Filled() bool {
	return parameter.filled
}

func (parameter AuthorAuthorIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *AuthorAuthorIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
