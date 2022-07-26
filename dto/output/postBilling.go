// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type PostBilling struct {
	ClientSecret  types.String `json:"clientSecret"`
	PaymentId     types.String `json:"paymentId"`
	PaymentStatus types.String `json:"paymentStatus"`
}
