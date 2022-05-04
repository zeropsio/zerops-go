// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*UserPut)(nil)

type UserPut struct {
	ClientId           uuid.ClientId `json:"clientId"`
	FirstName          types.String  `json:"firstName"`
	LastName           types.String  `json:"lastName"`
	LanguageId         types.String  `json:"languageId"`
	CountryCallingCode types.IntNull `json:"countryCallingCode"`
	PhoneNumber        types.IntNull `json:"phoneNumber"`
}

func (dto UserPut) GetClientId() uuid.ClientId {
	return dto.ClientId
}
func (dto UserPut) GetFirstName() types.String {
	return dto.FirstName
}
func (dto UserPut) GetLastName() types.String {
	return dto.LastName
}
func (dto UserPut) GetLanguageId() types.String {
	return dto.LanguageId
}
func (dto UserPut) GetCountryCallingCode() types.IntNull {
	return dto.CountryCallingCode
}
func (dto UserPut) GetPhoneNumber() types.IntNull {
	return dto.PhoneNumber
}

func (dto *UserPut) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ClientId           *uuid.ClientId
		FirstName          *types.String
		LastName           *types.String
		LanguageId         *types.String
		CountryCallingCode types.IntNull
		PhoneNumber        types.IntNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("UserPut", err)
	}
	var errorList validator.ErrorList
	if aux.ClientId == nil {
		errorList = errorList.With(validator.NewError("clientId", "field is required"))
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
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ClientId = *aux.ClientId
	dto.FirstName = *aux.FirstName
	dto.LastName = *aux.LastName
	dto.LanguageId = *aux.LanguageId
	dto.CountryCallingCode = aux.CountryCallingCode
	dto.PhoneNumber = aux.PhoneNumber

	return nil
}
