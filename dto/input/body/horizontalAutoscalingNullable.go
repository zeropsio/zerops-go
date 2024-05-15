// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*HorizontalAutoscalingNullable)(nil)

type HorizontalAutoscalingNullable struct {
	MaxContainerCount types.IntNull `json:"maxContainerCount"`
	MinContainerCount types.IntNull `json:"minContainerCount"`
}

func (dto HorizontalAutoscalingNullable) GetMaxContainerCount() types.IntNull {
	return dto.MaxContainerCount
}
func (dto HorizontalAutoscalingNullable) GetMinContainerCount() types.IntNull {
	return dto.MinContainerCount
}

func (dto *HorizontalAutoscalingNullable) UnmarshalJSON(b []byte) error {
	var aux = struct {
		MaxContainerCount types.IntNull
		MinContainerCount types.IntNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("HorizontalAutoscalingNullable", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.MaxContainerCount = aux.MaxContainerCount
	dto.MinContainerCount = aux.MinContainerCount

	return nil
}
