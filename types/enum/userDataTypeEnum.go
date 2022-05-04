// Generated ZEROPS sdk

package enum

type UserDataTypeEnum string

const (
	UserDataTypeEnumRestricted = UserDataTypeEnum("RESTRICTED")
	UserDataTypeEnumCommon     = UserDataTypeEnum("COMMON")
	UserDataTypeEnumUser       = UserDataTypeEnum("USER")
	UserDataTypeEnumInternal   = UserDataTypeEnum("INTERNAL")
)

func NewUserDataTypeEnumFromString(value string) (out UserDataTypeEnum, err error) {
	return UserDataTypeEnum(value), nil
}

func (enum UserDataTypeEnum) String() string {
	return string(enum)
}

func (enum UserDataTypeEnum) Native() string {
	return string(enum)
}

func (enum UserDataTypeEnum) Values() []UserDataTypeEnum {
	return UserDataTypeEnumAll()
}

func (enum UserDataTypeEnum) PublicValues() []UserDataTypeEnum {
	return UserDataTypeEnumAllPublic()
}

func (enum UserDataTypeEnum) PrivateValues() []UserDataTypeEnum {
	return UserDataTypeEnumAllPrivate()
}

func (enum UserDataTypeEnum) DefaultValue() UserDataTypeEnum {
	return UserDataTypeEnumDefault()
}

func (enum UserDataTypeEnum) Is(values ...UserDataTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func UserDataTypeEnumAllStrings() []string {
	return []string{
		string(UserDataTypeEnumRestricted), string(UserDataTypeEnumCommon), string(UserDataTypeEnumUser), string(UserDataTypeEnumInternal),
	}
}

func UserDataTypeEnumAll() []UserDataTypeEnum {
	return []UserDataTypeEnum{
		UserDataTypeEnumRestricted, UserDataTypeEnumCommon, UserDataTypeEnumUser, UserDataTypeEnumInternal,
	}
}

func UserDataTypeEnumAllPublic() []UserDataTypeEnum {
	return []UserDataTypeEnum{
		UserDataTypeEnumRestricted, UserDataTypeEnumCommon, UserDataTypeEnumUser, UserDataTypeEnumInternal,
	}
}

func UserDataTypeEnumAllPrivate() []UserDataTypeEnum {
	return []UserDataTypeEnum{}
}

func UserDataTypeEnumDefault() UserDataTypeEnum {
	return UserDataTypeEnumUser
}

func (enum UserDataTypeEnum) IsRestricted() bool {
	return enum.Is(UserDataTypeEnumRestricted)
}

func (enum UserDataTypeEnum) IsCommon() bool {
	return enum.Is(UserDataTypeEnumCommon)
}

func (enum UserDataTypeEnum) IsUser() bool {
	return enum.Is(UserDataTypeEnumUser)
}

func (enum UserDataTypeEnum) IsInternal() bool {
	return enum.Is(UserDataTypeEnumInternal)
}
