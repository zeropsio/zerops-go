// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsPublicPortRoutingResponse struct {
	Limit     types.Int                        `json:"limit"`
	Offset    types.Int                        `json:"offset"`
	TotalHits types.Int                        `json:"totalHits"`
	Items     EsPublicPortRoutingResponseItems `json:"items"`
}

type EsPublicPortRoutingResponseItems []EsPublicPortRouting

func (dto EsPublicPortRoutingResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsPublicPortRouting(dto))
}
