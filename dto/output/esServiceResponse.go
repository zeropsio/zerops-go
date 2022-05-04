// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsServiceResponse struct {
	Limit     types.Int              `json:"limit"`
	Offset    types.Int              `json:"offset"`
	TotalHits types.Int              `json:"totalHits"`
	Items     EsServiceResponseItems `json:"items"`
}

type EsServiceResponseItems []EsService

func (dto EsServiceResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsService(dto))
}
