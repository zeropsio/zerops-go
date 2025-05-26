// Generated ZEROPS sdk

package enum

type ProjectModeEnum string

const (
	ProjectModeEnumLegacy  = ProjectModeEnum("LEGACY")
	ProjectModeEnumLight   = ProjectModeEnum("LIGHT")
	ProjectModeEnumSerious = ProjectModeEnum("SERIOUS")
)

func NewProjectModeEnumFromString(value string) (out ProjectModeEnum, err error) {
	return ProjectModeEnum(value), nil
}

func (enum ProjectModeEnum) String() string {
	return string(enum)
}

func (enum ProjectModeEnum) Native() string {
	return string(enum)
}

func (enum ProjectModeEnum) Values() []ProjectModeEnum {
	return ProjectModeEnumAll()
}

func (enum ProjectModeEnum) PublicValues() []ProjectModeEnum {
	return ProjectModeEnumAllPublic()
}

func (enum ProjectModeEnum) PrivateValues() []ProjectModeEnum {
	return ProjectModeEnumAllPrivate()
}

func (enum ProjectModeEnum) DefaultValue() ProjectModeEnum {
	return ProjectModeEnumDefault()
}

func (enum ProjectModeEnum) Is(values ...ProjectModeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ProjectModeEnumAllStrings() []string {
	return []string{
		string(ProjectModeEnumLegacy), string(ProjectModeEnumLight), string(ProjectModeEnumSerious),
	}
}

func ProjectModeEnumAll() []ProjectModeEnum {
	return []ProjectModeEnum{
		ProjectModeEnumLegacy, ProjectModeEnumLight, ProjectModeEnumSerious,
	}
}

func ProjectModeEnumAllPublic() []ProjectModeEnum {
	return []ProjectModeEnum{
		ProjectModeEnumLight, ProjectModeEnumSerious,
	}
}

func ProjectModeEnumAllPrivate() []ProjectModeEnum {
	return []ProjectModeEnum{
		ProjectModeEnumLegacy,
	}
}

func ProjectModeEnumDefault() ProjectModeEnum {
	return ProjectModeEnumLegacy
}

func (enum ProjectModeEnum) IsLegacy() bool {
	return enum.Is(ProjectModeEnumLegacy)
}

func (enum ProjectModeEnum) IsLight() bool {
	return enum.Is(ProjectModeEnumLight)
}

func (enum ProjectModeEnum) IsSerious() bool {
	return enum.Is(ProjectModeEnumSerious)
}
