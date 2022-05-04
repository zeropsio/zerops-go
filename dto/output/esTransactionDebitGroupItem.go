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
	Cpu           types.Decimal  `json:"cpu"`
	Ram           types.Decimal  `json:"ram"`
	Disc          types.Decimal  `json:"disc"`
	Ceph_active   types.Decimal  `json:"ceph_active"`
	Ceph_inactive types.Decimal  `json:"ceph_inactive"`
	SumTotalPrice types.Decimal  `json:"sumTotalPrice"`
}
