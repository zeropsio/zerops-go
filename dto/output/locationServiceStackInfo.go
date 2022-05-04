// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type LocationServiceStackInfo struct {
	ServiceStackName            types.String `json:"serviceStackName"`
	ServiceStackTypeName        types.String `json:"serviceStackTypeName"`
	ServiceStackTypeVersionName types.String `json:"serviceStackTypeVersionName"`
}
