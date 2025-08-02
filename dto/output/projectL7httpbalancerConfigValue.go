// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ProjectL7httpbalancerConfigValue struct {
	Name    types.String     `json:"name"`
	Current types.StringNull `json:"current"`
	Default types.String     `json:"default"`
}
