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
var _ json.Unmarshaler = (*PostProjectVpn)(nil)

type PostProjectVpn struct {
	PublicKey  types.String        `json:"publicKey"`
	InstanceId uuid.InstanceIdNull `json:"instanceId"`
}

func (dto PostProjectVpn) GetPublicKey() types.String {
	return dto.PublicKey
}
func (dto PostProjectVpn) GetInstanceId() uuid.InstanceIdNull {
	return dto.InstanceId
}

func (dto *PostProjectVpn) UnmarshalJSON(b []byte) error {
	var aux = struct {
		PublicKey  *types.String
		InstanceId uuid.InstanceIdNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostProjectVpn", err)
	}
	var errorList validator.ErrorList
	if aux.PublicKey == nil {
		errorList = errorList.With(validator.NewError("publicKey", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.PublicKey = *aux.PublicKey
	dto.InstanceId = aux.InstanceId

	return nil
}
