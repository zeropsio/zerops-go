// Generated ZEROPS sdk

package query

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/queryDto.go.tmpl

import (
	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

type GitlabUrl struct {
	RedirectUrl types.String
	Action      enum.GitlabUrlActionEnum
	HaRecipe    types.StringNull
	NonHaRecipe types.StringNull
}
