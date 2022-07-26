// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type VpnRequest struct {
	AccessToken types.String   `json:"accessToken"`
	Expiry      types.DateTime `json:"expiry"`
}
