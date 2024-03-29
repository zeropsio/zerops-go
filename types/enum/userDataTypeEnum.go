// Generated ZEROPS sdk

package enum

type UserDataTypeEnum string

const (
	UserDataTypeEnumReadOnly = UserDataTypeEnum("READ_ONLY")
	UserDataTypeEnumEditable = UserDataTypeEnum("EDITABLE")
	UserDataTypeEnumSecret   = UserDataTypeEnum("SECRET")
	UserDataTypeEnumInternal = UserDataTypeEnum("INTERNAL")
	UserDataTypeEnumEnv      = UserDataTypeEnum("ENV")
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
		string(UserDataTypeEnumReadOnly), string(UserDataTypeEnumEditable), string(UserDataTypeEnumSecret), string(UserDataTypeEnumInternal), string(UserDataTypeEnumEnv),
	}
}

func UserDataTypeEnumAll() []UserDataTypeEnum {
	return []UserDataTypeEnum{
		UserDataTypeEnumReadOnly, UserDataTypeEnumEditable, UserDataTypeEnumSecret, UserDataTypeEnumInternal, UserDataTypeEnumEnv,
	}
}

func UserDataTypeEnumAllPublic() []UserDataTypeEnum {
	return []UserDataTypeEnum{
		UserDataTypeEnumReadOnly, UserDataTypeEnumEditable, UserDataTypeEnumSecret, UserDataTypeEnumInternal, UserDataTypeEnumEnv,
	}
}

func UserDataTypeEnumAllPrivate() []UserDataTypeEnum {
	return []UserDataTypeEnum{}
}

func UserDataTypeEnumDefault() UserDataTypeEnum {
	return UserDataTypeEnumSecret
}

func (enum UserDataTypeEnum) IsReadOnly() bool {
	return enum.Is(UserDataTypeEnumReadOnly)
}

func (enum UserDataTypeEnum) IsEditable() bool {
	return enum.Is(UserDataTypeEnumEditable)
}

func (enum UserDataTypeEnum) IsSecret() bool {
	return enum.Is(UserDataTypeEnumSecret)
}

func (enum UserDataTypeEnum) IsInternal() bool {
	return enum.Is(UserDataTypeEnumInternal)
}

func (enum UserDataTypeEnum) IsEnv() bool {
	return enum.Is(UserDataTypeEnumEnv)
}
