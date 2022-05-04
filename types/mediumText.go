// Generated ZEROPS sdk

package types

type MediumText string

func NewMediumText(value string) MediumText {
	return MediumText(value)
}

func (parameter MediumText) Native() string {
	return string(parameter)
}
