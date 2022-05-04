// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsClientResponse struct {
	Limit     types.Int             `json:"limit"`
	Offset    types.Int             `json:"offset"`
	TotalHits types.Int             `json:"totalHits"`
	Items     EsClientResponseItems `json:"items"`
}

type EsClientResponseItems []EsClient

func (dto EsClientResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsClient(dto))
}
