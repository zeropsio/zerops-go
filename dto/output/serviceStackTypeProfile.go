// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ServiceStackTypeProfile struct {
	Id                types.String                                                    `json:"id"`
	Name              types.String                                                    `json:"name"`
	Description       types.String                                                    `json:"description"`
	IsDefault         types.Bool                                                      `json:"isDefault"`
	CustomAutoscaling CustomAutoscaling                                               `json:"customAutoscaling"`
	Overrides         types.ObjectMap[types.String, ServiceStackTypeProfileOverrides] `json:"overrides"`
	Order             types.Int                                                       `json:"order"`
}
