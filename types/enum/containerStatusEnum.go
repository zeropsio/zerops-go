// Generated ZEROPS sdk

package enum

type ContainerStatusEnum string

const (
	ContainerStatusEnumCreating      = ContainerStatusEnum("CREATING")
	ContainerStatusEnumActive        = ContainerStatusEnum("ACTIVE")
	ContainerStatusEnumDeleting      = ContainerStatusEnum("DELETING")
	ContainerStatusEnumFailed        = ContainerStatusEnum("FAILED")
	ContainerStatusEnumStopping      = ContainerStatusEnum("STOPPING")
	ContainerStatusEnumStopped       = ContainerStatusEnum("STOPPED")
	ContainerStatusEnumStarting      = ContainerStatusEnum("STARTING")
	ContainerStatusEnumRestarting    = ContainerStatusEnum("RESTARTING")
	ContainerStatusEnumReloading     = ContainerStatusEnum("RELOADING")
	ContainerStatusEnumMoving        = ContainerStatusEnum("MOVING")
	ContainerStatusEnumActionFailed  = ContainerStatusEnum("ACTION_FAILED")
	ContainerStatusEnumMovingStopped = ContainerStatusEnum("MOVING_STOPPED")
)

func NewContainerStatusEnumFromString(value string) (out ContainerStatusEnum, err error) {
	return ContainerStatusEnum(value), nil
}

func (enum ContainerStatusEnum) String() string {
	return string(enum)
}

func (enum ContainerStatusEnum) Native() string {
	return string(enum)
}

func (enum ContainerStatusEnum) Values() []ContainerStatusEnum {
	return ContainerStatusEnumAll()
}

func (enum ContainerStatusEnum) PublicValues() []ContainerStatusEnum {
	return ContainerStatusEnumAllPublic()
}

func (enum ContainerStatusEnum) PrivateValues() []ContainerStatusEnum {
	return ContainerStatusEnumAllPrivate()
}

func (enum ContainerStatusEnum) DefaultValue() ContainerStatusEnum {
	return ContainerStatusEnumDefault()
}

func (enum ContainerStatusEnum) Is(values ...ContainerStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ContainerStatusEnumAllStrings() []string {
	return []string{
		string(ContainerStatusEnumCreating), string(ContainerStatusEnumActive), string(ContainerStatusEnumDeleting), string(ContainerStatusEnumFailed), string(ContainerStatusEnumStopping), string(ContainerStatusEnumStopped), string(ContainerStatusEnumStarting), string(ContainerStatusEnumRestarting), string(ContainerStatusEnumReloading), string(ContainerStatusEnumMoving), string(ContainerStatusEnumActionFailed), string(ContainerStatusEnumMovingStopped),
	}
}

func ContainerStatusEnumAll() []ContainerStatusEnum {
	return []ContainerStatusEnum{
		ContainerStatusEnumCreating, ContainerStatusEnumActive, ContainerStatusEnumDeleting, ContainerStatusEnumFailed, ContainerStatusEnumStopping, ContainerStatusEnumStopped, ContainerStatusEnumStarting, ContainerStatusEnumRestarting, ContainerStatusEnumReloading, ContainerStatusEnumMoving, ContainerStatusEnumActionFailed, ContainerStatusEnumMovingStopped,
	}
}

func ContainerStatusEnumAllPublic() []ContainerStatusEnum {
	return []ContainerStatusEnum{
		ContainerStatusEnumCreating, ContainerStatusEnumActive, ContainerStatusEnumDeleting, ContainerStatusEnumFailed, ContainerStatusEnumStopping, ContainerStatusEnumStopped, ContainerStatusEnumStarting, ContainerStatusEnumRestarting, ContainerStatusEnumReloading, ContainerStatusEnumMoving, ContainerStatusEnumActionFailed, ContainerStatusEnumMovingStopped,
	}
}

func ContainerStatusEnumAllPrivate() []ContainerStatusEnum {
	return []ContainerStatusEnum{}
}

func ContainerStatusEnumDefault() ContainerStatusEnum {
	return ContainerStatusEnumCreating
}

func (enum ContainerStatusEnum) IsCreating() bool {
	return enum.Is(ContainerStatusEnumCreating)
}

func (enum ContainerStatusEnum) IsActive() bool {
	return enum.Is(ContainerStatusEnumActive)
}

func (enum ContainerStatusEnum) IsDeleting() bool {
	return enum.Is(ContainerStatusEnumDeleting)
}

func (enum ContainerStatusEnum) IsFailed() bool {
	return enum.Is(ContainerStatusEnumFailed)
}

func (enum ContainerStatusEnum) IsStopping() bool {
	return enum.Is(ContainerStatusEnumStopping)
}

func (enum ContainerStatusEnum) IsStopped() bool {
	return enum.Is(ContainerStatusEnumStopped)
}

func (enum ContainerStatusEnum) IsStarting() bool {
	return enum.Is(ContainerStatusEnumStarting)
}

func (enum ContainerStatusEnum) IsRestarting() bool {
	return enum.Is(ContainerStatusEnumRestarting)
}

func (enum ContainerStatusEnum) IsReloading() bool {
	return enum.Is(ContainerStatusEnumReloading)
}

func (enum ContainerStatusEnum) IsMoving() bool {
	return enum.Is(ContainerStatusEnumMoving)
}

func (enum ContainerStatusEnum) IsActionFailed() bool {
	return enum.Is(ContainerStatusEnumActionFailed)
}

func (enum ContainerStatusEnum) IsMovingStopped() bool {
	return enum.Is(ContainerStatusEnumMovingStopped)
}
