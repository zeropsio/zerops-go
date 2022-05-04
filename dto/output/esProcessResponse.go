// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsProcessResponse struct {
	Limit     types.Int              `json:"limit"`
	Offset    types.Int              `json:"offset"`
	TotalHits types.Int              `json:"totalHits"`
	Items     EsProcessResponseItems `json:"items"`
}

type EsProcessResponseItems []EsProcess

func (dto EsProcessResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsProcess(dto))
}
