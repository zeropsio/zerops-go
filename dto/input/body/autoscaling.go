// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*Autoscaling)(nil)

type Autoscaling struct {
	CustomAutoscaling           *CustomAutoscaling   `json:"customAutoscaling"`
	AutoscalingProfileId        types.StringNull     `json:"autoscalingProfileId"`
	AutoscalingProfileOverrides types.JsonRawMessage `json:"autoscalingProfileOverrides"`
}

func (dto Autoscaling) GetCustomAutoscaling() *CustomAutoscaling {
	return dto.CustomAutoscaling
}
func (dto Autoscaling) GetAutoscalingProfileId() types.StringNull {
	return dto.AutoscalingProfileId
}
func (dto Autoscaling) GetAutoscalingProfileOverrides() types.JsonRawMessage {
	return dto.AutoscalingProfileOverrides
}

func (dto *Autoscaling) UnmarshalJSON(b []byte) error {
	var aux = struct {
		CustomAutoscaling           *CustomAutoscaling
		AutoscalingProfileId        types.StringNull
		AutoscalingProfileOverrides types.JsonRawMessage
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("Autoscaling", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.CustomAutoscaling = aux.CustomAutoscaling
	dto.AutoscalingProfileId = aux.AutoscalingProfileId
	dto.AutoscalingProfileOverrides = aux.AutoscalingProfileOverrides

	return nil
}
