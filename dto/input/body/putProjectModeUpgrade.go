// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutProjectModeUpgrade)(nil)

type PutProjectModeUpgrade struct {
	Mode enum.ProjectModeEnum `json:"mode"`
}

func (dto PutProjectModeUpgrade) GetMode() enum.ProjectModeEnum {
	return dto.Mode
}

func (dto *PutProjectModeUpgrade) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Mode *enum.ProjectModeEnum
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutProjectModeUpgrade", err)
	}
	var errorList validator.ErrorList
	if aux.Mode == nil {
		errorList = errorList.With(validator.NewError("mode", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Mode = *aux.Mode

	return nil
}
