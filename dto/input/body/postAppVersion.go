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
var _ json.Unmarshaler = (*PostAppVersion)(nil)

type PostAppVersion struct {
	ServiceStackId uuid.ServiceStackId `json:"serviceStackId"`
	Name           types.StringNull    `json:"name"`
}

func (dto PostAppVersion) GetServiceStackId() uuid.ServiceStackId {
	return dto.ServiceStackId
}
func (dto PostAppVersion) GetName() types.StringNull {
	return dto.Name
}

func (dto *PostAppVersion) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ServiceStackId *uuid.ServiceStackId
		Name           types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostAppVersion", err)
	}
	var errorList validator.ErrorList
	if aux.ServiceStackId == nil {
		errorList = errorList.With(validator.NewError("serviceStackId", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ServiceStackId = *aux.ServiceStackId
	dto.Name = aux.Name

	return nil
}
