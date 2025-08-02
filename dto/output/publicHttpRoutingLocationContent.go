// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type PublicHttpRoutingLocationContent struct {
	Enabled     types.Bool        `json:"enabled"`
	Code        types.Int         `json:"code"`
	Content     types.Text        `json:"content"`
	ContentType types.EmptyString `json:"contentType"`
}
