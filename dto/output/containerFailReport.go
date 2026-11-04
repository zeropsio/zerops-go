// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type ContainerFailReport struct {
	Failed           types.Bool                               `json:"failed"`
	FailedAt         types.DateTime                           `json:"failedAt"`
	Message          types.StringNull                         `json:"message"`
	LastReportedAt   types.DateTime                           `json:"lastReportedAt"`
	RecoveredAt      types.DateTimeNull                       `json:"recoveredAt"`
	RecoverStartedAt types.DateTimeNull                       `json:"recoverStartedAt"`
	RecoverType      *enum.ContainerFailReportRecoverTypeEnum `json:"recoverType"`
}
