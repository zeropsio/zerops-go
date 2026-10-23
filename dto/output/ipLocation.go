// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type IpLocation struct {
	Country      types.String     `json:"country"`
	CountryCode  types.String     `json:"countryCode"`
	City         types.String     `json:"city"`
	Organization types.StringNull `json:"organization"`
}
