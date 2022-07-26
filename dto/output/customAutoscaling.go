// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"
)

var _ strconv.NumError

type CustomAutoscaling struct {
	VerticalAutoscaling   *VerticalAutoscaling   `json:"verticalAutoscaling"`
	HorizontalAutoscaling *HorizontalAutoscaling `json:"horizontalAutoscaling"`
}
