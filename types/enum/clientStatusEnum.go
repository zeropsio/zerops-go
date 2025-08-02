// Generated ZEROPS sdk

package enum

type ClientStatusEnum string

const (
	ClientStatusEnumActive = ClientStatusEnum("ACTIVE")
	ClientStatusEnumBanned = ClientStatusEnum("BANNED")
)

func NewClientStatusEnumFromString(value string) (out ClientStatusEnum, err error) {
	return ClientStatusEnum(value), nil
}

func (enum ClientStatusEnum) String() string {
	return string(enum)
}

func (enum ClientStatusEnum) Native() string {
	return string(enum)
}

func (enum ClientStatusEnum) Values() []ClientStatusEnum {
	return ClientStatusEnumAll()
}

func (enum ClientStatusEnum) PublicValues() []ClientStatusEnum {
	return ClientStatusEnumAllPublic()
}

func (enum ClientStatusEnum) PrivateValues() []ClientStatusEnum {
	return ClientStatusEnumAllPrivate()
}

func (enum ClientStatusEnum) DefaultValue() ClientStatusEnum {
	return ClientStatusEnumDefault()
}

func (enum ClientStatusEnum) Is(values ...ClientStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ClientStatusEnumAllStrings() []string {
	return []string{
		string(ClientStatusEnumActive), string(ClientStatusEnumBanned),
	}
}

func ClientStatusEnumAll() []ClientStatusEnum {
	return []ClientStatusEnum{
		ClientStatusEnumActive, ClientStatusEnumBanned,
	}
}

func ClientStatusEnumAllPublic() []ClientStatusEnum {
	return []ClientStatusEnum{
		ClientStatusEnumActive, ClientStatusEnumBanned,
	}
}

func ClientStatusEnumAllPrivate() []ClientStatusEnum {
	return []ClientStatusEnum{}
}

func ClientStatusEnumDefault() ClientStatusEnum {
	return ""
}

func (enum ClientStatusEnum) IsActive() bool {
	return enum.Is(ClientStatusEnumActive)
}

func (enum ClientStatusEnum) IsBanned() bool {
	return enum.Is(ClientStatusEnumBanned)
}
