// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsLanguageResponse struct {
	Limit     types.Int               `json:"limit"`
	Offset    types.Int               `json:"offset"`
	TotalHits types.Int               `json:"totalHits"`
	Items     EsLanguageResponseItems `json:"items"`
}

type EsLanguageResponseItems []EsLanguage

func (dto EsLanguageResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsLanguage(dto))
}
