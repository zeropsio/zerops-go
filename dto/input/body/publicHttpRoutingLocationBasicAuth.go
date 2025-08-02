// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PublicHttpRoutingLocationBasicAuth)(nil)

type PublicHttpRoutingLocationBasicAuth struct {
	Enabled types.Bool                              `json:"enabled"`
	Realm   types.EmptyString                       `json:"realm"`
	Users   PublicHttpRoutingLocationBasicAuthUsers `json:"users"`
}

func (dto PublicHttpRoutingLocationBasicAuth) GetEnabled() types.Bool {
	return dto.Enabled
}
func (dto PublicHttpRoutingLocationBasicAuth) GetRealm() types.EmptyString {
	return dto.Realm
}
func (dto PublicHttpRoutingLocationBasicAuth) GetUsers() PublicHttpRoutingLocationBasicAuthUsers {
	return dto.Users
}

type PublicHttpRoutingLocationBasicAuthUsers []PublicHttpRoutingLocationBasicAuthUser

func (dto PublicHttpRoutingLocationBasicAuthUsers) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicHttpRoutingLocationBasicAuthUser(dto))
}

func (dto *PublicHttpRoutingLocationBasicAuth) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Enabled *types.Bool
		Realm   *types.EmptyString
		Users   *PublicHttpRoutingLocationBasicAuthUsers
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicHttpRoutingLocationBasicAuth", err)
	}
	var errorList validator.ErrorList
	if aux.Enabled == nil {
		errorList = errorList.With(validator.NewError("enabled", "field is required"))
	}
	if aux.Realm == nil {
		errorList = errorList.With(validator.NewError("realm", "field is required"))
	}
	if aux.Users == nil {
		errorList = errorList.With(validator.NewError("users", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Enabled = *aux.Enabled
	dto.Realm = *aux.Realm
	dto.Users = *aux.Users

	return nil
}
