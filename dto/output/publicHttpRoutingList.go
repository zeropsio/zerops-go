// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type PublicHttpRoutingList struct {
	List       PublicHttpRoutingListList `json:"list"`
	TotalCount types.Int                 `json:"totalCount"`
}

type PublicHttpRoutingListList []PublicHttpRouting

func (dto PublicHttpRoutingListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicHttpRouting(dto))
}
