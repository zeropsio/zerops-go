// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type Invoice struct {
	Id              uuid.InvoiceId         `json:"id"`
	ClientId        uuid.ClientId          `json:"clientId"`
	CurrencyId      stringId.CurrencyId    `json:"currencyId"`
	PaymentId       uuid.CardPaymentIdNull `json:"paymentId"`
	Created         types.DateTime         `json:"created"`
	LastUpdate      types.DateTime         `json:"lastUpdate"`
	Status          enum.InvoiceStatusEnum `json:"status"`
	Number          types.String           `json:"number"`
	BillingInfo     types.JsonRawMessage   `json:"billingInfo"`
	Supplier        types.JsonRawMessage   `json:"supplier"`
	VatRate         types.Float            `json:"vatRate"`
	Amount          types.Float            `json:"amount"`
	AmountVat       types.Float            `json:"amountVat"`
	AmountVatCzk    types.Float            `json:"amountVatCzk"`
	CurrencyRateCzk types.Float            `json:"currencyRateCzk"`
	Total           types.Float            `json:"total"`
	DueDate         types.DateTime         `json:"dueDate"`
	PaidDate        types.DateTimeNull     `json:"paidDate"`
}
