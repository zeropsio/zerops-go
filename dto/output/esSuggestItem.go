// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsSuggestItem struct {
	Id     uuid.EsSuggestItemId `json:"id"`
	Text   types.String         `json:"text"`
	Column types.String         `json:"column"`
	Score  types.Int            `json:"score"`
}
