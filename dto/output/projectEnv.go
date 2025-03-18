// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ProjectEnv struct {
	Id         uuid.EnvId       `json:"id"`
	Created    types.DateTime   `json:"created"`
	LastUpdate types.DateTime   `json:"lastUpdate"`
	ClientId   uuid.ClientId    `json:"clientId"`
	ProjectId  uuid.ProjectId   `json:"projectId"`
	Key        types.String     `json:"key"`
	Content    types.Text       `json:"content"`
	Type       enum.EnvTypeEnum `json:"type"`
}
