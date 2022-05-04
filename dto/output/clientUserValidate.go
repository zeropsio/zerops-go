// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ClientUserValidate struct {
	Id       uuid.UserId  `json:"id"`
	FullName types.String `json:"fullName"`
}
