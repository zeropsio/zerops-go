// Generated ZEROPS sdk

package enum

type ServiceStackConnectionStatusEnum string

const (
	ServiceStackConnectionStatusEnumCreating = ServiceStackConnectionStatusEnum("CREATING")
	ServiceStackConnectionStatusEnumActive   = ServiceStackConnectionStatusEnum("ACTIVE")
	ServiceStackConnectionStatusEnumDeleting = ServiceStackConnectionStatusEnum("DELETING")
)

func NewServiceStackConnectionStatusEnumFromString(value string) (out ServiceStackConnectionStatusEnum, err error) {
	return ServiceStackConnectionStatusEnum(value), nil
}

func (enum ServiceStackConnectionStatusEnum) String() string {
	return string(enum)
}

func (enum ServiceStackConnectionStatusEnum) Native() string {
	return string(enum)
}

func (enum ServiceStackConnectionStatusEnum) Values() []ServiceStackConnectionStatusEnum {
	return ServiceStackConnectionStatusEnumAll()
}

func (enum ServiceStackConnectionStatusEnum) PublicValues() []ServiceStackConnectionStatusEnum {
	return ServiceStackConnectionStatusEnumAllPublic()
}

func (enum ServiceStackConnectionStatusEnum) PrivateValues() []ServiceStackConnectionStatusEnum {
	return ServiceStackConnectionStatusEnumAllPrivate()
}

func (enum ServiceStackConnectionStatusEnum) DefaultValue() ServiceStackConnectionStatusEnum {
	return ServiceStackConnectionStatusEnumDefault()
}

func (enum ServiceStackConnectionStatusEnum) Is(values ...ServiceStackConnectionStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ServiceStackConnectionStatusEnumAllStrings() []string {
	return []string{
		string(ServiceStackConnectionStatusEnumCreating), string(ServiceStackConnectionStatusEnumActive), string(ServiceStackConnectionStatusEnumDeleting),
	}
}

func ServiceStackConnectionStatusEnumAll() []ServiceStackConnectionStatusEnum {
	return []ServiceStackConnectionStatusEnum{
		ServiceStackConnectionStatusEnumCreating, ServiceStackConnectionStatusEnumActive, ServiceStackConnectionStatusEnumDeleting,
	}
}

func ServiceStackConnectionStatusEnumAllPublic() []ServiceStackConnectionStatusEnum {
	return []ServiceStackConnectionStatusEnum{
		ServiceStackConnectionStatusEnumCreating, ServiceStackConnectionStatusEnumActive, ServiceStackConnectionStatusEnumDeleting,
	}
}

func ServiceStackConnectionStatusEnumAllPrivate() []ServiceStackConnectionStatusEnum {
	return []ServiceStackConnectionStatusEnum{}
}

func ServiceStackConnectionStatusEnumDefault() ServiceStackConnectionStatusEnum {
	return ServiceStackConnectionStatusEnumCreating
}

func (enum ServiceStackConnectionStatusEnum) IsCreating() bool {
	return enum.Is(ServiceStackConnectionStatusEnumCreating)
}

func (enum ServiceStackConnectionStatusEnum) IsActive() bool {
	return enum.Is(ServiceStackConnectionStatusEnumActive)
}

func (enum ServiceStackConnectionStatusEnum) IsDeleting() bool {
	return enum.Is(ServiceStackConnectionStatusEnumDeleting)
}
