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
var _ json.Unmarshaler = (*PostStandardServiceStack)(nil)

type PostStandardServiceStack struct {
	Name                        types.String                     `json:"name"`
	CustomAutoscaling           *CustomAutoscaling               `json:"customAutoscaling"`
	AutoscalingProfileId        types.StringNull                 `json:"autoscalingProfileId"`
	AutoscalingProfileOverrides types.JsonRawMessage             `json:"autoscalingProfileOverrides"`
	UserData                    PostStandardServiceStackUserData `json:"userData"`
	UserDataEnvFile             types.TextNull                   `json:"userDataEnvFile"`
	StartWithoutCode            types.BoolNull                   `json:"startWithoutCode"`
	BuildFromGit                types.StringNull                 `json:"buildFromGit"`
	ZeropsSetup                 types.StringNull                 `json:"zeropsSetup"`
	ZeropsYaml                  types.TextNull                   `json:"zeropsYaml"`
	EnvIsolation                types.StringNull                 `json:"envIsolation"`
	SshIsolation                types.StringNull                 `json:"sshIsolation"`
	EnableSubdomainAccess       types.BoolNull                   `json:"enableSubdomainAccess"`
	CdnEnabled                  types.BoolNull                   `json:"cdnEnabled"`
	Location                    stringId.LocationIdNull          `json:"location"`
	Mode                        types.StringNull                 `json:"mode"`
	Os                          types.StringNull                 `json:"os"`
}

func (dto PostStandardServiceStack) GetName() types.String {
	return dto.Name
}
func (dto PostStandardServiceStack) GetCustomAutoscaling() *CustomAutoscaling {
	return dto.CustomAutoscaling
}
func (dto PostStandardServiceStack) GetAutoscalingProfileId() types.StringNull {
	return dto.AutoscalingProfileId
}
func (dto PostStandardServiceStack) GetAutoscalingProfileOverrides() types.JsonRawMessage {
	return dto.AutoscalingProfileOverrides
}
func (dto PostStandardServiceStack) GetUserData() PostStandardServiceStackUserData {
	return dto.UserData
}
func (dto PostStandardServiceStack) GetUserDataEnvFile() types.TextNull {
	return dto.UserDataEnvFile
}
func (dto PostStandardServiceStack) GetStartWithoutCode() types.BoolNull {
	return dto.StartWithoutCode
}
func (dto PostStandardServiceStack) GetBuildFromGit() types.StringNull {
	return dto.BuildFromGit
}
func (dto PostStandardServiceStack) GetZeropsSetup() types.StringNull {
	return dto.ZeropsSetup
}
func (dto PostStandardServiceStack) GetZeropsYaml() types.TextNull {
	return dto.ZeropsYaml
}
func (dto PostStandardServiceStack) GetEnvIsolation() types.StringNull {
	return dto.EnvIsolation
}
func (dto PostStandardServiceStack) GetSshIsolation() types.StringNull {
	return dto.SshIsolation
}
func (dto PostStandardServiceStack) GetEnableSubdomainAccess() types.BoolNull {
	return dto.EnableSubdomainAccess
}
func (dto PostStandardServiceStack) GetCdnEnabled() types.BoolNull {
	return dto.CdnEnabled
}
func (dto PostStandardServiceStack) GetLocation() stringId.LocationIdNull {
	return dto.Location
}
func (dto PostStandardServiceStack) GetMode() types.StringNull {
	return dto.Mode
}
func (dto PostStandardServiceStack) GetOs() types.StringNull {
	return dto.Os
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
		Name                        *types.String
		CustomAutoscaling           *CustomAutoscaling
		AutoscalingProfileId        types.StringNull
		AutoscalingProfileOverrides types.JsonRawMessage
		UserData                    *PostStandardServiceStackUserData
		UserDataEnvFile             types.TextNull
		StartWithoutCode            types.BoolNull
		BuildFromGit                types.StringNull
		ZeropsSetup                 types.StringNull
		ZeropsYaml                  types.TextNull
		EnvIsolation                types.StringNull
		SshIsolation                types.StringNull
		EnableSubdomainAccess       types.BoolNull
		CdnEnabled                  types.BoolNull
		Location                    stringId.LocationIdNull
		Mode                        types.StringNull
		Os                          types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostStandardServiceStack", err)
	}
	var errorList validator.ErrorList
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if aux.UserData == nil {
		errorList = errorList.With(validator.NewError("userData", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Name = *aux.Name
	dto.CustomAutoscaling = aux.CustomAutoscaling
	dto.AutoscalingProfileId = aux.AutoscalingProfileId
	dto.AutoscalingProfileOverrides = aux.AutoscalingProfileOverrides
	dto.UserData = *aux.UserData
	dto.UserDataEnvFile = aux.UserDataEnvFile
	dto.StartWithoutCode = aux.StartWithoutCode
	dto.BuildFromGit = aux.BuildFromGit
	dto.ZeropsSetup = aux.ZeropsSetup
	dto.ZeropsYaml = aux.ZeropsYaml
	dto.EnvIsolation = aux.EnvIsolation
	dto.SshIsolation = aux.SshIsolation
	dto.EnableSubdomainAccess = aux.EnableSubdomainAccess
	dto.CdnEnabled = aux.CdnEnabled
	dto.Location = aux.Location
	dto.Mode = aux.Mode
	dto.Os = aux.Os

	return nil
}
