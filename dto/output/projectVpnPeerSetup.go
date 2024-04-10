// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ProjectVpnPeerSetup struct {
	PublicKey    types.String `json:"publicKey"`
	Ipv4         VpnIpConfig  `json:"ipv4"`
	Ipv6         VpnIpConfig  `json:"ipv6"`
	Setup        types.String `json:"setup"`
	SetupLinux   types.String `json:"setupLinux"`
	SetupMacOS   types.String `json:"setupMacOS"`
	SetupWindows types.String `json:"setupWindows"`
}
