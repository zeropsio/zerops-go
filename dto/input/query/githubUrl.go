// Generated ZEROPS sdk

package query

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/queryDto.go.tmpl

import (
	"github.com/zeropsio/zerops-go/types"
)

type GithubUrl struct {
	RedirectUrl types.String
	Action      types.String
	HaRecipe    types.StringNull
	NonHaRecipe types.StringNull
}
