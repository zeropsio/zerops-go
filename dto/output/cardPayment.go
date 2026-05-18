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

type CardPayment struct {
	Id              uuid.CardPaymentId              `json:"id"`
	ClientId        uuid.ClientId                   `json:"clientId"`
	UserId          uuid.UserIdNull                 `json:"userId"`
	CurrencyId      stringId.CurrencyId             `json:"currencyId"`
	Status          enum.PaymentStatusEnum          `json:"status"`
	InteractionType enum.PaymentInteractionTypeEnum `json:"interactionType"`
	Amount          types.Decimal                   `json:"amount"`
	AmountVat       types.Decimal                   `json:"amountVat"`
	PaymentMethod   types.StringNull                `json:"paymentMethod"`
	RequestId       types.StringNull                `json:"requestId"`
	Created         types.DateTime                  `json:"created"`
	LastUpdate      types.DateTime                  `json:"lastUpdate"`
	DatePaid        types.DateTimeNull              `json:"datePaid"`
	InvoiceId       uuid.InvoiceIdNull              `json:"invoiceId"`
}
