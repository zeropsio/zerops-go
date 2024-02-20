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
	CustomAutoscaling *CustomAutoscaling        `json:"customAutoscaling"`
	Mode              enum.ServiceStackModeEnum `json:"mode"`
}

func (dto PostStandardServiceStack) GetProjectId() uuid.ProjectId {
	return dto.ProjectId
}
func (dto PostStandardServiceStack) GetName() types.String {
	return dto.Name
}
func (dto PostStandardServiceStack) GetCustomAutoscaling() *CustomAutoscaling {
	return dto.CustomAutoscaling
}
func (dto PostStandardServiceStack) GetMode() enum.ServiceStackModeEnum {
	return dto.Mode
}

func (dto *PostStandardServiceStack) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ProjectId         *uuid.ProjectId
		Name              *types.String
		CustomAutoscaling *CustomAutoscaling
		Mode              *enum.ServiceStackModeEnum
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
	dto.CustomAutoscaling = aux.CustomAutoscaling
	dto.Mode = *aux.Mode

	return nil
}
