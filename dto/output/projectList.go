// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ProjectList struct {
	List       ProjectListList `json:"list"`
	TotalCount types.Int       `json:"totalCount"`
}

type ProjectListList []Project

func (dto ProjectListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Project(dto))
}
