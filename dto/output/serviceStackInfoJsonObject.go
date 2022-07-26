// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type ServiceStackInfoJsonObject struct {
	ServiceStackTypeName        types.String                      `json:"serviceStackTypeName"`
	ServiceStackTypeCategory    enum.ServiceStackTypeCategoryEnum `json:"serviceStackTypeCategory"`
	ServiceStackTypeVersionName types.String                      `json:"serviceStackTypeVersionName"`
}
