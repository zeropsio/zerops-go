// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*GithubIntegration)(nil)

type GithubIntegration struct {
	RepositoryFullName types.String                        `json:"repositoryFullName"`
	EventType          enum.GithubIntegrationEventTypeEnum `json:"eventType"`
	BranchName         types.StringNull                    `json:"branchName"`
	TagRegex           types.StringNull                    `json:"tagRegex"`
	IsActive           types.Bool                          `json:"isActive"`
	ZeropsYamlSetup    types.StringNull                    `json:"zeropsYamlSetup"`
	TriggerBuild       types.Bool                          `json:"triggerBuild"`
}

func (dto GithubIntegration) GetRepositoryFullName() types.String {
	return dto.RepositoryFullName
}
func (dto GithubIntegration) GetEventType() enum.GithubIntegrationEventTypeEnum {
	return dto.EventType
}
func (dto GithubIntegration) GetBranchName() types.StringNull {
	return dto.BranchName
}
func (dto GithubIntegration) GetTagRegex() types.StringNull {
	return dto.TagRegex
}
func (dto GithubIntegration) GetIsActive() types.Bool {
	return dto.IsActive
}
func (dto GithubIntegration) GetZeropsYamlSetup() types.StringNull {
	return dto.ZeropsYamlSetup
}
func (dto GithubIntegration) GetTriggerBuild() types.Bool {
	return dto.TriggerBuild
}

func (dto *GithubIntegration) UnmarshalJSON(b []byte) error {
	var aux = struct {
		RepositoryFullName *types.String
		EventType          *enum.GithubIntegrationEventTypeEnum
		BranchName         types.StringNull
		TagRegex           types.StringNull
		IsActive           *types.Bool
		ZeropsYamlSetup    types.StringNull
		TriggerBuild       *types.Bool
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("GithubIntegration", err)
	}
	var errorList validator.ErrorList
	if aux.RepositoryFullName == nil {
		errorList = errorList.With(validator.NewError("repositoryFullName", "field is required"))
	}
	if aux.EventType == nil {
		errorList = errorList.With(validator.NewError("eventType", "field is required"))
	}
	if aux.IsActive == nil {
		errorList = errorList.With(validator.NewError("isActive", "field is required"))
	}
	if aux.TriggerBuild == nil {
		errorList = errorList.With(validator.NewError("triggerBuild", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.RepositoryFullName = *aux.RepositoryFullName
	dto.EventType = *aux.EventType
	dto.BranchName = aux.BranchName
	dto.TagRegex = aux.TagRegex
	dto.IsActive = *aux.IsActive
	dto.ZeropsYamlSetup = aux.ZeropsYamlSetup
	dto.TriggerBuild = *aux.TriggerBuild

	return nil
}
