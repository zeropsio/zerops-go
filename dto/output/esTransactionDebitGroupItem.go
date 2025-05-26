// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsTransactionDebitGroupItem struct {
	From          types.DateTime `json:"from"`
	Till          types.DateTime `json:"till"`
	ProjectId     types.String   `json:"projectId"`
	StackId       types.String   `json:"stackId"`
	ProjectName   types.String   `json:"projectName"`
	StackName     types.String   `json:"stackName"`
	VCpu          types.Float    `json:"vCpu"`
	Cpu           types.Float    `json:"cpu"`
	Ram           types.Float    `json:"ram"`
	Disc          types.Float    `json:"disc"`
	CephActive    types.Float    `json:"cephActive"`
	CephInactive  types.Float    `json:"cephInactive"`
	SumTotalPrice types.Float    `json:"sumTotalPrice"`
}
