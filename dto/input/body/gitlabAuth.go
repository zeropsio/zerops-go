// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*GitlabAuth)(nil)

type GitlabAuth struct {
	Code  types.String `json:"code"`
	State types.String `json:"state"`
}

func (dto GitlabAuth) GetCode() types.String {
	return dto.Code
}
func (dto GitlabAuth) GetState() types.String {
	return dto.State
}

func (dto *GitlabAuth) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Code  *types.String
		State *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("GitlabAuth", err)
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

	return nil
}
