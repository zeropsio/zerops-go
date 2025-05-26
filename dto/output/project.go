// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type Project struct {
	Id                  uuid.ProjectId         `json:"id"`
	ClientId            uuid.ClientId          `json:"clientId"`
	Name                types.String           `json:"name"`
	Mode                enum.ProjectModeEnum   `json:"mode"`
	Description         types.TextNull         `json:"description"`
	TagList             types.StringArray      `json:"tagList"`
	Status              enum.ProjectStatusEnum `json:"status"`
	Created             types.DateTime         `json:"created"`
	LastUpdate          types.DateTime         `json:"lastUpdate"`
	RecipeSource        types.TextNull         `json:"recipeSource"`
	PublicIpV4          types.StringNull       `json:"publicIpV4"`
	PublicIpV4Shared    types.Bool             `json:"publicIpV4Shared"`
	PublicIpV6          types.StringNull       `json:"publicIpV6"`
	PublicZone          types.String           `json:"publicZone"`
	ZeropsSubdomainHost types.StringNull       `json:"zeropsSubdomainHost"`
	LogForwardingType   types.StringNull       `json:"logForwardingType"`
	AutoStartup         types.Bool             `json:"autoStartup"`
}
