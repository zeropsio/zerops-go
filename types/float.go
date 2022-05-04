// Generated ZEROPS sdk

package types

type Float float64

func NewFloat(value float64) Float {
	return Float(value)
}

func (parameter Float) Native() float64 {
	return float64(parameter)
}

func (parameter Float) Pointer() *float64 {
	return (*float64)(&parameter)
}

func (parameter Float) Decimal() Decimal {
	return NewDecimalFromFloat(parameter.Native())
}

func (parameter Float) FloatNull() FloatNull {
	return NewFloatNull(parameter.Native())
}
