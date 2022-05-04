// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type DeletePublicHttpRouting struct {
	Deleted           types.Bool         `json:"deleted"`
	PublicHttpRouting *PublicHttpRouting `json:"publicHttpRouting"`
}
