// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type AppVersionUserData struct {
	Key     types.String          `json:"key"`
	Content types.Text            `json:"content"`
	Type    enum.UserDataTypeEnum `json:"type"`
}
