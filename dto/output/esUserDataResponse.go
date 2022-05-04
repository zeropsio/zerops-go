// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsUserDataResponse struct {
	Limit     types.Int               `json:"limit"`
	Offset    types.Int               `json:"offset"`
	TotalHits types.Int               `json:"totalHits"`
	Items     EsUserDataResponseItems `json:"items"`
}

type EsUserDataResponseItems []EsUserData

func (dto EsUserDataResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsUserData(dto))
}
