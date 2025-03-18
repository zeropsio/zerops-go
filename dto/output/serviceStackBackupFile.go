// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ServiceStackBackupFile struct {
	Path     types.String `json:"path"`
	Name     types.String `json:"name"`
	Size     types.Int64  `json:"size"`
	Metadata types.Map    `json:"metadata"`
}
