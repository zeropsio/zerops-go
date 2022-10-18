// Generated ZEROPS sdk

package enum

type UserDataTypeEnum string

const (
	UserDataTypeEnumReadOnly = UserDataTypeEnum("READ_ONLY")
	UserDataTypeEnumEditable = UserDataTypeEnum("EDITABLE")
	UserDataTypeEnumUser     = UserDataTypeEnum("USER")
	UserDataTypeEnumInternal = UserDataTypeEnum("INTERNAL")
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
		string(UserDataTypeEnumReadOnly), string(UserDataTypeEnumEditable), string(UserDataTypeEnumUser), string(UserDataTypeEnumInternal),
	}
}

func UserDataTypeEnumAll() []UserDataTypeEnum {
	return []UserDataTypeEnum{
		UserDataTypeEnumReadOnly, UserDataTypeEnumEditable, UserDataTypeEnumUser, UserDataTypeEnumInternal,
	}
}

func UserDataTypeEnumAllPublic() []UserDataTypeEnum {
	return []UserDataTypeEnum{
		UserDataTypeEnumReadOnly, UserDataTypeEnumEditable, UserDataTypeEnumUser, UserDataTypeEnumInternal,
	}
}

func UserDataTypeEnumAllPrivate() []UserDataTypeEnum {
	return []UserDataTypeEnum{}
}

func UserDataTypeEnumDefault() UserDataTypeEnum {
	return UserDataTypeEnumUser
}

func (enum UserDataTypeEnum) IsReadOnly() bool {
	return enum.Is(UserDataTypeEnumReadOnly)
}

func (enum UserDataTypeEnum) IsEditable() bool {
	return enum.Is(UserDataTypeEnumEditable)
}

func (enum UserDataTypeEnum) IsUser() bool {
	return enum.Is(UserDataTypeEnumUser)
}

func (enum UserDataTypeEnum) IsInternal() bool {
	return enum.Is(UserDataTypeEnumInternal)
}
