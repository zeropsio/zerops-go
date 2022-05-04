// Generated ZEROPS sdk

package enum

type ProjectStatusEnum string

const (
	ProjectStatusEnumNew      = ProjectStatusEnum("NEW")
	ProjectStatusEnumCreating = ProjectStatusEnum("CREATING")
	ProjectStatusEnumActive   = ProjectStatusEnum("ACTIVE")
	ProjectStatusEnumDeleting = ProjectStatusEnum("DELETING")
	ProjectStatusEnumFailed   = ProjectStatusEnum("FAILED")
	ProjectStatusEnumStopping = ProjectStatusEnum("STOPPING")
	ProjectStatusEnumStopped  = ProjectStatusEnum("STOPPED")
	ProjectStatusEnumStarting = ProjectStatusEnum("STARTING")
)

func NewProjectStatusEnumFromString(value string) (out ProjectStatusEnum, err error) {
	return ProjectStatusEnum(value), nil
}

func (enum ProjectStatusEnum) String() string {
	return string(enum)
}

func (enum ProjectStatusEnum) Native() string {
	return string(enum)
}

func (enum ProjectStatusEnum) Values() []ProjectStatusEnum {
	return ProjectStatusEnumAll()
}

func (enum ProjectStatusEnum) PublicValues() []ProjectStatusEnum {
	return ProjectStatusEnumAllPublic()
}

func (enum ProjectStatusEnum) PrivateValues() []ProjectStatusEnum {
	return ProjectStatusEnumAllPrivate()
}

func (enum ProjectStatusEnum) DefaultValue() ProjectStatusEnum {
	return ProjectStatusEnumDefault()
}

func (enum ProjectStatusEnum) Is(values ...ProjectStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ProjectStatusEnumAllStrings() []string {
	return []string{
		string(ProjectStatusEnumNew), string(ProjectStatusEnumCreating), string(ProjectStatusEnumActive), string(ProjectStatusEnumDeleting), string(ProjectStatusEnumFailed), string(ProjectStatusEnumStopping), string(ProjectStatusEnumStopped), string(ProjectStatusEnumStarting),
	}
}

func ProjectStatusEnumAll() []ProjectStatusEnum {
	return []ProjectStatusEnum{
		ProjectStatusEnumNew, ProjectStatusEnumCreating, ProjectStatusEnumActive, ProjectStatusEnumDeleting, ProjectStatusEnumFailed, ProjectStatusEnumStopping, ProjectStatusEnumStopped, ProjectStatusEnumStarting,
	}
}

func ProjectStatusEnumAllPublic() []ProjectStatusEnum {
	return []ProjectStatusEnum{
		ProjectStatusEnumNew, ProjectStatusEnumCreating, ProjectStatusEnumActive, ProjectStatusEnumDeleting, ProjectStatusEnumFailed, ProjectStatusEnumStopping, ProjectStatusEnumStopped, ProjectStatusEnumStarting,
	}
}

func ProjectStatusEnumAllPrivate() []ProjectStatusEnum {
	return []ProjectStatusEnum{}
}

func ProjectStatusEnumDefault() ProjectStatusEnum {
	return ProjectStatusEnumNew
}

func (enum ProjectStatusEnum) IsNew() bool {
	return enum.Is(ProjectStatusEnumNew)
}

func (enum ProjectStatusEnum) IsCreating() bool {
	return enum.Is(ProjectStatusEnumCreating)
}

func (enum ProjectStatusEnum) IsActive() bool {
	return enum.Is(ProjectStatusEnumActive)
}

func (enum ProjectStatusEnum) IsDeleting() bool {
	return enum.Is(ProjectStatusEnumDeleting)
}

func (enum ProjectStatusEnum) IsFailed() bool {
	return enum.Is(ProjectStatusEnumFailed)
}

func (enum ProjectStatusEnum) IsStopping() bool {
	return enum.Is(ProjectStatusEnumStopping)
}

func (enum ProjectStatusEnum) IsStopped() bool {
	return enum.Is(ProjectStatusEnumStopped)
}

func (enum ProjectStatusEnum) IsStarting() bool {
	return enum.Is(ProjectStatusEnumStarting)
}
