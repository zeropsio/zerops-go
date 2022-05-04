// Generated ZEROPS sdk

package enum

type ServiceStackStatusEnum string

const (
	ServiceStackStatusEnumNew                    = ServiceStackStatusEnum("NEW")
	ServiceStackStatusEnumCreating               = ServiceStackStatusEnum("CREATING")
	ServiceStackStatusEnumActive                 = ServiceStackStatusEnum("ACTIVE")
	ServiceStackStatusEnumStopping               = ServiceStackStatusEnum("STOPPING")
	ServiceStackStatusEnumStopped                = ServiceStackStatusEnum("STOPPED")
	ServiceStackStatusEnumStarting               = ServiceStackStatusEnum("STARTING")
	ServiceStackStatusEnumRestarting             = ServiceStackStatusEnum("RESTARTING")
	ServiceStackStatusEnumReloading              = ServiceStackStatusEnum("RELOADING")
	ServiceStackStatusEnumDeleting               = ServiceStackStatusEnum("DELETING")
	ServiceStackStatusEnumDeleted                = ServiceStackStatusEnum("DELETED")
	ServiceStackStatusEnumFailed                 = ServiceStackStatusEnum("FAILED")
	ServiceStackStatusEnumActionFailed           = ServiceStackStatusEnum("ACTION_FAILED")
	ServiceStackStatusEnumUpgrading              = ServiceStackStatusEnum("UPGRADING")
	ServiceStackStatusEnumReadyToDeploy          = ServiceStackStatusEnum("READY_TO_DEPLOY")
	ServiceStackStatusEnumServiceCreating        = ServiceStackStatusEnum("SERVICE_CREATING")
	ServiceStackStatusEnumServiceActive          = ServiceStackStatusEnum("SERVICE_ACTIVE")
	ServiceStackStatusEnumServiceStopping        = ServiceStackStatusEnum("SERVICE_STOPPING")
	ServiceStackStatusEnumServiceStopped         = ServiceStackStatusEnum("SERVICE_STOPPED")
	ServiceStackStatusEnumServiceStarting        = ServiceStackStatusEnum("SERVICE_STARTING")
	ServiceStackStatusEnumServiceRestarting      = ServiceStackStatusEnum("SERVICE_RESTARTING")
	ServiceStackStatusEnumServiceReloading       = ServiceStackStatusEnum("SERVICE_RELOADING")
	ServiceStackStatusEnumServiceDeleting        = ServiceStackStatusEnum("SERVICE_DELETING")
	ServiceStackStatusEnumServiceDeleted         = ServiceStackStatusEnum("SERVICE_DELETED")
	ServiceStackStatusEnumServiceFailed          = ServiceStackStatusEnum("SERVICE_FAILED")
	ServiceStackStatusEnumServiceActionFailed    = ServiceStackStatusEnum("SERVICE_ACTION_FAILED")
	ServiceStackStatusEnumServiceRepairing       = ServiceStackStatusEnum("SERVICE_REPAIRING")
	ServiceStackStatusEnumServiceContainerFailed = ServiceStackStatusEnum("SERVICE_CONTAINER_FAILED")
	ServiceStackStatusEnumServiceMovingContainer = ServiceStackStatusEnum("SERVICE_MOVING_CONTAINER")
	ServiceStackStatusEnumServiceUpgrading       = ServiceStackStatusEnum("SERVICE_UPGRADING")
	ServiceStackStatusEnumServiceScaling         = ServiceStackStatusEnum("SERVICE_SCALING")
	ServiceStackStatusEnumServiceRepairFailed    = ServiceStackStatusEnum("SERVICE_REPAIR_FAILED")
	ServiceStackStatusEnumRepairing              = ServiceStackStatusEnum("REPAIRING")
	ServiceStackStatusEnumContainerFailed        = ServiceStackStatusEnum("CONTAINER_FAILED")
	ServiceStackStatusEnumMovingContainer        = ServiceStackStatusEnum("MOVING_CONTAINER")
	ServiceStackStatusEnumScaling                = ServiceStackStatusEnum("SCALING")
	ServiceStackStatusEnumRepairFailed           = ServiceStackStatusEnum("REPAIR_FAILED")
)

func NewServiceStackStatusEnumFromString(value string) (out ServiceStackStatusEnum, err error) {
	return ServiceStackStatusEnum(value), nil
}

func (enum ServiceStackStatusEnum) String() string {
	return string(enum)
}

func (enum ServiceStackStatusEnum) Native() string {
	return string(enum)
}

func (enum ServiceStackStatusEnum) Values() []ServiceStackStatusEnum {
	return ServiceStackStatusEnumAll()
}

func (enum ServiceStackStatusEnum) PublicValues() []ServiceStackStatusEnum {
	return ServiceStackStatusEnumAllPublic()
}

func (enum ServiceStackStatusEnum) PrivateValues() []ServiceStackStatusEnum {
	return ServiceStackStatusEnumAllPrivate()
}

func (enum ServiceStackStatusEnum) DefaultValue() ServiceStackStatusEnum {
	return ServiceStackStatusEnumDefault()
}

func (enum ServiceStackStatusEnum) Is(values ...ServiceStackStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ServiceStackStatusEnumAllStrings() []string {
	return []string{
		string(ServiceStackStatusEnumNew), string(ServiceStackStatusEnumCreating), string(ServiceStackStatusEnumActive), string(ServiceStackStatusEnumStopping), string(ServiceStackStatusEnumStopped), string(ServiceStackStatusEnumStarting), string(ServiceStackStatusEnumRestarting), string(ServiceStackStatusEnumReloading), string(ServiceStackStatusEnumDeleting), string(ServiceStackStatusEnumDeleted), string(ServiceStackStatusEnumFailed), string(ServiceStackStatusEnumActionFailed), string(ServiceStackStatusEnumUpgrading), string(ServiceStackStatusEnumReadyToDeploy), string(ServiceStackStatusEnumServiceCreating), string(ServiceStackStatusEnumServiceActive), string(ServiceStackStatusEnumServiceStopping), string(ServiceStackStatusEnumServiceStopped), string(ServiceStackStatusEnumServiceStarting), string(ServiceStackStatusEnumServiceRestarting), string(ServiceStackStatusEnumServiceReloading), string(ServiceStackStatusEnumServiceDeleting), string(ServiceStackStatusEnumServiceDeleted), string(ServiceStackStatusEnumServiceFailed), string(ServiceStackStatusEnumServiceActionFailed), string(ServiceStackStatusEnumServiceRepairing), string(ServiceStackStatusEnumServiceContainerFailed), string(ServiceStackStatusEnumServiceMovingContainer), string(ServiceStackStatusEnumServiceUpgrading), string(ServiceStackStatusEnumServiceScaling), string(ServiceStackStatusEnumServiceRepairFailed), string(ServiceStackStatusEnumRepairing), string(ServiceStackStatusEnumContainerFailed), string(ServiceStackStatusEnumMovingContainer), string(ServiceStackStatusEnumScaling), string(ServiceStackStatusEnumRepairFailed),
	}
}

func ServiceStackStatusEnumAll() []ServiceStackStatusEnum {
	return []ServiceStackStatusEnum{
		ServiceStackStatusEnumNew, ServiceStackStatusEnumCreating, ServiceStackStatusEnumActive, ServiceStackStatusEnumStopping, ServiceStackStatusEnumStopped, ServiceStackStatusEnumStarting, ServiceStackStatusEnumRestarting, ServiceStackStatusEnumReloading, ServiceStackStatusEnumDeleting, ServiceStackStatusEnumDeleted, ServiceStackStatusEnumFailed, ServiceStackStatusEnumActionFailed, ServiceStackStatusEnumUpgrading, ServiceStackStatusEnumReadyToDeploy, ServiceStackStatusEnumServiceCreating, ServiceStackStatusEnumServiceActive, ServiceStackStatusEnumServiceStopping, ServiceStackStatusEnumServiceStopped, ServiceStackStatusEnumServiceStarting, ServiceStackStatusEnumServiceRestarting, ServiceStackStatusEnumServiceReloading, ServiceStackStatusEnumServiceDeleting, ServiceStackStatusEnumServiceDeleted, ServiceStackStatusEnumServiceFailed, ServiceStackStatusEnumServiceActionFailed, ServiceStackStatusEnumServiceRepairing, ServiceStackStatusEnumServiceContainerFailed, ServiceStackStatusEnumServiceMovingContainer, ServiceStackStatusEnumServiceUpgrading, ServiceStackStatusEnumServiceScaling, ServiceStackStatusEnumServiceRepairFailed, ServiceStackStatusEnumRepairing, ServiceStackStatusEnumContainerFailed, ServiceStackStatusEnumMovingContainer, ServiceStackStatusEnumScaling, ServiceStackStatusEnumRepairFailed,
	}
}

func ServiceStackStatusEnumAllPublic() []ServiceStackStatusEnum {
	return []ServiceStackStatusEnum{
		ServiceStackStatusEnumNew, ServiceStackStatusEnumCreating, ServiceStackStatusEnumActive, ServiceStackStatusEnumStopping, ServiceStackStatusEnumStopped, ServiceStackStatusEnumStarting, ServiceStackStatusEnumRestarting, ServiceStackStatusEnumReloading, ServiceStackStatusEnumDeleting, ServiceStackStatusEnumDeleted, ServiceStackStatusEnumFailed, ServiceStackStatusEnumActionFailed, ServiceStackStatusEnumUpgrading, ServiceStackStatusEnumReadyToDeploy, ServiceStackStatusEnumServiceCreating, ServiceStackStatusEnumServiceActive, ServiceStackStatusEnumServiceStopping, ServiceStackStatusEnumServiceStopped, ServiceStackStatusEnumServiceStarting, ServiceStackStatusEnumServiceRestarting, ServiceStackStatusEnumServiceReloading, ServiceStackStatusEnumServiceDeleting, ServiceStackStatusEnumServiceDeleted, ServiceStackStatusEnumServiceFailed, ServiceStackStatusEnumServiceActionFailed, ServiceStackStatusEnumServiceRepairing, ServiceStackStatusEnumServiceContainerFailed, ServiceStackStatusEnumServiceMovingContainer, ServiceStackStatusEnumServiceUpgrading, ServiceStackStatusEnumServiceScaling, ServiceStackStatusEnumServiceRepairFailed, ServiceStackStatusEnumRepairing, ServiceStackStatusEnumContainerFailed, ServiceStackStatusEnumMovingContainer, ServiceStackStatusEnumScaling, ServiceStackStatusEnumRepairFailed,
	}
}

func ServiceStackStatusEnumAllPrivate() []ServiceStackStatusEnum {
	return []ServiceStackStatusEnum{}
}

func ServiceStackStatusEnumDefault() ServiceStackStatusEnum {
	return ServiceStackStatusEnumNew
}

func (enum ServiceStackStatusEnum) IsNew() bool {
	return enum.Is(ServiceStackStatusEnumNew)
}

func (enum ServiceStackStatusEnum) IsCreating() bool {
	return enum.Is(ServiceStackStatusEnumCreating)
}

func (enum ServiceStackStatusEnum) IsActive() bool {
	return enum.Is(ServiceStackStatusEnumActive)
}

func (enum ServiceStackStatusEnum) IsStopping() bool {
	return enum.Is(ServiceStackStatusEnumStopping)
}

func (enum ServiceStackStatusEnum) IsStopped() bool {
	return enum.Is(ServiceStackStatusEnumStopped)
}

func (enum ServiceStackStatusEnum) IsStarting() bool {
	return enum.Is(ServiceStackStatusEnumStarting)
}

func (enum ServiceStackStatusEnum) IsRestarting() bool {
	return enum.Is(ServiceStackStatusEnumRestarting)
}

func (enum ServiceStackStatusEnum) IsReloading() bool {
	return enum.Is(ServiceStackStatusEnumReloading)
}

func (enum ServiceStackStatusEnum) IsDeleting() bool {
	return enum.Is(ServiceStackStatusEnumDeleting)
}

func (enum ServiceStackStatusEnum) IsDeleted() bool {
	return enum.Is(ServiceStackStatusEnumDeleted)
}

func (enum ServiceStackStatusEnum) IsFailed() bool {
	return enum.Is(ServiceStackStatusEnumFailed)
}

func (enum ServiceStackStatusEnum) IsActionFailed() bool {
	return enum.Is(ServiceStackStatusEnumActionFailed)
}

func (enum ServiceStackStatusEnum) IsUpgrading() bool {
	return enum.Is(ServiceStackStatusEnumUpgrading)
}

func (enum ServiceStackStatusEnum) IsReadyToDeploy() bool {
	return enum.Is(ServiceStackStatusEnumReadyToDeploy)
}

func (enum ServiceStackStatusEnum) IsServiceCreating() bool {
	return enum.Is(ServiceStackStatusEnumServiceCreating)
}

func (enum ServiceStackStatusEnum) IsServiceActive() bool {
	return enum.Is(ServiceStackStatusEnumServiceActive)
}

func (enum ServiceStackStatusEnum) IsServiceStopping() bool {
	return enum.Is(ServiceStackStatusEnumServiceStopping)
}

func (enum ServiceStackStatusEnum) IsServiceStopped() bool {
	return enum.Is(ServiceStackStatusEnumServiceStopped)
}

func (enum ServiceStackStatusEnum) IsServiceStarting() bool {
	return enum.Is(ServiceStackStatusEnumServiceStarting)
}

func (enum ServiceStackStatusEnum) IsServiceRestarting() bool {
	return enum.Is(ServiceStackStatusEnumServiceRestarting)
}

func (enum ServiceStackStatusEnum) IsServiceReloading() bool {
	return enum.Is(ServiceStackStatusEnumServiceReloading)
}

func (enum ServiceStackStatusEnum) IsServiceDeleting() bool {
	return enum.Is(ServiceStackStatusEnumServiceDeleting)
}

func (enum ServiceStackStatusEnum) IsServiceDeleted() bool {
	return enum.Is(ServiceStackStatusEnumServiceDeleted)
}

func (enum ServiceStackStatusEnum) IsServiceFailed() bool {
	return enum.Is(ServiceStackStatusEnumServiceFailed)
}

func (enum ServiceStackStatusEnum) IsServiceActionFailed() bool {
	return enum.Is(ServiceStackStatusEnumServiceActionFailed)
}

func (enum ServiceStackStatusEnum) IsServiceRepairing() bool {
	return enum.Is(ServiceStackStatusEnumServiceRepairing)
}

func (enum ServiceStackStatusEnum) IsServiceContainerFailed() bool {
	return enum.Is(ServiceStackStatusEnumServiceContainerFailed)
}

func (enum ServiceStackStatusEnum) IsServiceMovingContainer() bool {
	return enum.Is(ServiceStackStatusEnumServiceMovingContainer)
}

func (enum ServiceStackStatusEnum) IsServiceUpgrading() bool {
	return enum.Is(ServiceStackStatusEnumServiceUpgrading)
}

func (enum ServiceStackStatusEnum) IsServiceScaling() bool {
	return enum.Is(ServiceStackStatusEnumServiceScaling)
}

func (enum ServiceStackStatusEnum) IsServiceRepairFailed() bool {
	return enum.Is(ServiceStackStatusEnumServiceRepairFailed)
}

func (enum ServiceStackStatusEnum) IsRepairing() bool {
	return enum.Is(ServiceStackStatusEnumRepairing)
}

func (enum ServiceStackStatusEnum) IsContainerFailed() bool {
	return enum.Is(ServiceStackStatusEnumContainerFailed)
}

func (enum ServiceStackStatusEnum) IsMovingContainer() bool {
	return enum.Is(ServiceStackStatusEnumMovingContainer)
}

func (enum ServiceStackStatusEnum) IsScaling() bool {
	return enum.Is(ServiceStackStatusEnumScaling)
}

func (enum ServiceStackStatusEnum) IsRepairFailed() bool {
	return enum.Is(ServiceStackStatusEnumRepairFailed)
}
