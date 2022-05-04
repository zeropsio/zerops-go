// Generated ZEROPS sdk

package enum

type ServiceStatusEnum string

const (
	ServiceStatusEnumCreating        = ServiceStatusEnum("CREATING")
	ServiceStatusEnumActive          = ServiceStatusEnum("ACTIVE")
	ServiceStatusEnumStopping        = ServiceStatusEnum("STOPPING")
	ServiceStatusEnumStopped         = ServiceStatusEnum("STOPPED")
	ServiceStatusEnumStarting        = ServiceStatusEnum("STARTING")
	ServiceStatusEnumRestarting      = ServiceStatusEnum("RESTARTING")
	ServiceStatusEnumReloading       = ServiceStatusEnum("RELOADING")
	ServiceStatusEnumDeleting        = ServiceStatusEnum("DELETING")
	ServiceStatusEnumDeleted         = ServiceStatusEnum("DELETED")
	ServiceStatusEnumFailed          = ServiceStatusEnum("FAILED")
	ServiceStatusEnumActionFailed    = ServiceStatusEnum("ACTION_FAILED")
	ServiceStatusEnumRepairing       = ServiceStatusEnum("REPAIRING")
	ServiceStatusEnumContainerFailed = ServiceStatusEnum("CONTAINER_FAILED")
	ServiceStatusEnumMovingContainer = ServiceStatusEnum("MOVING_CONTAINER")
	ServiceStatusEnumUpgrading       = ServiceStatusEnum("UPGRADING")
	ServiceStatusEnumScaling         = ServiceStatusEnum("SCALING")
	ServiceStatusEnumRepairFailed    = ServiceStatusEnum("REPAIR_FAILED")
)

func NewServiceStatusEnumFromString(value string) (out ServiceStatusEnum, err error) {
	return ServiceStatusEnum(value), nil
}

func (enum ServiceStatusEnum) String() string {
	return string(enum)
}

func (enum ServiceStatusEnum) Native() string {
	return string(enum)
}

func (enum ServiceStatusEnum) Values() []ServiceStatusEnum {
	return ServiceStatusEnumAll()
}

func (enum ServiceStatusEnum) PublicValues() []ServiceStatusEnum {
	return ServiceStatusEnumAllPublic()
}

func (enum ServiceStatusEnum) PrivateValues() []ServiceStatusEnum {
	return ServiceStatusEnumAllPrivate()
}

func (enum ServiceStatusEnum) DefaultValue() ServiceStatusEnum {
	return ServiceStatusEnumDefault()
}

func (enum ServiceStatusEnum) Is(values ...ServiceStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ServiceStatusEnumAllStrings() []string {
	return []string{
		string(ServiceStatusEnumCreating), string(ServiceStatusEnumActive), string(ServiceStatusEnumStopping), string(ServiceStatusEnumStopped), string(ServiceStatusEnumStarting), string(ServiceStatusEnumRestarting), string(ServiceStatusEnumReloading), string(ServiceStatusEnumDeleting), string(ServiceStatusEnumDeleted), string(ServiceStatusEnumFailed), string(ServiceStatusEnumActionFailed), string(ServiceStatusEnumRepairing), string(ServiceStatusEnumContainerFailed), string(ServiceStatusEnumMovingContainer), string(ServiceStatusEnumUpgrading), string(ServiceStatusEnumScaling), string(ServiceStatusEnumRepairFailed),
	}
}

func ServiceStatusEnumAll() []ServiceStatusEnum {
	return []ServiceStatusEnum{
		ServiceStatusEnumCreating, ServiceStatusEnumActive, ServiceStatusEnumStopping, ServiceStatusEnumStopped, ServiceStatusEnumStarting, ServiceStatusEnumRestarting, ServiceStatusEnumReloading, ServiceStatusEnumDeleting, ServiceStatusEnumDeleted, ServiceStatusEnumFailed, ServiceStatusEnumActionFailed, ServiceStatusEnumRepairing, ServiceStatusEnumContainerFailed, ServiceStatusEnumMovingContainer, ServiceStatusEnumUpgrading, ServiceStatusEnumScaling, ServiceStatusEnumRepairFailed,
	}
}

func ServiceStatusEnumAllPublic() []ServiceStatusEnum {
	return []ServiceStatusEnum{
		ServiceStatusEnumCreating, ServiceStatusEnumActive, ServiceStatusEnumStopping, ServiceStatusEnumStopped, ServiceStatusEnumStarting, ServiceStatusEnumRestarting, ServiceStatusEnumReloading, ServiceStatusEnumDeleting, ServiceStatusEnumDeleted, ServiceStatusEnumFailed, ServiceStatusEnumActionFailed, ServiceStatusEnumRepairing, ServiceStatusEnumContainerFailed, ServiceStatusEnumMovingContainer, ServiceStatusEnumUpgrading, ServiceStatusEnumScaling, ServiceStatusEnumRepairFailed,
	}
}

func ServiceStatusEnumAllPrivate() []ServiceStatusEnum {
	return []ServiceStatusEnum{}
}

func ServiceStatusEnumDefault() ServiceStatusEnum {
	return ServiceStatusEnumCreating
}

func (enum ServiceStatusEnum) IsCreating() bool {
	return enum.Is(ServiceStatusEnumCreating)
}

func (enum ServiceStatusEnum) IsActive() bool {
	return enum.Is(ServiceStatusEnumActive)
}

func (enum ServiceStatusEnum) IsStopping() bool {
	return enum.Is(ServiceStatusEnumStopping)
}

func (enum ServiceStatusEnum) IsStopped() bool {
	return enum.Is(ServiceStatusEnumStopped)
}

func (enum ServiceStatusEnum) IsStarting() bool {
	return enum.Is(ServiceStatusEnumStarting)
}

func (enum ServiceStatusEnum) IsRestarting() bool {
	return enum.Is(ServiceStatusEnumRestarting)
}

func (enum ServiceStatusEnum) IsReloading() bool {
	return enum.Is(ServiceStatusEnumReloading)
}

func (enum ServiceStatusEnum) IsDeleting() bool {
	return enum.Is(ServiceStatusEnumDeleting)
}

func (enum ServiceStatusEnum) IsDeleted() bool {
	return enum.Is(ServiceStatusEnumDeleted)
}

func (enum ServiceStatusEnum) IsFailed() bool {
	return enum.Is(ServiceStatusEnumFailed)
}

func (enum ServiceStatusEnum) IsActionFailed() bool {
	return enum.Is(ServiceStatusEnumActionFailed)
}

func (enum ServiceStatusEnum) IsRepairing() bool {
	return enum.Is(ServiceStatusEnumRepairing)
}

func (enum ServiceStatusEnum) IsContainerFailed() bool {
	return enum.Is(ServiceStatusEnumContainerFailed)
}

func (enum ServiceStatusEnum) IsMovingContainer() bool {
	return enum.Is(ServiceStatusEnumMovingContainer)
}

func (enum ServiceStatusEnum) IsUpgrading() bool {
	return enum.Is(ServiceStatusEnumUpgrading)
}

func (enum ServiceStatusEnum) IsScaling() bool {
	return enum.Is(ServiceStatusEnumScaling)
}

func (enum ServiceStatusEnum) IsRepairFailed() bool {
	return enum.Is(ServiceStatusEnumRepairFailed)
}
