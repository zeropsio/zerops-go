// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ProjectLightJsonObject struct {
	Id          uuid.ProjectId         `json:"id"`
	ClientId    uuid.ClientId          `json:"clientId"`
	Name        types.String           `json:"name"`
	Description types.TextNull         `json:"description"`
	TagList     types.StringArray      `json:"tagList"`
	Status      enum.ProjectStatusEnum `json:"status"`
	Created     types.DateTime         `json:"created"`
	LastUpdate  types.DateTime         `json:"lastUpdate"`
}
