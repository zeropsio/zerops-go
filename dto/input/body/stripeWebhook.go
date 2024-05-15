// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*StripeWebhook)(nil)

type StripeWebhook struct {
	Type   types.String         `json:"type"`
	Object types.JsonRawMessage `json:"object"`
}

func (dto StripeWebhook) GetType() types.String {
	return dto.Type
}
func (dto StripeWebhook) GetObject() types.JsonRawMessage {
	return dto.Object
}

func (dto *StripeWebhook) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Type   *types.String
		Object *types.JsonRawMessage
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("StripeWebhook", err)
	}
	var errorList validator.ErrorList
	if aux.Type == nil {
		errorList = errorList.With(validator.NewError("type", "field is required"))
	}
	if aux.Object == nil {
		errorList = errorList.With(validator.NewError("object", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Type = *aux.Type
	dto.Object = *aux.Object

	return nil
}
