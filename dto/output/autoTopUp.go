// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type AutoTopUp struct {
	Threshold            types.Decimal `json:"threshold"`
	Amount               types.Decimal `json:"amount"`
	CalendarMonthlyLimit types.Decimal `json:"calendarMonthlyLimit"`
}
