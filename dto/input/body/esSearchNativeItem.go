// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*EsSearchNativeItem)(nil)

type EsSearchNativeItem struct {
	Name     types.String         `json:"name"`
	Operator types.String         `json:"operator"`
	Value    types.JsonRawMessage `json:"value"`
}

func (dto EsSearchNativeItem) GetName() types.String {
	return dto.Name
}
func (dto EsSearchNativeItem) GetOperator() types.String {
	return dto.Operator
}
func (dto EsSearchNativeItem) GetValue() types.JsonRawMessage {
	return dto.Value
}

func (dto *EsSearchNativeItem) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Name     *types.String
		Operator *types.String
		Value    *types.JsonRawMessage
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("EsSearchNativeItem", err)
	}
	var errorList validator.ErrorList
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if aux.Operator == nil {
		errorList = errorList.With(validator.NewError("operator", "field is required"))
	}
	if aux.Value == nil {
		errorList = errorList.With(validator.NewError("value", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Name = *aux.Name
	dto.Operator = *aux.Operator
	dto.Value = *aux.Value

	return nil
}
