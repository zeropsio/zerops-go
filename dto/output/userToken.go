// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type UserToken struct {
	Id      uuid.UserTokenId `json:"id"`
	Name    types.String     `json:"name"`
	Created types.DateTime   `json:"created"`
}
