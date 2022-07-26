// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ServiceStackTypeVersionConfig struct {
	DefaultWebserverConfig types.StringNull       `json:"defaultWebserverConfig"`
	VerticalAutoscaling    *VerticalAutoscaling   `json:"verticalAutoscaling"`
	HorizontalAutoscaling  *HorizontalAutoscaling `json:"horizontalAutoscaling"`
}
