// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsCurrency struct {
	Id                  types.String `json:"id"`
	Symbol              types.String `json:"symbol"`
	DisplaySymbolBefore types.Bool   `json:"displaySymbolBefore"`
	RoundDigits         types.Int    `json:"roundDigits"`
}
