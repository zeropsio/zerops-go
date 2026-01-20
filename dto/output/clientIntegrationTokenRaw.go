// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ClientIntegrationTokenRaw struct {
	Id                uuid.UserId                       `json:"id"`
	Name              types.String                      `json:"name"`
	RoleCode          enum.ClientUserRoleCodeEnum       `json:"roleCode"`
	CanViewFinances   types.Bool                        `json:"canViewFinances"`
	CanEditFinances   types.Bool                        `json:"canEditFinances"`
	CanCreateProjects types.Bool                        `json:"canCreateProjects"`
	CreatedByUser     uuid.UserIdNull                   `json:"createdByUser"`
	Created           types.DateTime                    `json:"created"`
	LastUpdate        types.DateTime                    `json:"lastUpdate"`
	Projects          ClientIntegrationTokenRawProjects `json:"projects"`
	Token             types.String                      `json:"token"`
}

type ClientIntegrationTokenRawProjects []ClientIntegrationTokenProjectAccess

func (dto ClientIntegrationTokenRawProjects) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ClientIntegrationTokenProjectAccess(dto))
}
