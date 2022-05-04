// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsContainerResponse struct {
	Limit     types.Int                `json:"limit"`
	Offset    types.Int                `json:"offset"`
	TotalHits types.Int                `json:"totalHits"`
	Items     EsContainerResponseItems `json:"items"`
}

type EsContainerResponseItems []EsContainer

func (dto EsContainerResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsContainer(dto))
}
