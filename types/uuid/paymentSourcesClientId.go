// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type PaymentSourcesClientId types.UuidShort

func NewPaymentSourcesClientIdFromString(value string) (out PaymentSourcesClientId, err error) {
	return PaymentSourcesClientId(value), nil
}

func (parameter PaymentSourcesClientId) Native() string {
	return string(parameter)
}

func (parameter PaymentSourcesClientId) TypedString() types.String {
	return types.String(parameter)
}

type PaymentSourcesClientIdNull struct {
	value  PaymentSourcesClientId
	filled bool
}

func (parameter PaymentSourcesClientId) PaymentSourcesClientIdNull() PaymentSourcesClientIdNull {
	return NewPaymentSourcesClientIdNull(parameter)
}

func NewPaymentSourcesClientIdNull(value PaymentSourcesClientId) PaymentSourcesClientIdNull {
	return PaymentSourcesClientIdNull{
		filled: true,
		value:  value,
	}
}

func NewPaymentSourcesClientIdNullFromString(value string) PaymentSourcesClientIdNull {
	return PaymentSourcesClientIdNull{
		filled: true,
		value:  PaymentSourcesClientId(value),
	}
}

func (parameter PaymentSourcesClientIdNull) Get() (PaymentSourcesClientId, bool) {
	return parameter.value, parameter.filled
}

func (parameter PaymentSourcesClientIdNull) Filled() bool {
	return parameter.filled
}

func (parameter PaymentSourcesClientIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *PaymentSourcesClientIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
