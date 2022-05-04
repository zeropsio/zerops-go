// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsTransactionDebitGroupByResponse struct {
	Items       EsTransactionDebitGroupByResponseItems `json:"items"`
	TotalHits   types.Int                              `json:"totalHits"`
	TimeGroupBy types.String                           `json:"timeGroupBy"`
	Limit       types.IntNull                          `json:"limit"`
	From        types.DateTimeNull                     `json:"from"`
	Till        types.DateTimeNull                     `json:"till"`
	GroupBy     types.String                           `json:"groupBy"`
	ProjectId   types.StringNull                       `json:"projectId"`
	StackId     types.StringNull                       `json:"stackId"`
}

type EsTransactionDebitGroupByResponseItems []EsTransactionDebitGroupItem

func (dto EsTransactionDebitGroupByResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsTransactionDebitGroupItem(dto))
}
