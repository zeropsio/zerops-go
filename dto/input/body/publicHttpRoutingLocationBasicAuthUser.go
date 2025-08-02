// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PublicHttpRoutingLocationBasicAuthUser)(nil)

type PublicHttpRoutingLocationBasicAuthUser struct {
	Username types.EmptyString `json:"username"`
	Password types.EmptyString `json:"password"`
	Comment  types.EmptyString `json:"comment"`
}

func (dto PublicHttpRoutingLocationBasicAuthUser) GetUsername() types.EmptyString {
	return dto.Username
}
func (dto PublicHttpRoutingLocationBasicAuthUser) GetPassword() types.EmptyString {
	return dto.Password
}
func (dto PublicHttpRoutingLocationBasicAuthUser) GetComment() types.EmptyString {
	return dto.Comment
}

func (dto *PublicHttpRoutingLocationBasicAuthUser) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Username *types.EmptyString
		Password *types.EmptyString
		Comment  *types.EmptyString
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicHttpRoutingLocationBasicAuthUser", err)
	}
	var errorList validator.ErrorList
	if aux.Username == nil {
		errorList = errorList.With(validator.NewError("username", "field is required"))
	}
	if aux.Password == nil {
		errorList = errorList.With(validator.NewError("password", "field is required"))
	}
	if aux.Comment == nil {
		errorList = errorList.With(validator.NewError("comment", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Username = *aux.Username
	dto.Password = *aux.Password
	dto.Comment = *aux.Comment

	return nil
}
