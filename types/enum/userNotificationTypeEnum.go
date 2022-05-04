// Generated ZEROPS sdk

package enum

type UserNotificationTypeEnum string

const (
	UserNotificationTypeEnumSuccess = UserNotificationTypeEnum("SUCCESS")
	UserNotificationTypeEnumWarning = UserNotificationTypeEnum("WARNING")
	UserNotificationTypeEnumError   = UserNotificationTypeEnum("ERROR")
)

func NewUserNotificationTypeEnumFromString(value string) (out UserNotificationTypeEnum, err error) {
	return UserNotificationTypeEnum(value), nil
}

func (enum UserNotificationTypeEnum) String() string {
	return string(enum)
}

func (enum UserNotificationTypeEnum) Native() string {
	return string(enum)
}

func (enum UserNotificationTypeEnum) Values() []UserNotificationTypeEnum {
	return UserNotificationTypeEnumAll()
}

func (enum UserNotificationTypeEnum) PublicValues() []UserNotificationTypeEnum {
	return UserNotificationTypeEnumAllPublic()
}

func (enum UserNotificationTypeEnum) PrivateValues() []UserNotificationTypeEnum {
	return UserNotificationTypeEnumAllPrivate()
}

func (enum UserNotificationTypeEnum) DefaultValue() UserNotificationTypeEnum {
	return UserNotificationTypeEnumDefault()
}

func (enum UserNotificationTypeEnum) Is(values ...UserNotificationTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func UserNotificationTypeEnumAllStrings() []string {
	return []string{
		string(UserNotificationTypeEnumSuccess), string(UserNotificationTypeEnumWarning), string(UserNotificationTypeEnumError),
	}
}

func UserNotificationTypeEnumAll() []UserNotificationTypeEnum {
	return []UserNotificationTypeEnum{
		UserNotificationTypeEnumSuccess, UserNotificationTypeEnumWarning, UserNotificationTypeEnumError,
	}
}

func UserNotificationTypeEnumAllPublic() []UserNotificationTypeEnum {
	return []UserNotificationTypeEnum{
		UserNotificationTypeEnumSuccess, UserNotificationTypeEnumWarning, UserNotificationTypeEnumError,
	}
}

func UserNotificationTypeEnumAllPrivate() []UserNotificationTypeEnum {
	return []UserNotificationTypeEnum{}
}

func UserNotificationTypeEnumDefault() UserNotificationTypeEnum {
	return UserNotificationTypeEnumSuccess
}

func (enum UserNotificationTypeEnum) IsSuccess() bool {
	return enum.Is(UserNotificationTypeEnumSuccess)
}

func (enum UserNotificationTypeEnum) IsWarning() bool {
	return enum.Is(UserNotificationTypeEnumWarning)
}

func (enum UserNotificationTypeEnum) IsError() bool {
	return enum.Is(UserNotificationTypeEnumError)
}
