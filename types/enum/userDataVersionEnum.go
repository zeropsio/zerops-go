// Generated ZEROPS sdk

package enum

type UserDataVersionEnum string

const (
	UserDataVersionEnumUser           = UserDataVersionEnum("USER")
	UserDataVersionEnumInfrastructure = UserDataVersionEnum("INFRASTRUCTURE")
)

func NewUserDataVersionEnumFromString(value string) (out UserDataVersionEnum, err error) {
	return UserDataVersionEnum(value), nil
}

func (enum UserDataVersionEnum) String() string {
	return string(enum)
}

func (enum UserDataVersionEnum) Native() string {
	return string(enum)
}

func (enum UserDataVersionEnum) Values() []UserDataVersionEnum {
	return UserDataVersionEnumAll()
}

func (enum UserDataVersionEnum) PublicValues() []UserDataVersionEnum {
	return UserDataVersionEnumAllPublic()
}

func (enum UserDataVersionEnum) PrivateValues() []UserDataVersionEnum {
	return UserDataVersionEnumAllPrivate()
}

func (enum UserDataVersionEnum) DefaultValue() UserDataVersionEnum {
	return UserDataVersionEnumDefault()
}

func (enum UserDataVersionEnum) Is(values ...UserDataVersionEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func UserDataVersionEnumAllStrings() []string {
	return []string{
		string(UserDataVersionEnumUser), string(UserDataVersionEnumInfrastructure),
	}
}

func UserDataVersionEnumAll() []UserDataVersionEnum {
	return []UserDataVersionEnum{
		UserDataVersionEnumUser, UserDataVersionEnumInfrastructure,
	}
}

func UserDataVersionEnumAllPublic() []UserDataVersionEnum {
	return []UserDataVersionEnum{
		UserDataVersionEnumUser, UserDataVersionEnumInfrastructure,
	}
}

func UserDataVersionEnumAllPrivate() []UserDataVersionEnum {
	return []UserDataVersionEnum{}
}

func UserDataVersionEnumDefault() UserDataVersionEnum {
	return UserDataVersionEnumUser
}

func (enum UserDataVersionEnum) IsUser() bool {
	return enum.Is(UserDataVersionEnumUser)
}

func (enum UserDataVersionEnum) IsInfrastructure() bool {
	return enum.Is(UserDataVersionEnumInfrastructure)
}
