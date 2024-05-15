// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type Client struct {
	Id                      uuid.ClientId    `json:"id"`
	AccountName             types.String     `json:"accountName"`
	Avatar                  *ClientAvatar    `json:"avatar"`
	BillingInfo             *BillingInfo     `json:"billingInfo"`
	PaymentProviderClientId types.StringNull `json:"paymentProviderClientId"`
	VerifiedByTopUp         types.Bool       `json:"verifiedByTopUp"`
	VerifiedManually        types.Bool       `json:"verifiedManually"`
}
