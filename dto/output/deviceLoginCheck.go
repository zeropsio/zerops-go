// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type DeviceLoginCheck struct {
	Status enum.DeviceLoginStatusEnum `json:"status"`
	Auth   *AuthLogin                 `json:"auth"`
}
