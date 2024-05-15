// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type HorizontalAutoscalingNullable struct {
	MaxContainerCount types.IntNull `json:"maxContainerCount"`
	MinContainerCount types.IntNull `json:"minContainerCount"`
}
