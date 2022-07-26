// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type VCSRepositoryBranch struct {
	Name      types.String `json:"name"`
	IsDefault types.Bool   `json:"isDefault"`
}
