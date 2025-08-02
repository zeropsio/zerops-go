// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type PublicHttpRoutingLocationAccessPolicy struct {
	Enabled       types.Bool                                                  `json:"enabled"`
	DefaultPolicy enum.PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum `json:"defaultPolicy"`
	Cidr          types.StringArray                                           `json:"cidr"`
}
