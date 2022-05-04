// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*User)(nil)

type User struct {
	ClientId           uuid.ClientId                    `json:"clientId"`
	Email              types.Email                      `json:"email"`
	FirstName          types.String                     `json:"firstName"`
	LastName           types.String                     `json:"lastName"`
	LanguageId         types.String                     `json:"languageId"`
	CountryCallingCode types.IntNull                    `json:"countryCallingCode"`
	PhoneNumber        types.IntNull                    `json:"phoneNumber"`
	RoleCode           enum.ClientUserLightRoleCodeEnum `json:"roleCode"`
}

func (dto User) GetClientId() uuid.ClientId {
	return dto.ClientId
}
func (dto User) GetEmail() types.Email {
	return dto.Email
}
func (dto User) GetFirstName() types.String {
	return dto.FirstName
}
func (dto User) GetLastName() types.String {
	return dto.LastName
}
func (dto User) GetLanguageId() types.String {
	return dto.LanguageId
}
func (dto User) GetCountryCallingCode() types.IntNull {
	return dto.CountryCallingCode
}
func (dto User) GetPhoneNumber() types.IntNull {
	return dto.PhoneNumber
}
func (dto User) GetRoleCode() enum.ClientUserLightRoleCodeEnum {
	return dto.RoleCode
}

func (dto *User) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ClientId           *uuid.ClientId
		Email              *types.Email
		FirstName          *types.String
		LastName           *types.String
		LanguageId         *types.String
		CountryCallingCode types.IntNull
		PhoneNumber        types.IntNull
		RoleCode           *enum.ClientUserLightRoleCodeEnum
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("User", err)
	}
	var errorList validator.ErrorList
	if aux.ClientId == nil {
		errorList = errorList.With(validator.NewError("clientId", "field is required"))
	}
	if aux.Email == nil {
		errorList = errorList.With(validator.NewError("email", "field is required"))
	}
	if aux.FirstName == nil {
		errorList = errorList.With(validator.NewError("firstName", "field is required"))
	}
	if aux.LastName == nil {
		errorList = errorList.With(validator.NewError("lastName", "field is required"))
	}
	if aux.LanguageId == nil {
		errorList = errorList.With(validator.NewError("languageId", "field is required"))
	}
	if aux.RoleCode == nil {
		errorList = errorList.With(validator.NewError("roleCode", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ClientId = *aux.ClientId
	dto.Email = *aux.Email
	dto.FirstName = *aux.FirstName
	dto.LastName = *aux.LastName
	dto.LanguageId = *aux.LanguageId
	dto.CountryCallingCode = aux.CountryCallingCode
	dto.PhoneNumber = aux.PhoneNumber
	dto.RoleCode = *aux.RoleCode

	return nil
}
