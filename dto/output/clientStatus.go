// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ClientStatus struct {
	Credit      types.Float      `json:"credit"`
	PromoCredit types.Float      `json:"promoCredit"`
	ClientInfo  ClientInfoObject `json:"clientInfo"`
}
