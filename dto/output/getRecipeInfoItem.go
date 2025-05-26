// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type GetRecipeInfoItem struct {
	Name    types.String `json:"name"`
	Url     types.String `json:"url"`
	Content types.Text   `json:"content"`
}
