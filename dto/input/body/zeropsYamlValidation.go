// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ZeropsYamlValidation)(nil)

type ZeropsYamlValidation struct {
	Name               types.String                `json:"name"`
	ServiceStackTypeId stringId.ServiceStackTypeId `json:"serviceStackTypeId"`
	ZeropsYaml         types.Text                  `json:"zeropsYaml"`
}

func (dto ZeropsYamlValidation) GetName() types.String {
	return dto.Name
}
func (dto ZeropsYamlValidation) GetServiceStackTypeId() stringId.ServiceStackTypeId {
	return dto.ServiceStackTypeId
}
func (dto ZeropsYamlValidation) GetZeropsYaml() types.Text {
	return dto.ZeropsYaml
}

func (dto *ZeropsYamlValidation) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Name               *types.String
		ServiceStackTypeId *stringId.ServiceStackTypeId
		ZeropsYaml         *types.Text
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ZeropsYamlValidation", err)
	}
	var errorList validator.ErrorList
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if aux.ServiceStackTypeId == nil {
		errorList = errorList.With(validator.NewError("serviceStackTypeId", "field is required"))
	}
	if aux.ZeropsYaml == nil {
		errorList = errorList.With(validator.NewError("zeropsYaml", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Name = *aux.Name
	dto.ServiceStackTypeId = *aux.ServiceStackTypeId
	dto.ZeropsYaml = *aux.ZeropsYaml

	return nil
}
