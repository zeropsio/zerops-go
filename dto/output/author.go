// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type Author struct {
	AuthorId   uuid.AuthorAuthorId       `json:"authorId"`
	AuthorType enum.AuthorAuthorTypeEnum `json:"authorType"`
}
