// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*RecoverPassword)(nil)

type RecoverPassword struct {
	Email types.Email `json:"email"`
}

func (dto RecoverPassword) GetEmail() types.Email {
	return dto.Email
}

func (dto *RecoverPassword) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Email *types.Email
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("RecoverPassword", err)
	}
	var errorList validator.ErrorList
	if aux.Email == nil {
		errorList = errorList.With(validator.NewError("email", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Email = *aux.Email

	return nil
}
