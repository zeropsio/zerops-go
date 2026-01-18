// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*WebAuthnRegistrationOptions)(nil)

type WebAuthnRegistrationOptions struct {
	ResidentKey types.Bool `json:"residentKey"`
}

func (dto WebAuthnRegistrationOptions) GetResidentKey() types.Bool {
	return dto.ResidentKey
}

func (dto *WebAuthnRegistrationOptions) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ResidentKey *types.Bool
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("WebAuthnRegistrationOptions", err)
	}
	var errorList validator.ErrorList
	if aux.ResidentKey == nil {
		errorList = errorList.With(validator.NewError("residentKey", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ResidentKey = *aux.ResidentKey

	return nil
}
