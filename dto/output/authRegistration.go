// Generated ZEROPS sdk

package output

import (
	"strconv"
)

var _ strconv.NumError

type AuthRegistration struct {
	User     *User            `json:"user"`
	Auth     AuthFull         `json:"auth"`
	WebAuthn *WebAuthnOptions `json:"webAuthn"`
}
