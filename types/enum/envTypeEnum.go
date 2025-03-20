// Generated ZEROPS sdk

package enum

type EnvTypeEnum string

const (
	EnvTypeEnumUser   = EnvTypeEnum("USER")
	EnvTypeEnumSystem = EnvTypeEnum("SYSTEM")
)

func NewEnvTypeEnumFromString(value string) (out EnvTypeEnum, err error) {
	return EnvTypeEnum(value), nil
}

func (enum EnvTypeEnum) String() string {
	return string(enum)
}

func (enum EnvTypeEnum) Native() string {
	return string(enum)
}

func (enum EnvTypeEnum) Values() []EnvTypeEnum {
	return EnvTypeEnumAll()
}

func (enum EnvTypeEnum) PublicValues() []EnvTypeEnum {
	return EnvTypeEnumAllPublic()
}

func (enum EnvTypeEnum) PrivateValues() []EnvTypeEnum {
	return EnvTypeEnumAllPrivate()
}

func (enum EnvTypeEnum) DefaultValue() EnvTypeEnum {
	return EnvTypeEnumDefault()
}

func (enum EnvTypeEnum) Is(values ...EnvTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func EnvTypeEnumAllStrings() []string {
	return []string{
		string(EnvTypeEnumUser), string(EnvTypeEnumSystem),
	}
}

func EnvTypeEnumAll() []EnvTypeEnum {
	return []EnvTypeEnum{
		EnvTypeEnumUser, EnvTypeEnumSystem,
	}
}

func EnvTypeEnumAllPublic() []EnvTypeEnum {
	return []EnvTypeEnum{
		EnvTypeEnumUser, EnvTypeEnumSystem,
	}
}

func EnvTypeEnumAllPrivate() []EnvTypeEnum {
	return []EnvTypeEnum{}
}

func EnvTypeEnumDefault() EnvTypeEnum {
	return EnvTypeEnumUser
}

func (enum EnvTypeEnum) IsUser() bool {
	return enum.Is(EnvTypeEnumUser)
}

func (enum EnvTypeEnum) IsSystem() bool {
	return enum.Is(EnvTypeEnumSystem)
}
