// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/stringId"
)

var _ strconv.NumError

type EsCountry struct {
	Id   stringId.CountryId     `json:"id"`
	Name LanguageTextJsonObject `json:"name"`
	InEu types.Bool             `json:"inEu"`
}
