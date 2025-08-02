// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PublicHttpRoutingLocationRedirect)(nil)

type PublicHttpRoutingLocationRedirect struct {
	Enabled       types.Bool        `json:"enabled"`
	To            types.EmptyString `json:"to"`
	Code          types.Int         `json:"code"`
	PreservePath  types.Bool        `json:"preservePath"`
	PreserveQuery types.Bool        `json:"preserveQuery"`
}

func (dto PublicHttpRoutingLocationRedirect) GetEnabled() types.Bool {
	return dto.Enabled
}
func (dto PublicHttpRoutingLocationRedirect) GetTo() types.EmptyString {
	return dto.To
}
func (dto PublicHttpRoutingLocationRedirect) GetCode() types.Int {
	return dto.Code
}
func (dto PublicHttpRoutingLocationRedirect) GetPreservePath() types.Bool {
	return dto.PreservePath
}
func (dto PublicHttpRoutingLocationRedirect) GetPreserveQuery() types.Bool {
	return dto.PreserveQuery
}

func (dto *PublicHttpRoutingLocationRedirect) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Enabled       *types.Bool
		To            *types.EmptyString
		Code          *types.Int
		PreservePath  *types.Bool
		PreserveQuery *types.Bool
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicHttpRoutingLocationRedirect", err)
	}
	var errorList validator.ErrorList
	if aux.Enabled == nil {
		errorList = errorList.With(validator.NewError("enabled", "field is required"))
	}
	if aux.To == nil {
		errorList = errorList.With(validator.NewError("to", "field is required"))
	}
	if aux.Code == nil {
		errorList = errorList.With(validator.NewError("code", "field is required"))
	}
	if aux.PreservePath == nil {
		errorList = errorList.With(validator.NewError("preservePath", "field is required"))
	}
	if aux.PreserveQuery == nil {
		errorList = errorList.With(validator.NewError("preserveQuery", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Enabled = *aux.Enabled
	dto.To = *aux.To
	dto.Code = *aux.Code
	dto.PreservePath = *aux.PreservePath
	dto.PreserveQuery = *aux.PreserveQuery

	return nil
}
