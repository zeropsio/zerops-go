// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ServiceStackLightJsonObject struct {
	Id                        uuid.ServiceStackId                `json:"id"`
	Created                   types.DateTime                     `json:"created"`
	LastUpdate                types.DateTime                     `json:"lastUpdate"`
	ProjectId                 uuid.ProjectId                     `json:"projectId"`
	ServiceStackTypeId        stringId.ServiceStackTypeId        `json:"serviceStackTypeId"`
	ServiceStackTypeVersionId stringId.ServiceStackTypeVersionId `json:"serviceStackTypeVersionId"`
	DriverId                  stringId.DriverIdNull              `json:"driverId"`
	Name                      types.String                       `json:"name"`
	ServiceStackTypeInfo      ServiceStackInfoJsonObject         `json:"serviceStackTypeInfo"`
	Ports                     ServiceStackLightJsonObjectPorts   `json:"ports"`
}

type ServiceStackLightJsonObjectPorts []ServicePort

func (dto ServiceStackLightJsonObjectPorts) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServicePort(dto))
}
