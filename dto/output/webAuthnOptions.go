// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type WebAuthnOptions struct {
	SessionId uuid.WebAuthnOptionsSessionId `json:"sessionId"`
	Response  types.JsonRawMessage          `json:"response"`
}
