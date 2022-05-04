// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsClientUserResponse struct {
	Limit     types.Int                 `json:"limit"`
	Offset    types.Int                 `json:"offset"`
	TotalHits types.Int                 `json:"totalHits"`
	Items     EsClientUserResponseItems `json:"items"`
}

type EsClientUserResponseItems []EsClientUser

func (dto EsClientUserResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsClientUser(dto))
}
