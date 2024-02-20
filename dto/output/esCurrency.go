// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/stringId"
)

var _ strconv.NumError

type EsCurrency struct {
	Id                  stringId.CurrencyId `json:"id"`
	Symbol              types.String        `json:"symbol"`
	DisplaySymbolBefore types.Bool          `json:"displaySymbolBefore"`
	RoundDigits         types.Int           `json:"roundDigits"`
}
