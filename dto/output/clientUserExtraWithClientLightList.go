// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ClientUserExtraWithClientLightList struct {
	Count  types.Int                              `json:"count"`
	Total  types.Int                              `json:"total"`
	Limit  types.Int                              `json:"limit"`
	Offset types.Int                              `json:"offset"`
	List   ClientUserExtraWithClientLightListList `json:"list"`
}

type ClientUserExtraWithClientLightListList []ClientUserExtraWithClientLight

func (dto ClientUserExtraWithClientLightListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ClientUserExtraWithClientLight(dto))
}
