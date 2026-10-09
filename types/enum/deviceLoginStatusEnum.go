// Generated ZEROPS sdk

package enum

type DeviceLoginStatusEnum string

const (
	DeviceLoginStatusEnumPending  = DeviceLoginStatusEnum("PENDING")
	DeviceLoginStatusEnumApproved = DeviceLoginStatusEnum("APPROVED")
	DeviceLoginStatusEnumDenied   = DeviceLoginStatusEnum("DENIED")
)

func NewDeviceLoginStatusEnumFromString(value string) (out DeviceLoginStatusEnum, err error) {
	return DeviceLoginStatusEnum(value), nil
}

func (enum DeviceLoginStatusEnum) String() string {
	return string(enum)
}

func (enum DeviceLoginStatusEnum) Native() string {
	return string(enum)
}

func (enum DeviceLoginStatusEnum) Values() []DeviceLoginStatusEnum {
	return DeviceLoginStatusEnumAll()
}

func (enum DeviceLoginStatusEnum) PublicValues() []DeviceLoginStatusEnum {
	return DeviceLoginStatusEnumAllPublic()
}

func (enum DeviceLoginStatusEnum) PrivateValues() []DeviceLoginStatusEnum {
	return DeviceLoginStatusEnumAllPrivate()
}

func (enum DeviceLoginStatusEnum) DefaultValue() DeviceLoginStatusEnum {
	return DeviceLoginStatusEnumDefault()
}

func (enum DeviceLoginStatusEnum) Is(values ...DeviceLoginStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func DeviceLoginStatusEnumAllStrings() []string {
	return []string{
		string(DeviceLoginStatusEnumPending), string(DeviceLoginStatusEnumApproved), string(DeviceLoginStatusEnumDenied),
	}
}

func DeviceLoginStatusEnumAll() []DeviceLoginStatusEnum {
	return []DeviceLoginStatusEnum{
		DeviceLoginStatusEnumPending, DeviceLoginStatusEnumApproved, DeviceLoginStatusEnumDenied,
	}
}

func DeviceLoginStatusEnumAllPublic() []DeviceLoginStatusEnum {
	return []DeviceLoginStatusEnum{
		DeviceLoginStatusEnumPending, DeviceLoginStatusEnumApproved, DeviceLoginStatusEnumDenied,
	}
}

func DeviceLoginStatusEnumAllPrivate() []DeviceLoginStatusEnum {
	return []DeviceLoginStatusEnum{}
}

func DeviceLoginStatusEnumDefault() DeviceLoginStatusEnum {
	return DeviceLoginStatusEnumPending
}

func (enum DeviceLoginStatusEnum) IsPending() bool {
	return enum.Is(DeviceLoginStatusEnumPending)
}

func (enum DeviceLoginStatusEnum) IsApproved() bool {
	return enum.Is(DeviceLoginStatusEnumApproved)
}

func (enum DeviceLoginStatusEnum) IsDenied() bool {
	return enum.Is(DeviceLoginStatusEnumDenied)
}
