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
var _ json.Unmarshaler = (*ServiceStackImport)(nil)

type ServiceStackImport struct {
	ProjectId uuid.ProjectId `json:"projectId"`
	Yaml      types.Text     `json:"yaml"`
}

func (dto ServiceStackImport) GetProjectId() uuid.ProjectId {
	return dto.ProjectId
}
func (dto ServiceStackImport) GetYaml() types.Text {
	return dto.Yaml
}

func (dto *ServiceStackImport) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ProjectId *uuid.ProjectId
		Yaml      *types.Text
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ServiceStackImport", err)
	}
	var errorList validator.ErrorList
	if aux.ProjectId == nil {
		errorList = errorList.With(validator.NewError("projectId", "field is required"))
	}
	if aux.Yaml == nil {
		errorList = errorList.With(validator.NewError("yaml", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ProjectId = *aux.ProjectId
	dto.Yaml = *aux.Yaml

	return nil
}
