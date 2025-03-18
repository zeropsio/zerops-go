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
var _ json.Unmarshaler = (*ProjectEnvPutEnvFile)(nil)

type ProjectEnvPutEnvFile struct {
	ProjectId uuid.ProjectId `json:"projectId"`
	EnvFile   types.Text     `json:"envFile"`
}

func (dto ProjectEnvPutEnvFile) GetProjectId() uuid.ProjectId {
	return dto.ProjectId
}
func (dto ProjectEnvPutEnvFile) GetEnvFile() types.Text {
	return dto.EnvFile
}

func (dto *ProjectEnvPutEnvFile) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ProjectId *uuid.ProjectId
		EnvFile   *types.Text
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ProjectEnvPutEnvFile", err)
	}
	var errorList validator.ErrorList
	if aux.ProjectId == nil {
		errorList = errorList.With(validator.NewError("projectId", "field is required"))
	}
	if aux.EnvFile == nil {
		errorList = errorList.With(validator.NewError("envFile", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ProjectId = *aux.ProjectId
	dto.EnvFile = *aux.EnvFile

	return nil
}
