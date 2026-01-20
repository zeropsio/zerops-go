// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type WebAuthnOptionsSessionId types.UuidShort

func NewWebAuthnOptionsSessionIdFromString(value string) (out WebAuthnOptionsSessionId, err error) {
	return WebAuthnOptionsSessionId(value), nil
}

func (parameter WebAuthnOptionsSessionId) Native() string {
	return string(parameter)
}

func (parameter WebAuthnOptionsSessionId) TypedString() types.String {
	return types.String(parameter)
}

type WebAuthnOptionsSessionIdNull struct {
	value  WebAuthnOptionsSessionId
	filled bool
}

func (parameter WebAuthnOptionsSessionId) WebAuthnOptionsSessionIdNull() WebAuthnOptionsSessionIdNull {
	return NewWebAuthnOptionsSessionIdNull(parameter)
}

func NewWebAuthnOptionsSessionIdNull(value WebAuthnOptionsSessionId) WebAuthnOptionsSessionIdNull {
	return WebAuthnOptionsSessionIdNull{
		filled: true,
		value:  value,
	}
}

func NewWebAuthnOptionsSessionIdNullFromString(value string) WebAuthnOptionsSessionIdNull {
	return WebAuthnOptionsSessionIdNull{
		filled: true,
		value:  WebAuthnOptionsSessionId(value),
	}
}

func (parameter WebAuthnOptionsSessionIdNull) Get() (WebAuthnOptionsSessionId, bool) {
	return parameter.value, parameter.filled
}

func (parameter WebAuthnOptionsSessionIdNull) Filled() bool {
	return parameter.filled
}

func (parameter WebAuthnOptionsSessionIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *WebAuthnOptionsSessionIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
