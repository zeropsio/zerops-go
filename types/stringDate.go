// Generated ZEROPS sdk

package types

type StringDate string

func NewStringDate(value string) StringDate {
	return StringDate(value)
}

func NewStringDateFromString(value string) (out StringDate, err error) {
	return StringDate(value), nil
}

func (parameter StringDate) Native() string {
	return string(parameter)
}
