// Generated ZEROPS sdk

package types

type UuidShort string

func NewUuidShort(value string) UuidShort {
	return UuidShort(value)
}

func NewUuidShortFromString(value string) (out UuidShort, err error) {
	return UuidShort(value), nil
}

func (parameter UuidShort) UuidShortNull() UuidShortNull {
	return NewUuidShortNull(parameter.Native())
}

func (parameter UuidShort) Native() string {
	return string(parameter)
}
