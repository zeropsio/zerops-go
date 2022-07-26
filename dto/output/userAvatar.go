// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type UserAvatar struct {
	LargeAvatarUrl    types.StringNull `json:"largeAvatarUrl"`
	SmallAvatarUrl    types.StringNull `json:"smallAvatarUrl"`
	ExternalAvatarUrl types.StringNull `json:"externalAvatarUrl"`
}
