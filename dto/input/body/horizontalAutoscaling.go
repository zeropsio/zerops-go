// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*HorizontalAutoscaling)(nil)

type HorizontalAutoscaling struct {
	MaxContainerCount types.Int `json:"maxContainerCount"`
	MinContainerCount types.Int `json:"minContainerCount"`
}

func (dto HorizontalAutoscaling) GetMaxContainerCount() types.Int {
	return dto.MaxContainerCount
}
func (dto HorizontalAutoscaling) GetMinContainerCount() types.Int {
	return dto.MinContainerCount
}

func (dto *HorizontalAutoscaling) UnmarshalJSON(b []byte) error {
	var aux = struct {
		MaxContainerCount *types.Int
		MinContainerCount *types.Int
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("HorizontalAutoscaling", err)
	}
	var errorList validator.ErrorList
	if aux.MaxContainerCount == nil {
		errorList = errorList.With(validator.NewError("maxContainerCount", "field is required"))
	}
	if aux.MinContainerCount == nil {
		errorList = errorList.With(validator.NewError("minContainerCount", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.MaxContainerCount = *aux.MaxContainerCount
	dto.MinContainerCount = *aux.MinContainerCount

	return nil
}
