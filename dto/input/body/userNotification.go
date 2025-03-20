// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*UserNotification)(nil)

type UserNotification struct {
	ClientId       uuid.ClientId           `json:"clientId"`
	ProjectId      uuid.ProjectIdNull      `json:"projectId"`
	ServiceStackId uuid.ServiceStackIdNull `json:"serviceStackId"`
}

func (dto UserNotification) GetClientId() uuid.ClientId {
	return dto.ClientId
}
func (dto UserNotification) GetProjectId() uuid.ProjectIdNull {
	return dto.ProjectId
}
func (dto UserNotification) GetServiceStackId() uuid.ServiceStackIdNull {
	return dto.ServiceStackId
}

func (dto *UserNotification) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ClientId       *uuid.ClientId
		ProjectId      uuid.ProjectIdNull
		ServiceStackId uuid.ServiceStackIdNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("UserNotification", err)
	}
	var errorList validator.ErrorList
	if aux.ClientId == nil {
		errorList = errorList.With(validator.NewError("clientId", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ClientId = *aux.ClientId
	dto.ProjectId = aux.ProjectId
	dto.ServiceStackId = aux.ServiceStackId

	return nil
}
