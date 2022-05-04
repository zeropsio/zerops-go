// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ServicePort struct {
	Protocol    enum.ServicePortProtocolEnum `json:"protocol"`
	Port        types.Int                    `json:"port"`
	Description types.EmptyString            `json:"description"`
	Scheme      enum.ServicePortSchemeEnum   `json:"scheme"`
	ServiceId   uuid.ServiceIdNull           `json:"serviceId"`
}
