// Generated ZEROPS sdk

package types

type Text string

func NewText(value string) Text {
	return Text(value)
}

func (parameter Text) Native() string {
	return string(parameter)
}
