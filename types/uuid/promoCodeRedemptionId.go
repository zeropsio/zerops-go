// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type PromoCodeRedemptionId types.UuidShort

func NewPromoCodeRedemptionIdFromString(value string) (out PromoCodeRedemptionId, err error) {
	return PromoCodeRedemptionId(value), nil
}

func (parameter PromoCodeRedemptionId) Native() string {
	return string(parameter)
}

func (parameter PromoCodeRedemptionId) TypedString() types.String {
	return types.String(parameter)
}

type PromoCodeRedemptionIdNull struct {
	value  PromoCodeRedemptionId
	filled bool
}

func (parameter PromoCodeRedemptionId) PromoCodeRedemptionIdNull() PromoCodeRedemptionIdNull {
	return NewPromoCodeRedemptionIdNull(parameter)
}

func NewPromoCodeRedemptionIdNull(value PromoCodeRedemptionId) PromoCodeRedemptionIdNull {
	return PromoCodeRedemptionIdNull{
		filled: true,
		value:  value,
	}
}

func NewPromoCodeRedemptionIdNullFromString(value string) PromoCodeRedemptionIdNull {
	return PromoCodeRedemptionIdNull{
		filled: true,
		value:  PromoCodeRedemptionId(value),
	}
}

func (parameter PromoCodeRedemptionIdNull) Get() (PromoCodeRedemptionId, bool) {
	return parameter.value, parameter.filled
}

func (parameter PromoCodeRedemptionIdNull) Filled() bool {
	return parameter.filled
}

func (parameter PromoCodeRedemptionIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *PromoCodeRedemptionIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
