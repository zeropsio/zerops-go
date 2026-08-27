// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type IntegrationTokenDelegationPermission struct {
	RoleCode           enum.ClientUserRoleCodeEnum                            `json:"roleCode"`
	CanCreateProjects  types.Bool                                             `json:"canCreateProjects"`
	CanEditFinances    types.Bool                                             `json:"canEditFinances"`
	CanViewFinances    types.Bool                                             `json:"canViewFinances"`
	ProjectPermissions IntegrationTokenDelegationPermissionProjectPermissions `json:"projectPermissions"`
}

type IntegrationTokenDelegationPermissionProjectPermissions []IntegrationTokenDelegationProjectPermission

func (dto IntegrationTokenDelegationPermissionProjectPermissions) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]IntegrationTokenDelegationProjectPermission(dto))
}
