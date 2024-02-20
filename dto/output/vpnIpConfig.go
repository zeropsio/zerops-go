// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type VpnIpConfig struct {
	Network           VpnNetworkConfig `json:"network"`
	AssignedIpAddress types.String     `json:"assignedIpAddress"`
}
