// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutProjectPublicIpV4Shared)(nil)

type PutProjectPublicIpV4Shared struct {
	Enabled types.Bool `json:"enabled"`
}

func (dto PutProjectPublicIpV4Shared) GetEnabled() types.Bool {
	return dto.Enabled
}

func (dto *PutProjectPublicIpV4Shared) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Enabled *types.Bool
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutProjectPublicIpV4Shared", err)
	}
	var errorList validator.ErrorList
	if aux.Enabled == nil {
		errorList = errorList.With(validator.NewError("enabled", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Enabled = *aux.Enabled

	return nil
}
