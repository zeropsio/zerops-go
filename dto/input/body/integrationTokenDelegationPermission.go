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
var _ json.Unmarshaler = (*IntegrationTokenDelegationPermission)(nil)

type IntegrationTokenDelegationPermission struct {
	RoleCode           enum.ClientUserRoleCodeEnum                            `json:"roleCode"`
	CanCreateProjects  types.Bool                                             `json:"canCreateProjects"`
	CanEditFinances    types.Bool                                             `json:"canEditFinances"`
	CanViewFinances    types.Bool                                             `json:"canViewFinances"`
	ProjectPermissions IntegrationTokenDelegationPermissionProjectPermissions `json:"projectPermissions"`
}

func (dto IntegrationTokenDelegationPermission) GetRoleCode() enum.ClientUserRoleCodeEnum {
	return dto.RoleCode
}
func (dto IntegrationTokenDelegationPermission) GetCanCreateProjects() types.Bool {
	return dto.CanCreateProjects
}
func (dto IntegrationTokenDelegationPermission) GetCanEditFinances() types.Bool {
	return dto.CanEditFinances
}
func (dto IntegrationTokenDelegationPermission) GetCanViewFinances() types.Bool {
	return dto.CanViewFinances
}
func (dto IntegrationTokenDelegationPermission) GetProjectPermissions() IntegrationTokenDelegationPermissionProjectPermissions {
	return dto.ProjectPermissions
}

type IntegrationTokenDelegationPermissionProjectPermissions []IntegrationTokenDelegationProjectPermission

func (dto IntegrationTokenDelegationPermissionProjectPermissions) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]IntegrationTokenDelegationProjectPermission(dto))
}

func (dto *IntegrationTokenDelegationPermission) UnmarshalJSON(b []byte) error {
	var aux = struct {
		RoleCode           *enum.ClientUserRoleCodeEnum
		CanCreateProjects  *types.Bool
		CanEditFinances    *types.Bool
		CanViewFinances    *types.Bool
		ProjectPermissions *IntegrationTokenDelegationPermissionProjectPermissions
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("IntegrationTokenDelegationPermission", err)
	}
	var errorList validator.ErrorList
	if aux.RoleCode == nil {
		errorList = errorList.With(validator.NewError("roleCode", "field is required"))
	}
	if aux.CanCreateProjects == nil {
		errorList = errorList.With(validator.NewError("canCreateProjects", "field is required"))
	}
	if aux.CanEditFinances == nil {
		errorList = errorList.With(validator.NewError("canEditFinances", "field is required"))
	}
	if aux.CanViewFinances == nil {
		errorList = errorList.With(validator.NewError("canViewFinances", "field is required"))
	}
	if aux.ProjectPermissions == nil {
		errorList = errorList.With(validator.NewError("projectPermissions", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.RoleCode = *aux.RoleCode
	dto.CanCreateProjects = *aux.CanCreateProjects
	dto.CanEditFinances = *aux.CanEditFinances
	dto.CanViewFinances = *aux.CanViewFinances
	dto.ProjectPermissions = *aux.ProjectPermissions

	return nil
}
