// Generated ZEROPS sdk

package enum

type UserJsonObjectTypeEnum string

const (
	UserJsonObjectTypeEnumGithub = UserJsonObjectTypeEnum("GITHUB")
	UserJsonObjectTypeEnumGitlab = UserJsonObjectTypeEnum("GITLAB")
	UserJsonObjectTypeEnumUser   = UserJsonObjectTypeEnum("USER")
	UserJsonObjectTypeEnumSystem = UserJsonObjectTypeEnum("SYSTEM")
)

func NewUserJsonObjectTypeEnumFromString(value string) (out UserJsonObjectTypeEnum, err error) {
	return UserJsonObjectTypeEnum(value), nil
}

func (enum UserJsonObjectTypeEnum) String() string {
	return string(enum)
}

func (enum UserJsonObjectTypeEnum) Native() string {
	return string(enum)
}

func (enum UserJsonObjectTypeEnum) Values() []UserJsonObjectTypeEnum {
	return UserJsonObjectTypeEnumAll()
}

func (enum UserJsonObjectTypeEnum) PublicValues() []UserJsonObjectTypeEnum {
	return UserJsonObjectTypeEnumAllPublic()
}

func (enum UserJsonObjectTypeEnum) PrivateValues() []UserJsonObjectTypeEnum {
	return UserJsonObjectTypeEnumAllPrivate()
}

func (enum UserJsonObjectTypeEnum) DefaultValue() UserJsonObjectTypeEnum {
	return UserJsonObjectTypeEnumDefault()
}

func (enum UserJsonObjectTypeEnum) Is(values ...UserJsonObjectTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func UserJsonObjectTypeEnumAllStrings() []string {
	return []string{
		string(UserJsonObjectTypeEnumGithub), string(UserJsonObjectTypeEnumGitlab), string(UserJsonObjectTypeEnumUser), string(UserJsonObjectTypeEnumSystem),
	}
}

func UserJsonObjectTypeEnumAll() []UserJsonObjectTypeEnum {
	return []UserJsonObjectTypeEnum{
		UserJsonObjectTypeEnumGithub, UserJsonObjectTypeEnumGitlab, UserJsonObjectTypeEnumUser, UserJsonObjectTypeEnumSystem,
	}
}

func UserJsonObjectTypeEnumAllPublic() []UserJsonObjectTypeEnum {
	return []UserJsonObjectTypeEnum{
		UserJsonObjectTypeEnumGithub, UserJsonObjectTypeEnumGitlab, UserJsonObjectTypeEnumUser, UserJsonObjectTypeEnumSystem,
	}
}

func UserJsonObjectTypeEnumAllPrivate() []UserJsonObjectTypeEnum {
	return []UserJsonObjectTypeEnum{}
}

func UserJsonObjectTypeEnumDefault() UserJsonObjectTypeEnum {
	return UserJsonObjectTypeEnumGithub
}

func (enum UserJsonObjectTypeEnum) IsGithub() bool {
	return enum.Is(UserJsonObjectTypeEnumGithub)
}

func (enum UserJsonObjectTypeEnum) IsGitlab() bool {
	return enum.Is(UserJsonObjectTypeEnumGitlab)
}

func (enum UserJsonObjectTypeEnum) IsUser() bool {
	return enum.Is(UserJsonObjectTypeEnumUser)
}

func (enum UserJsonObjectTypeEnum) IsSystem() bool {
	return enum.Is(UserJsonObjectTypeEnumSystem)
}
