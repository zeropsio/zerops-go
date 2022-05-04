// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsCountryResponse struct {
	Limit     types.Int              `json:"limit"`
	Offset    types.Int              `json:"offset"`
	TotalHits types.Int              `json:"totalHits"`
	Items     EsCountryResponseItems `json:"items"`
}

type EsCountryResponseItems []EsCountry

func (dto EsCountryResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsCountry(dto))
}
