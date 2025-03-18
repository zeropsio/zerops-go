// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type ServiceStackDebug struct {
	DebugBuild          enum.ServiceStackDebugDebugEnum `json:"debugBuild"`
	DebugRuntimePrepare enum.ServiceStackDebugDebugEnum `json:"debugRuntimePrepare"`
}
