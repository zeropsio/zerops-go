// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ClientUserConnection)(nil)

type ClientUserConnection struct {
	RoleCode enum.ClientUserRoleCodeEnum `json:"roleCode"`
}

func (dto ClientUserConnection) GetRoleCode() enum.ClientUserRoleCodeEnum {
	return dto.RoleCode
}

func (dto *ClientUserConnection) UnmarshalJSON(b []byte) error {
	var aux = struct {
		RoleCode *enum.ClientUserRoleCodeEnum
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ClientUserConnection", err)
	}
	var errorList validator.ErrorList
	if aux.RoleCode == nil {
		errorList = errorList.With(validator.NewError("roleCode", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.RoleCode = *aux.RoleCode

	return nil
}
