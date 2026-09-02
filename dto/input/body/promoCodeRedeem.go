// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PromoCodeRedeem)(nil)

type PromoCodeRedeem struct {
	Code types.String `json:"code"`
}

func (dto PromoCodeRedeem) GetCode() types.String {
	return dto.Code
}

func (dto *PromoCodeRedeem) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Code *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PromoCodeRedeem", err)
	}
	var errorList validator.ErrorList
	if aux.Code == nil {
		errorList = errorList.With(validator.NewError("code", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Code = *aux.Code

	return nil
}
