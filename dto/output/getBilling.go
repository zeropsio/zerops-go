// Generated ZEROPS sdk

package output

import (
	"strconv"
)

var _ strconv.NumError

type GetBilling struct {
	Current BillingInfo  `json:"current"`
	Future  *BillingInfo `json:"future"`
}
