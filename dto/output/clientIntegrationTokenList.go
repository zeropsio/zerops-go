// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type ClientIntegrationTokenList struct {
	List ClientIntegrationTokenListList `json:"list"`
}

type ClientIntegrationTokenListList []ClientIntegrationToken

func (dto ClientIntegrationTokenListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ClientIntegrationToken(dto))
}
