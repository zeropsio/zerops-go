// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ProjectUserRole struct {
	Id           uuid.ClientUserProjectId    `json:"id"`
	ClientUserId uuid.ClientUserId           `json:"clientUserId"`
	RoleCode     enum.ClientUserRoleCodeEnum `json:"roleCode"`
}
