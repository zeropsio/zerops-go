// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsColumnsDefaultSort struct {
	Name      types.String `json:"name"`
	Ascending types.Bool   `json:"ascending"`
}
