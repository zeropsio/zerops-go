// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ProjectBackup struct {
	Path                types.String        `json:"path"`
	Name                types.String        `json:"name"`
	Created             types.DateTime      `json:"created"`
	Size                types.Int64         `json:"size"`
	Metadata            types.Map           `json:"metadata"`
	ServiceStackId      uuid.ServiceStackId `json:"serviceStackId"`
	ServiceStackName    types.String        `json:"serviceStackName"`
	ServiceStackDeleted types.Bool          `json:"serviceStackDeleted"`
}
