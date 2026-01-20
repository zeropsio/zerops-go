// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PrometheusServiceDiscovery)(nil)

type PrometheusServiceDiscovery struct {
	Targets types.StringArray `json:"targets"`
	Labels  types.Map         `json:"labels"`
}

func (dto PrometheusServiceDiscovery) GetTargets() types.StringArray {
	return dto.Targets
}
func (dto PrometheusServiceDiscovery) GetLabels() types.Map {
	return dto.Labels
}

func (dto *PrometheusServiceDiscovery) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Targets *types.StringArray
		Labels  *types.Map
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PrometheusServiceDiscovery", err)
	}
	var errorList validator.ErrorList
	if aux.Targets == nil {
		errorList = errorList.With(validator.NewError("targets", "field is required"))
	}
	if aux.Labels == nil {
		errorList = errorList.With(validator.NewError("labels", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Targets = *aux.Targets
	dto.Labels = *aux.Labels

	return nil
}
