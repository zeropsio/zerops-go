// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ContainerOsRelease struct {
	Id      types.StringNull `json:"id"`
	Version types.StringNull `json:"version"`
	Name    types.StringNull `json:"name"`
}
