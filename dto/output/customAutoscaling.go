// Generated ZEROPS sdk

package output

import (
	"strconv"
)

var _ strconv.NumError

type CustomAutoscaling struct {
	VerticalAutoscaling   *VerticalAutoscalingNullable   `json:"verticalAutoscaling"`
	HorizontalAutoscaling *HorizontalAutoscalingNullable `json:"horizontalAutoscaling"`
}
