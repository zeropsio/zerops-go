// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ErrorObject struct {
	Code    types.String         `json:"code"`
	Message types.String         `json:"message"`
	Meta    types.JsonRawMessage `json:"meta"`
}
