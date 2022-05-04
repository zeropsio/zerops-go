// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostBillingPayment)(nil)

type PostBillingPayment struct {
	AmountVat  types.Float         `json:"amountVat"`
	CurrencyId enum.CurrencyIdEnum `json:"currencyId"`
	ClientId   uuid.ClientId       `json:"clientId"`
	SourceId   types.StringNull    `json:"sourceId"`
}

func (dto PostBillingPayment) GetAmountVat() types.Float {
	return dto.AmountVat
}
func (dto PostBillingPayment) GetCurrencyId() enum.CurrencyIdEnum {
	return dto.CurrencyId
}
func (dto PostBillingPayment) GetClientId() uuid.ClientId {
	return dto.ClientId
}
func (dto PostBillingPayment) GetSourceId() types.StringNull {
	return dto.SourceId
}

func (dto *PostBillingPayment) UnmarshalJSON(b []byte) error {
	var aux = struct {
		AmountVat  *types.Float
		CurrencyId *enum.CurrencyIdEnum
		ClientId   *uuid.ClientId
		SourceId   types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostBillingPayment", err)
	}
	var errorList validator.ErrorList
	if aux.AmountVat == nil {
		errorList = errorList.With(validator.NewError("amountVat", "field is required"))
	}
	if aux.CurrencyId == nil {
		errorList = errorList.With(validator.NewError("currencyId", "field is required"))
	}
	if aux.ClientId == nil {
		errorList = errorList.With(validator.NewError("clientId", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.AmountVat = *aux.AmountVat
	dto.CurrencyId = *aux.CurrencyId
	dto.ClientId = *aux.ClientId
	dto.SourceId = aux.SourceId

	return nil
}
