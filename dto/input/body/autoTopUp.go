// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*AutoTopUp)(nil)

type AutoTopUp struct {
	Threshold            types.Decimal `json:"threshold"`
	Amount               types.Decimal `json:"amount"`
	CalendarMonthlyLimit types.Decimal `json:"calendarMonthlyLimit"`
}

func (dto AutoTopUp) GetThreshold() types.Decimal {
	return dto.Threshold
}
func (dto AutoTopUp) GetAmount() types.Decimal {
	return dto.Amount
}
func (dto AutoTopUp) GetCalendarMonthlyLimit() types.Decimal {
	return dto.CalendarMonthlyLimit
}

func (dto *AutoTopUp) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Threshold            *types.Decimal
		Amount               *types.Decimal
		CalendarMonthlyLimit *types.Decimal
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("AutoTopUp", err)
	}
	var errorList validator.ErrorList
	if aux.Threshold == nil {
		errorList = errorList.With(validator.NewError("threshold", "field is required"))
	}
	if aux.Amount == nil {
		errorList = errorList.With(validator.NewError("amount", "field is required"))
	}
	if aux.CalendarMonthlyLimit == nil {
		errorList = errorList.With(validator.NewError("calendarMonthlyLimit", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Threshold = *aux.Threshold
	dto.Amount = *aux.Amount
	dto.CalendarMonthlyLimit = *aux.CalendarMonthlyLimit

	return nil
}
