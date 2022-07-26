// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type UserJsonObject struct {
	Type      enum.UserJsonObjectTypeEnum `json:"type"`
	Id        uuid.UserIdNull             `json:"id"`
	Email     types.EmailNull             `json:"email"`
	FirstName types.StringNull            `json:"firstName"`
	FullName  types.StringNull            `json:"fullName"`
	Avatar    *UserAvatar                 `json:"avatar"`
}
