// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type CardPaymentList struct {
	List       CardPaymentListList `json:"list"`
	TotalCount types.Int           `json:"totalCount"`
}

type CardPaymentListList []CardPayment

func (dto CardPaymentListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]CardPayment(dto))
}
