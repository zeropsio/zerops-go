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
var _ json.Unmarshaler = (*DeviceLoginCheck)(nil)

type DeviceLoginCheck struct {
	RequestId  uuid.DeviceLoginRequestId `json:"requestId"`
	PollSecret types.String              `json:"pollSecret"`
	AppSecret  types.StringNull          `json:"appSecret"`
	Code       types.StringNull          `json:"code"`
}

func (dto DeviceLoginCheck) GetRequestId() uuid.DeviceLoginRequestId {
	return dto.RequestId
}
func (dto DeviceLoginCheck) GetPollSecret() types.String {
	return dto.PollSecret
}
func (dto DeviceLoginCheck) GetAppSecret() types.StringNull {
	return dto.AppSecret
}
func (dto DeviceLoginCheck) GetCode() types.StringNull {
	return dto.Code
}

func (dto *DeviceLoginCheck) UnmarshalJSON(b []byte) error {
	var aux = struct {
		RequestId  *uuid.DeviceLoginRequestId
		PollSecret *types.String
		AppSecret  types.StringNull
		Code       types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("DeviceLoginCheck", err)
	}
	var errorList validator.ErrorList
	if aux.RequestId == nil {
		errorList = errorList.With(validator.NewError("requestId", "field is required"))
	}
	if aux.PollSecret == nil {
		errorList = errorList.With(validator.NewError("pollSecret", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.RequestId = *aux.RequestId
	dto.PollSecret = *aux.PollSecret
	dto.AppSecret = aux.AppSecret
	dto.Code = aux.Code

	return nil
}
