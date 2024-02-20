// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type Region struct {
	Name      types.String `json:"name"`
	IsDefault types.Bool   `json:"isDefault"`
	Address   types.String `json:"address"`
}
