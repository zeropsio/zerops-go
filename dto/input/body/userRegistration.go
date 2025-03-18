// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*UserRegistration)(nil)

type UserRegistration struct {
	Email       types.Email         `json:"email"`
	Password    types.String        `json:"password"`
	LanguageId  stringId.LanguageId `json:"languageId"`
	AccountName types.String        `json:"accountName"`
	Name        types.String        `json:"name"`
	PromoCode   types.StringNull    `json:"promoCode"`
}

func (dto UserRegistration) GetEmail() types.Email {
	return dto.Email
}
func (dto UserRegistration) GetPassword() types.String {
	return dto.Password
}
func (dto UserRegistration) GetLanguageId() stringId.LanguageId {
	return dto.LanguageId
}
func (dto UserRegistration) GetAccountName() types.String {
	return dto.AccountName
}
func (dto UserRegistration) GetName() types.String {
	return dto.Name
}
func (dto UserRegistration) GetPromoCode() types.StringNull {
	return dto.PromoCode
}

func (dto *UserRegistration) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Email       *types.Email
		Password    *types.String
		LanguageId  *stringId.LanguageId
		AccountName *types.String
		Name        *types.String
		PromoCode   types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("UserRegistration", err)
	}
	var errorList validator.ErrorList
	if aux.Email == nil {
		errorList = errorList.With(validator.NewError("email", "field is required"))
	}
	if aux.Password == nil {
		errorList = errorList.With(validator.NewError("password", "field is required"))
	}
	if aux.LanguageId == nil {
		errorList = errorList.With(validator.NewError("languageId", "field is required"))
	}
	if aux.AccountName == nil {
		errorList = errorList.With(validator.NewError("accountName", "field is required"))
	}
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Email = *aux.Email
	dto.Password = *aux.Password
	dto.LanguageId = *aux.LanguageId
	dto.AccountName = *aux.AccountName
	dto.Name = *aux.Name
	dto.PromoCode = aux.PromoCode

	return nil
}
