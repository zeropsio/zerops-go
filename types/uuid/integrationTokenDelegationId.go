// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type IntegrationTokenDelegationId types.UuidShort

func NewIntegrationTokenDelegationIdFromString(value string) (out IntegrationTokenDelegationId, err error) {
	return IntegrationTokenDelegationId(value), nil
}

func (parameter IntegrationTokenDelegationId) Native() string {
	return string(parameter)
}

func (parameter IntegrationTokenDelegationId) TypedString() types.String {
	return types.String(parameter)
}

type IntegrationTokenDelegationIdNull struct {
	value  IntegrationTokenDelegationId
	filled bool
}

func (parameter IntegrationTokenDelegationId) IntegrationTokenDelegationIdNull() IntegrationTokenDelegationIdNull {
	return NewIntegrationTokenDelegationIdNull(parameter)
}

func NewIntegrationTokenDelegationIdNull(value IntegrationTokenDelegationId) IntegrationTokenDelegationIdNull {
	return IntegrationTokenDelegationIdNull{
		filled: true,
		value:  value,
	}
}

func NewIntegrationTokenDelegationIdNullFromString(value string) IntegrationTokenDelegationIdNull {
	return IntegrationTokenDelegationIdNull{
		filled: true,
		value:  IntegrationTokenDelegationId(value),
	}
}

func (parameter IntegrationTokenDelegationIdNull) Get() (IntegrationTokenDelegationId, bool) {
	return parameter.value, parameter.filled
}

func (parameter IntegrationTokenDelegationIdNull) Filled() bool {
	return parameter.filled
}

func (parameter IntegrationTokenDelegationIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *IntegrationTokenDelegationIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
