// Generated ZEROPS sdk

package enum

type ClientUserLightRoleCodeEnum string

const (
	ClientUserLightRoleCodeEnumManager = ClientUserLightRoleCodeEnum("MANAGER")
)

func NewClientUserLightRoleCodeEnumFromString(value string) (out ClientUserLightRoleCodeEnum, err error) {
	return ClientUserLightRoleCodeEnum(value), nil
}

func (enum ClientUserLightRoleCodeEnum) String() string {
	return string(enum)
}

func (enum ClientUserLightRoleCodeEnum) Native() string {
	return string(enum)
}

func (enum ClientUserLightRoleCodeEnum) Values() []ClientUserLightRoleCodeEnum {
	return ClientUserLightRoleCodeEnumAll()
}

func (enum ClientUserLightRoleCodeEnum) PublicValues() []ClientUserLightRoleCodeEnum {
	return ClientUserLightRoleCodeEnumAllPublic()
}

func (enum ClientUserLightRoleCodeEnum) PrivateValues() []ClientUserLightRoleCodeEnum {
	return ClientUserLightRoleCodeEnumAllPrivate()
}

func (enum ClientUserLightRoleCodeEnum) DefaultValue() ClientUserLightRoleCodeEnum {
	return ClientUserLightRoleCodeEnumDefault()
}

func (enum ClientUserLightRoleCodeEnum) Is(values ...ClientUserLightRoleCodeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ClientUserLightRoleCodeEnumAllStrings() []string {
	return []string{
		string(ClientUserLightRoleCodeEnumManager),
	}
}

func ClientUserLightRoleCodeEnumAll() []ClientUserLightRoleCodeEnum {
	return []ClientUserLightRoleCodeEnum{
		ClientUserLightRoleCodeEnumManager,
	}
}

func ClientUserLightRoleCodeEnumAllPublic() []ClientUserLightRoleCodeEnum {
	return []ClientUserLightRoleCodeEnum{}
}

func ClientUserLightRoleCodeEnumAllPrivate() []ClientUserLightRoleCodeEnum {
	return []ClientUserLightRoleCodeEnum{
		ClientUserLightRoleCodeEnumManager,
	}
}

func ClientUserLightRoleCodeEnumDefault() ClientUserLightRoleCodeEnum {
	return ClientUserLightRoleCodeEnumManager
}

func (enum ClientUserLightRoleCodeEnum) IsManager() bool {
	return enum.Is(ClientUserLightRoleCodeEnumManager)
}
