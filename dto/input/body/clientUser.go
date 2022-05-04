// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ClientUser)(nil)

type ClientUser struct {
	ClientId uuid.ClientId                    `json:"clientId"`
	RoleCode enum.ClientUserLightRoleCodeEnum `json:"roleCode"`
}

func (dto ClientUser) GetClientId() uuid.ClientId {
	return dto.ClientId
}
func (dto ClientUser) GetRoleCode() enum.ClientUserLightRoleCodeEnum {
	return dto.RoleCode
}

func (dto *ClientUser) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ClientId *uuid.ClientId
		RoleCode *enum.ClientUserLightRoleCodeEnum
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ClientUser", err)
	}
	var errorList validator.ErrorList
	if aux.ClientId == nil {
		errorList = errorList.With(validator.NewError("clientId", "field is required"))
	}
	if aux.RoleCode == nil {
		errorList = errorList.With(validator.NewError("roleCode", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ClientId = *aux.ClientId
	dto.RoleCode = *aux.RoleCode

	return nil
}
