// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ScalingResourceNullable)(nil)

type ScalingResourceNullable struct {
	CpuCoreCount types.IntNull   `json:"cpuCoreCount"`
	MemoryGBytes types.FloatNull `json:"memoryGBytes"`
	DiskGBytes   types.FloatNull `json:"diskGBytes"`
}

func (dto ScalingResourceNullable) GetCpuCoreCount() types.IntNull {
	return dto.CpuCoreCount
}
func (dto ScalingResourceNullable) GetMemoryGBytes() types.FloatNull {
	return dto.MemoryGBytes
}
func (dto ScalingResourceNullable) GetDiskGBytes() types.FloatNull {
	return dto.DiskGBytes
}

func (dto *ScalingResourceNullable) UnmarshalJSON(b []byte) error {
	var aux = struct {
		CpuCoreCount types.IntNull
		MemoryGBytes types.FloatNull
		DiskGBytes   types.FloatNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ScalingResourceNullable", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.CpuCoreCount = aux.CpuCoreCount
	dto.MemoryGBytes = aux.MemoryGBytes
	dto.DiskGBytes = aux.DiskGBytes

	return nil
}
