// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ProjectEnvItem)(nil)

type ProjectEnvItem struct {
	Key     types.String `json:"key"`
	Content types.Text   `json:"content"`
}

func (dto ProjectEnvItem) GetKey() types.String {
	return dto.Key
}
func (dto ProjectEnvItem) GetContent() types.Text {
	return dto.Content
}

func (dto *ProjectEnvItem) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Key     *types.String
		Content *types.Text
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ProjectEnvItem", err)
	}
	var errorList validator.ErrorList
	if aux.Key == nil {
		errorList = errorList.With(validator.NewError("key", "field is required"))
	}
	if aux.Content == nil {
		errorList = errorList.With(validator.NewError("content", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Key = *aux.Key
	dto.Content = *aux.Content

	return nil
}
