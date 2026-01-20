// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type WebAuthnLoginVerificationSessionId types.UuidShort

func NewWebAuthnLoginVerificationSessionIdFromString(value string) (out WebAuthnLoginVerificationSessionId, err error) {
	return WebAuthnLoginVerificationSessionId(value), nil
}

func (parameter WebAuthnLoginVerificationSessionId) Native() string {
	return string(parameter)
}

func (parameter WebAuthnLoginVerificationSessionId) TypedString() types.String {
	return types.String(parameter)
}

type WebAuthnLoginVerificationSessionIdNull struct {
	value  WebAuthnLoginVerificationSessionId
	filled bool
}

func (parameter WebAuthnLoginVerificationSessionId) WebAuthnLoginVerificationSessionIdNull() WebAuthnLoginVerificationSessionIdNull {
	return NewWebAuthnLoginVerificationSessionIdNull(parameter)
}

func NewWebAuthnLoginVerificationSessionIdNull(value WebAuthnLoginVerificationSessionId) WebAuthnLoginVerificationSessionIdNull {
	return WebAuthnLoginVerificationSessionIdNull{
		filled: true,
		value:  value,
	}
}

func NewWebAuthnLoginVerificationSessionIdNullFromString(value string) WebAuthnLoginVerificationSessionIdNull {
	return WebAuthnLoginVerificationSessionIdNull{
		filled: true,
		value:  WebAuthnLoginVerificationSessionId(value),
	}
}

func (parameter WebAuthnLoginVerificationSessionIdNull) Get() (WebAuthnLoginVerificationSessionId, bool) {
	return parameter.value, parameter.filled
}

func (parameter WebAuthnLoginVerificationSessionIdNull) Filled() bool {
	return parameter.filled
}

func (parameter WebAuthnLoginVerificationSessionIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *WebAuthnLoginVerificationSessionIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
