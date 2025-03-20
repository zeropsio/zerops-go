// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostBillingPayment)(nil)

type PostBillingPayment struct {
	Amount    types.Decimal                `json:"amount"`
	ClientId  uuid.ClientId                `json:"clientId"`
	SourceId  stringId.PaymentSourceIdNull `json:"sourceId"`
	PromoCode types.StringNull             `json:"promoCode"`
}

func (dto PostBillingPayment) GetAmount() types.Decimal {
	return dto.Amount
}
func (dto PostBillingPayment) GetClientId() uuid.ClientId {
	return dto.ClientId
}
func (dto PostBillingPayment) GetSourceId() stringId.PaymentSourceIdNull {
	return dto.SourceId
}
func (dto PostBillingPayment) GetPromoCode() types.StringNull {
	return dto.PromoCode
}

func (dto *PostBillingPayment) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Amount    *types.Decimal
		ClientId  *uuid.ClientId
		SourceId  stringId.PaymentSourceIdNull
		PromoCode types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostBillingPayment", err)
	}
	var errorList validator.ErrorList
	if aux.Amount == nil {
		errorList = errorList.With(validator.NewError("amount", "field is required"))
	}
	if aux.ClientId == nil {
		errorList = errorList.With(validator.NewError("clientId", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Amount = *aux.Amount
	dto.ClientId = *aux.ClientId
	dto.SourceId = aux.SourceId
	dto.PromoCode = aux.PromoCode

	return nil
}
