// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostProject)(nil)

type PostProject struct {
	Name             types.String            `json:"name"`
	Description      types.TextNull          `json:"description"`
	Mode             *enum.ProjectModeEnum   `json:"mode"`
	TagList          types.StringArray       `json:"tagList"`
	UserRoles        PostProjectUserRoles    `json:"userRoles"`
	EnvVariables     PostProjectEnvVariables `json:"envVariables"`
	EnvSecrets       PostProjectEnvSecrets   `json:"envSecrets"`
	PublicIpV4Shared types.Bool              `json:"publicIpV4Shared"`
	EnvIsolation     types.StringNull        `json:"envIsolation"`
	SshIsolation     types.StringNull        `json:"sshIsolation"`
	MaxCreditLimit   types.DecimalNull       `json:"maxCreditLimit"`
	Location         stringId.LocationIdNull `json:"location"`
}

func (dto PostProject) GetName() types.String {
	return dto.Name
}
func (dto PostProject) GetDescription() types.TextNull {
	return dto.Description
}
func (dto PostProject) GetMode() *enum.ProjectModeEnum {
	return dto.Mode
}
func (dto PostProject) GetTagList() types.StringArray {
	return dto.TagList
}
func (dto PostProject) GetUserRoles() PostProjectUserRoles {
	return dto.UserRoles
}
func (dto PostProject) GetEnvVariables() PostProjectEnvVariables {
	return dto.EnvVariables
}
func (dto PostProject) GetEnvSecrets() PostProjectEnvSecrets {
	return dto.EnvSecrets
}
func (dto PostProject) GetPublicIpV4Shared() types.Bool {
	return dto.PublicIpV4Shared
}
func (dto PostProject) GetEnvIsolation() types.StringNull {
	return dto.EnvIsolation
}
func (dto PostProject) GetSshIsolation() types.StringNull {
	return dto.SshIsolation
}
func (dto PostProject) GetMaxCreditLimit() types.DecimalNull {
	return dto.MaxCreditLimit
}
func (dto PostProject) GetLocation() stringId.LocationIdNull {
	return dto.Location
}

type PostProjectUserRoles []ProjectUserRole

func (dto PostProjectUserRoles) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProjectUserRole(dto))
}

type PostProjectEnvVariables []ProjectEnvItem

func (dto PostProjectEnvVariables) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProjectEnvItem(dto))
}

type PostProjectEnvSecrets []ProjectEnvItem

func (dto PostProjectEnvSecrets) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProjectEnvItem(dto))
}

func (dto *PostProject) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Name             *types.String
		Description      types.TextNull
		Mode             *enum.ProjectModeEnum
		TagList          *types.StringArray
		UserRoles        *PostProjectUserRoles
		EnvVariables     *PostProjectEnvVariables
		EnvSecrets       *PostProjectEnvSecrets
		PublicIpV4Shared *types.Bool
		EnvIsolation     types.StringNull
		SshIsolation     types.StringNull
		MaxCreditLimit   types.DecimalNull
		Location         stringId.LocationIdNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostProject", err)
	}
	var errorList validator.ErrorList
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if aux.TagList == nil {
		errorList = errorList.With(validator.NewError("tagList", "field is required"))
	}
	if aux.UserRoles == nil {
		errorList = errorList.With(validator.NewError("userRoles", "field is required"))
	}
	if aux.EnvVariables == nil {
		errorList = errorList.With(validator.NewError("envVariables", "field is required"))
	}
	if aux.EnvSecrets == nil {
		errorList = errorList.With(validator.NewError("envSecrets", "field is required"))
	}
	if aux.PublicIpV4Shared == nil {
		errorList = errorList.With(validator.NewError("publicIpV4Shared", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Name = *aux.Name
	dto.Description = aux.Description
	dto.Mode = aux.Mode
	dto.TagList = *aux.TagList
	dto.UserRoles = *aux.UserRoles
	dto.EnvVariables = *aux.EnvVariables
	dto.EnvSecrets = *aux.EnvSecrets
	dto.PublicIpV4Shared = *aux.PublicIpV4Shared
	dto.EnvIsolation = aux.EnvIsolation
	dto.SshIsolation = aux.SshIsolation
	dto.MaxCreditLimit = aux.MaxCreditLimit
	dto.Location = aux.Location

	return nil
}
