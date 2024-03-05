// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type ProjectVpnList struct {
	Peers   ProjectVpnListPeers `json:"peers"`
	Project ProjectVpnSetup     `json:"project"`
}

type ProjectVpnListPeers []ProjectVpnPeerSetup

func (dto ProjectVpnListPeers) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProjectVpnPeerSetup(dto))
}
