// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type DeletePublicPortRouting struct {
	Deleted           types.Bool         `json:"deleted"`
	PublicPortRouting *PublicPortRouting `json:"publicPortRouting"`
}
