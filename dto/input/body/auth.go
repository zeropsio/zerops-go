// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*Auth)(nil)

type Auth struct {
	Email    types.Email  `json:"email"`
	Password types.String `json:"password"`
}

func (dto Auth) GetEmail() types.Email {
	return dto.Email
}
func (dto Auth) GetPassword() types.String {
	return dto.Password
}

func (dto *Auth) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Email    *types.Email
		Password *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("Auth", err)
	}
	var errorList validator.ErrorList
	if aux.Email == nil {
		errorList = errorList.With(validator.NewError("email", "field is required"))
	}
	if aux.Password == nil {
		errorList = errorList.With(validator.NewError("password", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Email = *aux.Email
	dto.Password = *aux.Password

	return nil
}
