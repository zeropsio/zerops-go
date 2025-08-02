// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutProjectL7httpbalancerConfig)(nil)

type PutProjectL7httpbalancerConfig struct {
	Values PutProjectL7httpbalancerConfigValues `json:"values"`
}

func (dto PutProjectL7httpbalancerConfig) GetValues() PutProjectL7httpbalancerConfigValues {
	return dto.Values
}

type PutProjectL7httpbalancerConfigValues []PutProjectL7httpbalancerConfigValue

func (dto PutProjectL7httpbalancerConfigValues) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PutProjectL7httpbalancerConfigValue(dto))
}

func (dto *PutProjectL7httpbalancerConfig) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Values *PutProjectL7httpbalancerConfigValues
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutProjectL7httpbalancerConfig", err)
	}
	var errorList validator.ErrorList
	if aux.Values == nil {
		errorList = errorList.With(validator.NewError("values", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Values = *aux.Values

	return nil
}
