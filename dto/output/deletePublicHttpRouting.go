// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type DeletePublicHttpRouting struct {
	Deleted           types.Bool         `json:"deleted"`
	PublicHttpRouting *PublicHttpRouting `json:"publicHttpRouting"`
}
