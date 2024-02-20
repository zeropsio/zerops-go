// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type UserDataList struct {
	UserDataList UserDataListUserDataList `json:"userDataList"`
}

type UserDataListUserDataList []UserData

func (dto UserDataListUserDataList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserData(dto))
}
