// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"
)

var _ strconv.NumError

type VerticalAutoscaling struct {
	MaxResource ScalingResource `json:"maxResource"`
	MinResource ScalingResource `json:"minResource"`
}
