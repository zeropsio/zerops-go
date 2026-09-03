// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type CardPaymentList struct {
	Count  types.Int           `json:"count"`
	Total  types.Int           `json:"total"`
	Limit  types.Int           `json:"limit"`
	Offset types.Int           `json:"offset"`
	List   CardPaymentListList `json:"list"`
}

type CardPaymentListList []CardPayment

func (dto CardPaymentListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]CardPayment(dto))
}
