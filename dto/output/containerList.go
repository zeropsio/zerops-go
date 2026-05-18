// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ContainerList struct {
	List       ContainerListList `json:"list"`
	TotalCount types.Int         `json:"totalCount"`
}

type ContainerListList []Container

func (dto ContainerListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Container(dto))
}
