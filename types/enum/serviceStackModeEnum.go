// Generated ZEROPS sdk

package enum

type ServiceStackModeEnum string

const (
	ServiceStackModeEnumHa    = ServiceStackModeEnum("HA")
	ServiceStackModeEnumNonHa = ServiceStackModeEnum("NON_HA")
)

func NewServiceStackModeEnumFromString(value string) (out ServiceStackModeEnum, err error) {
	return ServiceStackModeEnum(value), nil
}

func (enum ServiceStackModeEnum) String() string {
	return string(enum)
}

func (enum ServiceStackModeEnum) Native() string {
	return string(enum)
}

func (enum ServiceStackModeEnum) Values() []ServiceStackModeEnum {
	return ServiceStackModeEnumAll()
}

func (enum ServiceStackModeEnum) PublicValues() []ServiceStackModeEnum {
	return ServiceStackModeEnumAllPublic()
}

func (enum ServiceStackModeEnum) PrivateValues() []ServiceStackModeEnum {
	return ServiceStackModeEnumAllPrivate()
}

func (enum ServiceStackModeEnum) DefaultValue() ServiceStackModeEnum {
	return ServiceStackModeEnumDefault()
}

func (enum ServiceStackModeEnum) Is(values ...ServiceStackModeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ServiceStackModeEnumAllStrings() []string {
	return []string{
		string(ServiceStackModeEnumHa), string(ServiceStackModeEnumNonHa),
	}
}

func ServiceStackModeEnumAll() []ServiceStackModeEnum {
	return []ServiceStackModeEnum{
		ServiceStackModeEnumHa, ServiceStackModeEnumNonHa,
	}
}

func ServiceStackModeEnumAllPublic() []ServiceStackModeEnum {
	return []ServiceStackModeEnum{
		ServiceStackModeEnumHa, ServiceStackModeEnumNonHa,
	}
}

func ServiceStackModeEnumAllPrivate() []ServiceStackModeEnum {
	return []ServiceStackModeEnum{}
}

func ServiceStackModeEnumDefault() ServiceStackModeEnum {
	return ServiceStackModeEnumHa
}

func (enum ServiceStackModeEnum) IsHa() bool {
	return enum.Is(ServiceStackModeEnumHa)
}

func (enum ServiceStackModeEnum) IsNonHa() bool {
	return enum.Is(ServiceStackModeEnumNonHa)
}
