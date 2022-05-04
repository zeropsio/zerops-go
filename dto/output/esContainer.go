// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsContainer struct {
	Id                      uuid.ContainerId         `json:"id"`
	ClientId                uuid.ClientId            `json:"clientId"`
	ProjectId               uuid.ProjectId           `json:"projectId"`
	ServiceStackId          uuid.ServiceStackId      `json:"serviceStackId"`
	ServiceId               uuid.ServiceId           `json:"serviceId"`
	Status                  enum.ContainerStatusEnum `json:"status"`
	Number                  types.Int                `json:"number"`
	Name                    types.StringNull         `json:"name"`
	Hostname                types.StringNull         `json:"hostname"`
	CurrentHardwareResource HardwareResource         `json:"currentHardwareResource"`
	Created                 types.DateTime           `json:"created"`
	LastUpdate              types.DateTime           `json:"lastUpdate"`
}
