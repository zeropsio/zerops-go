// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type TwoFAAuthLogin struct {
	User             *User            `json:"user"`
	Auth             AuthFull         `json:"auth"`
	NewRecoveryToken types.StringNull `json:"newRecoveryToken"`
}
