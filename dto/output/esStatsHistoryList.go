// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type EsStatsHistoryList struct {
	Limit       types.IntNull                  `json:"limit"`
	From        types.DateTimeNull             `json:"from"`
	Till        types.DateTimeNull             `json:"till"`
	TotalHits   types.Int                      `json:"totalHits"`
	GroupBy     enum.EsStatsHistoryGroupByEnum `json:"groupBy"`
	TimeGroupBy types.String                   `json:"timeGroupBy"`
	Items       EsStatsHistoryListItems        `json:"items"`
}

type EsStatsHistoryListItems []EsStatsHistory

func (dto EsStatsHistoryListItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsStatsHistory(dto))
}
