// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type PublicHttpRoutingLocationRateLimiting struct {
	Enabled  types.Bool                                        `json:"enabled"`
	Key      enum.PublicHttpRoutingLocationRateLimitingKeyEnum `json:"key"`
	Burst    types.Int                                         `json:"burst"`
	Rate     types.Int                                         `json:"rate"`
	Zone     types.EmptyString                                 `json:"zone"`
	ZoneSize types.Int                                         `json:"zoneSize"`
}
