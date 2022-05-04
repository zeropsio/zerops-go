// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type VCSRepository struct {
	Name       types.String                     `json:"name"`
	FullName   types.String                     `json:"fullName"`
	Visibility enum.VCSRepositoryVisibilityEnum `json:"visibility"`
	Owner      types.String                     `json:"owner"`
	OwnerUrl   types.String                     `json:"ownerUrl"`
	Admin      types.Bool                       `json:"admin"`
}
