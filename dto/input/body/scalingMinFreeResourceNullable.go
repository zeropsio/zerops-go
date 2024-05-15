// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ScalingMinFreeResourceNullable)(nil)

type ScalingMinFreeResourceNullable struct {
	CpuCoreCount   types.FloatNull `json:"cpuCoreCount"`
	CpuCorePercent types.FloatNull `json:"cpuCorePercent"`
	MemoryGBytes   types.FloatNull `json:"memoryGBytes"`
	MemoryPercent  types.FloatNull `json:"memoryPercent"`
}

func (dto ScalingMinFreeResourceNullable) GetCpuCoreCount() types.FloatNull {
	return dto.CpuCoreCount
}
func (dto ScalingMinFreeResourceNullable) GetCpuCorePercent() types.FloatNull {
	return dto.CpuCorePercent
}
func (dto ScalingMinFreeResourceNullable) GetMemoryGBytes() types.FloatNull {
	return dto.MemoryGBytes
}
func (dto ScalingMinFreeResourceNullable) GetMemoryPercent() types.FloatNull {
	return dto.MemoryPercent
}

func (dto *ScalingMinFreeResourceNullable) UnmarshalJSON(b []byte) error {
	var aux = struct {
		CpuCoreCount   types.FloatNull
		CpuCorePercent types.FloatNull
		MemoryGBytes   types.FloatNull
		MemoryPercent  types.FloatNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ScalingMinFreeResourceNullable", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.CpuCoreCount = aux.CpuCoreCount
	dto.CpuCorePercent = aux.CpuCorePercent
	dto.MemoryGBytes = aux.MemoryGBytes
	dto.MemoryPercent = aux.MemoryPercent

	return nil
}
