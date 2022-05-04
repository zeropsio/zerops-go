// Generated ZEROPS sdk

package types

type File string

func NewFile(value string) File {
	return File(value)
}

func (parameter File) Native() string {
	return string(parameter)
}
