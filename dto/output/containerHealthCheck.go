// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type ContainerHealthCheck struct {
	Status     enum.ContainerHealthCheckStatusEnum `json:"status"`
	Message    types.StringNull                    `json:"message"`
	ChangedAt  types.DateTime                      `json:"changedAt"`
	ReportedAt types.DateTime                      `json:"reportedAt"`
}
