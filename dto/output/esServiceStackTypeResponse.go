// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsServiceStackTypeResponse struct {
	Limit     types.Int                       `json:"limit"`
	Offset    types.Int                       `json:"offset"`
	TotalHits types.Int                       `json:"totalHits"`
	Items     EsServiceStackTypeResponseItems `json:"items"`
}

type EsServiceStackTypeResponseItems []EsServiceStackType

func (dto EsServiceStackTypeResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsServiceStackType(dto))
}
