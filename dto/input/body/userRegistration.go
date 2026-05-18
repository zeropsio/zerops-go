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
	Password    types.StringNull    `json:"password"`
	LanguageId  stringId.LanguageId `json:"languageId"`
	AccountName types.String        `json:"accountName"`
	Name        types.String        `json:"name"`
	PromoCode   types.StringNull    `json:"promoCode"`
	Reference   types.StringNull    `json:"reference"`
	Token       types.TextNull      `json:"token"`
	Passkey     types.BoolNull      `json:"passkey"`
}

func (dto UserRegistration) GetEmail() types.Email {
	return dto.Email
}
func (dto UserRegistration) GetPassword() types.StringNull {
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
func (dto UserRegistration) GetReference() types.StringNull {
	return dto.Reference
}
func (dto UserRegistration) GetToken() types.TextNull {
	return dto.Token
}
func (dto UserRegistration) GetPasskey() types.BoolNull {
	return dto.Passkey
}

func (dto *UserRegistration) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Email       *types.Email
		Password    types.StringNull
		LanguageId  *stringId.LanguageId
		AccountName *types.String
		Name        *types.String
		PromoCode   types.StringNull
		Reference   types.StringNull
		Token       types.TextNull
		Passkey     types.BoolNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("UserRegistration", err)
	}
	var errorList validator.ErrorList
	if aux.Email == nil {
		errorList = errorList.With(validator.NewError("email", "field is required"))
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
	dto.Password = aux.Password
	dto.LanguageId = *aux.LanguageId
	dto.AccountName = *aux.AccountName
	dto.Name = *aux.Name
	dto.PromoCode = aux.PromoCode
	dto.Reference = aux.Reference
	dto.Token = aux.Token
	dto.Passkey = aux.Passkey

	return nil
}
