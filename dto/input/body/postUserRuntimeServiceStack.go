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
var _ json.Unmarshaler = (*PostUserRuntimeServiceStack)(nil)

type PostUserRuntimeServiceStack struct {
	ProjectId         uuid.ProjectId                      `json:"projectId"`
	Name              types.String                        `json:"name"`
	Mode              enum.ServiceStackModeEnum           `json:"mode"`
	CustomAutoscaling *CustomAutoscaling                  `json:"customAutoscaling"`
	UserData          PostUserRuntimeServiceStackUserData `json:"userData"`
	Ports             PostUserRuntimeServiceStackPorts    `json:"ports"`
}

func (dto PostUserRuntimeServiceStack) GetProjectId() uuid.ProjectId {
	return dto.ProjectId
}
func (dto PostUserRuntimeServiceStack) GetName() types.String {
	return dto.Name
}
func (dto PostUserRuntimeServiceStack) GetMode() enum.ServiceStackModeEnum {
	return dto.Mode
}
func (dto PostUserRuntimeServiceStack) GetCustomAutoscaling() *CustomAutoscaling {
	return dto.CustomAutoscaling
}
func (dto PostUserRuntimeServiceStack) GetUserData() PostUserRuntimeServiceStackUserData {
	return dto.UserData
}
func (dto PostUserRuntimeServiceStack) GetPorts() PostUserRuntimeServiceStackPorts {
	return dto.Ports
}

type PostUserRuntimeServiceStackUserData []UserDataPut

func (dto PostUserRuntimeServiceStackUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserDataPut(dto))
}

type PostUserRuntimeServiceStackPorts []PutServiceStackServicePortLight

func (dto PostUserRuntimeServiceStackPorts) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PutServiceStackServicePortLight(dto))
}

func (dto *PostUserRuntimeServiceStack) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ProjectId         *uuid.ProjectId
		Name              *types.String
		Mode              *enum.ServiceStackModeEnum
		CustomAutoscaling *CustomAutoscaling
		UserData          *PostUserRuntimeServiceStackUserData
		Ports             *PostUserRuntimeServiceStackPorts
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostUserRuntimeServiceStack", err)
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
	if aux.UserData == nil {
		errorList = errorList.With(validator.NewError("userData", "field is required"))
	}
	if aux.Ports == nil {
		errorList = errorList.With(validator.NewError("ports", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ProjectId = *aux.ProjectId
	dto.Name = *aux.Name
	dto.Mode = *aux.Mode
	dto.CustomAutoscaling = aux.CustomAutoscaling
	dto.UserData = *aux.UserData
	dto.Ports = *aux.Ports

	return nil
}
