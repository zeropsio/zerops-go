// Generated ZEROPS sdk

package enum

type PublicHttpRoutingDomainDnsCheckStatusEnum string

const (
	PublicHttpRoutingDomainDnsCheckStatusEnumPending  = PublicHttpRoutingDomainDnsCheckStatusEnum("PENDING")
	PublicHttpRoutingDomainDnsCheckStatusEnumChecking = PublicHttpRoutingDomainDnsCheckStatusEnum("CHECKING")
	PublicHttpRoutingDomainDnsCheckStatusEnumOk       = PublicHttpRoutingDomainDnsCheckStatusEnum("OK")
	PublicHttpRoutingDomainDnsCheckStatusEnumFailed   = PublicHttpRoutingDomainDnsCheckStatusEnum("FAILED")
	PublicHttpRoutingDomainDnsCheckStatusEnumIgnored  = PublicHttpRoutingDomainDnsCheckStatusEnum("IGNORED")
)

func NewPublicHttpRoutingDomainDnsCheckStatusEnumFromString(value string) (out PublicHttpRoutingDomainDnsCheckStatusEnum, err error) {
	return PublicHttpRoutingDomainDnsCheckStatusEnum(value), nil
}

func (enum PublicHttpRoutingDomainDnsCheckStatusEnum) String() string {
	return string(enum)
}

func (enum PublicHttpRoutingDomainDnsCheckStatusEnum) Native() string {
	return string(enum)
}

func (enum PublicHttpRoutingDomainDnsCheckStatusEnum) Values() []PublicHttpRoutingDomainDnsCheckStatusEnum {
	return PublicHttpRoutingDomainDnsCheckStatusEnumAll()
}

func (enum PublicHttpRoutingDomainDnsCheckStatusEnum) PublicValues() []PublicHttpRoutingDomainDnsCheckStatusEnum {
	return PublicHttpRoutingDomainDnsCheckStatusEnumAllPublic()
}

func (enum PublicHttpRoutingDomainDnsCheckStatusEnum) PrivateValues() []PublicHttpRoutingDomainDnsCheckStatusEnum {
	return PublicHttpRoutingDomainDnsCheckStatusEnumAllPrivate()
}

func (enum PublicHttpRoutingDomainDnsCheckStatusEnum) DefaultValue() PublicHttpRoutingDomainDnsCheckStatusEnum {
	return PublicHttpRoutingDomainDnsCheckStatusEnumDefault()
}

func (enum PublicHttpRoutingDomainDnsCheckStatusEnum) Is(values ...PublicHttpRoutingDomainDnsCheckStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func PublicHttpRoutingDomainDnsCheckStatusEnumAllStrings() []string {
	return []string{
		string(PublicHttpRoutingDomainDnsCheckStatusEnumPending), string(PublicHttpRoutingDomainDnsCheckStatusEnumChecking), string(PublicHttpRoutingDomainDnsCheckStatusEnumOk), string(PublicHttpRoutingDomainDnsCheckStatusEnumFailed), string(PublicHttpRoutingDomainDnsCheckStatusEnumIgnored),
	}
}

func PublicHttpRoutingDomainDnsCheckStatusEnumAll() []PublicHttpRoutingDomainDnsCheckStatusEnum {
	return []PublicHttpRoutingDomainDnsCheckStatusEnum{
		PublicHttpRoutingDomainDnsCheckStatusEnumPending, PublicHttpRoutingDomainDnsCheckStatusEnumChecking, PublicHttpRoutingDomainDnsCheckStatusEnumOk, PublicHttpRoutingDomainDnsCheckStatusEnumFailed, PublicHttpRoutingDomainDnsCheckStatusEnumIgnored,
	}
}

func PublicHttpRoutingDomainDnsCheckStatusEnumAllPublic() []PublicHttpRoutingDomainDnsCheckStatusEnum {
	return []PublicHttpRoutingDomainDnsCheckStatusEnum{
		PublicHttpRoutingDomainDnsCheckStatusEnumPending, PublicHttpRoutingDomainDnsCheckStatusEnumChecking, PublicHttpRoutingDomainDnsCheckStatusEnumOk, PublicHttpRoutingDomainDnsCheckStatusEnumFailed, PublicHttpRoutingDomainDnsCheckStatusEnumIgnored,
	}
}

func PublicHttpRoutingDomainDnsCheckStatusEnumAllPrivate() []PublicHttpRoutingDomainDnsCheckStatusEnum {
	return []PublicHttpRoutingDomainDnsCheckStatusEnum{}
}

func PublicHttpRoutingDomainDnsCheckStatusEnumDefault() PublicHttpRoutingDomainDnsCheckStatusEnum {
	return PublicHttpRoutingDomainDnsCheckStatusEnumPending
}

func (enum PublicHttpRoutingDomainDnsCheckStatusEnum) IsPending() bool {
	return enum.Is(PublicHttpRoutingDomainDnsCheckStatusEnumPending)
}

func (enum PublicHttpRoutingDomainDnsCheckStatusEnum) IsChecking() bool {
	return enum.Is(PublicHttpRoutingDomainDnsCheckStatusEnumChecking)
}

func (enum PublicHttpRoutingDomainDnsCheckStatusEnum) IsOk() bool {
	return enum.Is(PublicHttpRoutingDomainDnsCheckStatusEnumOk)
}

func (enum PublicHttpRoutingDomainDnsCheckStatusEnum) IsFailed() bool {
	return enum.Is(PublicHttpRoutingDomainDnsCheckStatusEnumFailed)
}

func (enum PublicHttpRoutingDomainDnsCheckStatusEnum) IsIgnored() bool {
	return enum.Is(PublicHttpRoutingDomainDnsCheckStatusEnumIgnored)
}
