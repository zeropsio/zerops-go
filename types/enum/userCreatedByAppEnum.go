// Generated ZEROPS sdk

package enum

type UserCreatedByAppEnum string

const (
	UserCreatedByAppEnumPlatform = UserCreatedByAppEnum("PLATFORM")
	UserCreatedByAppEnumMate     = UserCreatedByAppEnum("MATE")
	UserCreatedByAppEnumBox      = UserCreatedByAppEnum("BOX")
)

func NewUserCreatedByAppEnumFromString(value string) (out UserCreatedByAppEnum, err error) {
	return UserCreatedByAppEnum(value), nil
}

func (enum UserCreatedByAppEnum) String() string {
	return string(enum)
}

func (enum UserCreatedByAppEnum) Native() string {
	return string(enum)
}

func (enum UserCreatedByAppEnum) Values() []UserCreatedByAppEnum {
	return UserCreatedByAppEnumAll()
}

func (enum UserCreatedByAppEnum) PublicValues() []UserCreatedByAppEnum {
	return UserCreatedByAppEnumAllPublic()
}

func (enum UserCreatedByAppEnum) PrivateValues() []UserCreatedByAppEnum {
	return UserCreatedByAppEnumAllPrivate()
}

func (enum UserCreatedByAppEnum) DefaultValue() UserCreatedByAppEnum {
	return UserCreatedByAppEnumDefault()
}

func (enum UserCreatedByAppEnum) Is(values ...UserCreatedByAppEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func UserCreatedByAppEnumAllStrings() []string {
	return []string{
		string(UserCreatedByAppEnumPlatform), string(UserCreatedByAppEnumMate), string(UserCreatedByAppEnumBox),
	}
}

func UserCreatedByAppEnumAll() []UserCreatedByAppEnum {
	return []UserCreatedByAppEnum{
		UserCreatedByAppEnumPlatform, UserCreatedByAppEnumMate, UserCreatedByAppEnumBox,
	}
}

func UserCreatedByAppEnumAllPublic() []UserCreatedByAppEnum {
	return []UserCreatedByAppEnum{
		UserCreatedByAppEnumPlatform, UserCreatedByAppEnumMate, UserCreatedByAppEnumBox,
	}
}

func UserCreatedByAppEnumAllPrivate() []UserCreatedByAppEnum {
	return []UserCreatedByAppEnum{}
}

func UserCreatedByAppEnumDefault() UserCreatedByAppEnum {
	return ""
}

func (enum UserCreatedByAppEnum) IsPlatform() bool {
	return enum.Is(UserCreatedByAppEnumPlatform)
}

func (enum UserCreatedByAppEnum) IsMate() bool {
	return enum.Is(UserCreatedByAppEnumMate)
}

func (enum UserCreatedByAppEnum) IsBox() bool {
	return enum.Is(UserCreatedByAppEnumBox)
}
