// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type PublicHttpRoutingLocationBasicAuthUser struct {
	Username types.EmptyString `json:"username"`
	Password types.EmptyString `json:"password"`
	Comment  types.EmptyString `json:"comment"`
}
