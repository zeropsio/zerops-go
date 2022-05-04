// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostProjectTagList)(nil)

type PostProjectTagList struct {
	ClientId uuid.ClientId `json:"clientId"`
}

func (dto PostProjectTagList) GetClientId() uuid.ClientId {
	return dto.ClientId
}

func (dto *PostProjectTagList) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ClientId *uuid.ClientId
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostProjectTagList", err)
	}
	var errorList validator.ErrorList
	if aux.ClientId == nil {
		errorList = errorList.With(validator.NewError("clientId", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ClientId = *aux.ClientId

	return nil
}
