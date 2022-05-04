// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type AppVersionLight struct {
	Id         uuid.AppVersionId         `json:"id"`
	Status     enum.AppVersionStatusEnum `json:"status"`
	Created    types.DateTime            `json:"created"`
	LastUpdate types.DateTime            `json:"lastUpdate"`
}
