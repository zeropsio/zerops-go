// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type Addon struct {
	ProjectId   uuid.ProjectIdNull `json:"projectId"`
	AddonName   types.String       `json:"addonName"`
	MetricId    types.String       `json:"metricId"`
	UserEnabled types.Bool         `json:"userEnabled"`
	ValidFrom   types.DateTime     `json:"validFrom"`
	ValidTill   types.DateTime     `json:"validTill"`
	UnitPrice   types.Float        `json:"unitPrice"`
}
