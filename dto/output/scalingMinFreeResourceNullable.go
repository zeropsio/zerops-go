// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ScalingMinFreeResourceNullable struct {
	CpuCoreCount   types.FloatNull `json:"cpuCoreCount"`
	CpuCorePercent types.FloatNull `json:"cpuCorePercent"`
	MemoryGBytes   types.FloatNull `json:"memoryGBytes"`
	MemoryPercent  types.FloatNull `json:"memoryPercent"`
}
