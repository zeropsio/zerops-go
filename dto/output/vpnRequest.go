// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type VpnRequest struct {
	AccessToken types.String   `json:"accessToken"`
	Expiry      types.DateTime `json:"expiry"`
}
