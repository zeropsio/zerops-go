// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsClientUser struct {
	Id       uuid.ClientUserId                `json:"id"`
	ClientId uuid.ClientId                    `json:"clientId"`
	UserId   uuid.UserId                      `json:"userId"`
	Status   enum.ClientUserLightStatusEnum   `json:"status"`
	RoleCode enum.ClientUserLightRoleCodeEnum `json:"roleCode"`
	Client   ClientLight                      `json:"client"`
	User     UserLight                        `json:"user"`
}
