// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ScalingResource)(nil)

type ScalingResource struct {
	CpuCoreCount types.Int   `json:"cpuCoreCount"`
	MemoryGBytes types.Float `json:"memoryGBytes"`
	DiskGBytes   types.Float `json:"diskGBytes"`
}

func (dto ScalingResource) GetCpuCoreCount() types.Int {
	return dto.CpuCoreCount
}
func (dto ScalingResource) GetMemoryGBytes() types.Float {
	return dto.MemoryGBytes
}
func (dto ScalingResource) GetDiskGBytes() types.Float {
	return dto.DiskGBytes
}

func (dto *ScalingResource) UnmarshalJSON(b []byte) error {
	var aux = struct {
		CpuCoreCount *types.Int
		MemoryGBytes *types.Float
		DiskGBytes   *types.Float
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ScalingResource", err)
	}
	var errorList validator.ErrorList
	if aux.CpuCoreCount == nil {
		errorList = errorList.With(validator.NewError("cpuCoreCount", "field is required"))
	}
	if aux.MemoryGBytes == nil {
		errorList = errorList.With(validator.NewError("memoryGBytes", "field is required"))
	}
	if aux.DiskGBytes == nil {
		errorList = errorList.With(validator.NewError("diskGBytes", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.CpuCoreCount = *aux.CpuCoreCount
	dto.MemoryGBytes = *aux.MemoryGBytes
	dto.DiskGBytes = *aux.DiskGBytes

	return nil
}
