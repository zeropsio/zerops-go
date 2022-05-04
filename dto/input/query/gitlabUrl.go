// Generated ZEROPS sdk

package query

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
