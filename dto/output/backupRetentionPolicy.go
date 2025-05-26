// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type BackupRetentionPolicy struct {
	MaxTotalFiles types.Int         `json:"maxTotalFiles"`
	MaxTotalGiB   types.Int         `json:"maxTotalGiB"`
	ProtectedTags types.StringArray `json:"protectedTags"`
	MinDaily      types.Int         `json:"minDaily"`
	MaxDaily      types.Int         `json:"maxDaily"`
	MinWeekly     types.Int         `json:"minWeekly"`
	MaxWeekly     types.Int         `json:"maxWeekly"`
	MinMonthly    types.Int         `json:"minMonthly"`
	MaxMonthly    types.Int         `json:"maxMonthly"`
}
