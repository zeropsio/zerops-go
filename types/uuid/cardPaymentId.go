// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type CardPaymentId types.UuidShort

func NewCardPaymentIdFromString(value string) (out CardPaymentId, err error) {
	return CardPaymentId(value), nil
}

func (parameter CardPaymentId) Native() string {
	return string(parameter)
}

func (parameter CardPaymentId) TypedString() types.String {
	return types.String(parameter)
}

type CardPaymentIdNull struct {
	value  CardPaymentId
	filled bool
}

func (parameter CardPaymentId) CardPaymentIdNull() CardPaymentIdNull {
	return NewCardPaymentIdNull(parameter)
}

func NewCardPaymentIdNull(value CardPaymentId) CardPaymentIdNull {
	return CardPaymentIdNull{
		filled: true,
		value:  value,
	}
}

func NewCardPaymentIdNullFromString(value string) CardPaymentIdNull {
	return CardPaymentIdNull{
		filled: true,
		value:  CardPaymentId(value),
	}
}

func (parameter CardPaymentIdNull) Get() (CardPaymentId, bool) {
	return parameter.value, parameter.filled
}

func (parameter CardPaymentIdNull) Filled() bool {
	return parameter.filled
}

func (parameter CardPaymentIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *CardPaymentIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
