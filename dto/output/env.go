// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type Env struct {
	Id         uuid.EnvId       `json:"id"`
	Key        types.String     `json:"key"`
	Content    types.Text       `json:"content"`
	Type       enum.EnvTypeEnum `json:"type"`
	ProjectId  uuid.ProjectId   `json:"projectId"`
	ClientId   uuid.ClientId    `json:"clientId"`
	Created    types.DateTime   `json:"created"`
	Editable   types.Bool       `json:"editable"`
	Sensitive  types.Bool       `json:"sensitive"`
	LastUpdate types.DateTime   `json:"lastUpdate"`
}
