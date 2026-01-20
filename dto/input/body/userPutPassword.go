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
	OldPassword types.StringNull `json:"oldPassword"`
	NewPassword types.StringNull `json:"newPassword"`
}

func (dto UserPutPassword) GetOldPassword() types.StringNull {
	return dto.OldPassword
}
func (dto UserPutPassword) GetNewPassword() types.StringNull {
	return dto.NewPassword
}

func (dto *UserPutPassword) UnmarshalJSON(b []byte) error {
	var aux = struct {
		OldPassword types.StringNull
		NewPassword types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("UserPutPassword", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.OldPassword = aux.OldPassword
	dto.NewPassword = aux.NewPassword

	return nil
}
