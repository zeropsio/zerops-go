// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*DeviceLoginInitiate)(nil)

type DeviceLoginInitiate struct {
	AppCode   types.String     `json:"appCode"`
	AppSecret types.StringNull `json:"appSecret"`
}

func (dto DeviceLoginInitiate) GetAppCode() types.String {
	return dto.AppCode
}
func (dto DeviceLoginInitiate) GetAppSecret() types.StringNull {
	return dto.AppSecret
}

func (dto *DeviceLoginInitiate) UnmarshalJSON(b []byte) error {
	var aux = struct {
		AppCode   *types.String
		AppSecret types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("DeviceLoginInitiate", err)
	}
	var errorList validator.ErrorList
	if aux.AppCode == nil {
		errorList = errorList.With(validator.NewError("appCode", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.AppCode = *aux.AppCode
	dto.AppSecret = aux.AppSecret

	return nil
}
