// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type HardwareResource struct {
	CpuCoreCount types.Int `json:"cpuCoreCount"`
	MemoryMBytes types.Int `json:"memoryMBytes"`
	DiskGBytes   types.Int `json:"diskGBytes"`
}
