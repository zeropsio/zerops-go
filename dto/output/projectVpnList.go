// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type ProjectVpnList struct {
	VpnUsers ProjectVpnListVpnUsers `json:"vpnUsers"`
	Project  ProjectVpnSetup        `json:"project"`
}

type ProjectVpnListVpnUsers []ProjectVpnUserSetup

func (dto ProjectVpnListVpnUsers) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProjectVpnUserSetup(dto))
}
