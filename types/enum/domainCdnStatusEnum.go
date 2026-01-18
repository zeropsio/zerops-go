// Generated ZEROPS sdk

package enum

type DomainCdnStatusEnum string

const (
	DomainCdnStatusEnumDisabled       = DomainCdnStatusEnum("DISABLED")
	DomainCdnStatusEnumBeingInstalled = DomainCdnStatusEnum("BEING_INSTALLED")
	DomainCdnStatusEnumActive         = DomainCdnStatusEnum("ACTIVE")
	DomainCdnStatusEnumBeingDisabled  = DomainCdnStatusEnum("BEING_DISABLED")
)

func NewDomainCdnStatusEnumFromString(value string) (out DomainCdnStatusEnum, err error) {
	return DomainCdnStatusEnum(value), nil
}

func (enum DomainCdnStatusEnum) String() string {
	return string(enum)
}

func (enum DomainCdnStatusEnum) Native() string {
	return string(enum)
}

func (enum DomainCdnStatusEnum) Values() []DomainCdnStatusEnum {
	return DomainCdnStatusEnumAll()
}

func (enum DomainCdnStatusEnum) PublicValues() []DomainCdnStatusEnum {
	return DomainCdnStatusEnumAllPublic()
}

func (enum DomainCdnStatusEnum) PrivateValues() []DomainCdnStatusEnum {
	return DomainCdnStatusEnumAllPrivate()
}

func (enum DomainCdnStatusEnum) DefaultValue() DomainCdnStatusEnum {
	return DomainCdnStatusEnumDefault()
}

func (enum DomainCdnStatusEnum) Is(values ...DomainCdnStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func DomainCdnStatusEnumAllStrings() []string {
	return []string{
		string(DomainCdnStatusEnumDisabled), string(DomainCdnStatusEnumBeingInstalled), string(DomainCdnStatusEnumActive), string(DomainCdnStatusEnumBeingDisabled),
	}
}

func DomainCdnStatusEnumAll() []DomainCdnStatusEnum {
	return []DomainCdnStatusEnum{
		DomainCdnStatusEnumDisabled, DomainCdnStatusEnumBeingInstalled, DomainCdnStatusEnumActive, DomainCdnStatusEnumBeingDisabled,
	}
}

func DomainCdnStatusEnumAllPublic() []DomainCdnStatusEnum {
	return []DomainCdnStatusEnum{
		DomainCdnStatusEnumDisabled, DomainCdnStatusEnumBeingInstalled, DomainCdnStatusEnumActive, DomainCdnStatusEnumBeingDisabled,
	}
}

func DomainCdnStatusEnumAllPrivate() []DomainCdnStatusEnum {
	return []DomainCdnStatusEnum{}
}

func DomainCdnStatusEnumDefault() DomainCdnStatusEnum {
	return DomainCdnStatusEnumDisabled
}

func (enum DomainCdnStatusEnum) IsDisabled() bool {
	return enum.Is(DomainCdnStatusEnumDisabled)
}

func (enum DomainCdnStatusEnum) IsBeingInstalled() bool {
	return enum.Is(DomainCdnStatusEnumBeingInstalled)
}

func (enum DomainCdnStatusEnum) IsActive() bool {
	return enum.Is(DomainCdnStatusEnumActive)
}

func (enum DomainCdnStatusEnum) IsBeingDisabled() bool {
	return enum.Is(DomainCdnStatusEnumBeingDisabled)
}
