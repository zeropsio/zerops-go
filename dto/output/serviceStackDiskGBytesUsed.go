// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ServiceStackDiskGBytesUsed struct {
	DiskGBytesUsed types.Float `json:"diskGBytesUsed"`
	RawPolicy      types.Text  `json:"rawPolicy"`
}
