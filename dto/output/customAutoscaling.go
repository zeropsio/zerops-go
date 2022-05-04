// Generated ZEROPS sdk

package output

import (
	"strconv"
)

var _ strconv.NumError

type CustomAutoscaling struct {
	VerticalAutoscaling   *VerticalAutoscaling   `json:"verticalAutoscaling"`
	HorizontalAutoscaling *HorizontalAutoscaling `json:"horizontalAutoscaling"`
}
