// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type InvoiceItem struct {
	Id               uuid.InvoiceItemId `json:"id"`
	Total            types.Float        `json:"total"`
	VatRate          types.Float        `json:"vatRate"`
	InvoiceId        uuid.InvoiceId     `json:"invoiceId"`
	Amount           types.Float        `json:"amount"`
	UnitPrice        types.Float        `json:"unitPrice"`
	IsReverseCharged types.Bool         `json:"isReverseCharged"`
	ItemName         types.String       `json:"itemName"`
}
