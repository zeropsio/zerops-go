// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ClientUserValidate)(nil)

type ClientUserValidate struct {
	ClientId uuid.ClientUserValidateClientId `json:"clientId"`
	Email    types.Email                     `json:"email"`
}

func (dto ClientUserValidate) GetClientId() uuid.ClientUserValidateClientId {
	return dto.ClientId
}
func (dto ClientUserValidate) GetEmail() types.Email {
	return dto.Email
}

func (dto *ClientUserValidate) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ClientId *uuid.ClientUserValidateClientId
		Email    *types.Email
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ClientUserValidate", err)
	}
	var errorList validator.ErrorList
	if aux.ClientId == nil {
		errorList = errorList.With(validator.NewError("clientId", "field is required"))
	}
	if aux.Email == nil {
		errorList = errorList.With(validator.NewError("email", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ClientId = *aux.ClientId
	dto.Email = *aux.Email

	return nil
}
