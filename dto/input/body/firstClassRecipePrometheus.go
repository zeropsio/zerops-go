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
var _ json.Unmarshaler = (*FirstClassRecipePrometheus)(nil)

type FirstClassRecipePrometheus struct {
	PrometheusProjectId         uuid.ProjectIdNull    `json:"prometheusProjectId"`
	ForwardMetricsFromProjectId uuid.ProjectId        `json:"forwardMetricsFromProjectId"`
	GrafanaDatabaseMode         types.StringNull      `json:"grafanaDatabaseMode"` // Deprecated
	ProjectCorePackage          *enum.ProjectModeEnum `json:"projectCorePackage"`
}

func (dto FirstClassRecipePrometheus) GetPrometheusProjectId() uuid.ProjectIdNull {
	return dto.PrometheusProjectId
}
func (dto FirstClassRecipePrometheus) GetForwardMetricsFromProjectId() uuid.ProjectId {
	return dto.ForwardMetricsFromProjectId
}
func (dto FirstClassRecipePrometheus) GetGrafanaDatabaseMode() types.StringNull {
	return dto.GrafanaDatabaseMode
}
func (dto FirstClassRecipePrometheus) GetProjectCorePackage() *enum.ProjectModeEnum {
	return dto.ProjectCorePackage
}

func (dto *FirstClassRecipePrometheus) UnmarshalJSON(b []byte) error {
	var aux = struct {
		PrometheusProjectId         uuid.ProjectIdNull
		ForwardMetricsFromProjectId *uuid.ProjectId
		GrafanaDatabaseMode         types.StringNull
		ProjectCorePackage          *enum.ProjectModeEnum
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("FirstClassRecipePrometheus", err)
	}
	var errorList validator.ErrorList
	if aux.ForwardMetricsFromProjectId == nil {
		errorList = errorList.With(validator.NewError("forwardMetricsFromProjectId", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.PrometheusProjectId = aux.PrometheusProjectId
	dto.ForwardMetricsFromProjectId = *aux.ForwardMetricsFromProjectId
	dto.GrafanaDatabaseMode = aux.GrafanaDatabaseMode
	dto.ProjectCorePackage = aux.ProjectCorePackage

	return nil
}
