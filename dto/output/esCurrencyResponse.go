// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsCurrencyResponse struct {
	Limit     types.Int               `json:"limit"`
	Offset    types.Int               `json:"offset"`
	TotalHits types.Int               `json:"totalHits"`
	Items     EsCurrencyResponseItems `json:"items"`
}

type EsCurrencyResponseItems []EsCurrency

func (dto EsCurrencyResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsCurrency(dto))
}
