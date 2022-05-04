// Generated ZEROPS sdk

package types

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

var _ json.Unmarshaler

type DecimalNull struct {
	value  Decimal
	filled bool
}

func NewDecimalNull(value decimal.Decimal) DecimalNull {
	return DecimalNull{
		value:  NewDecimal(value),
		filled: true,
	}
}

func (parameter DecimalNull) Get() (Decimal, bool) {
	return parameter.value, parameter.filled
}

func (parameter DecimalNull) Filled() bool {
	return parameter.filled
}

func (parameter DecimalNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *DecimalNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
