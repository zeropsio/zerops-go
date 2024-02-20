// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ProjectVpnUserSetup struct {
	PublicKey types.String `json:"publicKey"`
	Ipv4      VpnIpConfig  `json:"ipv4"`
	Ipv6      VpnIpConfig  `json:"ipv6"`
	Setup     types.String `json:"setup"`
}
