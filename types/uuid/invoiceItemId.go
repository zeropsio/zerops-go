// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type InvoiceItemId types.UuidShort

func NewInvoiceItemIdFromString(value string) (out InvoiceItemId, err error) {
	return InvoiceItemId(value), nil
}

func (parameter InvoiceItemId) Native() string {
	return string(parameter)
}

func (parameter InvoiceItemId) TypedString() types.String {
	return types.String(parameter)
}

type InvoiceItemIdNull struct {
	value  InvoiceItemId
	filled bool
}

func (parameter InvoiceItemId) InvoiceItemIdNull() InvoiceItemIdNull {
	return NewInvoiceItemIdNull(parameter)
}

func NewInvoiceItemIdNull(value InvoiceItemId) InvoiceItemIdNull {
	return InvoiceItemIdNull{
		filled: true,
		value:  value,
	}
}

func NewInvoiceItemIdNullFromString(value string) InvoiceItemIdNull {
	return InvoiceItemIdNull{
		filled: true,
		value:  InvoiceItemId(value),
	}
}

func (parameter InvoiceItemIdNull) Get() (InvoiceItemId, bool) {
	return parameter.value, parameter.filled
}

func (parameter InvoiceItemIdNull) Filled() bool {
	return parameter.filled
}

func (parameter InvoiceItemIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *InvoiceItemIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
