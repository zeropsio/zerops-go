// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ClientUserExtraWithClientLight struct {
	Id                uuid.ClientUserId           `json:"id"`
	ClientId          uuid.ClientId               `json:"clientId"`
	UserId            uuid.UserId                 `json:"userId"`
	Status            enum.ClientUserStatusEnum   `json:"status"`
	RoleCode          enum.ClientUserRoleCodeEnum `json:"roleCode"`
	CanViewFinances   types.Bool                  `json:"canViewFinances"`
	CanEditFinances   types.Bool                  `json:"canEditFinances"`
	CanCreateProjects types.Bool                  `json:"canCreateProjects"`
	User              UserLight                   `json:"user"`
	Client            ClientLight                 `json:"client"`
}
