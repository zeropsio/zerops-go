// Generated ZEROPS sdk

package enum

type GetProjectEnvFileOverrideEnvIsolationEnum string

const (
	GetProjectEnvFileOverrideEnvIsolationEnumNo      = GetProjectEnvFileOverrideEnvIsolationEnum("no")
	GetProjectEnvFileOverrideEnvIsolationEnumNone    = GetProjectEnvFileOverrideEnvIsolationEnum("none")
	GetProjectEnvFileOverrideEnvIsolationEnumService = GetProjectEnvFileOverrideEnvIsolationEnum("service")
)

func NewGetProjectEnvFileOverrideEnvIsolationEnumFromString(value string) (out GetProjectEnvFileOverrideEnvIsolationEnum, err error) {
	return GetProjectEnvFileOverrideEnvIsolationEnum(value), nil
}

func (enum GetProjectEnvFileOverrideEnvIsolationEnum) String() string {
	return string(enum)
}

func (enum GetProjectEnvFileOverrideEnvIsolationEnum) Native() string {
	return string(enum)
}

func (enum GetProjectEnvFileOverrideEnvIsolationEnum) Values() []GetProjectEnvFileOverrideEnvIsolationEnum {
	return GetProjectEnvFileOverrideEnvIsolationEnumAll()
}

func (enum GetProjectEnvFileOverrideEnvIsolationEnum) PublicValues() []GetProjectEnvFileOverrideEnvIsolationEnum {
	return GetProjectEnvFileOverrideEnvIsolationEnumAllPublic()
}

func (enum GetProjectEnvFileOverrideEnvIsolationEnum) PrivateValues() []GetProjectEnvFileOverrideEnvIsolationEnum {
	return GetProjectEnvFileOverrideEnvIsolationEnumAllPrivate()
}

func (enum GetProjectEnvFileOverrideEnvIsolationEnum) DefaultValue() GetProjectEnvFileOverrideEnvIsolationEnum {
	return GetProjectEnvFileOverrideEnvIsolationEnumDefault()
}

func (enum GetProjectEnvFileOverrideEnvIsolationEnum) Is(values ...GetProjectEnvFileOverrideEnvIsolationEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func GetProjectEnvFileOverrideEnvIsolationEnumAllStrings() []string {
	return []string{
		string(GetProjectEnvFileOverrideEnvIsolationEnumNo), string(GetProjectEnvFileOverrideEnvIsolationEnumNone), string(GetProjectEnvFileOverrideEnvIsolationEnumService),
	}
}

func GetProjectEnvFileOverrideEnvIsolationEnumAll() []GetProjectEnvFileOverrideEnvIsolationEnum {
	return []GetProjectEnvFileOverrideEnvIsolationEnum{
		GetProjectEnvFileOverrideEnvIsolationEnumNo, GetProjectEnvFileOverrideEnvIsolationEnumNone, GetProjectEnvFileOverrideEnvIsolationEnumService,
	}
}

func GetProjectEnvFileOverrideEnvIsolationEnumAllPublic() []GetProjectEnvFileOverrideEnvIsolationEnum {
	return []GetProjectEnvFileOverrideEnvIsolationEnum{
		GetProjectEnvFileOverrideEnvIsolationEnumNo, GetProjectEnvFileOverrideEnvIsolationEnumNone, GetProjectEnvFileOverrideEnvIsolationEnumService,
	}
}

func GetProjectEnvFileOverrideEnvIsolationEnumAllPrivate() []GetProjectEnvFileOverrideEnvIsolationEnum {
	return []GetProjectEnvFileOverrideEnvIsolationEnum{}
}

func GetProjectEnvFileOverrideEnvIsolationEnumDefault() GetProjectEnvFileOverrideEnvIsolationEnum {
	return GetProjectEnvFileOverrideEnvIsolationEnumNo
}

func (enum GetProjectEnvFileOverrideEnvIsolationEnum) IsNo() bool {
	return enum.Is(GetProjectEnvFileOverrideEnvIsolationEnumNo)
}

func (enum GetProjectEnvFileOverrideEnvIsolationEnum) IsNone() bool {
	return enum.Is(GetProjectEnvFileOverrideEnvIsolationEnumNone)
}

func (enum GetProjectEnvFileOverrideEnvIsolationEnum) IsService() bool {
	return enum.Is(GetProjectEnvFileOverrideEnvIsolationEnumService)
}
