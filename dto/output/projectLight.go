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

type ProjectLight struct {
	Id           uuid.ProjectId         `json:"id"`
	ClientId     uuid.ClientId          `json:"clientId"`
	Name         types.String           `json:"name"`
	Mode         enum.ProjectModeEnum   `json:"mode"`
	Description  types.TextNull         `json:"description"`
	TagList      types.StringArray      `json:"tagList"`
	Status       enum.ProjectStatusEnum `json:"status"`
	Created      types.DateTime         `json:"created"`
	LastUpdate   types.DateTime         `json:"lastUpdate"`
	RecipeSource types.TextNull         `json:"recipeSource"`
	UserRoles    ProjectLightUserRoles  `json:"userRoles"`
}

type ProjectLightUserRoles []ProjectUserRole

func (dto ProjectLightUserRoles) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProjectUserRole(dto))
}
