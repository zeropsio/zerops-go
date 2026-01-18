// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*AuthTOTP)(nil)

type AuthTOTP struct {
	Token types.String `json:"token"`
}

func (dto AuthTOTP) GetToken() types.String {
	return dto.Token
}

func (dto *AuthTOTP) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Token *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("AuthTOTP", err)
	}
	var errorList validator.ErrorList
	if aux.Token == nil {
		errorList = errorList.With(validator.NewError("token", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Token = *aux.Token

	return nil
}
