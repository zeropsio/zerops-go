// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*EsSearchItem)(nil)

type EsSearchItem struct {
	Name     types.String `json:"name"`
	Operator types.String `json:"operator"`
	Value    types.String `json:"value"`
}

func (dto EsSearchItem) GetName() types.String {
	return dto.Name
}
func (dto EsSearchItem) GetOperator() types.String {
	return dto.Operator
}
func (dto EsSearchItem) GetValue() types.String {
	return dto.Value
}

func (dto *EsSearchItem) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Name     *types.String
		Operator *types.String
		Value    *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("EsSearchItem", err)
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
