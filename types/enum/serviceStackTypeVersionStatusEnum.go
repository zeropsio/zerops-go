// Generated ZEROPS sdk

package enum

type ServiceStackTypeVersionStatusEnum string

const (
	ServiceStackTypeVersionStatusEnumActive   = ServiceStackTypeVersionStatusEnum("ACTIVE")
	ServiceStackTypeVersionStatusEnumDisabled = ServiceStackTypeVersionStatusEnum("DISABLED")
)

func NewServiceStackTypeVersionStatusEnumFromString(value string) (out ServiceStackTypeVersionStatusEnum, err error) {
	return ServiceStackTypeVersionStatusEnum(value), nil
}

func (enum ServiceStackTypeVersionStatusEnum) String() string {
	return string(enum)
}

func (enum ServiceStackTypeVersionStatusEnum) Native() string {
	return string(enum)
}

func (enum ServiceStackTypeVersionStatusEnum) Values() []ServiceStackTypeVersionStatusEnum {
	return ServiceStackTypeVersionStatusEnumAll()
}

func (enum ServiceStackTypeVersionStatusEnum) PublicValues() []ServiceStackTypeVersionStatusEnum {
	return ServiceStackTypeVersionStatusEnumAllPublic()
}

func (enum ServiceStackTypeVersionStatusEnum) PrivateValues() []ServiceStackTypeVersionStatusEnum {
	return ServiceStackTypeVersionStatusEnumAllPrivate()
}

func (enum ServiceStackTypeVersionStatusEnum) DefaultValue() ServiceStackTypeVersionStatusEnum {
	return ServiceStackTypeVersionStatusEnumDefault()
}

func (enum ServiceStackTypeVersionStatusEnum) Is(values ...ServiceStackTypeVersionStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ServiceStackTypeVersionStatusEnumAllStrings() []string {
	return []string{
		string(ServiceStackTypeVersionStatusEnumActive), string(ServiceStackTypeVersionStatusEnumDisabled),
	}
}

func ServiceStackTypeVersionStatusEnumAll() []ServiceStackTypeVersionStatusEnum {
	return []ServiceStackTypeVersionStatusEnum{
		ServiceStackTypeVersionStatusEnumActive, ServiceStackTypeVersionStatusEnumDisabled,
	}
}

func ServiceStackTypeVersionStatusEnumAllPublic() []ServiceStackTypeVersionStatusEnum {
	return []ServiceStackTypeVersionStatusEnum{}
}

func ServiceStackTypeVersionStatusEnumAllPrivate() []ServiceStackTypeVersionStatusEnum {
	return []ServiceStackTypeVersionStatusEnum{
		ServiceStackTypeVersionStatusEnumActive, ServiceStackTypeVersionStatusEnumDisabled,
	}
}

func ServiceStackTypeVersionStatusEnumDefault() ServiceStackTypeVersionStatusEnum {
	return ServiceStackTypeVersionStatusEnumActive
}

func (enum ServiceStackTypeVersionStatusEnum) IsActive() bool {
	return enum.Is(ServiceStackTypeVersionStatusEnumActive)
}

func (enum ServiceStackTypeVersionStatusEnum) IsDisabled() bool {
	return enum.Is(ServiceStackTypeVersionStatusEnumDisabled)
}
