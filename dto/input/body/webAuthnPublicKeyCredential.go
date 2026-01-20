// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*WebAuthnPublicKeyCredential)(nil)

type WebAuthnPublicKeyCredential struct {
	AuthenticatorAttachment types.String         `json:"authenticatorAttachment"`
	Id                      types.String         `json:"id"`
	RawId                   types.String         `json:"rawId"`
	Response                types.JsonRawMessage `json:"response"`
	Type                    types.String         `json:"type"`
	ClientExtensionResults  types.Map            `json:"clientExtensionResults"`
}

func (dto WebAuthnPublicKeyCredential) GetAuthenticatorAttachment() types.String {
	return dto.AuthenticatorAttachment
}
func (dto WebAuthnPublicKeyCredential) GetId() types.String {
	return dto.Id
}
func (dto WebAuthnPublicKeyCredential) GetRawId() types.String {
	return dto.RawId
}
func (dto WebAuthnPublicKeyCredential) GetResponse() types.JsonRawMessage {
	return dto.Response
}
func (dto WebAuthnPublicKeyCredential) GetType() types.String {
	return dto.Type
}
func (dto WebAuthnPublicKeyCredential) GetClientExtensionResults() types.Map {
	return dto.ClientExtensionResults
}

func (dto *WebAuthnPublicKeyCredential) UnmarshalJSON(b []byte) error {
	var aux = struct {
		AuthenticatorAttachment *types.String
		Id                      *types.String
		RawId                   *types.String
		Response                *types.JsonRawMessage
		Type                    *types.String
		ClientExtensionResults  *types.Map
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("WebAuthnPublicKeyCredential", err)
	}
	var errorList validator.ErrorList
	if aux.AuthenticatorAttachment == nil {
		errorList = errorList.With(validator.NewError("authenticatorAttachment", "field is required"))
	}
	if aux.Id == nil {
		errorList = errorList.With(validator.NewError("id", "field is required"))
	}
	if aux.RawId == nil {
		errorList = errorList.With(validator.NewError("rawId", "field is required"))
	}
	if aux.Response == nil {
		errorList = errorList.With(validator.NewError("response", "field is required"))
	}
	if aux.Type == nil {
		errorList = errorList.With(validator.NewError("type", "field is required"))
	}
	if aux.ClientExtensionResults == nil {
		errorList = errorList.With(validator.NewError("clientExtensionResults", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.AuthenticatorAttachment = *aux.AuthenticatorAttachment
	dto.Id = *aux.Id
	dto.RawId = *aux.RawId
	dto.Response = *aux.Response
	dto.Type = *aux.Type
	dto.ClientExtensionResults = *aux.ClientExtensionResults

	return nil
}
