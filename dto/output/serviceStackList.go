// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ServiceStackList struct {
	List       ServiceStackListList `json:"list"`
	TotalCount types.Int            `json:"totalCount"`
}

type ServiceStackListList []ServiceStack

func (dto ServiceStackListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStack(dto))
}
