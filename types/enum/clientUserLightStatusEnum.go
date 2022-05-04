// Generated ZEROPS sdk

package enum

type ClientUserLightStatusEnum string

const (
	ClientUserLightStatusEnumActive               = ClientUserLightStatusEnum("ACTIVE")
	ClientUserLightStatusEnumWaitingAuthorization = ClientUserLightStatusEnum("WAITING_AUTHORIZATION")
	ClientUserLightStatusEnumBeingDeleted         = ClientUserLightStatusEnum("BEING_DELETED")
)

func NewClientUserLightStatusEnumFromString(value string) (out ClientUserLightStatusEnum, err error) {
	return ClientUserLightStatusEnum(value), nil
}

func (enum ClientUserLightStatusEnum) String() string {
	return string(enum)
}

func (enum ClientUserLightStatusEnum) Native() string {
	return string(enum)
}

func (enum ClientUserLightStatusEnum) Values() []ClientUserLightStatusEnum {
	return ClientUserLightStatusEnumAll()
}

func (enum ClientUserLightStatusEnum) PublicValues() []ClientUserLightStatusEnum {
	return ClientUserLightStatusEnumAllPublic()
}

func (enum ClientUserLightStatusEnum) PrivateValues() []ClientUserLightStatusEnum {
	return ClientUserLightStatusEnumAllPrivate()
}

func (enum ClientUserLightStatusEnum) DefaultValue() ClientUserLightStatusEnum {
	return ClientUserLightStatusEnumDefault()
}

func (enum ClientUserLightStatusEnum) Is(values ...ClientUserLightStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ClientUserLightStatusEnumAllStrings() []string {
	return []string{
		string(ClientUserLightStatusEnumActive), string(ClientUserLightStatusEnumWaitingAuthorization), string(ClientUserLightStatusEnumBeingDeleted),
	}
}

func ClientUserLightStatusEnumAll() []ClientUserLightStatusEnum {
	return []ClientUserLightStatusEnum{
		ClientUserLightStatusEnumActive, ClientUserLightStatusEnumWaitingAuthorization, ClientUserLightStatusEnumBeingDeleted,
	}
}

func ClientUserLightStatusEnumAllPublic() []ClientUserLightStatusEnum {
	return []ClientUserLightStatusEnum{}
}

func ClientUserLightStatusEnumAllPrivate() []ClientUserLightStatusEnum {
	return []ClientUserLightStatusEnum{
		ClientUserLightStatusEnumActive, ClientUserLightStatusEnumWaitingAuthorization, ClientUserLightStatusEnumBeingDeleted,
	}
}

func ClientUserLightStatusEnumDefault() ClientUserLightStatusEnum {
	return ClientUserLightStatusEnumActive
}

func (enum ClientUserLightStatusEnum) IsActive() bool {
	return enum.Is(ClientUserLightStatusEnumActive)
}

func (enum ClientUserLightStatusEnum) IsWaitingAuthorization() bool {
	return enum.Is(ClientUserLightStatusEnumWaitingAuthorization)
}

func (enum ClientUserLightStatusEnum) IsBeingDeleted() bool {
	return enum.Is(ClientUserLightStatusEnumBeingDeleted)
}
