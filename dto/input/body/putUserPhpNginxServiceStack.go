// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutUserPhpNginxServiceStack)(nil)

type PutUserPhpNginxServiceStack struct {
	NginxConfig types.Text `json:"nginxConfig"`
}

func (dto PutUserPhpNginxServiceStack) GetNginxConfig() types.Text {
	return dto.NginxConfig
}

func (dto *PutUserPhpNginxServiceStack) UnmarshalJSON(b []byte) error {
	var aux = struct {
		NginxConfig *types.Text
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutUserPhpNginxServiceStack", err)
	}
	var errorList validator.ErrorList
	if aux.NginxConfig == nil {
		errorList = errorList.With(validator.NewError("nginxConfig", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.NginxConfig = *aux.NginxConfig

	return nil
}
