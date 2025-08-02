// Generated ZEROPS sdk

package enum

type ProjectConnectionSpeedAssignedEnum string

const (
	ProjectConnectionSpeedAssignedEnumLight     = ProjectConnectionSpeedAssignedEnum("LIGHT")
	ProjectConnectionSpeedAssignedEnumSerious   = ProjectConnectionSpeedAssignedEnum("SERIOUS")
	ProjectConnectionSpeedAssignedEnumThrottled = ProjectConnectionSpeedAssignedEnum("THROTTLED")
	ProjectConnectionSpeedAssignedEnumDisabled  = ProjectConnectionSpeedAssignedEnum("DISABLED")
	ProjectConnectionSpeedAssignedEnumUnlimited = ProjectConnectionSpeedAssignedEnum("UNLIMITED")
)

func NewProjectConnectionSpeedAssignedEnumFromString(value string) (out ProjectConnectionSpeedAssignedEnum, err error) {
	return ProjectConnectionSpeedAssignedEnum(value), nil
}

func (enum ProjectConnectionSpeedAssignedEnum) String() string {
	return string(enum)
}

func (enum ProjectConnectionSpeedAssignedEnum) Native() string {
	return string(enum)
}

func (enum ProjectConnectionSpeedAssignedEnum) Values() []ProjectConnectionSpeedAssignedEnum {
	return ProjectConnectionSpeedAssignedEnumAll()
}

func (enum ProjectConnectionSpeedAssignedEnum) PublicValues() []ProjectConnectionSpeedAssignedEnum {
	return ProjectConnectionSpeedAssignedEnumAllPublic()
}

func (enum ProjectConnectionSpeedAssignedEnum) PrivateValues() []ProjectConnectionSpeedAssignedEnum {
	return ProjectConnectionSpeedAssignedEnumAllPrivate()
}

func (enum ProjectConnectionSpeedAssignedEnum) DefaultValue() ProjectConnectionSpeedAssignedEnum {
	return ProjectConnectionSpeedAssignedEnumDefault()
}

func (enum ProjectConnectionSpeedAssignedEnum) Is(values ...ProjectConnectionSpeedAssignedEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ProjectConnectionSpeedAssignedEnumAllStrings() []string {
	return []string{
		string(ProjectConnectionSpeedAssignedEnumLight), string(ProjectConnectionSpeedAssignedEnumSerious), string(ProjectConnectionSpeedAssignedEnumThrottled), string(ProjectConnectionSpeedAssignedEnumDisabled), string(ProjectConnectionSpeedAssignedEnumUnlimited),
	}
}

func ProjectConnectionSpeedAssignedEnumAll() []ProjectConnectionSpeedAssignedEnum {
	return []ProjectConnectionSpeedAssignedEnum{
		ProjectConnectionSpeedAssignedEnumLight, ProjectConnectionSpeedAssignedEnumSerious, ProjectConnectionSpeedAssignedEnumThrottled, ProjectConnectionSpeedAssignedEnumDisabled, ProjectConnectionSpeedAssignedEnumUnlimited,
	}
}

func ProjectConnectionSpeedAssignedEnumAllPublic() []ProjectConnectionSpeedAssignedEnum {
	return []ProjectConnectionSpeedAssignedEnum{
		ProjectConnectionSpeedAssignedEnumLight, ProjectConnectionSpeedAssignedEnumSerious, ProjectConnectionSpeedAssignedEnumThrottled, ProjectConnectionSpeedAssignedEnumDisabled, ProjectConnectionSpeedAssignedEnumUnlimited,
	}
}

func ProjectConnectionSpeedAssignedEnumAllPrivate() []ProjectConnectionSpeedAssignedEnum {
	return []ProjectConnectionSpeedAssignedEnum{}
}

func ProjectConnectionSpeedAssignedEnumDefault() ProjectConnectionSpeedAssignedEnum {
	return ProjectConnectionSpeedAssignedEnumLight
}

func (enum ProjectConnectionSpeedAssignedEnum) IsLight() bool {
	return enum.Is(ProjectConnectionSpeedAssignedEnumLight)
}

func (enum ProjectConnectionSpeedAssignedEnum) IsSerious() bool {
	return enum.Is(ProjectConnectionSpeedAssignedEnumSerious)
}

func (enum ProjectConnectionSpeedAssignedEnum) IsThrottled() bool {
	return enum.Is(ProjectConnectionSpeedAssignedEnumThrottled)
}

func (enum ProjectConnectionSpeedAssignedEnum) IsDisabled() bool {
	return enum.Is(ProjectConnectionSpeedAssignedEnumDisabled)
}

func (enum ProjectConnectionSpeedAssignedEnum) IsUnlimited() bool {
	return enum.Is(ProjectConnectionSpeedAssignedEnumUnlimited)
}
