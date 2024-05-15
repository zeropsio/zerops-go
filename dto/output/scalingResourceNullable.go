// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ScalingResourceNullable struct {
	CpuCoreCount types.IntNull   `json:"cpuCoreCount"`
	MemoryGBytes types.FloatNull `json:"memoryGBytes"`
	DiskGBytes   types.FloatNull `json:"diskGBytes"`
}
