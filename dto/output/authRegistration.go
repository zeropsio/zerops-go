// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type AuthRegistration struct {
	User       *User            `json:"user"`
	Auth       AuthFull         `json:"auth"`
	WebAuthn   *WebAuthnOptions `json:"webAuthn"`
	ZcpClaimed types.Bool       `json:"zcpClaimed"`
}
