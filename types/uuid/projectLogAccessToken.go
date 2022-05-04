// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ProjectLogAccessToken types.UuidShort

func NewProjectLogAccessTokenFromString(value string) (out ProjectLogAccessToken, err error) {
	return ProjectLogAccessToken(value), nil
}

func (parameter ProjectLogAccessToken) Native() string {
	return string(parameter)
}

func (parameter ProjectLogAccessToken) TypedString() types.String {
	return types.String(parameter)
}

type ProjectLogAccessTokenNull struct {
	value  ProjectLogAccessToken
	filled bool
}

func (parameter ProjectLogAccessToken) ProjectLogAccessTokenNull() ProjectLogAccessTokenNull {
	return NewProjectLogAccessTokenNull(parameter)
}

func NewProjectLogAccessTokenNull(value ProjectLogAccessToken) ProjectLogAccessTokenNull {
	return ProjectLogAccessTokenNull{
		filled: true,
		value:  value,
	}
}

func NewProjectLogAccessTokenNullFromString(value string) ProjectLogAccessTokenNull {
	return ProjectLogAccessTokenNull{
		filled: true,
		value:  ProjectLogAccessToken(value),
	}
}

func (parameter ProjectLogAccessTokenNull) Get() (ProjectLogAccessToken, bool) {
	return parameter.value, parameter.filled
}

func (parameter ProjectLogAccessTokenNull) Filled() bool {
	return parameter.filled
}

func (parameter ProjectLogAccessTokenNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ProjectLogAccessTokenNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
