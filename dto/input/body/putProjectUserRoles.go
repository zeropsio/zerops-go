// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutProjectUserRoles)(nil)

type PutProjectUserRoles struct {
	UserRoles PutProjectUserRolesUserRoles `json:"userRoles"`
}

func (dto PutProjectUserRoles) GetUserRoles() PutProjectUserRolesUserRoles {
	return dto.UserRoles
}

type PutProjectUserRolesUserRoles []ProjectUserRole

func (dto PutProjectUserRolesUserRoles) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProjectUserRole(dto))
}

func (dto *PutProjectUserRoles) UnmarshalJSON(b []byte) error {
	var aux = struct {
		UserRoles *PutProjectUserRolesUserRoles
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutProjectUserRoles", err)
	}
	var errorList validator.ErrorList
	if aux.UserRoles == nil {
		errorList = errorList.With(validator.NewError("userRoles", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.UserRoles = *aux.UserRoles

	return nil
}
