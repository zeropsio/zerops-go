// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutBillingAddon)(nil)

type PutBillingAddon struct {
	RecurringEnabled types.Bool `json:"recurringEnabled"`
	Enabled          types.Bool `json:"enabled"`
}

func (dto PutBillingAddon) GetRecurringEnabled() types.Bool {
	return dto.RecurringEnabled
}
func (dto PutBillingAddon) GetEnabled() types.Bool {
	return dto.Enabled
}

func (dto *PutBillingAddon) UnmarshalJSON(b []byte) error {
	var aux = struct {
		RecurringEnabled *types.Bool
		Enabled          *types.Bool
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutBillingAddon", err)
	}
	var errorList validator.ErrorList
	if aux.RecurringEnabled == nil {
		errorList = errorList.With(validator.NewError("recurringEnabled", "field is required"))
	}
	if aux.Enabled == nil {
		errorList = errorList.With(validator.NewError("enabled", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.RecurringEnabled = *aux.RecurringEnabled
	dto.Enabled = *aux.Enabled

	return nil
}
