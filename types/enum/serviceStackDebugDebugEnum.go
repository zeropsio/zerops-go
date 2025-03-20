// Generated ZEROPS sdk

package enum

type ServiceStackDebugDebugEnum string

const (
	ServiceStackDebugDebugEnumNone   = ServiceStackDebugDebugEnum("NONE")
	ServiceStackDebugDebugEnumBefore = ServiceStackDebugDebugEnum("BEFORE")
	ServiceStackDebugDebugEnumOnFail = ServiceStackDebugDebugEnum("ON_FAIL")
	ServiceStackDebugDebugEnumAfter  = ServiceStackDebugDebugEnum("AFTER")
)

func NewServiceStackDebugDebugEnumFromString(value string) (out ServiceStackDebugDebugEnum, err error) {
	return ServiceStackDebugDebugEnum(value), nil
}

func (enum ServiceStackDebugDebugEnum) String() string {
	return string(enum)
}

func (enum ServiceStackDebugDebugEnum) Native() string {
	return string(enum)
}

func (enum ServiceStackDebugDebugEnum) Values() []ServiceStackDebugDebugEnum {
	return ServiceStackDebugDebugEnumAll()
}

func (enum ServiceStackDebugDebugEnum) PublicValues() []ServiceStackDebugDebugEnum {
	return ServiceStackDebugDebugEnumAllPublic()
}

func (enum ServiceStackDebugDebugEnum) PrivateValues() []ServiceStackDebugDebugEnum {
	return ServiceStackDebugDebugEnumAllPrivate()
}

func (enum ServiceStackDebugDebugEnum) DefaultValue() ServiceStackDebugDebugEnum {
	return ServiceStackDebugDebugEnumDefault()
}

func (enum ServiceStackDebugDebugEnum) Is(values ...ServiceStackDebugDebugEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ServiceStackDebugDebugEnumAllStrings() []string {
	return []string{
		string(ServiceStackDebugDebugEnumNone), string(ServiceStackDebugDebugEnumBefore), string(ServiceStackDebugDebugEnumOnFail), string(ServiceStackDebugDebugEnumAfter),
	}
}

func ServiceStackDebugDebugEnumAll() []ServiceStackDebugDebugEnum {
	return []ServiceStackDebugDebugEnum{
		ServiceStackDebugDebugEnumNone, ServiceStackDebugDebugEnumBefore, ServiceStackDebugDebugEnumOnFail, ServiceStackDebugDebugEnumAfter,
	}
}

func ServiceStackDebugDebugEnumAllPublic() []ServiceStackDebugDebugEnum {
	return []ServiceStackDebugDebugEnum{
		ServiceStackDebugDebugEnumNone, ServiceStackDebugDebugEnumBefore, ServiceStackDebugDebugEnumOnFail, ServiceStackDebugDebugEnumAfter,
	}
}

func ServiceStackDebugDebugEnumAllPrivate() []ServiceStackDebugDebugEnum {
	return []ServiceStackDebugDebugEnum{}
}

func ServiceStackDebugDebugEnumDefault() ServiceStackDebugDebugEnum {
	return ""
}

func (enum ServiceStackDebugDebugEnum) IsNone() bool {
	return enum.Is(ServiceStackDebugDebugEnumNone)
}

func (enum ServiceStackDebugDebugEnum) IsBefore() bool {
	return enum.Is(ServiceStackDebugDebugEnumBefore)
}

func (enum ServiceStackDebugDebugEnum) IsOnFail() bool {
	return enum.Is(ServiceStackDebugDebugEnumOnFail)
}

func (enum ServiceStackDebugDebugEnum) IsAfter() bool {
	return enum.Is(ServiceStackDebugDebugEnumAfter)
}
