// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type ServiceStackEnvList struct {
	Items ServiceStackEnvListItems `json:"items"`
}

type ServiceStackEnvListItems []ServiceStackEnv

func (dto ServiceStackEnvListItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackEnv(dto))
}
