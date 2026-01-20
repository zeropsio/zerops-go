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
var _ json.Unmarshaler = (*ClientUserConnection)(nil)

type ClientUserConnection struct {
	RoleCode          enum.ClientUserRoleCodeEnum `json:"roleCode"`
	CanViewFinances   types.Bool                  `json:"canViewFinances"`
	CanEditFinances   types.Bool                  `json:"canEditFinances"`
	CanCreateProjects types.Bool                  `json:"canCreateProjects"`
}

func (dto ClientUserConnection) GetRoleCode() enum.ClientUserRoleCodeEnum {
	return dto.RoleCode
}
func (dto ClientUserConnection) GetCanViewFinances() types.Bool {
	return dto.CanViewFinances
}
func (dto ClientUserConnection) GetCanEditFinances() types.Bool {
	return dto.CanEditFinances
}
func (dto ClientUserConnection) GetCanCreateProjects() types.Bool {
	return dto.CanCreateProjects
}

func (dto *ClientUserConnection) UnmarshalJSON(b []byte) error {
	var aux = struct {
		RoleCode          *enum.ClientUserRoleCodeEnum
		CanViewFinances   *types.Bool
		CanEditFinances   *types.Bool
		CanCreateProjects *types.Bool
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ClientUserConnection", err)
	}
	var errorList validator.ErrorList
	if aux.RoleCode == nil {
		errorList = errorList.With(validator.NewError("roleCode", "field is required"))
	}
	if aux.CanViewFinances == nil {
		errorList = errorList.With(validator.NewError("canViewFinances", "field is required"))
	}
	if aux.CanEditFinances == nil {
		errorList = errorList.With(validator.NewError("canEditFinances", "field is required"))
	}
	if aux.CanCreateProjects == nil {
		errorList = errorList.With(validator.NewError("canCreateProjects", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.RoleCode = *aux.RoleCode
	dto.CanViewFinances = *aux.CanViewFinances
	dto.CanEditFinances = *aux.CanEditFinances
	dto.CanCreateProjects = *aux.CanCreateProjects

	return nil
}
