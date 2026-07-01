// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutBillingClientAutocharge)(nil)

type PutBillingClientAutocharge struct {
	AutoTopUp *AutoTopUp `json:"autoTopUp"`
}

func (dto PutBillingClientAutocharge) GetAutoTopUp() *AutoTopUp {
	return dto.AutoTopUp
}

func (dto *PutBillingClientAutocharge) UnmarshalJSON(b []byte) error {
	var aux = struct {
		AutoTopUp *AutoTopUp
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutBillingClientAutocharge", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.AutoTopUp = aux.AutoTopUp

	return nil
}
