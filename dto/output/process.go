// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type Process struct {
	Id              uuid.ProcessId          `json:"id"`
	ClientId        uuid.ClientId           `json:"clientId"`
	ProjectId       uuid.ProjectId          `json:"projectId"`
	ServiceStackId  uuid.ServiceStackIdNull `json:"serviceStackId"`
	Project         ProjectLightJsonObject  `json:"project"`
	ServiceStacks   ProcessServiceStacks    `json:"serviceStacks"`
	Status          enum.ProcessStatusEnum  `json:"status"`
	Sequence        types.Int               `json:"sequence"`
	CreatedByUser   UserJsonObject          `json:"createdByUser"`
	CanceledByUser  *UserJsonObject         `json:"canceledByUser"`
	ActionName      types.String            `json:"actionName"`
	Created         types.DateTime          `json:"created"`
	LastUpdate      types.DateTime          `json:"lastUpdate"`
	Started         types.DateTimeNull      `json:"started"`
	Finished        types.DateTimeNull      `json:"finished"`
	CreatedBySystem types.Bool              `json:"createdBySystem"`
	AppVersion      *AppVersionJsonObject   `json:"appVersion"`
}

type ProcessServiceStacks []ServiceStackLightJsonObject

func (dto ProcessServiceStacks) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackLightJsonObject(dto))
}
