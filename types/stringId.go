// Generated ZEROPS sdk

package types

type StringId string

func NewStringId(value string) StringId {
	return StringId(value)
}

func NewStringIdFromString(value string) (out StringId, err error) {
	return StringId(value), nil
}

func (parameter StringId) StringIdNull() StringIdNull {
	return NewStringIdNull(parameter.Native())
}

func (parameter StringId) Native() string {
	return string(parameter)
}
