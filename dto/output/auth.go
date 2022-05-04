// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type Auth struct {
	AccessToken  uuid.AuthAccessTokenId  `json:"accessToken"`
	TokenType    types.String            `json:"tokenType"`
	ExpiresIn    types.Int               `json:"expiresIn"`
	ExpiresAt    types.DateTime          `json:"expiresAt"`
	RefreshToken uuid.AuthRefreshTokenId `json:"refreshToken"`
	UserId       uuid.AuthUserId         `json:"userId"`
	Author       Author                  `json:"author"`
}
