// Generated ZEROPS sdk

package enum

type ClientUserStatusEnum string

const (
	ClientUserStatusEnumWaitingAuthorization = ClientUserStatusEnum("WAITING_AUTHORIZATION")
	ClientUserStatusEnumActive               = ClientUserStatusEnum("ACTIVE")
	ClientUserStatusEnumBeingDeleted         = ClientUserStatusEnum("BEING_DELETED")
)

func NewClientUserStatusEnumFromString(value string) (out ClientUserStatusEnum, err error) {
	return ClientUserStatusEnum(value), nil
}

func (enum ClientUserStatusEnum) String() string {
	return string(enum)
}

func (enum ClientUserStatusEnum) Native() string {
	return string(enum)
}

func (enum ClientUserStatusEnum) Values() []ClientUserStatusEnum {
	return ClientUserStatusEnumAll()
}

func (enum ClientUserStatusEnum) PublicValues() []ClientUserStatusEnum {
	return ClientUserStatusEnumAllPublic()
}

func (enum ClientUserStatusEnum) PrivateValues() []ClientUserStatusEnum {
	return ClientUserStatusEnumAllPrivate()
}

func (enum ClientUserStatusEnum) DefaultValue() ClientUserStatusEnum {
	return ClientUserStatusEnumDefault()
}

func (enum ClientUserStatusEnum) Is(values ...ClientUserStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ClientUserStatusEnumAllStrings() []string {
	return []string{
		string(ClientUserStatusEnumWaitingAuthorization), string(ClientUserStatusEnumActive), string(ClientUserStatusEnumBeingDeleted),
	}
}

func ClientUserStatusEnumAll() []ClientUserStatusEnum {
	return []ClientUserStatusEnum{
		ClientUserStatusEnumWaitingAuthorization, ClientUserStatusEnumActive, ClientUserStatusEnumBeingDeleted,
	}
}

func ClientUserStatusEnumAllPublic() []ClientUserStatusEnum {
	return []ClientUserStatusEnum{
		ClientUserStatusEnumWaitingAuthorization, ClientUserStatusEnumActive, ClientUserStatusEnumBeingDeleted,
	}
}

func ClientUserStatusEnumAllPrivate() []ClientUserStatusEnum {
	return []ClientUserStatusEnum{}
}

func ClientUserStatusEnumDefault() ClientUserStatusEnum {
	return ""
}

func (enum ClientUserStatusEnum) IsWaitingAuthorization() bool {
	return enum.Is(ClientUserStatusEnumWaitingAuthorization)
}

func (enum ClientUserStatusEnum) IsActive() bool {
	return enum.Is(ClientUserStatusEnumActive)
}

func (enum ClientUserStatusEnum) IsBeingDeleted() bool {
	return enum.Is(ClientUserStatusEnumBeingDeleted)
}
