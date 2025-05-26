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
var _ json.Unmarshaler = (*PutStandardServiceStackTriggerPipeline)(nil)

type PutStandardServiceStackTriggerPipeline struct {
	CustomAutoscaling     *CustomAutoscaling                             `json:"customAutoscaling"`
	Mode                  *enum.ServiceStackModeEnum                     `json:"mode"`
	AppVersionId          uuid.AppVersionIdNull                          `json:"appVersionId"`
	UserData              PutStandardServiceStackTriggerPipelineUserData `json:"userData"`
	UserDataEnvFile       types.TextNull                                 `json:"userDataEnvFile"`
	StartWithoutCode      types.BoolNull                                 `json:"startWithoutCode"`
	BuildFromGit          types.StringNull                               `json:"buildFromGit"`
	ZeropsSetup           types.StringNull                               `json:"zeropsSetup"`
	ZeropsYaml            types.TextNull                                 `json:"zeropsYaml"`
	EnableSubdomainAccess types.BoolNull                                 `json:"enableSubdomainAccess"`
}

func (dto PutStandardServiceStackTriggerPipeline) GetCustomAutoscaling() *CustomAutoscaling {
	return dto.CustomAutoscaling
}
func (dto PutStandardServiceStackTriggerPipeline) GetMode() *enum.ServiceStackModeEnum {
	return dto.Mode
}
func (dto PutStandardServiceStackTriggerPipeline) GetAppVersionId() uuid.AppVersionIdNull {
	return dto.AppVersionId
}
func (dto PutStandardServiceStackTriggerPipeline) GetUserData() PutStandardServiceStackTriggerPipelineUserData {
	return dto.UserData
}
func (dto PutStandardServiceStackTriggerPipeline) GetUserDataEnvFile() types.TextNull {
	return dto.UserDataEnvFile
}
func (dto PutStandardServiceStackTriggerPipeline) GetStartWithoutCode() types.BoolNull {
	return dto.StartWithoutCode
}
func (dto PutStandardServiceStackTriggerPipeline) GetBuildFromGit() types.StringNull {
	return dto.BuildFromGit
}
func (dto PutStandardServiceStackTriggerPipeline) GetZeropsSetup() types.StringNull {
	return dto.ZeropsSetup
}
func (dto PutStandardServiceStackTriggerPipeline) GetZeropsYaml() types.TextNull {
	return dto.ZeropsYaml
}
func (dto PutStandardServiceStackTriggerPipeline) GetEnableSubdomainAccess() types.BoolNull {
	return dto.EnableSubdomainAccess
}

type PutStandardServiceStackTriggerPipelineUserData []UserDataPut

func (dto PutStandardServiceStackTriggerPipelineUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserDataPut(dto))
}

func (dto *PutStandardServiceStackTriggerPipeline) UnmarshalJSON(b []byte) error {
	var aux = struct {
		CustomAutoscaling     *CustomAutoscaling
		Mode                  *enum.ServiceStackModeEnum
		AppVersionId          uuid.AppVersionIdNull
		UserData              *PutStandardServiceStackTriggerPipelineUserData
		UserDataEnvFile       types.TextNull
		StartWithoutCode      types.BoolNull
		BuildFromGit          types.StringNull
		ZeropsSetup           types.StringNull
		ZeropsYaml            types.TextNull
		EnableSubdomainAccess types.BoolNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutStandardServiceStackTriggerPipeline", err)
	}
	var errorList validator.ErrorList
	if aux.UserData == nil {
		errorList = errorList.With(validator.NewError("userData", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.CustomAutoscaling = aux.CustomAutoscaling
	dto.Mode = aux.Mode
	dto.AppVersionId = aux.AppVersionId
	dto.UserData = *aux.UserData
	dto.UserDataEnvFile = aux.UserDataEnvFile
	dto.StartWithoutCode = aux.StartWithoutCode
	dto.BuildFromGit = aux.BuildFromGit
	dto.ZeropsSetup = aux.ZeropsSetup
	dto.ZeropsYaml = aux.ZeropsYaml
	dto.EnableSubdomainAccess = aux.EnableSubdomainAccess

	return nil
}
