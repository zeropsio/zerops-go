// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type DeviceLoginInitiate struct {
	RequestId       uuid.DeviceLoginRequestId `json:"requestId"`
	PollSecret      types.String              `json:"pollSecret"`
	UserCode        types.String              `json:"userCode"`
	VerificationUrl types.String              `json:"verificationUrl"`
	ExpiresIn       types.Int                 `json:"expiresIn"`
	Interval        types.Int                 `json:"interval"`
}
