// Generated ZEROPS sdk

package body

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonInputDto.go.tmpl

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostUserNginxServiceStack)(nil)

type PostUserNginxServiceStack struct {
	ProjectId         uuid.ProjectId                    `json:"projectId"`
	Name              types.String                      `json:"name"`
	Mode              enum.ServiceStackModeEnum         `json:"mode"`
	CustomAutoscaling *CustomAutoscaling                `json:"customAutoscaling"`
	UserData          PostUserNginxServiceStackUserData `json:"userData"`
	GithubIntegration *GithubIntegration                `json:"githubIntegration"`
	GitlabIntegration *GitlabIntegration                `json:"gitlabIntegration"`
	NginxConfig       types.Text                        `json:"nginxConfig"`
}

func (dto PostUserNginxServiceStack) GetProjectId() uuid.ProjectId {
	return dto.ProjectId
}
func (dto PostUserNginxServiceStack) GetName() types.String {
	return dto.Name
}
func (dto PostUserNginxServiceStack) GetMode() enum.ServiceStackModeEnum {
	return dto.Mode
}
func (dto PostUserNginxServiceStack) GetCustomAutoscaling() *CustomAutoscaling {
	return dto.CustomAutoscaling
}
func (dto PostUserNginxServiceStack) GetUserData() PostUserNginxServiceStackUserData {
	return dto.UserData
}
func (dto PostUserNginxServiceStack) GetGithubIntegration() *GithubIntegration {
	return dto.GithubIntegration
}
func (dto PostUserNginxServiceStack) GetGitlabIntegration() *GitlabIntegration {
	return dto.GitlabIntegration
}
func (dto PostUserNginxServiceStack) GetNginxConfig() types.Text {
	return dto.NginxConfig
}

type PostUserNginxServiceStackUserData []UserDataPut

func (dto PostUserNginxServiceStackUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserDataPut(dto))
}

func (dto *PostUserNginxServiceStack) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ProjectId         *uuid.ProjectId
		Name              *types.String
		Mode              *enum.ServiceStackModeEnum
		CustomAutoscaling *CustomAutoscaling
		UserData          *PostUserNginxServiceStackUserData
		GithubIntegration *GithubIntegration
		GitlabIntegration *GitlabIntegration
		NginxConfig       *types.Text
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostUserNginxServiceStack", err)
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
	if aux.NginxConfig == nil {
		errorList = errorList.With(validator.NewError("nginxConfig", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ProjectId = *aux.ProjectId
	dto.Name = *aux.Name
	dto.Mode = *aux.Mode
	dto.CustomAutoscaling = aux.CustomAutoscaling
	dto.UserData = *aux.UserData
	dto.GithubIntegration = aux.GithubIntegration
	dto.GitlabIntegration = aux.GitlabIntegration
	dto.NginxConfig = *aux.NginxConfig

	return nil
}
