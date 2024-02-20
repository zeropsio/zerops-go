// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ClientUserConnection)(nil)

type ClientUserConnection struct {
	RoleCode           enum.ClientUserRoleCodeEnum `json:"roleCode"`
	LanguageId         stringId.LanguageId         `json:"languageId"`
	Email              types.Email                 `json:"email"`
	FirstName          types.String                `json:"firstName"`
	LastName           types.EmptyString           `json:"lastName"`
	CountryCallingCode types.IntNull               `json:"countryCallingCode"`
	PhoneNumber        types.IntNull               `json:"phoneNumber"`
}

func (dto ClientUserConnection) GetRoleCode() enum.ClientUserRoleCodeEnum {
	return dto.RoleCode
}
func (dto ClientUserConnection) GetLanguageId() stringId.LanguageId {
	return dto.LanguageId
}
func (dto ClientUserConnection) GetEmail() types.Email {
	return dto.Email
}
func (dto ClientUserConnection) GetFirstName() types.String {
	return dto.FirstName
}
func (dto ClientUserConnection) GetLastName() types.EmptyString {
	return dto.LastName
}
func (dto ClientUserConnection) GetCountryCallingCode() types.IntNull {
	return dto.CountryCallingCode
}
func (dto ClientUserConnection) GetPhoneNumber() types.IntNull {
	return dto.PhoneNumber
}

func (dto *ClientUserConnection) UnmarshalJSON(b []byte) error {
	var aux = struct {
		RoleCode           *enum.ClientUserRoleCodeEnum
		LanguageId         *stringId.LanguageId
		Email              *types.Email
		FirstName          *types.String
		LastName           *types.EmptyString
		CountryCallingCode types.IntNull
		PhoneNumber        types.IntNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ClientUserConnection", err)
	}
	var errorList validator.ErrorList
	if aux.RoleCode == nil {
		errorList = errorList.With(validator.NewError("roleCode", "field is required"))
	}
	if aux.LanguageId == nil {
		errorList = errorList.With(validator.NewError("languageId", "field is required"))
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
	if errorList != nil {
		return errorList.GetError()
	}
	dto.RoleCode = *aux.RoleCode
	dto.LanguageId = *aux.LanguageId
	dto.Email = *aux.Email
	dto.FirstName = *aux.FirstName
	dto.LastName = *aux.LastName
	dto.CountryCallingCode = aux.CountryCallingCode
	dto.PhoneNumber = aux.PhoneNumber

	return nil
}
