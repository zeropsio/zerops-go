// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsUserData struct {
	Id               uuid.UserDataId       `json:"id"`
	ClientId         uuid.ClientId         `json:"clientId"`
	ProjectId        uuid.ProjectId        `json:"projectId"`
	ServiceStackId   uuid.ServiceStackId   `json:"serviceStackId"`
	Key              types.String          `json:"key"`
	Content          types.Text            `json:"content"`
	Type             enum.UserDataTypeEnum `json:"type"`
	Created          types.DateTime        `json:"created"`
	LastUpdate       types.DateTime        `json:"lastUpdate"`
	ServiceStackName types.String          `json:"serviceStackName"`
}
