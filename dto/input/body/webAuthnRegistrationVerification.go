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
var _ json.Unmarshaler = (*WebAuthnRegistrationVerification)(nil)

type WebAuthnRegistrationVerification struct {
	SessionId  uuid.WebAuthnLoginVerificationSessionId `json:"sessionId"`
	DeviceName types.String                            `json:"deviceName"`
	Response   WebAuthnPublicKeyCredential             `json:"response"`
}

func (dto WebAuthnRegistrationVerification) GetSessionId() uuid.WebAuthnLoginVerificationSessionId {
	return dto.SessionId
}
func (dto WebAuthnRegistrationVerification) GetDeviceName() types.String {
	return dto.DeviceName
}
func (dto WebAuthnRegistrationVerification) GetResponse() WebAuthnPublicKeyCredential {
	return dto.Response
}

func (dto *WebAuthnRegistrationVerification) UnmarshalJSON(b []byte) error {
	var aux = struct {
		SessionId  *uuid.WebAuthnLoginVerificationSessionId
		DeviceName *types.String
		Response   *WebAuthnPublicKeyCredential
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("WebAuthnRegistrationVerification", err)
	}
	var errorList validator.ErrorList
	if aux.SessionId == nil {
		errorList = errorList.With(validator.NewError("sessionId", "field is required"))
	}
	if aux.DeviceName == nil {
		errorList = errorList.With(validator.NewError("deviceName", "field is required"))
	}
	if aux.Response == nil {
		errorList = errorList.With(validator.NewError("response", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.SessionId = *aux.SessionId
	dto.DeviceName = *aux.DeviceName
	dto.Response = *aux.Response

	return nil
}
