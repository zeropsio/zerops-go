// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type VpnNetworkConfig struct {
	Network types.String `json:"network"`
	Gateway types.String `json:"gateway"`
}
