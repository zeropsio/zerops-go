// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsProjectResponse struct {
	Limit     types.Int              `json:"limit"`
	Offset    types.Int              `json:"offset"`
	TotalHits types.Int              `json:"totalHits"`
	Items     EsProjectResponseItems `json:"items"`
}

type EsProjectResponseItems []EsProject

func (dto EsProjectResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsProject(dto))
}
