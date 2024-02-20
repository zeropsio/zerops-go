// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/stringId"
)

var _ strconv.NumError

type EsLanguage struct {
	Id   stringId.LanguageId `json:"id"`
	Name types.String        `json:"name"`
}
