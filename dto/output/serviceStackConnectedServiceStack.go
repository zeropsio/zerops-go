// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type ServiceStackConnectedServiceStack struct {
	ServiceStack ServiceStackLight                     `json:"serviceStack"`
	Status       enum.ServiceStackConnectionStatusEnum `json:"status"`
}
