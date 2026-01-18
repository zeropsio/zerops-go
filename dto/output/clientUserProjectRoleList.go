// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type ClientUserProjectRoleList struct {
	ProjectRoleList ClientUserProjectRoleListProjectRoleList `json:"projectRoleList"`
}

type ClientUserProjectRoleListProjectRoleList []ClientUserProjectRole

func (dto ClientUserProjectRoleListProjectRoleList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ClientUserProjectRole(dto))
}
