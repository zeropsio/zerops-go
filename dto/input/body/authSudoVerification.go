// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*AuthSudoVerification)(nil)

type AuthSudoVerification struct {
	WebAuthn *WebAuthnLoginVerification `json:"webAuthn"`
	Token    types.StringNull           `json:"token"`
	Password types.StringNull           `json:"password"`
}

func (dto AuthSudoVerification) GetWebAuthn() *WebAuthnLoginVerification {
	return dto.WebAuthn
}
func (dto AuthSudoVerification) GetToken() types.StringNull {
	return dto.Token
}
func (dto AuthSudoVerification) GetPassword() types.StringNull {
	return dto.Password
}

func (dto *AuthSudoVerification) UnmarshalJSON(b []byte) error {
	var aux = struct {
		WebAuthn *WebAuthnLoginVerification
		Token    types.StringNull
		Password types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("AuthSudoVerification", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.WebAuthn = aux.WebAuthn
	dto.Token = aux.Token
	dto.Password = aux.Password

	return nil
}
