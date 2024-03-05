// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutProjectRemoteLogging)(nil)

type PutProjectRemoteLogging struct {
	SyslogConfig types.Text       `json:"syslogConfig"`
	Certificate  types.Text       `json:"certificate"`
	Type         types.StringNull `json:"type"`
}

func (dto PutProjectRemoteLogging) GetSyslogConfig() types.Text {
	return dto.SyslogConfig
}
func (dto PutProjectRemoteLogging) GetCertificate() types.Text {
	return dto.Certificate
}
func (dto PutProjectRemoteLogging) GetType() types.StringNull {
	return dto.Type
}

func (dto *PutProjectRemoteLogging) UnmarshalJSON(b []byte) error {
	var aux = struct {
		SyslogConfig *types.Text
		Certificate  *types.Text
		Type         types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutProjectRemoteLogging", err)
	}
	var errorList validator.ErrorList
	if aux.SyslogConfig == nil {
		errorList = errorList.With(validator.NewError("syslogConfig", "field is required"))
	}
	if aux.Certificate == nil {
		errorList = errorList.With(validator.NewError("certificate", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.SyslogConfig = *aux.SyslogConfig
	dto.Certificate = *aux.Certificate
	dto.Type = aux.Type

	return nil
}
