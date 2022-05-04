// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsCardPayment struct {
	Id              uuid.CardPaymentId              `json:"id"`
	ClientId        uuid.ClientId                   `json:"clientId"`
	Status          enum.CardPaymentStatusEnum      `json:"status"`
	RequestId       types.StringNull                `json:"requestId"`
	Created         types.DateTime                  `json:"created"`
	Amount          types.Float                     `json:"amount"`
	CurrencyId      uuid.CardPaymentCurrencyId      `json:"currencyId"`
	Application     enum.CardPaymentApplicationEnum `json:"application"`
	DatePaid        types.DateTimeNull              `json:"datePaid"`
	InteractionType types.String                    `json:"interactionType"`
	LastUpdate      types.DateTime                  `json:"lastUpdate"`
	PaymentMethod   types.String                    `json:"paymentMethod"`
	PaymentObject   types.JsonRawMessage            `json:"paymentObject"`
	Provider        enum.CardPaymentProviderEnum    `json:"provider"`
}
