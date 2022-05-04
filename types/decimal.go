// Generated ZEROPS sdk

package types

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/shopspring/decimal"
)

type Decimal decimal.Decimal

func NewDecimal(value decimal.Decimal) Decimal {
	return Decimal(value)
}

func NewDecimalFromString(value string) (Decimal, error) {
	d, err := decimal.NewFromString(value)
	if err != nil {
		return Decimal{}, err
	}
	return Decimal(d), nil
}

func NewDecimalFromFloat(value float64) Decimal {
	return Decimal(decimal.NewFromFloat(value))
}

func NewDecimalFromInt(value int) Decimal {
	return Decimal(decimal.NewFromInt32(int32(value)))
}

func (parameter Decimal) Native() decimal.Decimal {
	return decimal.Decimal(parameter)
}

func (parameter Decimal) DecimalNull() DecimalNull {
	return NewDecimalNull(parameter.Native())
}

func (parameter *Decimal) UnmarshalJSON(data []byte) error {
	var value *float64
	err := json.NewDecoder(bytes.NewReader(data)).Decode(&value)
	if err != nil {
		return errors.New("value is not valid decimal format, " + err.Error())
	}
	if value != nil {
		*parameter = NewDecimalFromFloat(*value)
	}
	return nil
}

func (parameter Decimal) MarshalJSON() ([]byte, error) {
	v, _ := parameter.Native().Float64()
	return json.Marshal(v)
}

// decimal funcs

func (parameter Decimal) Float() Float {
	f, _ := parameter.Native().Float64()
	return NewFloat(f)
}

func (parameter Decimal) Add(d2 Decimal) Decimal {
	return NewDecimal(parameter.Native().Add(d2.Native()))
}

func (parameter Decimal) Sub(d2 Decimal) Decimal {
	return NewDecimal(parameter.Native().Sub(d2.Native()))
}

func (parameter Decimal) Mul(d2 Decimal) Decimal {
	return NewDecimal(parameter.Native().Mul(d2.Native()))
}

func (parameter Decimal) Div(d2 Decimal) Decimal {
	return NewDecimal(parameter.Native().Div(d2.Native()))
}
