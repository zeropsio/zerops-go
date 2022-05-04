// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*Autoscaling)(nil)

type Autoscaling struct {
	Mode              enum.ServiceStackModeEnum `json:"mode"`
	CustomAutoscaling CustomAutoscaling         `json:"customAutoscaling"`
}

func (dto Autoscaling) GetMode() enum.ServiceStackModeEnum {
	return dto.Mode
}
func (dto Autoscaling) GetCustomAutoscaling() CustomAutoscaling {
	return dto.CustomAutoscaling
}

func (dto *Autoscaling) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Mode              *enum.ServiceStackModeEnum
		CustomAutoscaling *CustomAutoscaling
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("Autoscaling", err)
	}
	var errorList validator.ErrorList
	if aux.Mode == nil {
		errorList = errorList.With(validator.NewError("mode", "field is required"))
	}
	if aux.CustomAutoscaling == nil {
		errorList = errorList.With(validator.NewError("customAutoscaling", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Mode = *aux.Mode
	dto.CustomAutoscaling = *aux.CustomAutoscaling

	return nil
}
