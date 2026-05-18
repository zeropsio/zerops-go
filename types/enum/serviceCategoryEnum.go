// Generated ZEROPS sdk

package enum

type ServiceCategoryEnum string

const (
	ServiceCategoryEnumEmpty               = ServiceCategoryEnum("EMPTY")
	ServiceCategoryEnumMaster              = ServiceCategoryEnum("MASTER")
	ServiceCategoryEnumLogger              = ServiceCategoryEnum("LOGGER")
	ServiceCategoryEnumHttpL7Balancer      = ServiceCategoryEnum("HTTP_L7_BALANCER")
	ServiceCategoryEnumApplicationBalancer = ServiceCategoryEnum("APPLICATION_BALANCER")
	ServiceCategoryEnumDatabase            = ServiceCategoryEnum("DATABASE")
	ServiceCategoryEnumMessageBroker       = ServiceCategoryEnum("MESSAGE_BROKER")
	ServiceCategoryEnumUserApplication     = ServiceCategoryEnum("USER_APPLICATION")
	ServiceCategoryEnumSharedStorage       = ServiceCategoryEnum("SHARED_STORAGE")
	ServiceCategoryEnumSystem              = ServiceCategoryEnum("SYSTEM")
	ServiceCategoryEnumBuild               = ServiceCategoryEnum("BUILD")
	ServiceCategoryEnumPrepareRuntime      = ServiceCategoryEnum("PREPARE_RUNTIME")
	ServiceCategoryEnumMetric              = ServiceCategoryEnum("METRIC")
)

func NewServiceCategoryEnumFromString(value string) (out ServiceCategoryEnum, err error) {
	return ServiceCategoryEnum(value), nil
}

func (enum ServiceCategoryEnum) String() string {
	return string(enum)
}

func (enum ServiceCategoryEnum) Native() string {
	return string(enum)
}

func (enum ServiceCategoryEnum) Values() []ServiceCategoryEnum {
	return ServiceCategoryEnumAll()
}

func (enum ServiceCategoryEnum) PublicValues() []ServiceCategoryEnum {
	return ServiceCategoryEnumAllPublic()
}

func (enum ServiceCategoryEnum) PrivateValues() []ServiceCategoryEnum {
	return ServiceCategoryEnumAllPrivate()
}

func (enum ServiceCategoryEnum) DefaultValue() ServiceCategoryEnum {
	return ServiceCategoryEnumDefault()
}

func (enum ServiceCategoryEnum) Is(values ...ServiceCategoryEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ServiceCategoryEnumAllStrings() []string {
	return []string{
		string(ServiceCategoryEnumEmpty), string(ServiceCategoryEnumMaster), string(ServiceCategoryEnumLogger), string(ServiceCategoryEnumHttpL7Balancer), string(ServiceCategoryEnumApplicationBalancer), string(ServiceCategoryEnumDatabase), string(ServiceCategoryEnumMessageBroker), string(ServiceCategoryEnumUserApplication), string(ServiceCategoryEnumSharedStorage), string(ServiceCategoryEnumSystem), string(ServiceCategoryEnumBuild), string(ServiceCategoryEnumPrepareRuntime), string(ServiceCategoryEnumMetric),
	}
}

func ServiceCategoryEnumAll() []ServiceCategoryEnum {
	return []ServiceCategoryEnum{
		ServiceCategoryEnumEmpty, ServiceCategoryEnumMaster, ServiceCategoryEnumLogger, ServiceCategoryEnumHttpL7Balancer, ServiceCategoryEnumApplicationBalancer, ServiceCategoryEnumDatabase, ServiceCategoryEnumMessageBroker, ServiceCategoryEnumUserApplication, ServiceCategoryEnumSharedStorage, ServiceCategoryEnumSystem, ServiceCategoryEnumBuild, ServiceCategoryEnumPrepareRuntime, ServiceCategoryEnumMetric,
	}
}

func ServiceCategoryEnumAllPublic() []ServiceCategoryEnum {
	return []ServiceCategoryEnum{
		ServiceCategoryEnumEmpty, ServiceCategoryEnumMaster, ServiceCategoryEnumLogger, ServiceCategoryEnumHttpL7Balancer, ServiceCategoryEnumApplicationBalancer, ServiceCategoryEnumDatabase, ServiceCategoryEnumMessageBroker, ServiceCategoryEnumUserApplication, ServiceCategoryEnumSharedStorage, ServiceCategoryEnumSystem, ServiceCategoryEnumBuild, ServiceCategoryEnumPrepareRuntime, ServiceCategoryEnumMetric,
	}
}

func ServiceCategoryEnumAllPrivate() []ServiceCategoryEnum {
	return []ServiceCategoryEnum{}
}

func ServiceCategoryEnumDefault() ServiceCategoryEnum {
	return ServiceCategoryEnumEmpty
}

func (enum ServiceCategoryEnum) IsEmpty() bool {
	return enum.Is(ServiceCategoryEnumEmpty)
}

func (enum ServiceCategoryEnum) IsMaster() bool {
	return enum.Is(ServiceCategoryEnumMaster)
}

func (enum ServiceCategoryEnum) IsLogger() bool {
	return enum.Is(ServiceCategoryEnumLogger)
}

func (enum ServiceCategoryEnum) IsHttpL7Balancer() bool {
	return enum.Is(ServiceCategoryEnumHttpL7Balancer)
}

func (enum ServiceCategoryEnum) IsApplicationBalancer() bool {
	return enum.Is(ServiceCategoryEnumApplicationBalancer)
}

func (enum ServiceCategoryEnum) IsDatabase() bool {
	return enum.Is(ServiceCategoryEnumDatabase)
}

func (enum ServiceCategoryEnum) IsMessageBroker() bool {
	return enum.Is(ServiceCategoryEnumMessageBroker)
}

func (enum ServiceCategoryEnum) IsUserApplication() bool {
	return enum.Is(ServiceCategoryEnumUserApplication)
}

func (enum ServiceCategoryEnum) IsSharedStorage() bool {
	return enum.Is(ServiceCategoryEnumSharedStorage)
}

func (enum ServiceCategoryEnum) IsSystem() bool {
	return enum.Is(ServiceCategoryEnumSystem)
}

func (enum ServiceCategoryEnum) IsBuild() bool {
	return enum.Is(ServiceCategoryEnumBuild)
}

func (enum ServiceCategoryEnum) IsPrepareRuntime() bool {
	return enum.Is(ServiceCategoryEnumPrepareRuntime)
}

func (enum ServiceCategoryEnum) IsMetric() bool {
	return enum.Is(ServiceCategoryEnumMetric)
}
