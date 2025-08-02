// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsGroupBySearchResponse struct {
	Items       types.JsonRawMessageArray       `json:"items"`
	TotalHits   types.Int                       `json:"totalHits"`
	Limit       types.IntNull                   `json:"limit"`
	From        types.DateTimeNull              `json:"from"`
	Till        types.DateTimeNull              `json:"till"`
	ProjectId   uuid.ProjectIdNull              `json:"projectId"`
	StackId     uuid.ServiceStackIdNull         `json:"stackId"`
	GroupBy     enum.EsGroupBySearchGroupByEnum `json:"groupBy"`
	TimeGroupBy types.String                    `json:"timeGroupBy"`
}
