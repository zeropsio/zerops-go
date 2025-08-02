// Generated ZEROPS sdk

package enum

type ProjectConnectionSpeedSharedEnum string

const (
	ProjectConnectionSpeedSharedEnumLight     = ProjectConnectionSpeedSharedEnum("LIGHT")
	ProjectConnectionSpeedSharedEnumSerious   = ProjectConnectionSpeedSharedEnum("SERIOUS")
	ProjectConnectionSpeedSharedEnumThrottled = ProjectConnectionSpeedSharedEnum("THROTTLED")
	ProjectConnectionSpeedSharedEnumDisabled  = ProjectConnectionSpeedSharedEnum("DISABLED")
	ProjectConnectionSpeedSharedEnumUnlimited = ProjectConnectionSpeedSharedEnum("UNLIMITED")
)

func NewProjectConnectionSpeedSharedEnumFromString(value string) (out ProjectConnectionSpeedSharedEnum, err error) {
	return ProjectConnectionSpeedSharedEnum(value), nil
}

func (enum ProjectConnectionSpeedSharedEnum) String() string {
	return string(enum)
}

func (enum ProjectConnectionSpeedSharedEnum) Native() string {
	return string(enum)
}

func (enum ProjectConnectionSpeedSharedEnum) Values() []ProjectConnectionSpeedSharedEnum {
	return ProjectConnectionSpeedSharedEnumAll()
}

func (enum ProjectConnectionSpeedSharedEnum) PublicValues() []ProjectConnectionSpeedSharedEnum {
	return ProjectConnectionSpeedSharedEnumAllPublic()
}

func (enum ProjectConnectionSpeedSharedEnum) PrivateValues() []ProjectConnectionSpeedSharedEnum {
	return ProjectConnectionSpeedSharedEnumAllPrivate()
}

func (enum ProjectConnectionSpeedSharedEnum) DefaultValue() ProjectConnectionSpeedSharedEnum {
	return ProjectConnectionSpeedSharedEnumDefault()
}

func (enum ProjectConnectionSpeedSharedEnum) Is(values ...ProjectConnectionSpeedSharedEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ProjectConnectionSpeedSharedEnumAllStrings() []string {
	return []string{
		string(ProjectConnectionSpeedSharedEnumLight), string(ProjectConnectionSpeedSharedEnumSerious), string(ProjectConnectionSpeedSharedEnumThrottled), string(ProjectConnectionSpeedSharedEnumDisabled), string(ProjectConnectionSpeedSharedEnumUnlimited),
	}
}

func ProjectConnectionSpeedSharedEnumAll() []ProjectConnectionSpeedSharedEnum {
	return []ProjectConnectionSpeedSharedEnum{
		ProjectConnectionSpeedSharedEnumLight, ProjectConnectionSpeedSharedEnumSerious, ProjectConnectionSpeedSharedEnumThrottled, ProjectConnectionSpeedSharedEnumDisabled, ProjectConnectionSpeedSharedEnumUnlimited,
	}
}

func ProjectConnectionSpeedSharedEnumAllPublic() []ProjectConnectionSpeedSharedEnum {
	return []ProjectConnectionSpeedSharedEnum{
		ProjectConnectionSpeedSharedEnumLight, ProjectConnectionSpeedSharedEnumSerious, ProjectConnectionSpeedSharedEnumThrottled, ProjectConnectionSpeedSharedEnumDisabled, ProjectConnectionSpeedSharedEnumUnlimited,
	}
}

func ProjectConnectionSpeedSharedEnumAllPrivate() []ProjectConnectionSpeedSharedEnum {
	return []ProjectConnectionSpeedSharedEnum{}
}

func ProjectConnectionSpeedSharedEnumDefault() ProjectConnectionSpeedSharedEnum {
	return ProjectConnectionSpeedSharedEnumLight
}

func (enum ProjectConnectionSpeedSharedEnum) IsLight() bool {
	return enum.Is(ProjectConnectionSpeedSharedEnumLight)
}

func (enum ProjectConnectionSpeedSharedEnum) IsSerious() bool {
	return enum.Is(ProjectConnectionSpeedSharedEnumSerious)
}

func (enum ProjectConnectionSpeedSharedEnum) IsThrottled() bool {
	return enum.Is(ProjectConnectionSpeedSharedEnumThrottled)
}

func (enum ProjectConnectionSpeedSharedEnum) IsDisabled() bool {
	return enum.Is(ProjectConnectionSpeedSharedEnumDisabled)
}

func (enum ProjectConnectionSpeedSharedEnum) IsUnlimited() bool {
	return enum.Is(ProjectConnectionSpeedSharedEnumUnlimited)
}
