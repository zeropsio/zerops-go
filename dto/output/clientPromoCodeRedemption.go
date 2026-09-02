// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ClientPromoCodeRedemption struct {
	Id                uuid.PromoCodeRedemptionId `json:"id"`
	PromoCodeId       uuid.PromoCodeId           `json:"promoCodeId"`
	Code              types.String               `json:"code"`
	BonusCreditAmount types.Decimal              `json:"bonusCreditAmount"`
	Created           types.DateTime             `json:"created"`
}
