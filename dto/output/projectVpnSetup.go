// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ProjectVpnSetup struct {
	PublicKey types.String `json:"publicKey"`
	Ipv4      VpnConfig    `json:"ipv4"`
	Ipv6      VpnConfig    `json:"ipv6"`
}
