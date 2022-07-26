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

type ClientLight struct {
	Id            uuid.ClientId    `json:"id"`
	CompanyName   types.String     `json:"companyName"`
	CompanyNumber types.StringNull `json:"companyNumber"`
	VatNumber     types.StringNull `json:"vatNumber"`
	Avatar        *ClientAvatar    `json:"avatar"`
}
