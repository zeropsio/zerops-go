// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ProjectImport struct {
	ProjectId     uuid.ProjectId             `json:"projectId"`
	ProjectName   types.String               `json:"projectName"`
	ServiceStacks ProjectImportServiceStacks `json:"serviceStacks"`
}

type ProjectImportServiceStacks []ProjectImportServiceStack

func (dto ProjectImportServiceStacks) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProjectImportServiceStack(dto))
}
