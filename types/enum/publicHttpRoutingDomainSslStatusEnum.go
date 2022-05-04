// Generated ZEROPS sdk

package enum

type PublicHttpRoutingDomainSslStatusEnum string

const (
	PublicHttpRoutingDomainSslStatusEnumInactive           = PublicHttpRoutingDomainSslStatusEnum("INACTIVE")
	PublicHttpRoutingDomainSslStatusEnumBeingInstalled     = PublicHttpRoutingDomainSslStatusEnum("BEING_INSTALLED")
	PublicHttpRoutingDomainSslStatusEnumWaitingForDns      = PublicHttpRoutingDomainSslStatusEnum("WAITING_FOR_DNS")
	PublicHttpRoutingDomainSslStatusEnumActive             = PublicHttpRoutingDomainSslStatusEnum("ACTIVE")
	PublicHttpRoutingDomainSslStatusEnumInstallationFailed = PublicHttpRoutingDomainSslStatusEnum("INSTALLATION_FAILED")
)

func NewPublicHttpRoutingDomainSslStatusEnumFromString(value string) (out PublicHttpRoutingDomainSslStatusEnum, err error) {
	return PublicHttpRoutingDomainSslStatusEnum(value), nil
}

func (enum PublicHttpRoutingDomainSslStatusEnum) String() string {
	return string(enum)
}

func (enum PublicHttpRoutingDomainSslStatusEnum) Native() string {
	return string(enum)
}

func (enum PublicHttpRoutingDomainSslStatusEnum) Values() []PublicHttpRoutingDomainSslStatusEnum {
	return PublicHttpRoutingDomainSslStatusEnumAll()
}

func (enum PublicHttpRoutingDomainSslStatusEnum) PublicValues() []PublicHttpRoutingDomainSslStatusEnum {
	return PublicHttpRoutingDomainSslStatusEnumAllPublic()
}

func (enum PublicHttpRoutingDomainSslStatusEnum) PrivateValues() []PublicHttpRoutingDomainSslStatusEnum {
	return PublicHttpRoutingDomainSslStatusEnumAllPrivate()
}

func (enum PublicHttpRoutingDomainSslStatusEnum) DefaultValue() PublicHttpRoutingDomainSslStatusEnum {
	return PublicHttpRoutingDomainSslStatusEnumDefault()
}

func (enum PublicHttpRoutingDomainSslStatusEnum) Is(values ...PublicHttpRoutingDomainSslStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func PublicHttpRoutingDomainSslStatusEnumAllStrings() []string {
	return []string{
		string(PublicHttpRoutingDomainSslStatusEnumInactive), string(PublicHttpRoutingDomainSslStatusEnumBeingInstalled), string(PublicHttpRoutingDomainSslStatusEnumWaitingForDns), string(PublicHttpRoutingDomainSslStatusEnumActive), string(PublicHttpRoutingDomainSslStatusEnumInstallationFailed),
	}
}

func PublicHttpRoutingDomainSslStatusEnumAll() []PublicHttpRoutingDomainSslStatusEnum {
	return []PublicHttpRoutingDomainSslStatusEnum{
		PublicHttpRoutingDomainSslStatusEnumInactive, PublicHttpRoutingDomainSslStatusEnumBeingInstalled, PublicHttpRoutingDomainSslStatusEnumWaitingForDns, PublicHttpRoutingDomainSslStatusEnumActive, PublicHttpRoutingDomainSslStatusEnumInstallationFailed,
	}
}

func PublicHttpRoutingDomainSslStatusEnumAllPublic() []PublicHttpRoutingDomainSslStatusEnum {
	return []PublicHttpRoutingDomainSslStatusEnum{
		PublicHttpRoutingDomainSslStatusEnumInactive, PublicHttpRoutingDomainSslStatusEnumBeingInstalled, PublicHttpRoutingDomainSslStatusEnumWaitingForDns, PublicHttpRoutingDomainSslStatusEnumActive, PublicHttpRoutingDomainSslStatusEnumInstallationFailed,
	}
}

func PublicHttpRoutingDomainSslStatusEnumAllPrivate() []PublicHttpRoutingDomainSslStatusEnum {
	return []PublicHttpRoutingDomainSslStatusEnum{}
}

func PublicHttpRoutingDomainSslStatusEnumDefault() PublicHttpRoutingDomainSslStatusEnum {
	return PublicHttpRoutingDomainSslStatusEnumInactive
}

func (enum PublicHttpRoutingDomainSslStatusEnum) IsInactive() bool {
	return enum.Is(PublicHttpRoutingDomainSslStatusEnumInactive)
}

func (enum PublicHttpRoutingDomainSslStatusEnum) IsBeingInstalled() bool {
	return enum.Is(PublicHttpRoutingDomainSslStatusEnumBeingInstalled)
}

func (enum PublicHttpRoutingDomainSslStatusEnum) IsWaitingForDns() bool {
	return enum.Is(PublicHttpRoutingDomainSslStatusEnumWaitingForDns)
}

func (enum PublicHttpRoutingDomainSslStatusEnum) IsActive() bool {
	return enum.Is(PublicHttpRoutingDomainSslStatusEnumActive)
}

func (enum PublicHttpRoutingDomainSslStatusEnum) IsInstallationFailed() bool {
	return enum.Is(PublicHttpRoutingDomainSslStatusEnumInstallationFailed)
}
