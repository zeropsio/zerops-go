// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutProjectL7httpbalancerConfigValue)(nil)

type PutProjectL7httpbalancerConfigValue struct {
	Name  types.String     `json:"name"`
	Value types.StringNull `json:"value"`
}

func (dto PutProjectL7httpbalancerConfigValue) GetName() types.String {
	return dto.Name
}
func (dto PutProjectL7httpbalancerConfigValue) GetValue() types.StringNull {
	return dto.Value
}

func (dto *PutProjectL7httpbalancerConfigValue) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Name  *types.String
		Value types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutProjectL7httpbalancerConfigValue", err)
	}
	var errorList validator.ErrorList
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Name = *aux.Name
	dto.Value = aux.Value

	return nil
}
