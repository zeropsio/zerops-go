// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ClientLight struct {
	Id          uuid.ClientId         `json:"id"`
	Created     types.DateTime        `json:"created"`
	LastUpdate  types.DateTime        `json:"lastUpdate"`
	AccountName types.String          `json:"accountName"`
	Status      enum.ClientStatusEnum `json:"status"`
	Avatar      *ClientAvatar         `json:"avatar"`
}
