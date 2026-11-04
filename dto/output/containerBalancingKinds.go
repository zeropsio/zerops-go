// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ContainerBalancingKinds struct {
	L3  types.Bool `json:"l3"`
	L7  types.Bool `json:"l7"`
	Dns types.Bool `json:"dns"`
}
