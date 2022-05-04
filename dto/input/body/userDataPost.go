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
var _ json.Unmarshaler = (*UserDataPost)(nil)

type UserDataPost struct {
	ServiceStackId uuid.ServiceStackId `json:"serviceStackId"`
	Key            types.String        `json:"key"`
	Content        types.Text          `json:"content"`
}

func (dto UserDataPost) GetServiceStackId() uuid.ServiceStackId {
	return dto.ServiceStackId
}
func (dto UserDataPost) GetKey() types.String {
	return dto.Key
}
func (dto UserDataPost) GetContent() types.Text {
	return dto.Content
}

func (dto *UserDataPost) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ServiceStackId *uuid.ServiceStackId
		Key            *types.String
		Content        *types.Text
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("UserDataPost", err)
	}
	var errorList validator.ErrorList
	if aux.ServiceStackId == nil {
		errorList = errorList.With(validator.NewError("serviceStackId", "field is required"))
	}
	if aux.Key == nil {
		errorList = errorList.With(validator.NewError("key", "field is required"))
	}
	if aux.Content == nil {
		errorList = errorList.With(validator.NewError("content", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ServiceStackId = *aux.ServiceStackId
	dto.Key = *aux.Key
	dto.Content = *aux.Content

	return nil
}
