// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type Invoice struct {
	TestField    types.Int           `json:"testField"`
	InvoiceItems InvoiceInvoiceItems `json:"invoiceItems"`
}

type InvoiceInvoiceItems []InvoiceItem

func (dto InvoiceInvoiceItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]InvoiceItem(dto))
}
