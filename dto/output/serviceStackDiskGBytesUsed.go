// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ServiceStackDiskGBytesUsed struct {
	DiskGBytesUsed types.Float `json:"diskGBytesUsed"`
	Objects        types.Int64 `json:"objects"`
	RawPolicy      types.Text  `json:"rawPolicy"`
}
