// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutServiceStackPorts)(nil)

type PutServiceStackPorts struct {
	Ports PutServiceStackPortsPorts `json:"ports"`
}

func (dto PutServiceStackPorts) GetPorts() PutServiceStackPortsPorts {
	return dto.Ports
}

type PutServiceStackPortsPorts []PutServiceStackServicePortLight

func (dto PutServiceStackPortsPorts) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PutServiceStackServicePortLight(dto))
}

func (dto *PutServiceStackPorts) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Ports *PutServiceStackPortsPorts
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutServiceStackPorts", err)
	}
	var errorList validator.ErrorList
	if aux.Ports == nil {
		errorList = errorList.With(validator.NewError("ports", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Ports = *aux.Ports

	return nil
}
