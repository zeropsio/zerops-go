// Generated ZEROPS sdk

package stringId

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type PaymentSourceId types.StringId

func (parameter PaymentSourceId) PaymentSourceIdNull() PaymentSourceIdNull {
	return NewPaymentSourceIdNull(parameter)
}

func NewPaymentSourceIdNullFromString(value string) PaymentSourceIdNull {
	return PaymentSourceIdNull{
		filled: true,
		value:  PaymentSourceId(value),
	}
}

func NewPaymentSourceIdNull(value PaymentSourceId) PaymentSourceIdNull {
	return PaymentSourceIdNull{
		filled: true,
		value:  value,
	}
}

func NewPaymentSourceIdFromString(value string) (out PaymentSourceId, err error) {
	return PaymentSourceId(value), nil
}

func (parameter PaymentSourceId) Native() string {
	return string(parameter)
}

type PaymentSourceIdNull struct {
	value  PaymentSourceId
	filled bool
}

func (parameter PaymentSourceIdNull) Get() (PaymentSourceId, bool) {
	return parameter.value, parameter.filled
}

func (parameter PaymentSourceIdNull) Filled() bool {
	return parameter.filled
}

func (parameter PaymentSourceIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *PaymentSourceIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
