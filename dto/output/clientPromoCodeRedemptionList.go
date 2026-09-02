// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ClientPromoCodeRedemptionList struct {
	List       ClientPromoCodeRedemptionListList `json:"list"`
	TotalCount types.Int                         `json:"totalCount"`
}

type ClientPromoCodeRedemptionListList []ClientPromoCodeRedemption

func (dto ClientPromoCodeRedemptionListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ClientPromoCodeRedemption(dto))
}
