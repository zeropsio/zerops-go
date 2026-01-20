// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ClientLight struct {
	Id          uuid.ClientId `json:"id"`
	AccountName types.String  `json:"accountName"`
	Avatar      *ClientAvatar `json:"avatar"`
}
