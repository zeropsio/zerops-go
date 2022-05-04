// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type AddressJsonObject struct {
	CompanyName types.String `json:"companyName"`
	Street      types.String `json:"street"`
	City        types.String `json:"city"`
	Postcode    types.String `json:"postcode"`
	CountryId   types.String `json:"countryId"`
}
