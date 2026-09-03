// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ProcessList struct {
	Count  types.Int       `json:"count"`
	Total  types.Int       `json:"total"`
	Limit  types.Int       `json:"limit"`
	Offset types.Int       `json:"offset"`
	List   ProcessListList `json:"list"`
}

type ProcessListList []Process

func (dto ProcessListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Process(dto))
}
