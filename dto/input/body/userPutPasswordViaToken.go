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
var _ json.Unmarshaler = (*UserPutPasswordViaToken)(nil)

type UserPutPasswordViaToken struct {
	Email    types.Email                       `json:"email"`
	Password types.String                      `json:"password"`
	Token    uuid.UserPutPasswordViaTokenToken `json:"token"`
}

func (dto UserPutPasswordViaToken) GetEmail() types.Email {
	return dto.Email
}
func (dto UserPutPasswordViaToken) GetPassword() types.String {
	return dto.Password
}
func (dto UserPutPasswordViaToken) GetToken() uuid.UserPutPasswordViaTokenToken {
	return dto.Token
}

func (dto *UserPutPasswordViaToken) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Email    *types.Email
		Password *types.String
		Token    *uuid.UserPutPasswordViaTokenToken
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("UserPutPasswordViaToken", err)
	}
	var errorList validator.ErrorList
	if aux.Email == nil {
		errorList = errorList.With(validator.NewError("email", "field is required"))
	}
	if aux.Password == nil {
		errorList = errorList.With(validator.NewError("password", "field is required"))
	}
	if aux.Token == nil {
		errorList = errorList.With(validator.NewError("token", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Email = *aux.Email
	dto.Password = *aux.Password
	dto.Token = *aux.Token

	return nil
}
