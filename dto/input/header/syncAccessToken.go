// Generated ZEROPS sdk

package header

import (
	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

type SyncAccessToken struct {
	Authorization uuid.AccessTokenKey
	Sync          types.BoolNull
}
