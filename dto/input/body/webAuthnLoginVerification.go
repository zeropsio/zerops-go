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
var _ json.Unmarshaler = (*WebAuthnLoginVerification)(nil)

type WebAuthnLoginVerification struct {
	SessionId uuid.WebAuthnLoginVerificationSessionId `json:"sessionId"`
	Response  types.JsonRawMessage                    `json:"response"`
}

func (dto WebAuthnLoginVerification) GetSessionId() uuid.WebAuthnLoginVerificationSessionId {
	return dto.SessionId
}
func (dto WebAuthnLoginVerification) GetResponse() types.JsonRawMessage {
	return dto.Response
}

func (dto *WebAuthnLoginVerification) UnmarshalJSON(b []byte) error {
	var aux = struct {
		SessionId *uuid.WebAuthnLoginVerificationSessionId
		Response  *types.JsonRawMessage
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("WebAuthnLoginVerification", err)
	}
	var errorList validator.ErrorList
	if aux.SessionId == nil {
		errorList = errorList.With(validator.NewError("sessionId", "field is required"))
	}
	if aux.Response == nil {
		errorList = errorList.With(validator.NewError("response", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.SessionId = *aux.SessionId
	dto.Response = *aux.Response

	return nil
}
