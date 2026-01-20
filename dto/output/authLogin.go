// Generated ZEROPS sdk

package output

import (
	"strconv"
)

var _ strconv.NumError

type AuthLogin struct {
	User *User    `json:"user"`
	Auth AuthFull `json:"auth"`
}
