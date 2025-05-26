// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*VerticalAutoscalingNullable)(nil)

type VerticalAutoscalingNullable struct {
	MaxResource       *ScalingResourceNullable             `json:"maxResource"`
	MinResource       *ScalingResourceNullable             `json:"minResource"`
	MinFreeResource   *ScalingMinFreeResourceNullable      `json:"minFreeResource"`
	CpuMode           *enum.VerticalAutoscalingCpuModeEnum `json:"cpuMode"`
	StartCpuCoreCount types.IntNull                        `json:"startCpuCoreCount"`
	SwapEnabled       types.BoolNull                       `json:"swapEnabled"`
}

func (dto VerticalAutoscalingNullable) GetMaxResource() *ScalingResourceNullable {
	return dto.MaxResource
}
func (dto VerticalAutoscalingNullable) GetMinResource() *ScalingResourceNullable {
	return dto.MinResource
}
func (dto VerticalAutoscalingNullable) GetMinFreeResource() *ScalingMinFreeResourceNullable {
	return dto.MinFreeResource
}
func (dto VerticalAutoscalingNullable) GetCpuMode() *enum.VerticalAutoscalingCpuModeEnum {
	return dto.CpuMode
}
func (dto VerticalAutoscalingNullable) GetStartCpuCoreCount() types.IntNull {
	return dto.StartCpuCoreCount
}
func (dto VerticalAutoscalingNullable) GetSwapEnabled() types.BoolNull {
	return dto.SwapEnabled
}

func (dto *VerticalAutoscalingNullable) UnmarshalJSON(b []byte) error {
	var aux = struct {
		MaxResource       *ScalingResourceNullable
		MinResource       *ScalingResourceNullable
		MinFreeResource   *ScalingMinFreeResourceNullable
		CpuMode           *enum.VerticalAutoscalingCpuModeEnum
		StartCpuCoreCount types.IntNull
		SwapEnabled       types.BoolNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("VerticalAutoscalingNullable", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.MaxResource = aux.MaxResource
	dto.MinResource = aux.MinResource
	dto.MinFreeResource = aux.MinFreeResource
	dto.CpuMode = aux.CpuMode
	dto.StartCpuCoreCount = aux.StartCpuCoreCount
	dto.SwapEnabled = aux.SwapEnabled

	return nil
}
