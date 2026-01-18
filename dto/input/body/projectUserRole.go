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
var _ json.Unmarshaler = (*ProjectUserRole)(nil)

type ProjectUserRole struct {
	Id       uuid.ClientUserId           `json:"id"`
	RoleCode enum.ClientUserRoleCodeEnum `json:"roleCode"`
}

func (dto ProjectUserRole) GetId() uuid.ClientUserId {
	return dto.Id
}
func (dto ProjectUserRole) GetRoleCode() enum.ClientUserRoleCodeEnum {
	return dto.RoleCode
}

func (dto *ProjectUserRole) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Id       *uuid.ClientUserId
		RoleCode *enum.ClientUserRoleCodeEnum
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ProjectUserRole", err)
	}
	var errorList validator.ErrorList
	if aux.Id == nil {
		errorList = errorList.With(validator.NewError("id", "field is required"))
	}
	if aux.RoleCode == nil {
		errorList = errorList.With(validator.NewError("roleCode", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Id = *aux.Id
	dto.RoleCode = *aux.RoleCode

	return nil
}
