// Generated ZEROPS sdk

package enum

type PublicPortRoutingPublicIpTypeEnum string

const (
	PublicPortRoutingPublicIpTypeEnumIpV4 = PublicPortRoutingPublicIpTypeEnum("IP_V4")
	PublicPortRoutingPublicIpTypeEnumIpV6 = PublicPortRoutingPublicIpTypeEnum("IP_V6")
)

func NewPublicPortRoutingPublicIpTypeEnumFromString(value string) (out PublicPortRoutingPublicIpTypeEnum, err error) {
	return PublicPortRoutingPublicIpTypeEnum(value), nil
}

func (enum PublicPortRoutingPublicIpTypeEnum) String() string {
	return string(enum)
}

func (enum PublicPortRoutingPublicIpTypeEnum) Native() string {
	return string(enum)
}

func (enum PublicPortRoutingPublicIpTypeEnum) Values() []PublicPortRoutingPublicIpTypeEnum {
	return PublicPortRoutingPublicIpTypeEnumAll()
}

func (enum PublicPortRoutingPublicIpTypeEnum) PublicValues() []PublicPortRoutingPublicIpTypeEnum {
	return PublicPortRoutingPublicIpTypeEnumAllPublic()
}

func (enum PublicPortRoutingPublicIpTypeEnum) PrivateValues() []PublicPortRoutingPublicIpTypeEnum {
	return PublicPortRoutingPublicIpTypeEnumAllPrivate()
}

func (enum PublicPortRoutingPublicIpTypeEnum) DefaultValue() PublicPortRoutingPublicIpTypeEnum {
	return PublicPortRoutingPublicIpTypeEnumDefault()
}

func (enum PublicPortRoutingPublicIpTypeEnum) Is(values ...PublicPortRoutingPublicIpTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func PublicPortRoutingPublicIpTypeEnumAllStrings() []string {
	return []string{
		string(PublicPortRoutingPublicIpTypeEnumIpV4), string(PublicPortRoutingPublicIpTypeEnumIpV6),
	}
}

func PublicPortRoutingPublicIpTypeEnumAll() []PublicPortRoutingPublicIpTypeEnum {
	return []PublicPortRoutingPublicIpTypeEnum{
		PublicPortRoutingPublicIpTypeEnumIpV4, PublicPortRoutingPublicIpTypeEnumIpV6,
	}
}

func PublicPortRoutingPublicIpTypeEnumAllPublic() []PublicPortRoutingPublicIpTypeEnum {
	return []PublicPortRoutingPublicIpTypeEnum{
		PublicPortRoutingPublicIpTypeEnumIpV4, PublicPortRoutingPublicIpTypeEnumIpV6,
	}
}

func PublicPortRoutingPublicIpTypeEnumAllPrivate() []PublicPortRoutingPublicIpTypeEnum {
	return []PublicPortRoutingPublicIpTypeEnum{}
}

func PublicPortRoutingPublicIpTypeEnumDefault() PublicPortRoutingPublicIpTypeEnum {
	return PublicPortRoutingPublicIpTypeEnumIpV4
}

func (enum PublicPortRoutingPublicIpTypeEnum) IsIpV4() bool {
	return enum.Is(PublicPortRoutingPublicIpTypeEnumIpV4)
}

func (enum PublicPortRoutingPublicIpTypeEnum) IsIpV6() bool {
	return enum.Is(PublicPortRoutingPublicIpTypeEnumIpV6)
}
