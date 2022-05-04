// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type InvoiceId types.UuidShort

func NewInvoiceIdFromString(value string) (out InvoiceId, err error) {
	return InvoiceId(value), nil
}

func (parameter InvoiceId) Native() string {
	return string(parameter)
}

func (parameter InvoiceId) TypedString() types.String {
	return types.String(parameter)
}

type InvoiceIdNull struct {
	value  InvoiceId
	filled bool
}

func (parameter InvoiceId) InvoiceIdNull() InvoiceIdNull {
	return NewInvoiceIdNull(parameter)
}

func NewInvoiceIdNull(value InvoiceId) InvoiceIdNull {
	return InvoiceIdNull{
		filled: true,
		value:  value,
	}
}

func NewInvoiceIdNullFromString(value string) InvoiceIdNull {
	return InvoiceIdNull{
		filled: true,
		value:  InvoiceId(value),
	}
}

func (parameter InvoiceIdNull) Get() (InvoiceId, bool) {
	return parameter.value, parameter.filled
}

func (parameter InvoiceIdNull) Filled() bool {
	return parameter.filled
}

func (parameter InvoiceIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *InvoiceIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
