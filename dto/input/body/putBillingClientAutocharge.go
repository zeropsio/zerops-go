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
var _ json.Unmarshaler = (*PutBillingClientAutocharge)(nil)

type PutBillingClientAutocharge struct {
	AutoCharge          types.Bool                           `json:"autoCharge"`
	AutoChargePeriod    enum.PaymentInfoAutoChargePeriodEnum `json:"autoChargePeriod"`
	MaximumChargeAmount types.Decimal                        `json:"maximumChargeAmount"`
}

func (dto PutBillingClientAutocharge) GetAutoCharge() types.Bool {
	return dto.AutoCharge
}
func (dto PutBillingClientAutocharge) GetAutoChargePeriod() enum.PaymentInfoAutoChargePeriodEnum {
	return dto.AutoChargePeriod
}
func (dto PutBillingClientAutocharge) GetMaximumChargeAmount() types.Decimal {
	return dto.MaximumChargeAmount
}

func (dto *PutBillingClientAutocharge) UnmarshalJSON(b []byte) error {
	var aux = struct {
		AutoCharge          *types.Bool
		AutoChargePeriod    *enum.PaymentInfoAutoChargePeriodEnum
		MaximumChargeAmount *types.Decimal
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutBillingClientAutocharge", err)
	}
	var errorList validator.ErrorList
	if aux.AutoCharge == nil {
		errorList = errorList.With(validator.NewError("autoCharge", "field is required"))
	}
	if aux.AutoChargePeriod == nil {
		errorList = errorList.With(validator.NewError("autoChargePeriod", "field is required"))
	}
	if aux.MaximumChargeAmount == nil {
		errorList = errorList.With(validator.NewError("maximumChargeAmount", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.AutoCharge = *aux.AutoCharge
	dto.AutoChargePeriod = *aux.AutoChargePeriod
	dto.MaximumChargeAmount = *aux.MaximumChargeAmount

	return nil
}
