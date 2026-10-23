// Generated ZEROPS sdk

package enum

type PublicHttpRoutingDomainHostMatchEnum string

const (
	PublicHttpRoutingDomainHostMatchEnumExact         = PublicHttpRoutingDomainHostMatchEnum("EXACT")
	PublicHttpRoutingDomainHostMatchEnumSandboxPrefix = PublicHttpRoutingDomainHostMatchEnum("SANDBOX_PREFIX")
)

func NewPublicHttpRoutingDomainHostMatchEnumFromString(value string) (out PublicHttpRoutingDomainHostMatchEnum, err error) {
	return PublicHttpRoutingDomainHostMatchEnum(value), nil
}

func (enum PublicHttpRoutingDomainHostMatchEnum) String() string {
	return string(enum)
}

func (enum PublicHttpRoutingDomainHostMatchEnum) Native() string {
	return string(enum)
}

func (enum PublicHttpRoutingDomainHostMatchEnum) Values() []PublicHttpRoutingDomainHostMatchEnum {
	return PublicHttpRoutingDomainHostMatchEnumAll()
}

func (enum PublicHttpRoutingDomainHostMatchEnum) PublicValues() []PublicHttpRoutingDomainHostMatchEnum {
	return PublicHttpRoutingDomainHostMatchEnumAllPublic()
}

func (enum PublicHttpRoutingDomainHostMatchEnum) PrivateValues() []PublicHttpRoutingDomainHostMatchEnum {
	return PublicHttpRoutingDomainHostMatchEnumAllPrivate()
}

func (enum PublicHttpRoutingDomainHostMatchEnum) DefaultValue() PublicHttpRoutingDomainHostMatchEnum {
	return PublicHttpRoutingDomainHostMatchEnumDefault()
}

func (enum PublicHttpRoutingDomainHostMatchEnum) Is(values ...PublicHttpRoutingDomainHostMatchEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func PublicHttpRoutingDomainHostMatchEnumAllStrings() []string {
	return []string{
		string(PublicHttpRoutingDomainHostMatchEnumExact), string(PublicHttpRoutingDomainHostMatchEnumSandboxPrefix),
	}
}

func PublicHttpRoutingDomainHostMatchEnumAll() []PublicHttpRoutingDomainHostMatchEnum {
	return []PublicHttpRoutingDomainHostMatchEnum{
		PublicHttpRoutingDomainHostMatchEnumExact, PublicHttpRoutingDomainHostMatchEnumSandboxPrefix,
	}
}

func PublicHttpRoutingDomainHostMatchEnumAllPublic() []PublicHttpRoutingDomainHostMatchEnum {
	return []PublicHttpRoutingDomainHostMatchEnum{
		PublicHttpRoutingDomainHostMatchEnumExact, PublicHttpRoutingDomainHostMatchEnumSandboxPrefix,
	}
}

func PublicHttpRoutingDomainHostMatchEnumAllPrivate() []PublicHttpRoutingDomainHostMatchEnum {
	return []PublicHttpRoutingDomainHostMatchEnum{}
}

func PublicHttpRoutingDomainHostMatchEnumDefault() PublicHttpRoutingDomainHostMatchEnum {
	return ""
}

func (enum PublicHttpRoutingDomainHostMatchEnum) IsExact() bool {
	return enum.Is(PublicHttpRoutingDomainHostMatchEnumExact)
}

func (enum PublicHttpRoutingDomainHostMatchEnum) IsSandboxPrefix() bool {
	return enum.Is(PublicHttpRoutingDomainHostMatchEnumSandboxPrefix)
}
