// Generated ZEROPS sdk

package enum

type ServicePortSchemeEnum string

const (
	ServicePortSchemeEnumTcp        = ServicePortSchemeEnum("tcp")
	ServicePortSchemeEnumUdp        = ServicePortSchemeEnum("udp")
	ServicePortSchemeEnumHttp       = ServicePortSchemeEnum("http")
	ServicePortSchemeEnumHttps      = ServicePortSchemeEnum("https")
	ServicePortSchemeEnumRedis      = ServicePortSchemeEnum("redis")
	ServicePortSchemeEnumMysql      = ServicePortSchemeEnum("mysql")
	ServicePortSchemeEnumUdpinflux  = ServicePortSchemeEnum("udpinflux")
	ServicePortSchemeEnumMongodb    = ServicePortSchemeEnum("mongodb")
	ServicePortSchemeEnumPostgresql = ServicePortSchemeEnum("postgresql")
	ServicePortSchemeEnumAmqp       = ServicePortSchemeEnum("amqp")
	ServicePortSchemeEnumStomp      = ServicePortSchemeEnum("stomp")
	ServicePortSchemeEnumMqtt       = ServicePortSchemeEnum("mqtt")
)

func NewServicePortSchemeEnumFromString(value string) (out ServicePortSchemeEnum, err error) {
	return ServicePortSchemeEnum(value), nil
}

func (enum ServicePortSchemeEnum) String() string {
	return string(enum)
}

func (enum ServicePortSchemeEnum) Native() string {
	return string(enum)
}

func (enum ServicePortSchemeEnum) Values() []ServicePortSchemeEnum {
	return ServicePortSchemeEnumAll()
}

func (enum ServicePortSchemeEnum) PublicValues() []ServicePortSchemeEnum {
	return ServicePortSchemeEnumAllPublic()
}

func (enum ServicePortSchemeEnum) PrivateValues() []ServicePortSchemeEnum {
	return ServicePortSchemeEnumAllPrivate()
}

func (enum ServicePortSchemeEnum) DefaultValue() ServicePortSchemeEnum {
	return ServicePortSchemeEnumDefault()
}

func (enum ServicePortSchemeEnum) Is(values ...ServicePortSchemeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ServicePortSchemeEnumAllStrings() []string {
	return []string{
		string(ServicePortSchemeEnumTcp), string(ServicePortSchemeEnumUdp), string(ServicePortSchemeEnumHttp), string(ServicePortSchemeEnumHttps), string(ServicePortSchemeEnumRedis), string(ServicePortSchemeEnumMysql), string(ServicePortSchemeEnumUdpinflux), string(ServicePortSchemeEnumMongodb), string(ServicePortSchemeEnumPostgresql), string(ServicePortSchemeEnumAmqp), string(ServicePortSchemeEnumStomp), string(ServicePortSchemeEnumMqtt),
	}
}

func ServicePortSchemeEnumAll() []ServicePortSchemeEnum {
	return []ServicePortSchemeEnum{
		ServicePortSchemeEnumTcp, ServicePortSchemeEnumUdp, ServicePortSchemeEnumHttp, ServicePortSchemeEnumHttps, ServicePortSchemeEnumRedis, ServicePortSchemeEnumMysql, ServicePortSchemeEnumUdpinflux, ServicePortSchemeEnumMongodb, ServicePortSchemeEnumPostgresql, ServicePortSchemeEnumAmqp, ServicePortSchemeEnumStomp, ServicePortSchemeEnumMqtt,
	}
}

func ServicePortSchemeEnumAllPublic() []ServicePortSchemeEnum {
	return []ServicePortSchemeEnum{
		ServicePortSchemeEnumTcp, ServicePortSchemeEnumUdp, ServicePortSchemeEnumHttp, ServicePortSchemeEnumHttps, ServicePortSchemeEnumRedis, ServicePortSchemeEnumMysql, ServicePortSchemeEnumUdpinflux, ServicePortSchemeEnumMongodb, ServicePortSchemeEnumPostgresql, ServicePortSchemeEnumAmqp, ServicePortSchemeEnumStomp, ServicePortSchemeEnumMqtt,
	}
}

func ServicePortSchemeEnumAllPrivate() []ServicePortSchemeEnum {
	return []ServicePortSchemeEnum{}
}

func ServicePortSchemeEnumDefault() ServicePortSchemeEnum {
	return ServicePortSchemeEnumTcp
}

func (enum ServicePortSchemeEnum) IsTcp() bool {
	return enum.Is(ServicePortSchemeEnumTcp)
}

func (enum ServicePortSchemeEnum) IsUdp() bool {
	return enum.Is(ServicePortSchemeEnumUdp)
}

func (enum ServicePortSchemeEnum) IsHttp() bool {
	return enum.Is(ServicePortSchemeEnumHttp)
}

func (enum ServicePortSchemeEnum) IsHttps() bool {
	return enum.Is(ServicePortSchemeEnumHttps)
}

func (enum ServicePortSchemeEnum) IsRedis() bool {
	return enum.Is(ServicePortSchemeEnumRedis)
}

func (enum ServicePortSchemeEnum) IsMysql() bool {
	return enum.Is(ServicePortSchemeEnumMysql)
}

func (enum ServicePortSchemeEnum) IsUdpinflux() bool {
	return enum.Is(ServicePortSchemeEnumUdpinflux)
}

func (enum ServicePortSchemeEnum) IsMongodb() bool {
	return enum.Is(ServicePortSchemeEnumMongodb)
}

func (enum ServicePortSchemeEnum) IsPostgresql() bool {
	return enum.Is(ServicePortSchemeEnumPostgresql)
}

func (enum ServicePortSchemeEnum) IsAmqp() bool {
	return enum.Is(ServicePortSchemeEnumAmqp)
}

func (enum ServicePortSchemeEnum) IsStomp() bool {
	return enum.Is(ServicePortSchemeEnumStomp)
}

func (enum ServicePortSchemeEnum) IsMqtt() bool {
	return enum.Is(ServicePortSchemeEnumMqtt)
}
