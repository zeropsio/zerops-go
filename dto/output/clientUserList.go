// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type ClientUserList struct {
	ClientUserList ClientUserListClientUserList `json:"clientUserList"`
}

type ClientUserListClientUserList []ClientUserExtra

func (dto ClientUserListClientUserList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ClientUserExtra(dto))
}
