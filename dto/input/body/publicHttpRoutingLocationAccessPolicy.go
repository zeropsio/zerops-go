// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PublicHttpRoutingLocationAccessPolicy)(nil)

type PublicHttpRoutingLocationAccessPolicy struct {
	Enabled       types.Bool                                                  `json:"enabled"`
	DefaultPolicy enum.PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum `json:"defaultPolicy"`
	Cidr          types.StringArray                                           `json:"cidr"`
}

func (dto PublicHttpRoutingLocationAccessPolicy) GetEnabled() types.Bool {
	return dto.Enabled
}
func (dto PublicHttpRoutingLocationAccessPolicy) GetDefaultPolicy() enum.PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum {
	return dto.DefaultPolicy
}
func (dto PublicHttpRoutingLocationAccessPolicy) GetCidr() types.StringArray {
	return dto.Cidr
}

func (dto *PublicHttpRoutingLocationAccessPolicy) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Enabled       *types.Bool
		DefaultPolicy *enum.PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum
		Cidr          *types.StringArray
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicHttpRoutingLocationAccessPolicy", err)
	}
	var errorList validator.ErrorList
	if aux.Enabled == nil {
		errorList = errorList.With(validator.NewError("enabled", "field is required"))
	}
	if aux.DefaultPolicy == nil {
		errorList = errorList.With(validator.NewError("defaultPolicy", "field is required"))
	}
	if aux.Cidr == nil {
		errorList = errorList.With(validator.NewError("cidr", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Enabled = *aux.Enabled
	dto.DefaultPolicy = *aux.DefaultPolicy
	dto.Cidr = *aux.Cidr

	return nil
}
