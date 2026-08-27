// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type ClientIntegrationTokenDelegationList struct {
	List ClientIntegrationTokenDelegationListList `json:"list"`
}

type ClientIntegrationTokenDelegationListList []ClientIntegrationTokenDelegation

func (dto ClientIntegrationTokenDelegationListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ClientIntegrationTokenDelegation(dto))
}
