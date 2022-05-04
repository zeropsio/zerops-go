// Generated ZEROPS sdk

package types

type String string

func NewString(value string) String {
	return String(value)
}

func NewStringFromString(value string) (out String, err error) {
	return String(value), nil
}

func (parameter String) StringNull() StringNull {
	return NewStringNull(parameter.Native())
}

func (parameter String) Native() string {
	return string(parameter)
}

func (parameter String) String() string {
	return string(parameter)
}

func (parameter String) Pointer() *string {
	return (*string)(&parameter)
}

func (parameter String) Bytes() []byte {
	return []byte(parameter)
}
