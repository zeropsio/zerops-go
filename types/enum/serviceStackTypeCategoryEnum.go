// Generated ZEROPS sdk

package enum

type ServiceStackTypeCategoryEnum string

const (
	ServiceStackTypeCategoryEnumCore           = ServiceStackTypeCategoryEnum("CORE")
	ServiceStackTypeCategoryEnumHttpL7Balancer = ServiceStackTypeCategoryEnum("HTTP_L7_BALANCER")
	ServiceStackTypeCategoryEnumStandard       = ServiceStackTypeCategoryEnum("STANDARD")
	ServiceStackTypeCategoryEnumUser           = ServiceStackTypeCategoryEnum("USER")
	ServiceStackTypeCategoryEnumSharedStorage  = ServiceStackTypeCategoryEnum("SHARED_STORAGE")
	ServiceStackTypeCategoryEnumObjectStorage  = ServiceStackTypeCategoryEnum("OBJECT_STORAGE")
	ServiceStackTypeCategoryEnumBuild          = ServiceStackTypeCategoryEnum("BUILD")
	ServiceStackTypeCategoryEnumInternal       = ServiceStackTypeCategoryEnum("INTERNAL")
	ServiceStackTypeCategoryEnumPrepareRuntime = ServiceStackTypeCategoryEnum("PREPARE_RUNTIME")
)

func NewServiceStackTypeCategoryEnumFromString(value string) (out ServiceStackTypeCategoryEnum, err error) {
	return ServiceStackTypeCategoryEnum(value), nil
}

func (enum ServiceStackTypeCategoryEnum) String() string {
	return string(enum)
}

func (enum ServiceStackTypeCategoryEnum) Native() string {
	return string(enum)
}

func (enum ServiceStackTypeCategoryEnum) Values() []ServiceStackTypeCategoryEnum {
	return ServiceStackTypeCategoryEnumAll()
}

func (enum ServiceStackTypeCategoryEnum) PublicValues() []ServiceStackTypeCategoryEnum {
	return ServiceStackTypeCategoryEnumAllPublic()
}

func (enum ServiceStackTypeCategoryEnum) PrivateValues() []ServiceStackTypeCategoryEnum {
	return ServiceStackTypeCategoryEnumAllPrivate()
}

func (enum ServiceStackTypeCategoryEnum) DefaultValue() ServiceStackTypeCategoryEnum {
	return ServiceStackTypeCategoryEnumDefault()
}

func (enum ServiceStackTypeCategoryEnum) Is(values ...ServiceStackTypeCategoryEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ServiceStackTypeCategoryEnumAllStrings() []string {
	return []string{
		string(ServiceStackTypeCategoryEnumCore), string(ServiceStackTypeCategoryEnumHttpL7Balancer), string(ServiceStackTypeCategoryEnumStandard), string(ServiceStackTypeCategoryEnumUser), string(ServiceStackTypeCategoryEnumSharedStorage), string(ServiceStackTypeCategoryEnumObjectStorage), string(ServiceStackTypeCategoryEnumBuild), string(ServiceStackTypeCategoryEnumInternal), string(ServiceStackTypeCategoryEnumPrepareRuntime),
	}
}

func ServiceStackTypeCategoryEnumAll() []ServiceStackTypeCategoryEnum {
	return []ServiceStackTypeCategoryEnum{
		ServiceStackTypeCategoryEnumCore, ServiceStackTypeCategoryEnumHttpL7Balancer, ServiceStackTypeCategoryEnumStandard, ServiceStackTypeCategoryEnumUser, ServiceStackTypeCategoryEnumSharedStorage, ServiceStackTypeCategoryEnumObjectStorage, ServiceStackTypeCategoryEnumBuild, ServiceStackTypeCategoryEnumInternal, ServiceStackTypeCategoryEnumPrepareRuntime,
	}
}

func ServiceStackTypeCategoryEnumAllPublic() []ServiceStackTypeCategoryEnum {
	return []ServiceStackTypeCategoryEnum{
		ServiceStackTypeCategoryEnumCore, ServiceStackTypeCategoryEnumHttpL7Balancer, ServiceStackTypeCategoryEnumStandard, ServiceStackTypeCategoryEnumUser, ServiceStackTypeCategoryEnumSharedStorage, ServiceStackTypeCategoryEnumObjectStorage, ServiceStackTypeCategoryEnumBuild, ServiceStackTypeCategoryEnumInternal, ServiceStackTypeCategoryEnumPrepareRuntime,
	}
}

func ServiceStackTypeCategoryEnumAllPrivate() []ServiceStackTypeCategoryEnum {
	return []ServiceStackTypeCategoryEnum{}
}

func ServiceStackTypeCategoryEnumDefault() ServiceStackTypeCategoryEnum {
	return ServiceStackTypeCategoryEnumCore
}

func (enum ServiceStackTypeCategoryEnum) IsCore() bool {
	return enum.Is(ServiceStackTypeCategoryEnumCore)
}

func (enum ServiceStackTypeCategoryEnum) IsHttpL7Balancer() bool {
	return enum.Is(ServiceStackTypeCategoryEnumHttpL7Balancer)
}

func (enum ServiceStackTypeCategoryEnum) IsStandard() bool {
	return enum.Is(ServiceStackTypeCategoryEnumStandard)
}

func (enum ServiceStackTypeCategoryEnum) IsUser() bool {
	return enum.Is(ServiceStackTypeCategoryEnumUser)
}

func (enum ServiceStackTypeCategoryEnum) IsSharedStorage() bool {
	return enum.Is(ServiceStackTypeCategoryEnumSharedStorage)
}

func (enum ServiceStackTypeCategoryEnum) IsObjectStorage() bool {
	return enum.Is(ServiceStackTypeCategoryEnumObjectStorage)
}

func (enum ServiceStackTypeCategoryEnum) IsBuild() bool {
	return enum.Is(ServiceStackTypeCategoryEnumBuild)
}

func (enum ServiceStackTypeCategoryEnum) IsInternal() bool {
	return enum.Is(ServiceStackTypeCategoryEnumInternal)
}

func (enum ServiceStackTypeCategoryEnum) IsPrepareRuntime() bool {
	return enum.Is(ServiceStackTypeCategoryEnumPrepareRuntime)
}
