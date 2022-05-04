// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostClientCheckVAT)(nil)

type PostClientCheckVAT struct {
	VatNumber types.String `json:"vatNumber"`
}

func (dto PostClientCheckVAT) GetVatNumber() types.String {
	return dto.VatNumber
}

func (dto *PostClientCheckVAT) UnmarshalJSON(b []byte) error {
	var aux = struct {
		VatNumber *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostClientCheckVAT", err)
	}
	var errorList validator.ErrorList
	if aux.VatNumber == nil {
		errorList = errorList.With(validator.NewError("vatNumber", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.VatNumber = *aux.VatNumber

	return nil
}
