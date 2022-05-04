// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsCardPaymentResponse struct {
	Limit     types.Int                  `json:"limit"`
	Offset    types.Int                  `json:"offset"`
	TotalHits types.Int                  `json:"totalHits"`
	Items     EsCardPaymentResponseItems `json:"items"`
}

type EsCardPaymentResponseItems []EsCardPayment

func (dto EsCardPaymentResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsCardPayment(dto))
}
