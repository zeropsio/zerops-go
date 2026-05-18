// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ServiceList struct {
	List       ServiceListList `json:"list"`
	TotalCount types.Int       `json:"totalCount"`
}

type ServiceListList []Service

func (dto ServiceListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Service(dto))
}
