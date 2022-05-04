// Generated ZEROPS sdk

package types

type Ipv6 string

func NewIpv6(value string) Ipv6 {
	return Ipv6(value)
}

func (parameter Ipv6) Native() string {
	return string(parameter)
}
