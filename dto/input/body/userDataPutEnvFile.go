// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*UserDataPutEnvFile)(nil)

type UserDataPutEnvFile struct {
	ServiceStackId uuid.ServiceStackId `json:"serviceStackId"`
	EnvFile        types.Text          `json:"envFile"`
}

func (dto UserDataPutEnvFile) GetServiceStackId() uuid.ServiceStackId {
	return dto.ServiceStackId
}
func (dto UserDataPutEnvFile) GetEnvFile() types.Text {
	return dto.EnvFile
}

func (dto *UserDataPutEnvFile) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ServiceStackId *uuid.ServiceStackId
		EnvFile        *types.Text
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("UserDataPutEnvFile", err)
	}
	var errorList validator.ErrorList
	if aux.ServiceStackId == nil {
		errorList = errorList.With(validator.NewError("serviceStackId", "field is required"))
	}
	if aux.EnvFile == nil {
		errorList = errorList.With(validator.NewError("envFile", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ServiceStackId = *aux.ServiceStackId
	dto.EnvFile = *aux.EnvFile

	return nil
}
