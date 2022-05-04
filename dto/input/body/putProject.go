// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutProject)(nil)

type PutProject struct {
	Name           types.String      `json:"name"`
	Description    types.TextNull    `json:"description"`
	TagList        types.StringArray `json:"tagList"`
	MaxCreditLimit types.DecimalNull `json:"maxCreditLimit"`
}

func (dto PutProject) GetName() types.String {
	return dto.Name
}
func (dto PutProject) GetDescription() types.TextNull {
	return dto.Description
}
func (dto PutProject) GetTagList() types.StringArray {
	return dto.TagList
}
func (dto PutProject) GetMaxCreditLimit() types.DecimalNull {
	return dto.MaxCreditLimit
}

func (dto *PutProject) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Name           *types.String
		Description    types.TextNull
		TagList        *types.StringArray
		MaxCreditLimit types.DecimalNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutProject", err)
	}
	var errorList validator.ErrorList
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if aux.TagList == nil {
		errorList = errorList.With(validator.NewError("tagList", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Name = *aux.Name
	dto.Description = aux.Description
	dto.TagList = *aux.TagList
	dto.MaxCreditLimit = aux.MaxCreditLimit

	return nil
}
