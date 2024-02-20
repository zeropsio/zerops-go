// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/stringId"
)

var _ strconv.NumError

type Country struct {
	Id      stringId.CountryId `json:"id"`
	Name    types.String       `json:"name"`
	InEu    types.Bool         `json:"inEu"`
	VatRate types.Float        `json:"vatRate"`
}
