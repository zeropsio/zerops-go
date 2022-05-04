// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*AuthRefresh)(nil)

type AuthRefresh struct {
	RefreshTokenId uuid.AuthRefreshTokenId `json:"refreshTokenId"`
}

func (dto AuthRefresh) GetRefreshTokenId() uuid.AuthRefreshTokenId {
	return dto.RefreshTokenId
}

func (dto *AuthRefresh) UnmarshalJSON(b []byte) error {
	var aux = struct {
		RefreshTokenId *uuid.AuthRefreshTokenId
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("AuthRefresh", err)
	}
	var errorList validator.ErrorList
	if aux.RefreshTokenId == nil {
		errorList = errorList.With(validator.NewError("refreshTokenId", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.RefreshTokenId = *aux.RefreshTokenId

	return nil
}
