// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsAppVersionResponse struct {
	Limit     types.Int                 `json:"limit"`
	Offset    types.Int                 `json:"offset"`
	TotalHits types.Int                 `json:"totalHits"`
	Items     EsAppVersionResponseItems `json:"items"`
}

type EsAppVersionResponseItems []EsAppVersion

func (dto EsAppVersionResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsAppVersion(dto))
}
