// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type UserDataLight struct {
	Id             uuid.UserDataId          `json:"id"`
	ClientId       uuid.ClientId            `json:"clientId"`
	ProjectId      uuid.ProjectId           `json:"projectId"`
	ServiceStackId uuid.ServiceStackId      `json:"serviceStackId"`
	Key            types.String             `json:"key"`
	Content        types.Text               `json:"content"`
	Type           enum.UserDataTypeEnum    `json:"type"`
	Created        types.DateTime           `json:"created"`
	LastUpdate     types.DateTime           `json:"lastUpdate"`
	IsSynced       types.Bool               `json:"isSynced"`
	DeleteOnSync   types.Bool               `json:"deleteOnSync"`
	Version        enum.UserDataVersionEnum `json:"version"`
	LastSync       types.DateTimeNull       `json:"lastSync"`
}
