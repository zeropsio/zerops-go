// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type UserNotification struct {
	Id              uuid.UserNotificationId       `json:"id"`
	ClientId        uuid.ClientId                 `json:"clientId"`
	ProjectId       uuid.ProjectIdNull            `json:"projectId"`
	ServiceStackId  uuid.ServiceStackIdNull       `json:"serviceStackId"`
	Project         *ProjectLightJsonObject       `json:"project"`
	ServiceStacks   UserNotificationServiceStacks `json:"serviceStacks"`
	ServiceName     types.StringNull              `json:"serviceName"`
	ContainerId     uuid.ContainerIdNull          `json:"containerId"`
	ContainerNumber types.IntNull                 `json:"containerNumber"`
	Type            enum.UserNotificationTypeEnum `json:"type"`
	CreatedByUser   UserJsonObject                `json:"createdByUser"`
	CanceledByUser  *UserJsonObject               `json:"canceledByUser"`
	ActionName      types.String                  `json:"actionName"`
	ActionCreated   types.DateTime                `json:"actionCreated"`
	ActionStarted   types.DateTimeNull            `json:"actionStarted"`
	ActionFinished  types.DateTimeNull            `json:"actionFinished"`
	Error           *ErrorObject                  `json:"error"`
	AppVersion      *AppVersionJsonObject         `json:"appVersion"`
	Created         types.DateTime                `json:"created"`
	ExecutorTag     *enum.ProcessExecutorTagEnum  `json:"executorTag"`
}

type UserNotificationServiceStacks []ServiceStackLightJsonObject

func (dto UserNotificationServiceStacks) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackLightJsonObject(dto))
}
