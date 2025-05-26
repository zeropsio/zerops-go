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
var _ json.Unmarshaler = (*PostProject)(nil)

type PostProject struct {
	ClientId       uuid.ClientId           `json:"clientId"`
	Name           types.String            `json:"name"`
	Description    types.TextNull          `json:"description"`
	Mode           *enum.ProjectModeEnum   `json:"mode"`
	TagList        types.StringArray       `json:"tagList"`
	EnvVariables   PostProjectEnvVariables `json:"envVariables"`
	EnvIsolation   types.StringNull        `json:"envIsolation"`
	SshIsolation   types.StringNull        `json:"sshIsolation"`
	MaxCreditLimit types.DecimalNull       `json:"maxCreditLimit"`
}

func (dto PostProject) GetClientId() uuid.ClientId {
	return dto.ClientId
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
func (dto PostProject) GetEnvVariables() PostProjectEnvVariables {
	return dto.EnvVariables
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

type PostProjectEnvVariables []ProjectEnvPut

func (dto PostProjectEnvVariables) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProjectEnvPut(dto))
}

func (dto *PostProject) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ClientId       *uuid.ClientId
		Name           *types.String
		Description    types.TextNull
		Mode           *enum.ProjectModeEnum
		TagList        *types.StringArray
		EnvVariables   *PostProjectEnvVariables
		EnvIsolation   types.StringNull
		SshIsolation   types.StringNull
		MaxCreditLimit types.DecimalNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostProject", err)
	}
	var errorList validator.ErrorList
	if aux.ClientId == nil {
		errorList = errorList.With(validator.NewError("clientId", "field is required"))
	}
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if aux.TagList == nil {
		errorList = errorList.With(validator.NewError("tagList", "field is required"))
	}
	if aux.EnvVariables == nil {
		errorList = errorList.With(validator.NewError("envVariables", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ClientId = *aux.ClientId
	dto.Name = *aux.Name
	dto.Description = aux.Description
	dto.Mode = aux.Mode
	dto.TagList = *aux.TagList
	dto.EnvVariables = *aux.EnvVariables
	dto.EnvIsolation = aux.EnvIsolation
	dto.SshIsolation = aux.SshIsolation
	dto.MaxCreditLimit = aux.MaxCreditLimit

	return nil
}
