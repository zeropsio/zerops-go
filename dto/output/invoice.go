// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type Invoice struct {
	Id                uuid.InvoiceId      `json:"id"`
	Created           types.DateTime      `json:"created"`
	LastUpdate        types.DateTime      `json:"lastUpdate"`
	Number            types.Int64         `json:"number"`
	Status            types.String        `json:"status"`
	Client            ClientLight         `json:"client"`
	ClientAddress     AddressJsonObject   `json:"clientAddress"`
	Contractor        ContractorLight     `json:"contractor"`
	ContractorAddress AddressJsonObject   `json:"contractorAddress"`
	DueDate           types.DateTime      `json:"dueDate"`
	Total             types.Float         `json:"total"`
	VatRate           types.Float         `json:"vatRate"`
	VatAmount         types.Float         `json:"vatAmount"`
	Currency          Currency            `json:"currency"`
	TotalVat          types.Float         `json:"totalVat"`
	PaidDate          types.DateTime      `json:"paidDate"`
	TotalDue          types.Float         `json:"totalDue"`
	InvoiceItems      InvoiceInvoiceItems `json:"invoiceItems"`
}

type InvoiceInvoiceItems []InvoiceItem

func (dto InvoiceInvoiceItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]InvoiceItem(dto))
}
