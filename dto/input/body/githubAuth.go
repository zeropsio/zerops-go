// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*GithubAuth)(nil)

type GithubAuth struct {
	Code         types.String     `json:"code"`
	State        types.String     `json:"state"`
	Email        types.StringNull `json:"email"`
	PromoCode    types.StringNull `json:"promoCode"`
	ClaimZcpPool types.BoolNull   `json:"claimZcpPool"`
}

func (dto GithubAuth) GetCode() types.String {
	return dto.Code
}
func (dto GithubAuth) GetState() types.String {
	return dto.State
}
func (dto GithubAuth) GetEmail() types.StringNull {
	return dto.Email
}
func (dto GithubAuth) GetPromoCode() types.StringNull {
	return dto.PromoCode
}
func (dto GithubAuth) GetClaimZcpPool() types.BoolNull {
	return dto.ClaimZcpPool
}

func (dto *GithubAuth) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Code         *types.String
		State        *types.String
		Email        types.StringNull
		PromoCode    types.StringNull
		ClaimZcpPool types.BoolNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("GithubAuth", err)
	}
	var errorList validator.ErrorList
	if aux.Code == nil {
		errorList = errorList.With(validator.NewError("code", "field is required"))
	}
	if aux.State == nil {
		errorList = errorList.With(validator.NewError("state", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Code = *aux.Code
	dto.State = *aux.State
	dto.Email = aux.Email
	dto.PromoCode = aux.PromoCode
	dto.ClaimZcpPool = aux.ClaimZcpPool

	return nil
}
