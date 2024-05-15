// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type BillingInfoId types.UuidShort

func NewBillingInfoIdFromString(value string) (out BillingInfoId, err error) {
	return BillingInfoId(value), nil
}

func (parameter BillingInfoId) Native() string {
	return string(parameter)
}

func (parameter BillingInfoId) TypedString() types.String {
	return types.String(parameter)
}

type BillingInfoIdNull struct {
	value  BillingInfoId
	filled bool
}

func (parameter BillingInfoId) BillingInfoIdNull() BillingInfoIdNull {
	return NewBillingInfoIdNull(parameter)
}

func NewBillingInfoIdNull(value BillingInfoId) BillingInfoIdNull {
	return BillingInfoIdNull{
		filled: true,
		value:  value,
	}
}

func NewBillingInfoIdNullFromString(value string) BillingInfoIdNull {
	return BillingInfoIdNull{
		filled: true,
		value:  BillingInfoId(value),
	}
}

func (parameter BillingInfoIdNull) Get() (BillingInfoId, bool) {
	return parameter.value, parameter.filled
}

func (parameter BillingInfoIdNull) Filled() bool {
	return parameter.filled
}

func (parameter BillingInfoIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *BillingInfoIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
