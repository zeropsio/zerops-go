// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ClientUserProjectRole)(nil)

type ClientUserProjectRole struct {
	ProjectId uuid.ProjectId              `json:"projectId"`
	RoleCode  enum.ClientUserRoleCodeEnum `json:"roleCode"`
}

func (dto ClientUserProjectRole) GetProjectId() uuid.ProjectId {
	return dto.ProjectId
}
func (dto ClientUserProjectRole) GetRoleCode() enum.ClientUserRoleCodeEnum {
	return dto.RoleCode
}

func (dto *ClientUserProjectRole) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ProjectId *uuid.ProjectId
		RoleCode  *enum.ClientUserRoleCodeEnum
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ClientUserProjectRole", err)
	}
	var errorList validator.ErrorList
	if aux.ProjectId == nil {
		errorList = errorList.With(validator.NewError("projectId", "field is required"))
	}
	if aux.RoleCode == nil {
		errorList = errorList.With(validator.NewError("roleCode", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ProjectId = *aux.ProjectId
	dto.RoleCode = *aux.RoleCode

	return nil
}
