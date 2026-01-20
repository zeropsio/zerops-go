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
var _ json.Unmarshaler = (*ClientIntegrationToken)(nil)

type ClientIntegrationToken struct {
	Name              types.String                   `json:"name"`
	RoleCode          enum.ClientUserRoleCodeEnum    `json:"roleCode"`
	CanViewFinances   types.Bool                     `json:"canViewFinances"`
	CanEditFinances   types.Bool                     `json:"canEditFinances"`
	CanCreateProjects types.Bool                     `json:"canCreateProjects"`
	Projects          ClientIntegrationTokenProjects `json:"projects"`
}

func (dto ClientIntegrationToken) GetName() types.String {
	return dto.Name
}
func (dto ClientIntegrationToken) GetRoleCode() enum.ClientUserRoleCodeEnum {
	return dto.RoleCode
}
func (dto ClientIntegrationToken) GetCanViewFinances() types.Bool {
	return dto.CanViewFinances
}
func (dto ClientIntegrationToken) GetCanEditFinances() types.Bool {
	return dto.CanEditFinances
}
func (dto ClientIntegrationToken) GetCanCreateProjects() types.Bool {
	return dto.CanCreateProjects
}
func (dto ClientIntegrationToken) GetProjects() ClientIntegrationTokenProjects {
	return dto.Projects
}

type ClientIntegrationTokenProjects []ClientIntegrationTokenProjectAccess

func (dto ClientIntegrationTokenProjects) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ClientIntegrationTokenProjectAccess(dto))
}

func (dto *ClientIntegrationToken) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Name              *types.String
		RoleCode          *enum.ClientUserRoleCodeEnum
		CanViewFinances   *types.Bool
		CanEditFinances   *types.Bool
		CanCreateProjects *types.Bool
		Projects          *ClientIntegrationTokenProjects
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ClientIntegrationToken", err)
	}
	var errorList validator.ErrorList
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
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
	if aux.Projects == nil {
		errorList = errorList.With(validator.NewError("projects", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Name = *aux.Name
	dto.RoleCode = *aux.RoleCode
	dto.CanViewFinances = *aux.CanViewFinances
	dto.CanEditFinances = *aux.CanEditFinances
	dto.CanCreateProjects = *aux.CanCreateProjects
	dto.Projects = *aux.Projects

	return nil
}
