// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ServiceStackTypeProfileOverrides struct {
	Name        types.String `json:"name"`
	DataType    types.String `json:"dataType"`
	Description types.String `json:"description"`
}
