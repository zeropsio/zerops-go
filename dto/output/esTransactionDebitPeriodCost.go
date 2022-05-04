// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsTransactionDebitPeriodCost struct {
	Last24hours                types.Float `json:"last24hours"`
	Last7days                  types.Float `json:"last7days"`
	Last30days                 types.Float `json:"last30days"`
	Last12month                types.Float `json:"last12month"`
	Today                      types.Float `json:"today"`
	Yesterday                  types.Float `json:"yesterday"`
	LastMonth                  types.Float `json:"lastMonth"`
	ThisMonth                  types.Float `json:"thisMonth"`
	LastYear                   types.Float `json:"lastYear"`
	ThisYear                   types.Float `json:"thisYear"`
	AverageLast30Days          types.Float `json:"averageLast30Days"`
	AverageLast30DaysPrecision types.Int   `json:"averageLast30DaysPrecision"`
}
