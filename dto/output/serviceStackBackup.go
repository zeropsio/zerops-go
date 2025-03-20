// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ServiceStackBackup struct {
	Success types.Bool `json:"success"`
	Async   types.Bool `json:"async"`
}
