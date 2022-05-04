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
var _ json.Unmarshaler = (*PostProject)(nil)

type PostProject struct {
	ClientId       uuid.ClientId     `json:"clientId"`
	Name           types.String      `json:"name"`
	Description    types.TextNull    `json:"description"`
	TagList        types.StringArray `json:"tagList"`
	MaxCreditLimit types.DecimalNull `json:"maxCreditLimit"`
}

func (dto PostProject) GetClientId() uuid.ClientId {
	return dto.ClientId
}
func (dto PostProject) GetName() types.String {
	return dto.Name
}
func (dto PostProject) GetDescription() types.TextNull {
	return dto.Description
}
func (dto PostProject) GetTagList() types.StringArray {
	return dto.TagList
}
func (dto PostProject) GetMaxCreditLimit() types.DecimalNull {
	return dto.MaxCreditLimit
}

func (dto *PostProject) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ClientId       *uuid.ClientId
		Name           *types.String
		Description    types.TextNull
		TagList        *types.StringArray
		MaxCreditLimit types.DecimalNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostProject", err)
	}
	var errorList validator.ErrorList
	if aux.ClientId == nil {
		errorList = errorList.With(validator.NewError("clientId", "field is required"))
	}
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if aux.TagList == nil {
		errorList = errorList.With(validator.NewError("tagList", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ClientId = *aux.ClientId
	dto.Name = *aux.Name
	dto.Description = aux.Description
	dto.TagList = *aux.TagList
	dto.MaxCreditLimit = aux.MaxCreditLimit

	return nil
}
