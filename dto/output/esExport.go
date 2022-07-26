// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsExport struct {
	Value uuid.EsExportValue `json:"value"`
}
