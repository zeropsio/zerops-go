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
	ProjectId         uuid.ProjectId                   `json:"projectId"`
	Name              types.String                     `json:"name"`
	StartWithoutCode  types.Bool                       `json:"startWithoutCode"`
	CustomAutoscaling *CustomAutoscaling               `json:"customAutoscaling"`
	Mode              *enum.ServiceStackModeEnum       `json:"mode"`
	UserData          PostStandardServiceStackUserData `json:"userData"`
	UserDataEnvFile   types.TextNull                   `json:"userDataEnvFile"`
}

func (dto PostStandardServiceStack) GetProjectId() uuid.ProjectId {
	return dto.ProjectId
}
func (dto PostStandardServiceStack) GetName() types.String {
	return dto.Name
}
func (dto PostStandardServiceStack) GetStartWithoutCode() types.Bool {
	return dto.StartWithoutCode
}
func (dto PostStandardServiceStack) GetCustomAutoscaling() *CustomAutoscaling {
	return dto.CustomAutoscaling
}
func (dto PostStandardServiceStack) GetMode() *enum.ServiceStackModeEnum {
	return dto.Mode
}
func (dto PostStandardServiceStack) GetUserData() PostStandardServiceStackUserData {
	return dto.UserData
}
func (dto PostStandardServiceStack) GetUserDataEnvFile() types.TextNull {
	return dto.UserDataEnvFile
}

type PostStandardServiceStackUserData []UserDataPut

func (dto PostStandardServiceStackUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserDataPut(dto))
}

func (dto *PostStandardServiceStack) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ProjectId         *uuid.ProjectId
		Name              *types.String
		StartWithoutCode  *types.Bool
		CustomAutoscaling *CustomAutoscaling
		Mode              *enum.ServiceStackModeEnum
		UserData          *PostStandardServiceStackUserData
		UserDataEnvFile   types.TextNull
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
	if aux.StartWithoutCode == nil {
		errorList = errorList.With(validator.NewError("startWithoutCode", "field is required"))
	}
	if aux.UserData == nil {
		errorList = errorList.With(validator.NewError("userData", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ProjectId = *aux.ProjectId
	dto.Name = *aux.Name
	dto.StartWithoutCode = *aux.StartWithoutCode
	dto.CustomAutoscaling = aux.CustomAutoscaling
	dto.Mode = aux.Mode
	dto.UserData = *aux.UserData
	dto.UserDataEnvFile = aux.UserDataEnvFile

	return nil
}
