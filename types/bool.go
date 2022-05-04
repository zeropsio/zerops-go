// Generated ZEROPS sdk

package types

import (
	"strconv"
)

type Bool bool

func NewBool(value bool) Bool {
	return Bool(value)
}

func NewBoolFromString(value string) (out Bool, err error) {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return out, err
	}
	return Bool(b), nil
}

func (parameter Bool) Native() bool {
	return bool(parameter)
}
