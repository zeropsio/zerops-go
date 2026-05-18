// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type StackTypeConfig struct {
	DefaultWebserverConfig types.StringNull                   `json:"defaultWebserverConfig"`
	DefaultAutoscaling     *CustomAutoscaling                 `json:"defaultAutoscaling"`
	VerticalAutoscaling    *VerticalAutoscaling               `json:"verticalAutoscaling"`
	HorizontalAutoscaling  *HorizontalAutoscaling             `json:"horizontalAutoscaling"`
	AutoscalingProfiles    StackTypeConfigAutoscalingProfiles `json:"autoscalingProfiles"`
}

type StackTypeConfigAutoscalingProfiles []ServiceStackTypeProfile

func (dto StackTypeConfigAutoscalingProfiles) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackTypeProfile(dto))
}
