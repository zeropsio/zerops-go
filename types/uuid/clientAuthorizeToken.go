// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ClientAuthorizeToken types.UuidShort

func NewClientAuthorizeTokenFromString(value string) (out ClientAuthorizeToken, err error) {
	return ClientAuthorizeToken(value), nil
}

func (parameter ClientAuthorizeToken) Native() string {
	return string(parameter)
}

func (parameter ClientAuthorizeToken) TypedString() types.String {
	return types.String(parameter)
}

type ClientAuthorizeTokenNull struct {
	value  ClientAuthorizeToken
	filled bool
}

func (parameter ClientAuthorizeToken) ClientAuthorizeTokenNull() ClientAuthorizeTokenNull {
	return NewClientAuthorizeTokenNull(parameter)
}

func NewClientAuthorizeTokenNull(value ClientAuthorizeToken) ClientAuthorizeTokenNull {
	return ClientAuthorizeTokenNull{
		filled: true,
		value:  value,
	}
}

func NewClientAuthorizeTokenNullFromString(value string) ClientAuthorizeTokenNull {
	return ClientAuthorizeTokenNull{
		filled: true,
		value:  ClientAuthorizeToken(value),
	}
}

func (parameter ClientAuthorizeTokenNull) Get() (ClientAuthorizeToken, bool) {
	return parameter.value, parameter.filled
}

func (parameter ClientAuthorizeTokenNull) Filled() bool {
	return parameter.filled
}

func (parameter ClientAuthorizeTokenNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ClientAuthorizeTokenNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
