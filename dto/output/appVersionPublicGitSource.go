// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type AppVersionPublicGitSource struct {
	GitUrl     types.String `json:"gitUrl"`
	BranchName types.String `json:"branchName"`
}
