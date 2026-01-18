// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ClientUserValidate)(nil)

type ClientUserValidate struct {
	Email types.Email `json:"email"`
}

func (dto ClientUserValidate) GetEmail() types.Email {
	return dto.Email
}

func (dto *ClientUserValidate) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Email *types.Email
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ClientUserValidate", err)
	}
	var errorList validator.ErrorList
	if aux.Email == nil {
		errorList = errorList.With(validator.NewError("email", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Email = *aux.Email

	return nil
}
