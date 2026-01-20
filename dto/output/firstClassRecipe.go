// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type FirstClassRecipe struct {
	ProjectId   uuid.ProjectId            `json:"projectId"`
	ProjectName types.String              `json:"projectName"`
	Processes   FirstClassRecipeProcesses `json:"processes"`
}

type FirstClassRecipeProcesses []Process

func (dto FirstClassRecipeProcesses) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Process(dto))
}
