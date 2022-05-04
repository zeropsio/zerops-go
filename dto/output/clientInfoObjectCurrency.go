// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ClientInfoObjectCurrency struct {
	Id                  types.String `json:"id"`
	Symbol              types.String `json:"symbol"`
	DisplaySymbolBefore types.Bool   `json:"displaySymbolBefore"`
	RoundDigits         types.Int    `json:"roundDigits"`
}
