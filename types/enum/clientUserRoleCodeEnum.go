// Generated ZEROPS sdk

package enum

type ClientUserRoleCodeEnum string

const (
	ClientUserRoleCodeEnumManager = ClientUserRoleCodeEnum("MANAGER")
)

func NewClientUserRoleCodeEnumFromString(value string) (out ClientUserRoleCodeEnum, err error) {
	return ClientUserRoleCodeEnum(value), nil
}

func (enum ClientUserRoleCodeEnum) String() string {
	return string(enum)
}

func (enum ClientUserRoleCodeEnum) Native() string {
	return string(enum)
}

func (enum ClientUserRoleCodeEnum) Values() []ClientUserRoleCodeEnum {
	return ClientUserRoleCodeEnumAll()
}

func (enum ClientUserRoleCodeEnum) PublicValues() []ClientUserRoleCodeEnum {
	return ClientUserRoleCodeEnumAllPublic()
}

func (enum ClientUserRoleCodeEnum) PrivateValues() []ClientUserRoleCodeEnum {
	return ClientUserRoleCodeEnumAllPrivate()
}

func (enum ClientUserRoleCodeEnum) DefaultValue() ClientUserRoleCodeEnum {
	return ClientUserRoleCodeEnumDefault()
}

func (enum ClientUserRoleCodeEnum) Is(values ...ClientUserRoleCodeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ClientUserRoleCodeEnumAllStrings() []string {
	return []string{
		string(ClientUserRoleCodeEnumManager),
	}
}

func ClientUserRoleCodeEnumAll() []ClientUserRoleCodeEnum {
	return []ClientUserRoleCodeEnum{
		ClientUserRoleCodeEnumManager,
	}
}

func ClientUserRoleCodeEnumAllPublic() []ClientUserRoleCodeEnum {
	return []ClientUserRoleCodeEnum{
		ClientUserRoleCodeEnumManager,
	}
}

func ClientUserRoleCodeEnumAllPrivate() []ClientUserRoleCodeEnum {
	return []ClientUserRoleCodeEnum{}
}

func ClientUserRoleCodeEnumDefault() ClientUserRoleCodeEnum {
	return ""
}

func (enum ClientUserRoleCodeEnum) IsManager() bool {
	return enum.Is(ClientUserRoleCodeEnumManager)
}
