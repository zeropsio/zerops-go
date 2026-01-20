// Generated ZEROPS sdk

package output

import (
	"strconv"
)

var _ strconv.NumError

type WebAuthnRegistrationVerification struct {
	Credential WebAuthnCredential `json:"credential"`
	Auth       Auth               `json:"auth"`
}
