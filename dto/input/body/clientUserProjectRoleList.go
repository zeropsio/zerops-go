// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ClientUserProjectRoleList)(nil)

type ClientUserProjectRoleList struct {
	ProjectRoleList ClientUserProjectRoleListProjectRoleList `json:"projectRoleList"`
}

func (dto ClientUserProjectRoleList) GetProjectRoleList() ClientUserProjectRoleListProjectRoleList {
	return dto.ProjectRoleList
}

type ClientUserProjectRoleListProjectRoleList []ClientUserProjectRole

func (dto ClientUserProjectRoleListProjectRoleList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ClientUserProjectRole(dto))
}

func (dto *ClientUserProjectRoleList) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ProjectRoleList *ClientUserProjectRoleListProjectRoleList
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ClientUserProjectRoleList", err)
	}
	var errorList validator.ErrorList
	if aux.ProjectRoleList == nil {
		errorList = errorList.With(validator.NewError("projectRoleList", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ProjectRoleList = *aux.ProjectRoleList

	return nil
}
