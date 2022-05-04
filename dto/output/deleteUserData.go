// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type DeleteUserData struct {
	Deleted  types.Bool `json:"deleted"`
	UserData *UserData  `json:"userData"`
}
