// Generated ZEROPS sdk

package output

import (
	"strconv"
)

var _ strconv.NumError

type ProjectVpnItem struct {
	Project ProjectVpnSetup     `json:"project"`
	Peer    ProjectVpnPeerSetup `json:"peer"`
}
