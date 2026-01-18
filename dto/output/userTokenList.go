// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type UserTokenList struct {
	List UserTokenListList `json:"list"`
}

type UserTokenListList []UserToken

func (dto UserTokenListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserToken(dto))
}
