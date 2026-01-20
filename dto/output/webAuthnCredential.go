// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type WebAuthnCredential struct {
	Id         types.String   `json:"id"`
	Name       types.String   `json:"name"`
	Created    types.DateTime `json:"created"`
	LastUpdate types.DateTime `json:"lastUpdate"`
	IsResident types.Bool     `json:"isResident"`
}
