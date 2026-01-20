// Generated ZEROPS sdk

package enum

type UserStatusEnum string

const (
	UserStatusEnumActive = UserStatusEnum("ACTIVE")
	UserStatusEnumBanned = UserStatusEnum("BANNED")
)

func NewUserStatusEnumFromString(value string) (out UserStatusEnum, err error) {
	return UserStatusEnum(value), nil
}

func (enum UserStatusEnum) String() string {
	return string(enum)
}

func (enum UserStatusEnum) Native() string {
	return string(enum)
}

func (enum UserStatusEnum) Values() []UserStatusEnum {
	return UserStatusEnumAll()
}

func (enum UserStatusEnum) PublicValues() []UserStatusEnum {
	return UserStatusEnumAllPublic()
}

func (enum UserStatusEnum) PrivateValues() []UserStatusEnum {
	return UserStatusEnumAllPrivate()
}

func (enum UserStatusEnum) DefaultValue() UserStatusEnum {
	return UserStatusEnumDefault()
}

func (enum UserStatusEnum) Is(values ...UserStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func UserStatusEnumAllStrings() []string {
	return []string{
		string(UserStatusEnumActive), string(UserStatusEnumBanned),
	}
}

func UserStatusEnumAll() []UserStatusEnum {
	return []UserStatusEnum{
		UserStatusEnumActive, UserStatusEnumBanned,
	}
}

func UserStatusEnumAllPublic() []UserStatusEnum {
	return []UserStatusEnum{
		UserStatusEnumActive, UserStatusEnumBanned,
	}
}

func UserStatusEnumAllPrivate() []UserStatusEnum {
	return []UserStatusEnum{}
}

func UserStatusEnumDefault() UserStatusEnum {
	return ""
}

func (enum UserStatusEnum) IsActive() bool {
	return enum.Is(UserStatusEnumActive)
}

func (enum UserStatusEnum) IsBanned() bool {
	return enum.Is(UserStatusEnumBanned)
}
