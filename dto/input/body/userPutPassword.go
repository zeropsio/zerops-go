// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*UserPutPassword)(nil)

type UserPutPassword struct {
	OldPassword types.String `json:"oldPassword"`
	NewPassword types.String `json:"newPassword"`
}

func (dto UserPutPassword) GetOldPassword() types.String {
	return dto.OldPassword
}
func (dto UserPutPassword) GetNewPassword() types.String {
	return dto.NewPassword
}

func (dto *UserPutPassword) UnmarshalJSON(b []byte) error {
	var aux = struct {
		OldPassword *types.String
		NewPassword *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("UserPutPassword", err)
	}
	var errorList validator.ErrorList
	if aux.OldPassword == nil {
		errorList = errorList.With(validator.NewError("oldPassword", "field is required"))
	}
	if aux.NewPassword == nil {
		errorList = errorList.With(validator.NewError("newPassword", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.OldPassword = *aux.OldPassword
	dto.NewPassword = *aux.NewPassword

	return nil
}
