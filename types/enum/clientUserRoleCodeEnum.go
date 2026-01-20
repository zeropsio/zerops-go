// Generated ZEROPS sdk

package enum

type ClientUserRoleCodeEnum string

const (
	ClientUserRoleCodeEnumOwner     = ClientUserRoleCodeEnum("OWNER")
	ClientUserRoleCodeEnumAdmin     = ClientUserRoleCodeEnum("ADMIN")
	ClientUserRoleCodeEnumBasicUser = ClientUserRoleCodeEnum("BASIC_USER")
	ClientUserRoleCodeEnumReadOnly  = ClientUserRoleCodeEnum("READ_ONLY")
	ClientUserRoleCodeEnumNoAccess  = ClientUserRoleCodeEnum("NO_ACCESS")
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
		string(ClientUserRoleCodeEnumOwner), string(ClientUserRoleCodeEnumAdmin), string(ClientUserRoleCodeEnumBasicUser), string(ClientUserRoleCodeEnumReadOnly), string(ClientUserRoleCodeEnumNoAccess),
	}
}

func ClientUserRoleCodeEnumAll() []ClientUserRoleCodeEnum {
	return []ClientUserRoleCodeEnum{
		ClientUserRoleCodeEnumOwner, ClientUserRoleCodeEnumAdmin, ClientUserRoleCodeEnumBasicUser, ClientUserRoleCodeEnumReadOnly, ClientUserRoleCodeEnumNoAccess,
	}
}

func ClientUserRoleCodeEnumAllPublic() []ClientUserRoleCodeEnum {
	return []ClientUserRoleCodeEnum{
		ClientUserRoleCodeEnumOwner, ClientUserRoleCodeEnumAdmin, ClientUserRoleCodeEnumBasicUser, ClientUserRoleCodeEnumReadOnly, ClientUserRoleCodeEnumNoAccess,
	}
}

func ClientUserRoleCodeEnumAllPrivate() []ClientUserRoleCodeEnum {
	return []ClientUserRoleCodeEnum{}
}

func ClientUserRoleCodeEnumDefault() ClientUserRoleCodeEnum {
	return ClientUserRoleCodeEnumBasicUser
}

func (enum ClientUserRoleCodeEnum) IsOwner() bool {
	return enum.Is(ClientUserRoleCodeEnumOwner)
}

func (enum ClientUserRoleCodeEnum) IsAdmin() bool {
	return enum.Is(ClientUserRoleCodeEnumAdmin)
}

func (enum ClientUserRoleCodeEnum) IsBasicUser() bool {
	return enum.Is(ClientUserRoleCodeEnumBasicUser)
}

func (enum ClientUserRoleCodeEnum) IsReadOnly() bool {
	return enum.Is(ClientUserRoleCodeEnumReadOnly)
}

func (enum ClientUserRoleCodeEnum) IsNoAccess() bool {
	return enum.Is(ClientUserRoleCodeEnumNoAccess)
}
