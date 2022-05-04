// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsColumnsItem struct {
	Name            types.String      `json:"name"`
	OutputName      types.String      `json:"outputName"`
	Type            types.String      `json:"type"`
	SubType         types.String      `json:"subType"`
	Suggest         types.Bool        `json:"suggest"`
	DefaultView     types.Bool        `json:"defaultView"`
	Operators       types.StringArray `json:"operators"`
	DefaultOperator types.String      `json:"defaultOperator"`
	Reference       types.String      `json:"reference"`
	ReferenceColumn types.String      `json:"referenceColumn"`
	TypeData        types.String      `json:"typeData"`
}
