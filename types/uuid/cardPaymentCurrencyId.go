// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type CardPaymentCurrencyId types.UuidShort

func NewCardPaymentCurrencyIdFromString(value string) (out CardPaymentCurrencyId, err error) {
	return CardPaymentCurrencyId(value), nil
}

func (parameter CardPaymentCurrencyId) Native() string {
	return string(parameter)
}

func (parameter CardPaymentCurrencyId) TypedString() types.String {
	return types.String(parameter)
}

type CardPaymentCurrencyIdNull struct {
	value  CardPaymentCurrencyId
	filled bool
}

func (parameter CardPaymentCurrencyId) CardPaymentCurrencyIdNull() CardPaymentCurrencyIdNull {
	return NewCardPaymentCurrencyIdNull(parameter)
}

func NewCardPaymentCurrencyIdNull(value CardPaymentCurrencyId) CardPaymentCurrencyIdNull {
	return CardPaymentCurrencyIdNull{
		filled: true,
		value:  value,
	}
}

func NewCardPaymentCurrencyIdNullFromString(value string) CardPaymentCurrencyIdNull {
	return CardPaymentCurrencyIdNull{
		filled: true,
		value:  CardPaymentCurrencyId(value),
	}
}

func (parameter CardPaymentCurrencyIdNull) Get() (CardPaymentCurrencyId, bool) {
	return parameter.value, parameter.filled
}

func (parameter CardPaymentCurrencyIdNull) Filled() bool {
	return parameter.filled
}

func (parameter CardPaymentCurrencyIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *CardPaymentCurrencyIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
