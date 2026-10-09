// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type ProjectEnvList struct {
	Items ProjectEnvListItems `json:"items"`
}

type ProjectEnvListItems []ProjectEnv

func (dto ProjectEnvListItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProjectEnv(dto))
}
