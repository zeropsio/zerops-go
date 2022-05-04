// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutServiceStackServicePortLight)(nil)

type PutServiceStackServicePortLight struct {
	Port     types.Int                    `json:"port"`
	Protocol enum.ServicePortProtocolEnum `json:"protocol"`
	Scheme   enum.ServicePortSchemeEnum   `json:"scheme"`
}

func (dto PutServiceStackServicePortLight) GetPort() types.Int {
	return dto.Port
}
func (dto PutServiceStackServicePortLight) GetProtocol() enum.ServicePortProtocolEnum {
	return dto.Protocol
}
func (dto PutServiceStackServicePortLight) GetScheme() enum.ServicePortSchemeEnum {
	return dto.Scheme
}

func (dto *PutServiceStackServicePortLight) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Port     *types.Int
		Protocol *enum.ServicePortProtocolEnum
		Scheme   *enum.ServicePortSchemeEnum
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutServiceStackServicePortLight", err)
	}
	var errorList validator.ErrorList
	if aux.Port == nil {
		errorList = errorList.With(validator.NewError("port", "field is required"))
	}
	if aux.Protocol == nil {
		errorList = errorList.With(validator.NewError("protocol", "field is required"))
	}
	if aux.Scheme == nil {
		errorList = errorList.With(validator.NewError("scheme", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Port = *aux.Port
	dto.Protocol = *aux.Protocol
	dto.Scheme = *aux.Scheme

	return nil
}
