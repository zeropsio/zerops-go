// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*CustomAutoscaling)(nil)

type CustomAutoscaling struct {
	VerticalAutoscaling   *VerticalAutoscalingNullable   `json:"verticalAutoscaling"`
	HorizontalAutoscaling *HorizontalAutoscalingNullable `json:"horizontalAutoscaling"`
}

func (dto CustomAutoscaling) GetVerticalAutoscaling() *VerticalAutoscalingNullable {
	return dto.VerticalAutoscaling
}
func (dto CustomAutoscaling) GetHorizontalAutoscaling() *HorizontalAutoscalingNullable {
	return dto.HorizontalAutoscaling
}

func (dto *CustomAutoscaling) UnmarshalJSON(b []byte) error {
	var aux = struct {
		VerticalAutoscaling   *VerticalAutoscalingNullable
		HorizontalAutoscaling *HorizontalAutoscalingNullable
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("CustomAutoscaling", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.VerticalAutoscaling = aux.VerticalAutoscaling
	dto.HorizontalAutoscaling = aux.HorizontalAutoscaling

	return nil
}
