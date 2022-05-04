// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*VerticalAutoscaling)(nil)

type VerticalAutoscaling struct {
	MaxResource ScalingResource `json:"maxResource"`
	MinResource ScalingResource `json:"minResource"`
}

func (dto VerticalAutoscaling) GetMaxResource() ScalingResource {
	return dto.MaxResource
}
func (dto VerticalAutoscaling) GetMinResource() ScalingResource {
	return dto.MinResource
}

func (dto *VerticalAutoscaling) UnmarshalJSON(b []byte) error {
	var aux = struct {
		MaxResource *ScalingResource
		MinResource *ScalingResource
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("VerticalAutoscaling", err)
	}
	var errorList validator.ErrorList
	if aux.MaxResource == nil {
		errorList = errorList.With(validator.NewError("maxResource", "field is required"))
	}
	if aux.MinResource == nil {
		errorList = errorList.With(validator.NewError("minResource", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.MaxResource = *aux.MaxResource
	dto.MinResource = *aux.MinResource

	return nil
}
