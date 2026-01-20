// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type TotpRegistration struct {
	SessionId uuid.TotpRegistrationSessionId `json:"sessionId"`
	Qrcode    types.String                   `json:"qrcode"`
	Code      types.String                   `json:"code"`
}
