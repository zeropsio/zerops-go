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

type EsCardPayment struct {
	Id                uuid.CardPaymentId              `json:"id"`
	ClientId          uuid.ClientId                   `json:"clientId"`
	Status            enum.PaymentStatusEnum          `json:"status"`
	RequestId         types.StringNull                `json:"requestId"`
	Created           types.DateTime                  `json:"created"`
	Amount            types.Decimal                   `json:"amount"`
	AmountVat         types.Decimal                   `json:"amountVat"`
	CurrencyId        stringId.CurrencyId             `json:"currencyId"`
	DatePaid          types.DateTimeNull              `json:"datePaid"`
	InteractionType   enum.PaymentInteractionTypeEnum `json:"interactionType"`
	LastUpdate        types.DateTime                  `json:"lastUpdate"`
	PaymentMethod     types.StringNull                `json:"paymentMethod"`
	PaymentObject     types.JsonRawMessage            `json:"paymentObject"`
	BillingInfoObject types.JsonRawMessage            `json:"billingInfoObject"`
}
