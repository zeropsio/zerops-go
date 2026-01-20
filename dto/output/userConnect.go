// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type UserConnect struct {
	Id    uuid.UserId `json:"id"`
	Email types.Email `json:"email"`
}
