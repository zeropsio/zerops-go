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
	VerticalAutoscaling   *VerticalAutoscaling   `json:"verticalAutoscaling"`
	HorizontalAutoscaling *HorizontalAutoscaling `json:"horizontalAutoscaling"`
}

func (dto CustomAutoscaling) GetVerticalAutoscaling() *VerticalAutoscaling {
	return dto.VerticalAutoscaling
}
func (dto CustomAutoscaling) GetHorizontalAutoscaling() *HorizontalAutoscaling {
	return dto.HorizontalAutoscaling
}

func (dto *CustomAutoscaling) UnmarshalJSON(b []byte) error {
	var aux = struct {
		VerticalAutoscaling   *VerticalAutoscaling
		HorizontalAutoscaling *HorizontalAutoscaling
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
