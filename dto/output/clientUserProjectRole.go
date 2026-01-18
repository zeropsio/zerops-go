// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ClientUserProjectRole struct {
	Id        uuid.ClientUserProjectId    `json:"id"`
	ProjectId uuid.ProjectId              `json:"projectId"`
	RoleCode  enum.ClientUserRoleCodeEnum `json:"roleCode"`
}
