// Generated ZEROPS sdk

package enum

type ServicePortProtocolEnum string

const (
	ServicePortProtocolEnumTcp = ServicePortProtocolEnum("tcp")
	ServicePortProtocolEnumUdp = ServicePortProtocolEnum("udp")
)

func NewServicePortProtocolEnumFromString(value string) (out ServicePortProtocolEnum, err error) {
	return ServicePortProtocolEnum(value), nil
}

func (enum ServicePortProtocolEnum) String() string {
	return string(enum)
}

func (enum ServicePortProtocolEnum) Native() string {
	return string(enum)
}

func (enum ServicePortProtocolEnum) Values() []ServicePortProtocolEnum {
	return ServicePortProtocolEnumAll()
}

func (enum ServicePortProtocolEnum) PublicValues() []ServicePortProtocolEnum {
	return ServicePortProtocolEnumAllPublic()
}

func (enum ServicePortProtocolEnum) PrivateValues() []ServicePortProtocolEnum {
	return ServicePortProtocolEnumAllPrivate()
}

func (enum ServicePortProtocolEnum) DefaultValue() ServicePortProtocolEnum {
	return ServicePortProtocolEnumDefault()
}

func (enum ServicePortProtocolEnum) Is(values ...ServicePortProtocolEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ServicePortProtocolEnumAllStrings() []string {
	return []string{
		string(ServicePortProtocolEnumTcp), string(ServicePortProtocolEnumUdp),
	}
}

func ServicePortProtocolEnumAll() []ServicePortProtocolEnum {
	return []ServicePortProtocolEnum{
		ServicePortProtocolEnumTcp, ServicePortProtocolEnumUdp,
	}
}

func ServicePortProtocolEnumAllPublic() []ServicePortProtocolEnum {
	return []ServicePortProtocolEnum{
		ServicePortProtocolEnumTcp, ServicePortProtocolEnumUdp,
	}
}

func ServicePortProtocolEnumAllPrivate() []ServicePortProtocolEnum {
	return []ServicePortProtocolEnum{}
}

func ServicePortProtocolEnumDefault() ServicePortProtocolEnum {
	return ServicePortProtocolEnumTcp
}

func (enum ServicePortProtocolEnum) IsTcp() bool {
	return enum.Is(ServicePortProtocolEnumTcp)
}

func (enum ServicePortProtocolEnum) IsUdp() bool {
	return enum.Is(ServicePortProtocolEnumUdp)
}
