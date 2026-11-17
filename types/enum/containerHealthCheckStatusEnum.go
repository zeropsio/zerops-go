// Generated ZEROPS sdk

package enum

type ContainerHealthCheckStatusEnum string

const (
	ContainerHealthCheckStatusEnumNotConfigured = ContainerHealthCheckStatusEnum("NOT_CONFIGURED")
	ContainerHealthCheckStatusEnumDisabled      = ContainerHealthCheckStatusEnum("DISABLED")
	ContainerHealthCheckStatusEnumPending       = ContainerHealthCheckStatusEnum("PENDING")
	ContainerHealthCheckStatusEnumPassing       = ContainerHealthCheckStatusEnum("PASSING")
	ContainerHealthCheckStatusEnumFailing       = ContainerHealthCheckStatusEnum("FAILING")
	ContainerHealthCheckStatusEnumFailed        = ContainerHealthCheckStatusEnum("FAILED")
)

func NewContainerHealthCheckStatusEnumFromString(value string) (out ContainerHealthCheckStatusEnum, err error) {
	return ContainerHealthCheckStatusEnum(value), nil
}

func (enum ContainerHealthCheckStatusEnum) String() string {
	return string(enum)
}

func (enum ContainerHealthCheckStatusEnum) Native() string {
	return string(enum)
}

func (enum ContainerHealthCheckStatusEnum) Values() []ContainerHealthCheckStatusEnum {
	return ContainerHealthCheckStatusEnumAll()
}

func (enum ContainerHealthCheckStatusEnum) PublicValues() []ContainerHealthCheckStatusEnum {
	return ContainerHealthCheckStatusEnumAllPublic()
}

func (enum ContainerHealthCheckStatusEnum) PrivateValues() []ContainerHealthCheckStatusEnum {
	return ContainerHealthCheckStatusEnumAllPrivate()
}

func (enum ContainerHealthCheckStatusEnum) DefaultValue() ContainerHealthCheckStatusEnum {
	return ContainerHealthCheckStatusEnumDefault()
}

func (enum ContainerHealthCheckStatusEnum) Is(values ...ContainerHealthCheckStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ContainerHealthCheckStatusEnumAllStrings() []string {
	return []string{
		string(ContainerHealthCheckStatusEnumNotConfigured), string(ContainerHealthCheckStatusEnumDisabled), string(ContainerHealthCheckStatusEnumPending), string(ContainerHealthCheckStatusEnumPassing), string(ContainerHealthCheckStatusEnumFailing), string(ContainerHealthCheckStatusEnumFailed),
	}
}

func ContainerHealthCheckStatusEnumAll() []ContainerHealthCheckStatusEnum {
	return []ContainerHealthCheckStatusEnum{
		ContainerHealthCheckStatusEnumNotConfigured, ContainerHealthCheckStatusEnumDisabled, ContainerHealthCheckStatusEnumPending, ContainerHealthCheckStatusEnumPassing, ContainerHealthCheckStatusEnumFailing, ContainerHealthCheckStatusEnumFailed,
	}
}

func ContainerHealthCheckStatusEnumAllPublic() []ContainerHealthCheckStatusEnum {
	return []ContainerHealthCheckStatusEnum{
		ContainerHealthCheckStatusEnumNotConfigured, ContainerHealthCheckStatusEnumDisabled, ContainerHealthCheckStatusEnumPending, ContainerHealthCheckStatusEnumPassing, ContainerHealthCheckStatusEnumFailing, ContainerHealthCheckStatusEnumFailed,
	}
}

func ContainerHealthCheckStatusEnumAllPrivate() []ContainerHealthCheckStatusEnum {
	return []ContainerHealthCheckStatusEnum{}
}

func ContainerHealthCheckStatusEnumDefault() ContainerHealthCheckStatusEnum {
	return ""
}

func (enum ContainerHealthCheckStatusEnum) IsNotConfigured() bool {
	return enum.Is(ContainerHealthCheckStatusEnumNotConfigured)
}

func (enum ContainerHealthCheckStatusEnum) IsDisabled() bool {
	return enum.Is(ContainerHealthCheckStatusEnumDisabled)
}

func (enum ContainerHealthCheckStatusEnum) IsPending() bool {
	return enum.Is(ContainerHealthCheckStatusEnumPending)
}

func (enum ContainerHealthCheckStatusEnum) IsPassing() bool {
	return enum.Is(ContainerHealthCheckStatusEnumPassing)
}

func (enum ContainerHealthCheckStatusEnum) IsFailing() bool {
	return enum.Is(ContainerHealthCheckStatusEnumFailing)
}

func (enum ContainerHealthCheckStatusEnum) IsFailed() bool {
	return enum.Is(ContainerHealthCheckStatusEnumFailed)
}
