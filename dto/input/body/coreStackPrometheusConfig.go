// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*CoreStackPrometheusConfig)(nil)

type CoreStackPrometheusConfig struct {
	Configs    CoreStackPrometheusConfigConfigs `json:"configs"`
	InstanceId uuid.InstanceIdNull              `json:"instanceId"`
}

func (dto CoreStackPrometheusConfig) GetConfigs() CoreStackPrometheusConfigConfigs {
	return dto.Configs
}
func (dto CoreStackPrometheusConfig) GetInstanceId() uuid.InstanceIdNull {
	return dto.InstanceId
}

type CoreStackPrometheusConfigConfigs []PrometheusServiceDiscovery

func (dto CoreStackPrometheusConfigConfigs) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PrometheusServiceDiscovery(dto))
}

func (dto *CoreStackPrometheusConfig) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Configs    *CoreStackPrometheusConfigConfigs
		InstanceId uuid.InstanceIdNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("CoreStackPrometheusConfig", err)
	}
	var errorList validator.ErrorList
	if aux.Configs == nil {
		errorList = errorList.With(validator.NewError("configs", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Configs = *aux.Configs
	dto.InstanceId = aux.InstanceId

	return nil
}
