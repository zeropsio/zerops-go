// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type VerticalAutoscalingNullable struct {
	MaxResource       *ScalingResourceNullable             `json:"maxResource"`
	MinResource       *ScalingResourceNullable             `json:"minResource"`
	MinFreeResource   *ScalingMinFreeResourceNullable      `json:"minFreeResource"`
	CpuMode           *enum.VerticalAutoscalingCpuModeEnum `json:"cpuMode"`
	StartCpuCoreCount types.IntNull                        `json:"startCpuCoreCount"`
	SwapEnabled       types.BoolNull                       `json:"swapEnabled"`
}
