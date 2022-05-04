// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ProjectLog struct {
	AccessToken uuid.ProjectLogAccessToken `json:"accessToken"`
	Expiration  types.DateTime             `json:"expiration"`
	Url         types.String               `json:"url"`
}
