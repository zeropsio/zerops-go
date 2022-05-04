// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ScalingResource struct {
	CpuCoreCount types.Int   `json:"cpuCoreCount"`
	MemoryGBytes types.Float `json:"memoryGBytes"`
	DiskGBytes   types.Float `json:"diskGBytes"`
}
