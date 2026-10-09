// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type DeviceLoginInfo struct {
	RequestId   uuid.DeviceLoginRequestId  `json:"requestId"`
	AppName     types.String               `json:"appName"`
	AppType     enum.LoginAppTypeEnum      `json:"appType"`
	UserCode    types.String               `json:"userCode"`
	Kind        enum.DeviceLoginKindEnum   `json:"kind"`
	Status      enum.DeviceLoginStatusEnum `json:"status"`
	ExpiresIn   types.Int                  `json:"expiresIn"`
	IpMatches   types.Bool                 `json:"ipMatches"`
	InitiatorIp types.StringNull           `json:"initiatorIp"`
	CallbackUrl types.TextNull             `json:"callbackUrl"`
}
