// Generated ZEROPS sdk

package enum

type DomainSslStatusEnum string

const (
	DomainSslStatusEnumInactive       = DomainSslStatusEnum("INACTIVE")
	DomainSslStatusEnumWaitingForDns  = DomainSslStatusEnum("WAITING_FOR_DNS")
	DomainSslStatusEnumBeingInstalled = DomainSslStatusEnum("BEING_INSTALLED")
	DomainSslStatusEnumBeingRenewed   = DomainSslStatusEnum("BEING_RENEWED")
	DomainSslStatusEnumActive         = DomainSslStatusEnum("ACTIVE")
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
		string(DomainSslStatusEnumInactive), string(DomainSslStatusEnumWaitingForDns), string(DomainSslStatusEnumBeingInstalled), string(DomainSslStatusEnumBeingRenewed), string(DomainSslStatusEnumActive),
	}
}

func DomainSslStatusEnumAll() []DomainSslStatusEnum {
	return []DomainSslStatusEnum{
		DomainSslStatusEnumInactive, DomainSslStatusEnumWaitingForDns, DomainSslStatusEnumBeingInstalled, DomainSslStatusEnumBeingRenewed, DomainSslStatusEnumActive,
	}
}

func DomainSslStatusEnumAllPublic() []DomainSslStatusEnum {
	return []DomainSslStatusEnum{
		DomainSslStatusEnumInactive, DomainSslStatusEnumWaitingForDns, DomainSslStatusEnumBeingInstalled, DomainSslStatusEnumBeingRenewed, DomainSslStatusEnumActive,
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

func (enum DomainSslStatusEnum) IsBeingRenewed() bool {
	return enum.Is(DomainSslStatusEnumBeingRenewed)
}

func (enum DomainSslStatusEnum) IsActive() bool {
	return enum.Is(DomainSslStatusEnumActive)
}
