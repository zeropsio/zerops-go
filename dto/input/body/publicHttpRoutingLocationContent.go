// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PublicHttpRoutingLocationContent)(nil)

type PublicHttpRoutingLocationContent struct {
	Enabled     types.Bool        `json:"enabled"`
	Code        types.Int         `json:"code"`
	Content     types.Text        `json:"content"`
	ContentType types.EmptyString `json:"contentType"`
}

func (dto PublicHttpRoutingLocationContent) GetEnabled() types.Bool {
	return dto.Enabled
}
func (dto PublicHttpRoutingLocationContent) GetCode() types.Int {
	return dto.Code
}
func (dto PublicHttpRoutingLocationContent) GetContent() types.Text {
	return dto.Content
}
func (dto PublicHttpRoutingLocationContent) GetContentType() types.EmptyString {
	return dto.ContentType
}

func (dto *PublicHttpRoutingLocationContent) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Enabled     *types.Bool
		Code        *types.Int
		Content     *types.Text
		ContentType *types.EmptyString
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicHttpRoutingLocationContent", err)
	}
	var errorList validator.ErrorList
	if aux.Enabled == nil {
		errorList = errorList.With(validator.NewError("enabled", "field is required"))
	}
	if aux.Code == nil {
		errorList = errorList.With(validator.NewError("code", "field is required"))
	}
	if aux.Content == nil {
		errorList = errorList.With(validator.NewError("content", "field is required"))
	}
	if aux.ContentType == nil {
		errorList = errorList.With(validator.NewError("contentType", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Enabled = *aux.Enabled
	dto.Code = *aux.Code
	dto.Content = *aux.Content
	dto.ContentType = *aux.ContentType

	return nil
}
