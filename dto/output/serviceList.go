// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ServiceList struct {
	Count  types.Int       `json:"count"`
	Total  types.Int       `json:"total"`
	Limit  types.Int       `json:"limit"`
	Offset types.Int       `json:"offset"`
	List   ServiceListList `json:"list"`
}

type ServiceListList []Service

func (dto ServiceListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Service(dto))
}
