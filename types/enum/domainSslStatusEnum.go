// Generated ZEROPS sdk

package enum

type DomainSslStatusEnum string

const (
	DomainSslStatusEnumInactive           = DomainSslStatusEnum("INACTIVE")
	DomainSslStatusEnumWaitingForDns      = DomainSslStatusEnum("WAITING_FOR_DNS")
	DomainSslStatusEnumBeingInstalled     = DomainSslStatusEnum("BEING_INSTALLED")
	DomainSslStatusEnumActive             = DomainSslStatusEnum("ACTIVE")
	DomainSslStatusEnumInstallationFailed = DomainSslStatusEnum("INSTALLATION_FAILED")
)

func NewDomainSslStatusEnumFromString(value string) (out DomainSslStatusEnum, err error) {
	return DomainSslStatusEnum(value), nil
}

func (enum DomainSslStatusEnum) String() string {
	return string(enum)
}

func (enum DomainSslStatusEnum) Native() string {
	return string(enum)
}

func (enum DomainSslStatusEnum) Values() []DomainSslStatusEnum {
	return DomainSslStatusEnumAll()
}

func (enum DomainSslStatusEnum) PublicValues() []DomainSslStatusEnum {
	return DomainSslStatusEnumAllPublic()
}

func (enum DomainSslStatusEnum) PrivateValues() []DomainSslStatusEnum {
	return DomainSslStatusEnumAllPrivate()
}

func (enum DomainSslStatusEnum) DefaultValue() DomainSslStatusEnum {
	return DomainSslStatusEnumDefault()
}

func (enum DomainSslStatusEnum) Is(values ...DomainSslStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func DomainSslStatusEnumAllStrings() []string {
	return []string{
		string(DomainSslStatusEnumInactive), string(DomainSslStatusEnumWaitingForDns), string(DomainSslStatusEnumBeingInstalled), string(DomainSslStatusEnumActive), string(DomainSslStatusEnumInstallationFailed),
	}
}

func DomainSslStatusEnumAll() []DomainSslStatusEnum {
	return []DomainSslStatusEnum{
		DomainSslStatusEnumInactive, DomainSslStatusEnumWaitingForDns, DomainSslStatusEnumBeingInstalled, DomainSslStatusEnumActive, DomainSslStatusEnumInstallationFailed,
	}
}

func DomainSslStatusEnumAllPublic() []DomainSslStatusEnum {
	return []DomainSslStatusEnum{
		DomainSslStatusEnumInactive, DomainSslStatusEnumWaitingForDns, DomainSslStatusEnumBeingInstalled, DomainSslStatusEnumActive, DomainSslStatusEnumInstallationFailed,
	}
}

func DomainSslStatusEnumAllPrivate() []DomainSslStatusEnum {
	return []DomainSslStatusEnum{}
}

func DomainSslStatusEnumDefault() DomainSslStatusEnum {
	return DomainSslStatusEnumInactive
}

func (enum DomainSslStatusEnum) IsInactive() bool {
	return enum.Is(DomainSslStatusEnumInactive)
}

func (enum DomainSslStatusEnum) IsWaitingForDns() bool {
	return enum.Is(DomainSslStatusEnumWaitingForDns)
}

func (enum DomainSslStatusEnum) IsBeingInstalled() bool {
	return enum.Is(DomainSslStatusEnumBeingInstalled)
}

func (enum DomainSslStatusEnum) IsActive() bool {
	return enum.Is(DomainSslStatusEnumActive)
}

func (enum DomainSslStatusEnum) IsInstallationFailed() bool {
	return enum.Is(DomainSslStatusEnumInstallationFailed)
}
