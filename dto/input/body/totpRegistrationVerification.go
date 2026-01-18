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
var _ json.Unmarshaler = (*TotpRegistrationVerification)(nil)

type TotpRegistrationVerification struct {
	SessionId uuid.TotpRegistrationVerificationSessionId `json:"sessionId"`
	Token     types.String                               `json:"token"`
}

func (dto TotpRegistrationVerification) GetSessionId() uuid.TotpRegistrationVerificationSessionId {
	return dto.SessionId
}
func (dto TotpRegistrationVerification) GetToken() types.String {
	return dto.Token
}

func (dto *TotpRegistrationVerification) UnmarshalJSON(b []byte) error {
	var aux = struct {
		SessionId *uuid.TotpRegistrationVerificationSessionId
		Token     *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("TotpRegistrationVerification", err)
	}
	var errorList validator.ErrorList
	if aux.SessionId == nil {
		errorList = errorList.With(validator.NewError("sessionId", "field is required"))
	}
	if aux.Token == nil {
		errorList = errorList.With(validator.NewError("token", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.SessionId = *aux.SessionId
	dto.Token = *aux.Token

	return nil
}
