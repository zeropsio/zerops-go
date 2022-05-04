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
var _ json.Unmarshaler = (*PostVpnRequest)(nil)

type PostVpnRequest struct {
	Id              uuid.ProjectId `json:"id"`
	ClientPublicKey types.String   `json:"clientPublicKey"`
}

func (dto PostVpnRequest) GetId() uuid.ProjectId {
	return dto.Id
}
func (dto PostVpnRequest) GetClientPublicKey() types.String {
	return dto.ClientPublicKey
}

func (dto *PostVpnRequest) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Id              *uuid.ProjectId
		ClientPublicKey *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostVpnRequest", err)
	}
	var errorList validator.ErrorList
	if aux.Id == nil {
		errorList = errorList.With(validator.NewError("id", "field is required"))
	}
	if aux.ClientPublicKey == nil {
		errorList = errorList.With(validator.NewError("clientPublicKey", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Id = *aux.Id
	dto.ClientPublicKey = *aux.ClientPublicKey

	return nil
}
