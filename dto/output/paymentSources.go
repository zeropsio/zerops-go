// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type PaymentSources struct {
	ClientId uuid.PaymentSourcesClientId `json:"clientId"`
	Sources  PaymentSourcesSources       `json:"sources"`
}

type PaymentSourcesSources []PaymentSource

func (dto PaymentSourcesSources) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PaymentSource(dto))
}
