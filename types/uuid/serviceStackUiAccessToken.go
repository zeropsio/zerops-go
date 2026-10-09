// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ServiceStackUiAccessToken types.UuidShort

func NewServiceStackUiAccessTokenFromString(value string) (out ServiceStackUiAccessToken, err error) {
	return ServiceStackUiAccessToken(value), nil
}

func (parameter ServiceStackUiAccessToken) Native() string {
	return string(parameter)
}

func (parameter ServiceStackUiAccessToken) TypedString() types.String {
	return types.String(parameter)
}

type ServiceStackUiAccessTokenNull struct {
	value  ServiceStackUiAccessToken
	filled bool
}

func (parameter ServiceStackUiAccessToken) ServiceStackUiAccessTokenNull() ServiceStackUiAccessTokenNull {
	return NewServiceStackUiAccessTokenNull(parameter)
}

func NewServiceStackUiAccessTokenNull(value ServiceStackUiAccessToken) ServiceStackUiAccessTokenNull {
	return ServiceStackUiAccessTokenNull{
		filled: true,
		value:  value,
	}
}

func NewServiceStackUiAccessTokenNullFromString(value string) ServiceStackUiAccessTokenNull {
	return ServiceStackUiAccessTokenNull{
		filled: true,
		value:  ServiceStackUiAccessToken(value),
	}
}

func (parameter ServiceStackUiAccessTokenNull) Get() (ServiceStackUiAccessToken, bool) {
	return parameter.value, parameter.filled
}

func (parameter ServiceStackUiAccessTokenNull) Filled() bool {
	return parameter.filled
}

func (parameter ServiceStackUiAccessTokenNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ServiceStackUiAccessTokenNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
