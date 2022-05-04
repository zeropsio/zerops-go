// Generated ZEROPS sdk

package types

type Object string

func NewObject(value string) Object {
	return Object(value)
}

func (parameter Object) Native() string {
	return string(parameter)
}
