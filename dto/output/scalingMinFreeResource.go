// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ScalingMinFreeResource struct {
	CpuCoreCount   types.Float `json:"cpuCoreCount"`
	CpuCorePercent types.Float `json:"cpuCorePercent"`
	MemoryGBytes   types.Float `json:"memoryGBytes"`
	MemoryPercent  types.Float `json:"memoryPercent"`
}
