// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*UserNotification)(nil)

type UserNotification struct {
	ProjectId      types.StringNull `json:"projectId"`
	ServiceStackId types.StringNull `json:"serviceStackId"`
}

func (dto UserNotification) GetProjectId() types.StringNull {
	return dto.ProjectId
}
func (dto UserNotification) GetServiceStackId() types.StringNull {
	return dto.ServiceStackId
}

func (dto *UserNotification) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ProjectId      types.StringNull
		ServiceStackId types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("UserNotification", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.ProjectId = aux.ProjectId
	dto.ServiceStackId = aux.ServiceStackId

	return nil
}
