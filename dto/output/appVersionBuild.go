// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type AppVersionBuild struct {
	ServiceStackId            uuid.ServiceStackIdNull                `json:"serviceStackId"`
	ServiceStackName          types.StringNull                       `json:"serviceStackName"`
	ServiceStackTypeVersionId stringId.ServiceStackTypeVersionIdNull `json:"serviceStackTypeVersionId"`
	PipelineStart             types.DateTimeNull                     `json:"pipelineStart"`
	PipelineFinish            types.DateTimeNull                     `json:"pipelineFinish"`
	PipelineFailed            types.DateTimeNull                     `json:"pipelineFailed"`
	ContainerCreationStart    types.DateTimeNull                     `json:"containerCreationStart"`
	StartDate                 types.DateTimeNull                     `json:"startDate"`
	EndDate                   types.DateTimeNull                     `json:"endDate"`
	CacheUsed                 types.Bool                             `json:"cacheUsed"`
	HasCurrentCache           types.Bool                             `json:"hasCurrentCache"`
}
