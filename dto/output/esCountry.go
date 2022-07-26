// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsCountry struct {
	Id      types.String           `json:"id"`
	Name    LanguageTextJsonObject `json:"name"`
	InEu    types.Bool             `json:"inEu"`
	VatRate types.Float            `json:"vatRate"`
}
