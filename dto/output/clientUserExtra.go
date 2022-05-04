// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ClientUserExtra struct {
	Id       uuid.ClientUserId                `json:"id"`
	ClientId uuid.ClientId                    `json:"clientId"`
	UserId   uuid.UserId                      `json:"userId"`
	Status   enum.ClientUserLightStatusEnum   `json:"status"`
	RoleCode enum.ClientUserLightRoleCodeEnum `json:"roleCode"`
	Client   ClientExtra                      `json:"client"`
	User     UserLight                        `json:"user"`
}
