// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type VpnConfig struct {
	Network        VpnNetworkConfig `json:"network"`
	SharedEndpoint types.String     `json:"sharedEndpoint"`
	Endpoint       types.String     `json:"endpoint"`
}
