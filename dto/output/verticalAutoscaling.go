// Generated ZEROPS sdk

package output

import (
	"strconv"
)

var _ strconv.NumError

type VerticalAutoscaling struct {
	MaxResource ScalingResource `json:"maxResource"`
	MinResource ScalingResource `json:"minResource"`
}
