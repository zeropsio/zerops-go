// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostBillingCreditExport)(nil)

type PostBillingCreditExport struct {
	Month types.String `json:"month"`
}

func (dto PostBillingCreditExport) GetMonth() types.String {
	return dto.Month
}

func (dto *PostBillingCreditExport) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Month *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostBillingCreditExport", err)
	}
	var errorList validator.ErrorList
	if aux.Month == nil {
		errorList = errorList.With(validator.NewError("month", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Month = *aux.Month

	return nil
}
