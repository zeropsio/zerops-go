// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*DeleteBillingClientPaymentSource)(nil)

type DeleteBillingClientPaymentSource struct {
	SourceId stringId.PaymentSourceId `json:"sourceId"`
}

func (dto DeleteBillingClientPaymentSource) GetSourceId() stringId.PaymentSourceId {
	return dto.SourceId
}

func (dto *DeleteBillingClientPaymentSource) UnmarshalJSON(b []byte) error {
	var aux = struct {
		SourceId *stringId.PaymentSourceId
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("DeleteBillingClientPaymentSource", err)
	}
	var errorList validator.ErrorList
	if aux.SourceId == nil {
		errorList = errorList.With(validator.NewError("sourceId", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.SourceId = *aux.SourceId

	return nil
}
