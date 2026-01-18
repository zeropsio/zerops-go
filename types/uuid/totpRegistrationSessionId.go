// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type TotpRegistrationSessionId types.UuidShort

func NewTotpRegistrationSessionIdFromString(value string) (out TotpRegistrationSessionId, err error) {
	return TotpRegistrationSessionId(value), nil
}

func (parameter TotpRegistrationSessionId) Native() string {
	return string(parameter)
}

func (parameter TotpRegistrationSessionId) TypedString() types.String {
	return types.String(parameter)
}

type TotpRegistrationSessionIdNull struct {
	value  TotpRegistrationSessionId
	filled bool
}

func (parameter TotpRegistrationSessionId) TotpRegistrationSessionIdNull() TotpRegistrationSessionIdNull {
	return NewTotpRegistrationSessionIdNull(parameter)
}

func NewTotpRegistrationSessionIdNull(value TotpRegistrationSessionId) TotpRegistrationSessionIdNull {
	return TotpRegistrationSessionIdNull{
		filled: true,
		value:  value,
	}
}

func NewTotpRegistrationSessionIdNullFromString(value string) TotpRegistrationSessionIdNull {
	return TotpRegistrationSessionIdNull{
		filled: true,
		value:  TotpRegistrationSessionId(value),
	}
}

func (parameter TotpRegistrationSessionIdNull) Get() (TotpRegistrationSessionId, bool) {
	return parameter.value, parameter.filled
}

func (parameter TotpRegistrationSessionIdNull) Filled() bool {
	return parameter.filled
}

func (parameter TotpRegistrationSessionIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *TotpRegistrationSessionIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
