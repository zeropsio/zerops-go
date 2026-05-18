// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type PublicPortRoutingList struct {
	List       PublicPortRoutingListList `json:"list"`
	TotalCount types.Int                 `json:"totalCount"`
}

type PublicPortRoutingListList []PublicPortRouting

func (dto PublicPortRoutingListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicPortRouting(dto))
}
