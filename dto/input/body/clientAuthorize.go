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
var _ json.Unmarshaler = (*ClientAuthorize)(nil)

type ClientAuthorize struct {
	ClientId uuid.ClientId               `json:"clientId"`
	Email    types.UrlEncodedBase64Email `json:"email"`
	Token    uuid.ClientAuthorizeToken   `json:"token"`
}

func (dto ClientAuthorize) GetClientId() uuid.ClientId {
	return dto.ClientId
}
func (dto ClientAuthorize) GetEmail() types.UrlEncodedBase64Email {
	return dto.Email
}
func (dto ClientAuthorize) GetToken() uuid.ClientAuthorizeToken {
	return dto.Token
}

func (dto *ClientAuthorize) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ClientId *uuid.ClientId
		Email    *types.UrlEncodedBase64Email
		Token    *uuid.ClientAuthorizeToken
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ClientAuthorize", err)
	}
	var errorList validator.ErrorList
	if aux.ClientId == nil {
		errorList = errorList.With(validator.NewError("clientId", "field is required"))
	}
	if aux.Email == nil {
		errorList = errorList.With(validator.NewError("email", "field is required"))
	}
	if aux.Token == nil {
		errorList = errorList.With(validator.NewError("token", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ClientId = *aux.ClientId
	dto.Email = *aux.Email
	dto.Token = *aux.Token

	return nil
}
