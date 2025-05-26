// Generated ZEROPS sdk

package enum

type PublicHttpRoutingDomainCdnStatusEnum string

const (
	PublicHttpRoutingDomainCdnStatusEnumDisabled       = PublicHttpRoutingDomainCdnStatusEnum("DISABLED")
	PublicHttpRoutingDomainCdnStatusEnumBeingInstalled = PublicHttpRoutingDomainCdnStatusEnum("BEING_INSTALLED")
	PublicHttpRoutingDomainCdnStatusEnumActive         = PublicHttpRoutingDomainCdnStatusEnum("ACTIVE")
	PublicHttpRoutingDomainCdnStatusEnumBeingDisabled  = PublicHttpRoutingDomainCdnStatusEnum("BEING_DISABLED")
)

func NewPublicHttpRoutingDomainCdnStatusEnumFromString(value string) (out PublicHttpRoutingDomainCdnStatusEnum, err error) {
	return PublicHttpRoutingDomainCdnStatusEnum(value), nil
}

func (enum PublicHttpRoutingDomainCdnStatusEnum) String() string {
	return string(enum)
}

func (enum PublicHttpRoutingDomainCdnStatusEnum) Native() string {
	return string(enum)
}

func (enum PublicHttpRoutingDomainCdnStatusEnum) Values() []PublicHttpRoutingDomainCdnStatusEnum {
	return PublicHttpRoutingDomainCdnStatusEnumAll()
}

func (enum PublicHttpRoutingDomainCdnStatusEnum) PublicValues() []PublicHttpRoutingDomainCdnStatusEnum {
	return PublicHttpRoutingDomainCdnStatusEnumAllPublic()
}

func (enum PublicHttpRoutingDomainCdnStatusEnum) PrivateValues() []PublicHttpRoutingDomainCdnStatusEnum {
	return PublicHttpRoutingDomainCdnStatusEnumAllPrivate()
}

func (enum PublicHttpRoutingDomainCdnStatusEnum) DefaultValue() PublicHttpRoutingDomainCdnStatusEnum {
	return PublicHttpRoutingDomainCdnStatusEnumDefault()
}

func (enum PublicHttpRoutingDomainCdnStatusEnum) Is(values ...PublicHttpRoutingDomainCdnStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func PublicHttpRoutingDomainCdnStatusEnumAllStrings() []string {
	return []string{
		string(PublicHttpRoutingDomainCdnStatusEnumDisabled), string(PublicHttpRoutingDomainCdnStatusEnumBeingInstalled), string(PublicHttpRoutingDomainCdnStatusEnumActive), string(PublicHttpRoutingDomainCdnStatusEnumBeingDisabled),
	}
}

func PublicHttpRoutingDomainCdnStatusEnumAll() []PublicHttpRoutingDomainCdnStatusEnum {
	return []PublicHttpRoutingDomainCdnStatusEnum{
		PublicHttpRoutingDomainCdnStatusEnumDisabled, PublicHttpRoutingDomainCdnStatusEnumBeingInstalled, PublicHttpRoutingDomainCdnStatusEnumActive, PublicHttpRoutingDomainCdnStatusEnumBeingDisabled,
	}
}

func PublicHttpRoutingDomainCdnStatusEnumAllPublic() []PublicHttpRoutingDomainCdnStatusEnum {
	return []PublicHttpRoutingDomainCdnStatusEnum{
		PublicHttpRoutingDomainCdnStatusEnumDisabled, PublicHttpRoutingDomainCdnStatusEnumBeingInstalled, PublicHttpRoutingDomainCdnStatusEnumActive, PublicHttpRoutingDomainCdnStatusEnumBeingDisabled,
	}
}

func PublicHttpRoutingDomainCdnStatusEnumAllPrivate() []PublicHttpRoutingDomainCdnStatusEnum {
	return []PublicHttpRoutingDomainCdnStatusEnum{}
}

func PublicHttpRoutingDomainCdnStatusEnumDefault() PublicHttpRoutingDomainCdnStatusEnum {
	return ""
}

func (enum PublicHttpRoutingDomainCdnStatusEnum) IsDisabled() bool {
	return enum.Is(PublicHttpRoutingDomainCdnStatusEnumDisabled)
}

func (enum PublicHttpRoutingDomainCdnStatusEnum) IsBeingInstalled() bool {
	return enum.Is(PublicHttpRoutingDomainCdnStatusEnumBeingInstalled)
}

func (enum PublicHttpRoutingDomainCdnStatusEnum) IsActive() bool {
	return enum.Is(PublicHttpRoutingDomainCdnStatusEnumActive)
}

func (enum PublicHttpRoutingDomainCdnStatusEnum) IsBeingDisabled() bool {
	return enum.Is(PublicHttpRoutingDomainCdnStatusEnumBeingDisabled)
}
