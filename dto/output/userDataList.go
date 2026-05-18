// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type UserDataList struct {
	List       UserDataListList `json:"list"`
	TotalCount types.Int        `json:"totalCount"`
}

type UserDataListList []UserData

func (dto UserDataListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserData(dto))
}
