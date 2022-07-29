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
var _ json.Unmarshaler = (*PostUserServiceStack)(nil)

type PostUserServiceStack struct {
	ProjectId         uuid.ProjectId               `json:"projectId"`
	Name              types.String                 `json:"name"`
	Mode              enum.ServiceStackModeEnum    `json:"mode"`
	CustomAutoscaling *CustomAutoscaling           `json:"customAutoscaling"`
	UserData          PostUserServiceStackUserData `json:"userData"`
}

func (dto PostUserServiceStack) GetProjectId() uuid.ProjectId {
	return dto.ProjectId
}
func (dto PostUserServiceStack) GetName() types.String {
	return dto.Name
}
func (dto PostUserServiceStack) GetMode() enum.ServiceStackModeEnum {
	return dto.Mode
}
func (dto PostUserServiceStack) GetCustomAutoscaling() *CustomAutoscaling {
	return dto.CustomAutoscaling
}
func (dto PostUserServiceStack) GetUserData() PostUserServiceStackUserData {
	return dto.UserData
}

type PostUserServiceStackUserData []UserDataPut

func (dto PostUserServiceStackUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserDataPut(dto))
}

func (dto *PostUserServiceStack) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ProjectId         *uuid.ProjectId
		Name              *types.String
		Mode              *enum.ServiceStackModeEnum
		CustomAutoscaling *CustomAutoscaling
		UserData          *PostUserServiceStackUserData
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostUserServiceStack", err)
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
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ProjectId = *aux.ProjectId
	dto.Name = *aux.Name
	dto.Mode = *aux.Mode
	dto.CustomAutoscaling = aux.CustomAutoscaling
	dto.UserData = *aux.UserData

	return nil
}
