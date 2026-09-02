// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type PromoCodeId types.UuidShort

func NewPromoCodeIdFromString(value string) (out PromoCodeId, err error) {
	return PromoCodeId(value), nil
}

func (parameter PromoCodeId) Native() string {
	return string(parameter)
}

func (parameter PromoCodeId) TypedString() types.String {
	return types.String(parameter)
}

type PromoCodeIdNull struct {
	value  PromoCodeId
	filled bool
}

func (parameter PromoCodeId) PromoCodeIdNull() PromoCodeIdNull {
	return NewPromoCodeIdNull(parameter)
}

func NewPromoCodeIdNull(value PromoCodeId) PromoCodeIdNull {
	return PromoCodeIdNull{
		filled: true,
		value:  value,
	}
}

func NewPromoCodeIdNullFromString(value string) PromoCodeIdNull {
	return PromoCodeIdNull{
		filled: true,
		value:  PromoCodeId(value),
	}
}

func (parameter PromoCodeIdNull) Get() (PromoCodeId, bool) {
	return parameter.value, parameter.filled
}

func (parameter PromoCodeIdNull) Filled() bool {
	return parameter.filled
}

func (parameter PromoCodeIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *PromoCodeIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
