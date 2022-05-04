// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsTransactionDebitCostItem struct {
	Id         types.String                 `json:"id"`
	PeriodCost EsTransactionDebitPeriodCost `json:"periodCost"`
}
