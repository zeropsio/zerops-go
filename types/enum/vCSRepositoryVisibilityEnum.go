// Generated ZEROPS sdk

package enum

type VCSRepositoryVisibilityEnum string

const (
	VCSRepositoryVisibilityEnumPublic  = VCSRepositoryVisibilityEnum("PUBLIC")
	VCSRepositoryVisibilityEnumPrivate = VCSRepositoryVisibilityEnum("PRIVATE")
)

func NewVCSRepositoryVisibilityEnumFromString(value string) (out VCSRepositoryVisibilityEnum, err error) {
	return VCSRepositoryVisibilityEnum(value), nil
}

func (enum VCSRepositoryVisibilityEnum) String() string {
	return string(enum)
}

func (enum VCSRepositoryVisibilityEnum) Native() string {
	return string(enum)
}

func (enum VCSRepositoryVisibilityEnum) Values() []VCSRepositoryVisibilityEnum {
	return VCSRepositoryVisibilityEnumAll()
}

func (enum VCSRepositoryVisibilityEnum) PublicValues() []VCSRepositoryVisibilityEnum {
	return VCSRepositoryVisibilityEnumAllPublic()
}

func (enum VCSRepositoryVisibilityEnum) PrivateValues() []VCSRepositoryVisibilityEnum {
	return VCSRepositoryVisibilityEnumAllPrivate()
}

func (enum VCSRepositoryVisibilityEnum) DefaultValue() VCSRepositoryVisibilityEnum {
	return VCSRepositoryVisibilityEnumDefault()
}

func (enum VCSRepositoryVisibilityEnum) Is(values ...VCSRepositoryVisibilityEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func VCSRepositoryVisibilityEnumAllStrings() []string {
	return []string{
		string(VCSRepositoryVisibilityEnumPublic), string(VCSRepositoryVisibilityEnumPrivate),
	}
}

func VCSRepositoryVisibilityEnumAll() []VCSRepositoryVisibilityEnum {
	return []VCSRepositoryVisibilityEnum{
		VCSRepositoryVisibilityEnumPublic, VCSRepositoryVisibilityEnumPrivate,
	}
}

func VCSRepositoryVisibilityEnumAllPublic() []VCSRepositoryVisibilityEnum {
	return []VCSRepositoryVisibilityEnum{
		VCSRepositoryVisibilityEnumPublic, VCSRepositoryVisibilityEnumPrivate,
	}
}

func VCSRepositoryVisibilityEnumAllPrivate() []VCSRepositoryVisibilityEnum {
	return []VCSRepositoryVisibilityEnum{}
}

func VCSRepositoryVisibilityEnumDefault() VCSRepositoryVisibilityEnum {
	return VCSRepositoryVisibilityEnumPublic
}

func (enum VCSRepositoryVisibilityEnum) IsPublic() bool {
	return enum.Is(VCSRepositoryVisibilityEnumPublic)
}

func (enum VCSRepositoryVisibilityEnum) IsPrivate() bool {
	return enum.Is(VCSRepositoryVisibilityEnumPrivate)
}
