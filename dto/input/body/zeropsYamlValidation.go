// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ZeropsYamlValidation)(nil)

type ZeropsYamlValidation struct {
	ServiceStackName            types.String                           `json:"serviceStackName"`
	ServiceStackTypeId          stringId.ServiceStackTypeId            `json:"serviceStackTypeId"`
	ServiceStackTypeVersionName types.String                           `json:"serviceStackTypeVersionName"`
	ZeropsYaml                  types.MediumText                       `json:"zeropsYaml"`
	Operation                   enum.ZeropsYamlValidationOperationEnum `json:"operation"`
}

func (dto ZeropsYamlValidation) GetServiceStackName() types.String {
	return dto.ServiceStackName
}
func (dto ZeropsYamlValidation) GetServiceStackTypeId() stringId.ServiceStackTypeId {
	return dto.ServiceStackTypeId
}
func (dto ZeropsYamlValidation) GetServiceStackTypeVersionName() types.String {
	return dto.ServiceStackTypeVersionName
}
func (dto ZeropsYamlValidation) GetZeropsYaml() types.MediumText {
	return dto.ZeropsYaml
}
func (dto ZeropsYamlValidation) GetOperation() enum.ZeropsYamlValidationOperationEnum {
	return dto.Operation
}

func (dto *ZeropsYamlValidation) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ServiceStackName            *types.String
		ServiceStackTypeId          *stringId.ServiceStackTypeId
		ServiceStackTypeVersionName *types.String
		ZeropsYaml                  *types.MediumText
		Operation                   *enum.ZeropsYamlValidationOperationEnum
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ZeropsYamlValidation", err)
	}
	var errorList validator.ErrorList
	if aux.ServiceStackName == nil {
		errorList = errorList.With(validator.NewError("serviceStackName", "field is required"))
	}
	if aux.ServiceStackTypeId == nil {
		errorList = errorList.With(validator.NewError("serviceStackTypeId", "field is required"))
	}
	if aux.ServiceStackTypeVersionName == nil {
		errorList = errorList.With(validator.NewError("serviceStackTypeVersionName", "field is required"))
	}
	if aux.ZeropsYaml == nil {
		errorList = errorList.With(validator.NewError("zeropsYaml", "field is required"))
	}
	if aux.Operation == nil {
		errorList = errorList.With(validator.NewError("operation", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ServiceStackName = *aux.ServiceStackName
	dto.ServiceStackTypeId = *aux.ServiceStackTypeId
	dto.ServiceStackTypeVersionName = *aux.ServiceStackTypeVersionName
	dto.ZeropsYaml = *aux.ZeropsYaml
	dto.Operation = *aux.Operation

	return nil
}
