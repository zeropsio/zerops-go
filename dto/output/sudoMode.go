// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type SudoMode struct {
	InSudo    types.Bool         `json:"inSudo"`
	SudoUntil types.DateTimeNull `json:"sudoUntil"`
}
