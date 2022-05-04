// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ServiceStackNextAppVersion)(nil)

type ServiceStackNextAppVersion struct {
	UseCustomRuntime types.Bool `json:"useCustomRuntime"`
}

func (dto ServiceStackNextAppVersion) GetUseCustomRuntime() types.Bool {
	return dto.UseCustomRuntime
}

func (dto *ServiceStackNextAppVersion) UnmarshalJSON(b []byte) error {
	var aux = struct {
		UseCustomRuntime *types.Bool
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ServiceStackNextAppVersion", err)
	}
	var errorList validator.ErrorList
	if aux.UseCustomRuntime == nil {
		errorList = errorList.With(validator.NewError("useCustomRuntime", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.UseCustomRuntime = *aux.UseCustomRuntime

	return nil
}
