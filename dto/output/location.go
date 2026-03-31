// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/stringId"
)

var _ strconv.NumError

type Location struct {
	Id      stringId.LocationId `json:"id"`
	Name    types.String        `json:"name"`
	PingUrl types.String        `json:"pingUrl"`
}
