// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type TwoFARecoveryCodes struct {
	Codes types.StringArray `json:"codes"`
}
