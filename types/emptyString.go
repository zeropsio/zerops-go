// Generated ZEROPS sdk

package types

type EmptyString string

func NewEmptyString(value string) EmptyString {
	return EmptyString(value)
}

func NewEmptyStringFromString(value string) (out EmptyString, err error) {
	return EmptyString(value), nil
}

func (parameter EmptyString) EmptyStringNull() EmptyStringNull {
	return NewEmptyStringNull(parameter.Native())
}

func (parameter EmptyString) Native() string {
	return string(parameter)
}
