// Generated ZEROPS sdk

package enum

type DeviceLoginKindEnum string

const (
	DeviceLoginKindEnumLogin = DeviceLoginKindEnum("LOGIN")
	DeviceLoginKindEnumSudo  = DeviceLoginKindEnum("SUDO")
)

func NewDeviceLoginKindEnumFromString(value string) (out DeviceLoginKindEnum, err error) {
	return DeviceLoginKindEnum(value), nil
}

func (enum DeviceLoginKindEnum) String() string {
	return string(enum)
}

func (enum DeviceLoginKindEnum) Native() string {
	return string(enum)
}

func (enum DeviceLoginKindEnum) Values() []DeviceLoginKindEnum {
	return DeviceLoginKindEnumAll()
}

func (enum DeviceLoginKindEnum) PublicValues() []DeviceLoginKindEnum {
	return DeviceLoginKindEnumAllPublic()
}

func (enum DeviceLoginKindEnum) PrivateValues() []DeviceLoginKindEnum {
	return DeviceLoginKindEnumAllPrivate()
}

func (enum DeviceLoginKindEnum) DefaultValue() DeviceLoginKindEnum {
	return DeviceLoginKindEnumDefault()
}

func (enum DeviceLoginKindEnum) Is(values ...DeviceLoginKindEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func DeviceLoginKindEnumAllStrings() []string {
	return []string{
		string(DeviceLoginKindEnumLogin), string(DeviceLoginKindEnumSudo),
	}
}

func DeviceLoginKindEnumAll() []DeviceLoginKindEnum {
	return []DeviceLoginKindEnum{
		DeviceLoginKindEnumLogin, DeviceLoginKindEnumSudo,
	}
}

func DeviceLoginKindEnumAllPublic() []DeviceLoginKindEnum {
	return []DeviceLoginKindEnum{
		DeviceLoginKindEnumLogin, DeviceLoginKindEnumSudo,
	}
}

func DeviceLoginKindEnumAllPrivate() []DeviceLoginKindEnum {
	return []DeviceLoginKindEnum{}
}

func DeviceLoginKindEnumDefault() DeviceLoginKindEnum {
	return DeviceLoginKindEnumLogin
}

func (enum DeviceLoginKindEnum) IsLogin() bool {
	return enum.Is(DeviceLoginKindEnumLogin)
}

func (enum DeviceLoginKindEnum) IsSudo() bool {
	return enum.Is(DeviceLoginKindEnumSudo)
}
