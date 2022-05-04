// Generated ZEROPS sdk

package types

type Ipv4 string

func NewIpv4(value string) Ipv4 {
	return Ipv4(value)
}

func (parameter Ipv4) Native() string {
	return string(parameter)
}
