// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsServiceStackResponse struct {
	Limit     types.Int                   `json:"limit"`
	Offset    types.Int                   `json:"offset"`
	TotalHits types.Int                   `json:"totalHits"`
	Items     EsServiceStackResponseItems `json:"items"`
}

type EsServiceStackResponseItems []EsServiceStack

func (dto EsServiceStackResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsServiceStack(dto))
}
