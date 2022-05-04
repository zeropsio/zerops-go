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

type EsUserNotification struct {
	Id             uuid.UserNotificationId         `json:"id"`
	ClientId       uuid.ClientIdNull               `json:"clientId"`
	Project        *ProjectLightJsonObject         `json:"project"`
	ServiceStacks  EsUserNotificationServiceStacks `json:"serviceStacks"`
	Type           enum.UserNotificationTypeEnum   `json:"type"`
	CreatedByUser  UserJsonObject                  `json:"createdByUser"`
	CanceledByUser *UserJsonObject                 `json:"canceledByUser"`
	ActionName     types.String                    `json:"actionName"`
	ActionCreated  types.DateTime                  `json:"actionCreated"`
	ActionFinished types.DateTimeNull              `json:"actionFinished"`
	Error          *ErrorObject                    `json:"error"`
	UserId         uuid.UserId                     `json:"userId"`
	Acknowledged   types.Bool                      `json:"acknowledged"`
	AppVersion     *AppVersionJsonObject           `json:"appVersion"`
}

type EsUserNotificationServiceStacks []ServiceStackLightJsonObject

func (dto EsUserNotificationServiceStacks) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackLightJsonObject(dto))
}
