// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutStandardServiceStackTriggerExternalRepositoryIntegration)(nil)

type PutStandardServiceStackTriggerExternalRepositoryIntegration struct {
	UserData        PutStandardServiceStackTriggerExternalRepositoryIntegrationUserData `json:"userData"`
	UserDataEnvFile types.TextNull                                                      `json:"userDataEnvFile"` // Deprecated
	EnvVarsFile     types.TextNull                                                      `json:"envVarsFile"`
	EnvSecretsFile  types.TextNull                                                      `json:"envSecretsFile"`
}

func (dto PutStandardServiceStackTriggerExternalRepositoryIntegration) GetUserData() PutStandardServiceStackTriggerExternalRepositoryIntegrationUserData {
	return dto.UserData
}
func (dto PutStandardServiceStackTriggerExternalRepositoryIntegration) GetUserDataEnvFile() types.TextNull {
	return dto.UserDataEnvFile
}
func (dto PutStandardServiceStackTriggerExternalRepositoryIntegration) GetEnvVarsFile() types.TextNull {
	return dto.EnvVarsFile
}
func (dto PutStandardServiceStackTriggerExternalRepositoryIntegration) GetEnvSecretsFile() types.TextNull {
	return dto.EnvSecretsFile
}

type PutStandardServiceStackTriggerExternalRepositoryIntegrationUserData []UserDataPut

func (dto PutStandardServiceStackTriggerExternalRepositoryIntegrationUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserDataPut(dto))
}

func (dto *PutStandardServiceStackTriggerExternalRepositoryIntegration) UnmarshalJSON(b []byte) error {
	var aux = struct {
		UserData        *PutStandardServiceStackTriggerExternalRepositoryIntegrationUserData
		UserDataEnvFile types.TextNull
		EnvVarsFile     types.TextNull
		EnvSecretsFile  types.TextNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutStandardServiceStackTriggerExternalRepositoryIntegration", err)
	}
	var errorList validator.ErrorList
	if aux.UserData == nil {
		errorList = errorList.With(validator.NewError("userData", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.UserData = *aux.UserData
	dto.UserDataEnvFile = aux.UserDataEnvFile
	dto.EnvVarsFile = aux.EnvVarsFile
	dto.EnvSecretsFile = aux.EnvSecretsFile

	return nil
}
