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
var _ json.Unmarshaler = (*PublicHttpRoutingLocationPut)(nil)

type PublicHttpRoutingLocationPut struct {
	Path           types.String        `json:"path"`
	Port           types.Int           `json:"port"`
	ServiceStackId uuid.ServiceStackId `json:"serviceStackId"`
}

func (dto PublicHttpRoutingLocationPut) GetPath() types.String {
	return dto.Path
}
func (dto PublicHttpRoutingLocationPut) GetPort() types.Int {
	return dto.Port
}
func (dto PublicHttpRoutingLocationPut) GetServiceStackId() uuid.ServiceStackId {
	return dto.ServiceStackId
}

func (dto *PublicHttpRoutingLocationPut) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Path           *types.String
		Port           *types.Int
		ServiceStackId *uuid.ServiceStackId
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicHttpRoutingLocationPut", err)
	}
	var errorList validator.ErrorList
	if aux.Path == nil {
		errorList = errorList.With(validator.NewError("path", "field is required"))
	}
	if aux.Port == nil {
		errorList = errorList.With(validator.NewError("port", "field is required"))
	}
	if aux.ServiceStackId == nil {
		errorList = errorList.With(validator.NewError("serviceStackId", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Path = *aux.Path
	dto.Port = *aux.Port
	dto.ServiceStackId = *aux.ServiceStackId

	return nil
}
