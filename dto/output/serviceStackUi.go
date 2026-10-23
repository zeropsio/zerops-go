// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ServiceStackUi struct {
	AccessToken uuid.ServiceStackUiAccessToken `json:"accessToken"`
	Expiration  types.DateTime                 `json:"expiration"`
	UrlUi       types.String                   `json:"urlUi"`
	UrlApi      types.String                   `json:"urlApi"`
}
