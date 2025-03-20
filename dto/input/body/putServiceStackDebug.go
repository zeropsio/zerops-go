// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutServiceStackDebug)(nil)

type PutServiceStackDebug struct {
	DebugBuild          *enum.ServiceStackDebugDebugEnum `json:"debugBuild"`
	DebugRuntimePrepare *enum.ServiceStackDebugDebugEnum `json:"debugRuntimePrepare"`
}

func (dto PutServiceStackDebug) GetDebugBuild() *enum.ServiceStackDebugDebugEnum {
	return dto.DebugBuild
}
func (dto PutServiceStackDebug) GetDebugRuntimePrepare() *enum.ServiceStackDebugDebugEnum {
	return dto.DebugRuntimePrepare
}

func (dto *PutServiceStackDebug) UnmarshalJSON(b []byte) error {
	var aux = struct {
		DebugBuild          *enum.ServiceStackDebugDebugEnum
		DebugRuntimePrepare *enum.ServiceStackDebugDebugEnum
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutServiceStackDebug", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.DebugBuild = aux.DebugBuild
	dto.DebugRuntimePrepare = aux.DebugRuntimePrepare

	return nil
}
