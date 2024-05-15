// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type VerticalAutoscaling struct {
	MaxResource       ScalingResource                     `json:"maxResource"`
	MinResource       ScalingResource                     `json:"minResource"`
	MinFreeResource   ScalingMinFreeResource              `json:"minFreeResource"`
	CpuMode           enum.VerticalAutoscalingCpuModeEnum `json:"cpuMode"`
	StartCpuCoreCount types.Int                           `json:"startCpuCoreCount"`
}
