// Generated ZEROPS sdk

package enum

type LoginAppTypeEnum string

const (
	LoginAppTypeEnumPublic       = LoginAppTypeEnum("PUBLIC")
	LoginAppTypeEnumConfidential = LoginAppTypeEnum("CONFIDENTIAL")
)

func NewLoginAppTypeEnumFromString(value string) (out LoginAppTypeEnum, err error) {
	return LoginAppTypeEnum(value), nil
}

func (enum LoginAppTypeEnum) String() string {
	return string(enum)
}

func (enum LoginAppTypeEnum) Native() string {
	return string(enum)
}

func (enum LoginAppTypeEnum) Values() []LoginAppTypeEnum {
	return LoginAppTypeEnumAll()
}

func (enum LoginAppTypeEnum) PublicValues() []LoginAppTypeEnum {
	return LoginAppTypeEnumAllPublic()
}

func (enum LoginAppTypeEnum) PrivateValues() []LoginAppTypeEnum {
	return LoginAppTypeEnumAllPrivate()
}

func (enum LoginAppTypeEnum) DefaultValue() LoginAppTypeEnum {
	return LoginAppTypeEnumDefault()
}

func (enum LoginAppTypeEnum) Is(values ...LoginAppTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func LoginAppTypeEnumAllStrings() []string {
	return []string{
		string(LoginAppTypeEnumPublic), string(LoginAppTypeEnumConfidential),
	}
}

func LoginAppTypeEnumAll() []LoginAppTypeEnum {
	return []LoginAppTypeEnum{
		LoginAppTypeEnumPublic, LoginAppTypeEnumConfidential,
	}
}

func LoginAppTypeEnumAllPublic() []LoginAppTypeEnum {
	return []LoginAppTypeEnum{
		LoginAppTypeEnumPublic, LoginAppTypeEnumConfidential,
	}
}

func LoginAppTypeEnumAllPrivate() []LoginAppTypeEnum {
	return []LoginAppTypeEnum{}
}

func LoginAppTypeEnumDefault() LoginAppTypeEnum {
	return LoginAppTypeEnumPublic
}

func (enum LoginAppTypeEnum) IsPublic() bool {
	return enum.Is(LoginAppTypeEnumPublic)
}

func (enum LoginAppTypeEnum) IsConfidential() bool {
	return enum.Is(LoginAppTypeEnumConfidential)
}
