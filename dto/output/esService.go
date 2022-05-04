// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsService struct {
	Id               uuid.ServiceId         `json:"id"`
	ClientId         uuid.ClientId          `json:"clientId"`
	ProjectId        uuid.ProjectId         `json:"projectId"`
	ServiceStackId   uuid.ServiceStackId    `json:"serviceStackId"`
	Name             types.String           `json:"name"`
	Status           enum.ServiceStatusEnum `json:"status"`
	StackAccessPoint types.Bool             `json:"stackAccessPoint"`
	Created          types.DateTime         `json:"created"`
	LastUpdate       types.DateTime         `json:"lastUpdate"`
}
