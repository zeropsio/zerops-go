// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type UserLight struct {
	Id        uuid.UserId  `json:"id"`
	Email     types.Email  `json:"email"`
	FullName  types.String `json:"fullName"`
	FirstName types.String `json:"firstName"`
	LastName  types.String `json:"lastName"`
	Avatar    *UserAvatar  `json:"avatar"`
}
