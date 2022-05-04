// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostStandardServiceStack)(nil)

type PostStandardServiceStack struct {
	ProjectId         uuid.ProjectId            `json:"projectId"`
	Name              types.String              `json:"name"`
	Mode              enum.ServiceStackModeEnum `json:"mode"`
	CustomAutoscaling *CustomAutoscaling        `json:"customAutoscaling"`
}

func (dto PostStandardServiceStack) GetProjectId() uuid.ProjectId {
	return dto.ProjectId
}
func (dto PostStandardServiceStack) GetName() types.String {
	return dto.Name
}
func (dto PostStandardServiceStack) GetMode() enum.ServiceStackModeEnum {
	return dto.Mode
}
func (dto PostStandardServiceStack) GetCustomAutoscaling() *CustomAutoscaling {
	return dto.CustomAutoscaling
}

func (dto *PostStandardServiceStack) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ProjectId         *uuid.ProjectId
		Name              *types.String
		Mode              *enum.ServiceStackModeEnum
		CustomAutoscaling *CustomAutoscaling
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostStandardServiceStack", err)
	}
	var errorList validator.ErrorList
	if aux.ProjectId == nil {
		errorList = errorList.With(validator.NewError("projectId", "field is required"))
	}
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if aux.Mode == nil {
		errorList = errorList.With(validator.NewError("mode", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ProjectId = *aux.ProjectId
	dto.Name = *aux.Name
	dto.Mode = *aux.Mode
	dto.CustomAutoscaling = aux.CustomAutoscaling

	return nil
}
