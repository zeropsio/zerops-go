// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ProjectEnvPost)(nil)

type ProjectEnvPost struct {
	Key       types.String `json:"key"`
	Content   types.Text   `json:"content"`
	Sensitive types.Bool   `json:"sensitive"`
}

func (dto ProjectEnvPost) GetKey() types.String {
	return dto.Key
}
func (dto ProjectEnvPost) GetContent() types.Text {
	return dto.Content
}
func (dto ProjectEnvPost) GetSensitive() types.Bool {
	return dto.Sensitive
}

func (dto *ProjectEnvPost) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Key       *types.String
		Content   *types.Text
		Sensitive *types.Bool
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ProjectEnvPost", err)
	}
	var errorList validator.ErrorList
	if aux.Key == nil {
		errorList = errorList.With(validator.NewError("key", "field is required"))
	}
	if aux.Content == nil {
		errorList = errorList.With(validator.NewError("content", "field is required"))
	}
	if aux.Sensitive == nil {
		errorList = errorList.With(validator.NewError("sensitive", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Key = *aux.Key
	dto.Content = *aux.Content
	dto.Sensitive = *aux.Sensitive

	return nil
}
