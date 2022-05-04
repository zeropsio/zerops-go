// Generated ZEROPS sdk

package types

import (
	"strconv"
)

type Int64 int64

func NewInt64(value int64) Int64 {
	return Int64(value)
}

func NewInt64FromString(value string) (out Int64, err error) {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return out, err
	}
	out = Int64(i)

	return
}

func (parameter Int64) Int64Null() Int64Null {
	return NewInt64Null(parameter.Native())
}

func (parameter Int64) Native() int64 {
	return int64(parameter)
}
