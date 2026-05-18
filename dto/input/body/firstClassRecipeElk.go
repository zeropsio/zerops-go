// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*FirstClassRecipeElk)(nil)

type FirstClassRecipeElk struct {
	ElkProjectId             uuid.ProjectIdNull    `json:"elkProjectId"`
	ForwardLogsFromProjectId uuid.ProjectIdNull    `json:"forwardLogsFromProjectId"`
	ElasticsearchMode        types.StringNull      `json:"elasticsearchMode"`
	ProjectCorePackage       *enum.ProjectModeEnum `json:"projectCorePackage"`
	IncludeLogstash          types.Bool            `json:"includeLogstash"`
	IncludeApm               types.Bool            `json:"includeApm"`
}

func (dto FirstClassRecipeElk) GetElkProjectId() uuid.ProjectIdNull {
	return dto.ElkProjectId
}
func (dto FirstClassRecipeElk) GetForwardLogsFromProjectId() uuid.ProjectIdNull {
	return dto.ForwardLogsFromProjectId
}
func (dto FirstClassRecipeElk) GetElasticsearchMode() types.StringNull {
	return dto.ElasticsearchMode
}
func (dto FirstClassRecipeElk) GetProjectCorePackage() *enum.ProjectModeEnum {
	return dto.ProjectCorePackage
}
func (dto FirstClassRecipeElk) GetIncludeLogstash() types.Bool {
	return dto.IncludeLogstash
}
func (dto FirstClassRecipeElk) GetIncludeApm() types.Bool {
	return dto.IncludeApm
}

func (dto *FirstClassRecipeElk) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ElkProjectId             uuid.ProjectIdNull
		ForwardLogsFromProjectId uuid.ProjectIdNull
		ElasticsearchMode        types.StringNull
		ProjectCorePackage       *enum.ProjectModeEnum
		IncludeLogstash          *types.Bool
		IncludeApm               *types.Bool
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("FirstClassRecipeElk", err)
	}
	var errorList validator.ErrorList
	if aux.IncludeLogstash == nil {
		errorList = errorList.With(validator.NewError("includeLogstash", "field is required"))
	}
	if aux.IncludeApm == nil {
		errorList = errorList.With(validator.NewError("includeApm", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ElkProjectId = aux.ElkProjectId
	dto.ForwardLogsFromProjectId = aux.ForwardLogsFromProjectId
	dto.ElasticsearchMode = aux.ElasticsearchMode
	dto.ProjectCorePackage = aux.ProjectCorePackage
	dto.IncludeLogstash = *aux.IncludeLogstash
	dto.IncludeApm = *aux.IncludeApm

	return nil
}
