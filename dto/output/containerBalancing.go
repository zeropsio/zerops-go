// Generated ZEROPS sdk

package output

import (
	"strconv"
)

var _ strconv.NumError

type ContainerBalancing struct {
	Available ContainerBalancingKinds `json:"available"`
	User      ContainerBalancingKinds `json:"user"`
	App       ContainerBalancingKinds `json:"app"`
}
