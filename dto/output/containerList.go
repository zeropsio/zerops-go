// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ContainerList struct {
	Count  types.Int         `json:"count"`
	Total  types.Int         `json:"total"`
	Limit  types.Int         `json:"limit"`
	Offset types.Int         `json:"offset"`
	List   ContainerListList `json:"list"`
}

type ContainerListList []Container

func (dto ContainerListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Container(dto))
}
