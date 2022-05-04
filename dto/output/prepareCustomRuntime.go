// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type PrepareCustomRuntime struct {
	ContainerCreationStart types.DateTimeNull      `json:"containerCreationStart"`
	StartDate              types.DateTimeNull      `json:"startDate"`
	EndDate                types.DateTimeNull      `json:"endDate"`
	ServiceStackId         uuid.ServiceStackIdNull `json:"serviceStackId"`
	ServiceStackName       types.StringNull        `json:"serviceStackName"`
}
