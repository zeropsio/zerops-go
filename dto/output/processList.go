// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ProcessList struct {
	List       ProcessListList `json:"list"`
	TotalCount types.Int       `json:"totalCount"`
}

type ProcessListList []Process

func (dto ProcessListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Process(dto))
}
