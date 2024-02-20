// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsClientUser struct {
	Id       uuid.ClientUserId           `json:"id"`
	ClientId uuid.ClientId               `json:"clientId"`
	UserId   uuid.UserId                 `json:"userId"`
	Status   enum.ClientUserStatusEnum   `json:"status"`
	RoleCode enum.ClientUserRoleCodeEnum `json:"roleCode"`
	Client   ClientLight                 `json:"client"`
	User     UserLight                   `json:"user"`
}
