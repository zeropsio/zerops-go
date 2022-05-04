// Generated ZEROPS sdk

package types

import (
	"strconv"
)

type Int int

func NewInt(value int) Int {
	return Int(value)
}

func NewIntFromString(value string) (out Int, err error) {
	i, err := strconv.Atoi(value)
	if err != nil {
		return out, err
	}
	out = Int(i)

	return
}

func (parameter Int) IntNull() IntNull {
	return NewIntNull(parameter.Native())
}

func (parameter Int) Native() int {
	return int(parameter)
}

func (parameter Int) Pointer() *int {
	return (*int)(&parameter)
}
