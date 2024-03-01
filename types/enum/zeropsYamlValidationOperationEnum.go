// Generated ZEROPS sdk

package enum

type ZeropsYamlValidationOperationEnum string

const (
	ZeropsYamlValidationOperationEnumDeploy         = ZeropsYamlValidationOperationEnum("DEPLOY")
	ZeropsYamlValidationOperationEnumBuildAndDeploy = ZeropsYamlValidationOperationEnum("BUILD_AND_DEPLOY")
)

func NewZeropsYamlValidationOperationEnumFromString(value string) (out ZeropsYamlValidationOperationEnum, err error) {
	return ZeropsYamlValidationOperationEnum(value), nil
}

func (enum ZeropsYamlValidationOperationEnum) String() string {
	return string(enum)
}

func (enum ZeropsYamlValidationOperationEnum) Native() string {
	return string(enum)
}

func (enum ZeropsYamlValidationOperationEnum) Values() []ZeropsYamlValidationOperationEnum {
	return ZeropsYamlValidationOperationEnumAll()
}

func (enum ZeropsYamlValidationOperationEnum) PublicValues() []ZeropsYamlValidationOperationEnum {
	return ZeropsYamlValidationOperationEnumAllPublic()
}

func (enum ZeropsYamlValidationOperationEnum) PrivateValues() []ZeropsYamlValidationOperationEnum {
	return ZeropsYamlValidationOperationEnumAllPrivate()
}

func (enum ZeropsYamlValidationOperationEnum) DefaultValue() ZeropsYamlValidationOperationEnum {
	return ZeropsYamlValidationOperationEnumDefault()
}

func (enum ZeropsYamlValidationOperationEnum) Is(values ...ZeropsYamlValidationOperationEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ZeropsYamlValidationOperationEnumAllStrings() []string {
	return []string{
		string(ZeropsYamlValidationOperationEnumDeploy), string(ZeropsYamlValidationOperationEnumBuildAndDeploy),
	}
}

func ZeropsYamlValidationOperationEnumAll() []ZeropsYamlValidationOperationEnum {
	return []ZeropsYamlValidationOperationEnum{
		ZeropsYamlValidationOperationEnumDeploy, ZeropsYamlValidationOperationEnumBuildAndDeploy,
	}
}

func ZeropsYamlValidationOperationEnumAllPublic() []ZeropsYamlValidationOperationEnum {
	return []ZeropsYamlValidationOperationEnum{
		ZeropsYamlValidationOperationEnumDeploy, ZeropsYamlValidationOperationEnumBuildAndDeploy,
	}
}

func ZeropsYamlValidationOperationEnumAllPrivate() []ZeropsYamlValidationOperationEnum {
	return []ZeropsYamlValidationOperationEnum{}
}

func ZeropsYamlValidationOperationEnumDefault() ZeropsYamlValidationOperationEnum {
	return ZeropsYamlValidationOperationEnumDeploy
}

func (enum ZeropsYamlValidationOperationEnum) IsDeploy() bool {
	return enum.Is(ZeropsYamlValidationOperationEnumDeploy)
}

func (enum ZeropsYamlValidationOperationEnum) IsBuildAndDeploy() bool {
	return enum.Is(ZeropsYamlValidationOperationEnumBuildAndDeploy)
}
