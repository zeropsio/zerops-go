// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type PublicHttpRoutingLocationRedirect struct {
	Enabled       types.Bool        `json:"enabled"`
	To            types.EmptyString `json:"to"`
	Code          types.Int         `json:"code"`
	PreservePath  types.Bool        `json:"preservePath"`
	PreserveQuery types.Bool        `json:"preserveQuery"`
}
