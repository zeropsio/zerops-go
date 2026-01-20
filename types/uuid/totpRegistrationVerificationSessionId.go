// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type TotpRegistrationVerificationSessionId types.UuidShort

func NewTotpRegistrationVerificationSessionIdFromString(value string) (out TotpRegistrationVerificationSessionId, err error) {
	return TotpRegistrationVerificationSessionId(value), nil
}

func (parameter TotpRegistrationVerificationSessionId) Native() string {
	return string(parameter)
}

func (parameter TotpRegistrationVerificationSessionId) TypedString() types.String {
	return types.String(parameter)
}

type TotpRegistrationVerificationSessionIdNull struct {
	value  TotpRegistrationVerificationSessionId
	filled bool
}

func (parameter TotpRegistrationVerificationSessionId) TotpRegistrationVerificationSessionIdNull() TotpRegistrationVerificationSessionIdNull {
	return NewTotpRegistrationVerificationSessionIdNull(parameter)
}

func NewTotpRegistrationVerificationSessionIdNull(value TotpRegistrationVerificationSessionId) TotpRegistrationVerificationSessionIdNull {
	return TotpRegistrationVerificationSessionIdNull{
		filled: true,
		value:  value,
	}
}

func NewTotpRegistrationVerificationSessionIdNullFromString(value string) TotpRegistrationVerificationSessionIdNull {
	return TotpRegistrationVerificationSessionIdNull{
		filled: true,
		value:  TotpRegistrationVerificationSessionId(value),
	}
}

func (parameter TotpRegistrationVerificationSessionIdNull) Get() (TotpRegistrationVerificationSessionId, bool) {
	return parameter.value, parameter.filled
}

func (parameter TotpRegistrationVerificationSessionIdNull) Filled() bool {
	return parameter.filled
}

func (parameter TotpRegistrationVerificationSessionIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *TotpRegistrationVerificationSessionIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
