// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ContainerTechnologyVersion struct {
	Name    types.String     `json:"name"`
	Version types.StringNull `json:"version"`
}
