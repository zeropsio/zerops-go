// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type AuthFull struct {
	AccessToken   types.String          `json:"accessToken"`
	TokenType     types.String          `json:"tokenType"`
	ExpiresIn     types.Int             `json:"expiresIn"`
	ExpiresAt     types.DateTime        `json:"expiresAt"`
	TwoFAMethods  types.StringArrayNull `json:"twoFAMethods"`
	TwoFAVerified types.BoolNull        `json:"twoFAVerified"`
	SudoUntil     types.DateTimeNull    `json:"sudoUntil"`
	UserId        uuid.AuthUserId       `json:"userId"`
	Author        Author                `json:"author"`
	RefreshToken  types.StringNull      `json:"refreshToken"`
}
