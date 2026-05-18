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
var _ json.Unmarshaler = (*PostServiceStack)(nil)

type PostServiceStack struct {
	Name                        types.String             `json:"name"`
	CustomAutoscaling           *CustomAutoscaling       `json:"customAutoscaling"`
	AutoscalingProfileId        types.StringNull         `json:"autoscalingProfileId"`
	AutoscalingProfileOverrides types.JsonRawMessage     `json:"autoscalingProfileOverrides"`
	UserData                    PostServiceStackUserData `json:"userData"`
	UserDataEnvFile             types.TextNull           `json:"userDataEnvFile"`
	StartWithoutCode            types.BoolNull           `json:"startWithoutCode"`
	BuildFromGit                types.StringNull         `json:"buildFromGit"`
	ZeropsSetup                 types.StringNull         `json:"zeropsSetup"`
	ZeropsYaml                  types.TextNull           `json:"zeropsYaml"`
	EnvIsolation                types.StringNull         `json:"envIsolation"`
	SshIsolation                types.StringNull         `json:"sshIsolation"`
	EnableSubdomainAccess       types.BoolNull           `json:"enableSubdomainAccess"`
	CdnEnabled                  types.BoolNull           `json:"cdnEnabled"`
	Location                    stringId.LocationIdNull  `json:"location"`
	ServiceStackVersionName     types.String             `json:"serviceStackVersionName"`
}

func (dto PostServiceStack) GetName() types.String {
	return dto.Name
}
func (dto PostServiceStack) GetCustomAutoscaling() *CustomAutoscaling {
	return dto.CustomAutoscaling
}
func (dto PostServiceStack) GetAutoscalingProfileId() types.StringNull {
	return dto.AutoscalingProfileId
}
func (dto PostServiceStack) GetAutoscalingProfileOverrides() types.JsonRawMessage {
	return dto.AutoscalingProfileOverrides
}
func (dto PostServiceStack) GetUserData() PostServiceStackUserData {
	return dto.UserData
}
func (dto PostServiceStack) GetUserDataEnvFile() types.TextNull {
	return dto.UserDataEnvFile
}
func (dto PostServiceStack) GetStartWithoutCode() types.BoolNull {
	return dto.StartWithoutCode
}
func (dto PostServiceStack) GetBuildFromGit() types.StringNull {
	return dto.BuildFromGit
}
func (dto PostServiceStack) GetZeropsSetup() types.StringNull {
	return dto.ZeropsSetup
}
func (dto PostServiceStack) GetZeropsYaml() types.TextNull {
	return dto.ZeropsYaml
}
func (dto PostServiceStack) GetEnvIsolation() types.StringNull {
	return dto.EnvIsolation
}
func (dto PostServiceStack) GetSshIsolation() types.StringNull {
	return dto.SshIsolation
}
func (dto PostServiceStack) GetEnableSubdomainAccess() types.BoolNull {
	return dto.EnableSubdomainAccess
}
func (dto PostServiceStack) GetCdnEnabled() types.BoolNull {
	return dto.CdnEnabled
}
func (dto PostServiceStack) GetLocation() stringId.LocationIdNull {
	return dto.Location
}
func (dto PostServiceStack) GetServiceStackVersionName() types.String {
	return dto.ServiceStackVersionName
}

type PostServiceStackUserData []UserDataPut

func (dto PostServiceStackUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserDataPut(dto))
}

func (dto *PostServiceStack) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Name                        *types.String
		CustomAutoscaling           *CustomAutoscaling
		AutoscalingProfileId        types.StringNull
		AutoscalingProfileOverrides types.JsonRawMessage
		UserData                    *PostServiceStackUserData
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
		ServiceStackVersionName     *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostServiceStack", err)
	}
	var errorList validator.ErrorList
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if aux.UserData == nil {
		errorList = errorList.With(validator.NewError("userData", "field is required"))
	}
	if aux.ServiceStackVersionName == nil {
		errorList = errorList.With(validator.NewError("serviceStackVersionName", "field is required"))
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
	dto.ServiceStackVersionName = *aux.ServiceStackVersionName

	return nil
}
